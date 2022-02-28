package sv21

import (
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
	nowDateTime := utils.GetCurDateTimeFormat()
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

		priceRedis, err := db.RedisGetInt64("token_info:price:" + t.Token + "_" + t.ChainId)
		if err != nil {
			if err.Error() == "redigo: nil returned" {
				err = db.RedisSetInt64("token_info:price:"+t.Token+"_"+t.ChainId, price, 120)
			} else {
				log.Logger.Sugar().Error("UpdateContractPrice get price from redis err ", t.Symbol, t.ChainId, err)
			}
		} else {
			if priceRedis == price {
				continue
			}
			err = db.RedisSetInt64("token_info:price:"+t.Token+"_"+t.ChainId, price, 120)
		}

		err = db.Mysql.Table("token_info").Where("id=?", t.Id).Updates(map[string]interface{}{
			"price":      price,
			"updated_at": nowDateTime,
		}).Debug().Error
		if err != nil {
			log.Logger.Sugar().Error("UpdateContractPrice err ", t.Symbol, t.ChainId, err)
			continue
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
