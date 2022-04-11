package sv21

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"pledge-backend/config"
	tokengo "pledge-backend/contract/bindings/tokenv22"
	"pledge-backend/db"
	"pledge-backend/log"
	"pledge-backend/schedule/models"
	"pledge-backend/utils"
	"strings"
)

type poolService struct{}

func NewPool() *poolService {
	return &poolService{}
}

func (s *poolService) UpdateAllPoolInfo() {

	s.UpdatePoolInfo(config.Config.TestNet.PledgePoolToken, config.Config.TestNet.NetUrl, config.Config.TestNet.ChainId)

	s.UpdatePoolInfo(config.Config.MainNet.PledgePoolToken, config.Config.MainNet.NetUrl, config.Config.MainNet.ChainId)

}

func (s *poolService) UpdatePoolInfo(contractAddress, network, chainId string) {

	log.Logger.Sugar().Info("UpdatePoolInfo ", contractAddress+" "+network)
	ethereumConn, err := ethclient.Dial(network)
	if nil != err {
		log.Logger.Error(err.Error())
		return
	}
	pledgePoolToken, err := tokengo.NewPledgePoolToken(common.HexToAddress(contractAddress), ethereumConn)
	if nil != err {
		log.Logger.Error(err.Error())
		return
	}

	// borrowFee
	borrowFee, err := pledgePoolToken.PledgePoolTokenCaller.BorrowFee(nil)

	// lendFee
	lendFee, err := pledgePoolToken.PledgePoolTokenCaller.LendFee(nil)

	//poolLength
	pLength, err := pledgePoolToken.PledgePoolTokenCaller.PoolLength(nil)
	if nil != err {
		log.Logger.Error(err.Error())
		return
	}

	for i := 0; i <= int(pLength.Int64())-1; i++ {

		log.Logger.Sugar().Info("UpdatePoolInfo ", i)
		poolId := utils.IntToString(i + 1)
		baseInfo, err := pledgePoolToken.PledgePoolTokenCaller.PoolBaseInfo(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Logger.Sugar().Info("UpdatePoolInfo PoolBaseInfo err", poolId, err)
			continue
		}

		_, borrowToken := models.NewTokenInfo().GetTokenInfo(baseInfo.BorrowToken.String(), chainId)
		_, lendToken := models.NewTokenInfo().GetTokenInfo(baseInfo.LendToken.String(), chainId)

		lendTokenJson, _ := json.Marshal(models.LendToken{
			LendFee:    lendFee.String(),
			TokenLogo:  lendToken.Logo,
			TokenName:  lendToken.Symbol,
			TokenPrice: lendToken.Price,
		})
		borrowTokenJson, _ := json.Marshal(models.BorrowToken{
			BorrowFee:  borrowFee.String(),
			TokenLogo:  borrowToken.Logo,
			TokenName:  borrowToken.Symbol,
			TokenPrice: borrowToken.Price,
		})

		poolBase := models.PoolBase{
			SettleTime:             baseInfo.SettleTime.String(),
			PoolId:                 utils.StringToInt(poolId),
			ChainId:                chainId,
			EndTime:                baseInfo.EndTime.String(),
			InterestRate:           baseInfo.InterestRate.String(),
			MaxSupply:              baseInfo.MaxSupply.String(),
			LendSupply:             baseInfo.LendSupply.String(),
			BorrowSupply:           baseInfo.BorrowSupply.String(),
			MartgageRate:           baseInfo.MartgageRate.String(),
			LendToken:              baseInfo.LendToken.String(),
			LendTokenSymbol:        lendToken.Symbol,
			LendTokenInfo:          string(lendTokenJson),
			BorrowToken:            baseInfo.BorrowToken.String(),
			BorrowTokenSymbol:      borrowToken.Symbol,
			BorrowTokenInfo:        string(borrowTokenJson),
			State:                  utils.IntToString(int(baseInfo.State)),
			SpCoin:                 baseInfo.SpCoin.String(),
			JpCoin:                 baseInfo.JpCoin.String(),
			AutoLiquidateThreshold: baseInfo.AutoLiquidateThreshold.String(),
		}

		hasInfoData, byteBaseInfoStr, baseInfoMd5Str := s.GetPoolMd5(&poolBase, "base_info:pool_"+chainId+"_"+poolId)
		if !hasInfoData || (baseInfoMd5Str != byteBaseInfoStr) { // have new data
			//tokenInfo
			err = models.NewPoolBase().SavePoolBase(chainId, poolId, &poolBase)
			if err != nil {
				log.Logger.Sugar().Error("SavePoolBase err ", chainId, poolId)
			}
			db.RedisSet("base_info:pool_"+chainId+"_"+poolId, baseInfoMd5Str, 60*30) //The expiration time is set to prevent hsah collision
		}

		dataInfo, err := pledgePoolToken.PledgePoolTokenCaller.PoolDataInfo(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Logger.Sugar().Info("UpdatePoolInfo PoolBaseInfo err", poolId, err)
			continue
		}

		hasPoolData, byteDataInfoStr, dataInfoMd5Str := s.GetPoolMd5(&poolBase, "data_info:pool_"+chainId+"_"+poolId)
		if !hasPoolData || (dataInfoMd5Str != byteDataInfoStr) { // have new data
			poolData := models.PoolData{
				PoolId:                 poolId,
				ChainId:                chainId,
				FinishAmountBorrow:     dataInfo.FinishAmountBorrow.String(),
				FinishAmountLend:       dataInfo.FinishAmountLend.String(),
				LiquidationAmounBorrow: dataInfo.LiquidationAmounBorrow.String(),
				LiquidationAmounLend:   dataInfo.LiquidationAmounLend.String(),
				SettleAmountBorrow:     dataInfo.SettleAmountBorrow.String(),
				SettleAmountLend:       dataInfo.SettleAmountLend.String(),
			}
			err = models.NewPoolData().SavePoolData(chainId, poolId, &poolData)
			if err != nil {
				log.Logger.Sugar().Error("SavePoolData err ", chainId, poolId)
			}
			db.RedisSet("data_info:pool_"+chainId+"_"+poolId, dataInfoMd5Str, 60*30) //The expiration time is set to prevent hsah collision
		}
	}
}

func (s *poolService) GetPoolMd5(baseInfo *models.PoolBase, key string) (bool, string, string) {
	baseInfoBytes, _ := json.Marshal(baseInfo)
	baseInfoMd5Str := utils.Md5(string(baseInfoBytes))
	resInfoBytes, _ := db.RedisGet(key)
	if len(resInfoBytes) > 0 {
		return true, strings.Trim(string(resInfoBytes), `"`), baseInfoMd5Str
	} else {
		return false, strings.Trim(string(resInfoBytes), `"`), baseInfoMd5Str
	}
}
