package sv21

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/log"
	"pledge-backend/schedule/models"
	"pledge-backend/utils"
	"regexp"
	"strings"
)

type TokenLogo struct{}

func NewTokenLogo() *TokenLogo {
	return &TokenLogo{}
}

func (s *TokenLogo) UpdateTokenLogo() {

	// update remote logo
	res, err := utils.HttpGet(config.Config.Token.LogoUrl)
	if err != nil {
		log.Logger.Sugar().Info("UpdateTokenLogo HttpGet err", err)
	} else {
		tokenLogoRemote := models.TokenLogoRemote{}

		err = json.Unmarshal(res, &tokenLogoRemote)
		if err != nil {
			log.Logger.Sugar().Error("UpdateTokenLogo json.Unmarshal err ", err)
			return
		}
		for _, t := range tokenLogoRemote.Tokens {

			hasNewData, err := s.CheckLogoData(t.Address, utils.IntToString(t.ChainID), t.LogoURI, t.Symbol)
			if err != nil {
				log.Logger.Sugar().Error("UpdateTokenLogo CheckLogoData err ", err)
				continue
			}

			if hasNewData {
				err = s.SaveLogoData(t.Address, utils.IntToString(t.ChainID), t.LogoURI, t.Symbol, t.Decimals)
				if err != nil {
					log.Logger.Sugar().Error("UpdateTokenLogo SaveLogoData err ", err)
					continue
				}
			}
		}
	}

	//update local logo,Local logos have high weight,so execute later, local logos are divided by name
	for _, v := range LocalTokenLogo {
		for _, t := range v {
			if t["token"] == "" {
				continue
			}
			hasNewData, err := s.CheckLogoData(t["token"], t["chain_id"], t["logo"], t["symbol"])
			if err != nil {
				continue
			}

			if hasNewData {
				err = s.SaveLogoData(t["token"], t["chain_id"], t["logo"], t["symbol"], utils.StringToInt(t["decimals"]))
				if err != nil {
					log.Logger.Sugar().Error("UpdateTokenLogo SaveLogoData err ", err)
					continue
				}
			}
		}
	}
}

// CheckLogoData Saving logo data to redis if it has new logo
func (s *TokenLogo) CheckLogoData(token, chainId, logoUrl, symbol string) (bool, error) {
	redisKey := "token_info:" + chainId + ":" + token
	redisTokenInfoBytes, err := db.RedisGet(redisKey)
	if len(redisTokenInfoBytes) <= 0 {
		err = s.CheckTokenInfo(token, chainId)
		if err != nil {
			log.Logger.Error(err.Error())
		}
		err = db.RedisSet(redisKey, models.RedisTokenInfo{
			Token:   token,
			ChainId: chainId,
			Logo:    logoUrl,
			Symbol:  symbol,
		}, 0)
		if err != nil {
			log.Logger.Error(err.Error())
			return false, err
		}
	} else {
		redisTokenInfo := models.RedisTokenInfo{}
		err = json.Unmarshal(redisTokenInfoBytes, &redisTokenInfo)
		if err != nil {
			log.Logger.Error(err.Error())
			return false, err
		}

		if redisTokenInfo.Logo == logoUrl {
			return false, nil
		}

		redisTokenInfo.Logo = logoUrl
		redisTokenInfo.Symbol = symbol
		err = db.RedisSet(redisKey, redisTokenInfo, 0)
		if err != nil {
			log.Logger.Error(err.Error())
			return true, err
		}
	}
	return true, nil
}

