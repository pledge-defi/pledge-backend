package sv22

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"pledge-backend/config"
	tokengo "pledge-backend/contract/tokengo/tokenv22"
	"pledge-backend/db"
	"pledge-backend/log"
	"pledge-backend/schedule/models"
	"pledge-backend/utils"
)

type TokenPrice struct{}

func NewTokenPrice() *TokenPrice {
	return &TokenPrice{}
}

// UpdateContractPrice update contract price
func (s *TokenPrice) UpdateContractPrice() {
	var tokens []models.TokenInfo
	db.Mysql.Table("token_info").Find(&tokens)
	for _, t := range tokens {
		if t.Token == "" {
			log.Logger.Sugar().Error("UpdateContractPrice token empty ", t.Symbol, t.ChainId)
			continue
		}
		err := errors.New("")
		var price int64 = 0
		if t.ChainId == "97" {
			err, price = s.GetTestNetTokenPrice(t.Token)
		} else if t.ChainId == "56" {
			err, price = s.GetMainNetTokenPrice(t.Token)
		} else {
			log.Logger.Sugar().Error("UpdateContractPrice chain_id err ", t.Symbol, t.ChainId)
			continue
		}
		if err != nil {
			log.Logger.Sugar().Error("UpdateContractPrice err ", t.Symbol, t.ChainId, err)
			continue
		}

		hasNewData, err := s.CheckPriceData(t.Token, t.ChainId, utils.Int64ToString(price))
		if err != nil {
			log.Logger.Sugar().Error("UpdateContractPrice CheckPriceData err ", err)
			continue
		}

		if hasNewData {
			err = s.SavePriceData(t.Token, t.ChainId, utils.Int64ToString(price))
			if err != nil {
				log.Logger.Sugar().Error("UpdateContractPrice SavePriceData err ", err)
				continue
			}
		}
	}
}

// GetMainNetTokenPrice get contract price on main net
func (s *TokenPrice) GetMainNetTokenPrice(token string) (error, int64) {
	ethereumConn, err := ethclient.Dial(config.Config.MainNet.NetUrl)
	if nil != err {
		log.Logger.Error(err.Error())
		return err, 0
	}

	bscPledgeOracleMainnetToken, err := tokengo.NewBscPledgeOracleMainnetToken(common.HexToAddress(config.Config.MainNet.BscPledgeOracleToken), ethereumConn)
	if nil != err {
		log.Logger.Error(err.Error())
		return err, 0
	}

	price, _ := bscPledgeOracleMainnetToken.GetPrice(nil, common.HexToAddress(token))
	if nil != err {
		log.Logger.Error(err.Error())
		return err, 0
	}

	return nil, price.Int64()
}

// GetTestNetTokenPrice get contract price on test net
func (s *TokenPrice) GetTestNetTokenPrice(token string) (error, int64) {
	ethereumConn, err := ethclient.Dial(config.Config.TestNet.NetUrl)
	if nil != err {
		log.Logger.Error(err.Error())
		return err, 0
	}

	bscPledgeOracleTestnetToken, err := tokengo.NewBscPledgeOracleTestnetToken(common.HexToAddress(config.Config.TestNet.BscPledgeOracleToken), ethereumConn)
	if nil != err {
		log.Logger.Error(err.Error())
		return err, 0
	}

	price, err := bscPledgeOracleTestnetToken.GetPrice(nil, common.HexToAddress(token))
	if nil != err {
		log.Logger.Error(err.Error())
		return err, 0
	}

	return nil, price.Int64()
}

// CheckPriceData Saving price data to redis if it has new price
func (s *TokenPrice) CheckPriceData(token, chainId, price string) (bool, error) {
	redisTokenInfoBytes, err := db.RedisGet("token_info:" + token + "_" + chainId)
	if err != nil {
		if err.Error() == "redigo: nil returned" {
			err = db.RedisSet("token_info:"+token+"_"+chainId, models.RedisTokenInfo{
				Token:   token,
				ChainId: chainId,
				Price:   price,
			}, 120)
			if err != nil {
				log.Logger.Error(err.Error())
				return false, err
			}
		} else {
			log.Logger.Sugar().Error("UpdateContractSymbol get symbol from redis err ", token, chainId, err)
			return false, err
		}
	} else {
		redisTokenInfo := models.RedisTokenInfo{}
		err = json.Unmarshal(redisTokenInfoBytes, &redisTokenInfo)
		if err != nil {
			log.Logger.Error(err.Error())
			return false, err
		}

		if redisTokenInfo.Price == price {
			return false, nil
		}

		err = db.RedisSet("token_info:"+token+"_"+chainId, models.RedisTokenInfo{
			Logo:    redisTokenInfo.Logo,
			Token:   redisTokenInfo.Token,
			Symbol:  redisTokenInfo.Symbol,
			ChainId: redisTokenInfo.ChainId,
			Price:   price,
		}, 120)
		if err != nil {
			log.Logger.Error(err.Error())
			return true, err
		}
	}
	return true, nil
}

// SavePriceData Saving price data to mysql if it has new price
func (s *TokenPrice) SavePriceData(token, chainId, price string) error {

	nowDateTime := utils.GetCurDateTimeFormat()

	err := db.Mysql.Table("token_info").Where("token=? and chain_id=? ", token, chainId).Updates(map[string]interface{}{
		"price":      price,
		"updated_at": nowDateTime,
	}).Debug().Error
	if err != nil {
		log.Logger.Sugar().Error("UpdateContractPrice SavePriceData err ", err)
		return err
	}

	return nil
}