// CheckTokenInfo  Insert token information if it was not in mysql
func (s *TokenLogo) CheckTokenInfo(token, chainId string) error {
	tokenInfo := models.TokenInfo{}
	err := db.Mysql.Table("token_info").Where("token=? and chain_id=?", token, chainId).First(&tokenInfo).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tokenInfo = models.TokenInfo{}
			nowDateTime := utils.GetCurDateTimeFormat()
			tokenInfo.Token = token
			tokenInfo.ChainId = chainId
			tokenInfo.UpdatedAt = nowDateTime
			tokenInfo.CreatedAt = nowDateTime
			err = db.Mysql.Table("token_info").Create(&tokenInfo).Debug().Error
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

// SaveLogoData Saving logo data to mysql if it has new logo
func (s *TokenLogo) SaveLogoData(token, chainId, logoUrl, symbol string, decimals int) error {
	nowDateTime := utils.GetCurDateTimeFormat()

	err := db.Mysql.Table("token_info").Where("token=? and chain_id=? ", token, chainId).Updates(map[string]interface{}{
		"symbol":     symbol,
		"logo":       logoUrl,
		"decimals":   decimals,
		"updated_at": nowDateTime,
	}).Debug().Error
	if err != nil {
		log.Logger.Sugar().Error("UpdateTokenLogo SaveLogoData err ", err)
		return err
	}

	return nil
}

func GetBaseUrl() string {

	domainName := config.Config.Env.DomainName
	domainNameSlice := strings.Split(domainName, "")
	pattern := "\\d+" //反斜杠要转义
	isNumber, _ := regexp.MatchString(pattern, domainNameSlice[0])
	if isNumber {
		return config.Config.Env.Protocol + "://" + config.Config.Env.DomainName + ":" + config.Config.Env.Port + "/"
	}
	return config.Config.Env.Protocol + "://" + config.Config.Env.DomainName + "/"
}

var BaseUrl = GetBaseUrl()

var LocalTokenLogo = map[string]map[string]map[string]string{
	"BNB": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "0x0000000000000000000000000000000000000000",
			"symbol":   "BNB",
			"logo":     BaseUrl + "storage/img/BNB.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x0000000000000000000000000000000000000000",
			"symbol":   "BNB",
			"logo":     BaseUrl + "storage/img/BNB.png",
		},
	},
	"BTC": {
		"test_net": {
			"chain_id": "97",
			"decimals": "8",
			"token":    "0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658",
			"symbol":   "BTC",
			"logo":     BaseUrl + "storage/img/BTC.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "8",
			"token":    "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
			"symbol":   "BTC",
			"logo":     BaseUrl + "storage/img/BTC.png",
		},
	},
	"BTCB": {
		"test_net": {
			"chain_id": "97",
			"decimals": "8",
			"token":    "0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658",
			"symbol":   "BTC",
			"logo":     BaseUrl + "storage/img/BTC.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "8",
			"token":    "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
			"symbol":   "BTC",
			"logo":     BaseUrl + "storage/img/BTC.png",
		},
	},
	"BUSD": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "0xE676Dcd74f44023b95E0E2C6436C97991A7497DA",
			"symbol":   "BUSD",
			"logo":     BaseUrl + "storage/img/BUSD.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
			"symbol":   "BUSD",
			"logo":     BaseUrl + "storage/img/BUSD.png",
		},
	},
	"DAI": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "0x490BC3FCc845d37C1686044Cd2d6589585DE9B8B",
			"symbol":   "DAI",
			"logo":     BaseUrl + "storage/img/DAI.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3",
			"symbol":   "DAI",
			"logo":     BaseUrl + "storage/img/DAI.png",
		},
	},
	"ETH": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "",
			"symbol":   "ETH",
			"logo":     BaseUrl + "storage/img/ETH.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
			"symbol":   "ETH",
			"logo":     BaseUrl + "storage/img/ETH.png",
		},
	},
	"USDT": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "",
			"symbol":   "USDT",
			"logo":     BaseUrl + "storage/img/USDT.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x55d398326f99059ff775485246999027b3197955",
			"symbol":   "USDT",
			"logo":     BaseUrl + "storage/img/USDT.png",
		},
	},
	"CAKE": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "0xEAEd08168a2D34Ae2B9ea1c1f920E0BC00F9fA67",
			"symbol":   "CAKE",
			"logo":     BaseUrl + "storage/img/CAKE.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
			"symbol":   "CAKE",
			"logo":     BaseUrl + "storage/img/CAKE.png",
		},
	},
	"PLGR": {
		"test_net": {
			"chain_id": "97",
			"decimals": "18",
			"token":    "",
			"symbol":   "PLGR",
			"logo":     BaseUrl + "storage/img/PLGR.png",
		},
		"main_net": {
			"chain_id": "56",
			"decimals": "18",
			"token":    "0x6Aa91CbfE045f9D154050226fCc830ddbA886CED",
			"symbol":   "PLGR",
			"logo":     BaseUrl + "storage/img/PLGR.png",
		},
	},
}
