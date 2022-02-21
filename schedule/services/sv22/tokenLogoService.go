package sv22

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

	nowDateTime := utils.GetCurDateTimeFormat()

	// update remote logo
	res, err := utils.HttpGet(config.Config.Token.LogoUrl)
	if err != nil {
		log.Logger.Sugar().Info("UpdateTokenLogo HttpGet err", err)
	} else {
		tokenLogoRemote := models.TokenLogoRemote{}

		json.Unmarshal([]byte(res), &tokenLogoRemote)
		for _, t := range tokenLogoRemote.Tokens {
			var tokenInfo models.TokenInfo
			err = db.Mysql.Table("token_info").Where("token=?", t.Address).First(&tokenInfo).Debug().Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = db.Mysql.Table("token_info").Create(&models.TokenInfo{
					Token:     t.Address,
					Logo:      t.LogoURI,
					Symbol:    t.Symbol,
					ChainId:   utils.IntToString(t.ChainID),
					CreatedAt: nowDateTime,
					UpdatedAt: nowDateTime,
				}).Debug().Error
				if err != nil {
					log.Logger.Sugar().Error("UpdateTokenLogo token_info create err ", err)
					continue
				}
			} else {
				log.Logger.Sugar().Error("UpdateTokenLogo token_info find err ", err)
				continue
			}

			err = db.Mysql.Table("token_info").Where("token=? and chain_id=? ", t.Address, t.ChainID).Updates(map[string]interface{}{
				"logo":       t.LogoURI,
				"symbol":     t.Symbol,
				"updated_at": nowDateTime,
			}).Debug().Error
			if err != nil {
				log.Logger.Sugar().Error("UpdateTokenLogo token_info update err ", err)
				continue
			}
		}
	}

	//update local logo,Local logos have high weight,so execute later, local logos are divided by name
	for _, v := range LocalTokenLogo {
		for _, t := range v {
			if t["token"] == "" {
				continue
			}
			err = db.Mysql.Table("token_info").Where("chain_id=? and token=?", t["chain_id"], t["token"]).First(&models.TokenInfo{}).Debug().Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				nowDateTime := utils.GetCurDateTimeFormat()
				err = db.Mysql.Table("token_info").Create(&models.TokenInfo{
					Token:     t["token"],
					Symbol:    t["symbol"],
					ChainId:   t["chain_id"],
					Logo:      t["logo"],
					CreatedAt: nowDateTime,
					UpdatedAt: nowDateTime,
				}).Debug().Error
				if err != nil {
					log.Logger.Sugar().Error("UpdateTokenLogo token_info create err ", err)
					continue
				}
			} else {
				err = db.Mysql.Table("token_info").Where("chain_id=? and token=?", t["chain_id"], t["token"]).Updates(map[string]interface{}{
					"logo":       t["logo"],
					"symbol":     t["symbol"],
					"updated_at": nowDateTime,
				}).Debug().Error
				if err != nil {
					log.Logger.Sugar().Error("UpdateTokenLogo token_info update err ", err)
				}
				continue
			}
		}
	}
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

var LocalTokenLogo map[string]map[string]map[string]string = map[string]map[string]map[string]string{
	"BNB": map[string]map[string]string{
		"test_net": map[string]string{
			"chain_id": "97",
			"token":    "0x0000000000000000000000000000000000000000",
			"symbol":   "BNB",
			"logo":     BaseUrl + "storage/img/BNB.png",
		},
		"main_net": map[string]string{
			"chain_id": "56",
			"token":    "0x0000000000000000000000000000000000000000",
			"symbol":   "BNB",
			"logo":     BaseUrl + "storage/img/BNB.png",
		},
	},
	"BTC": map[string]map[string]string{
		"test_net": map[string]string{
			"chain_id": "97",
			"token":    "0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658",
			"symbol":   "BTC",
			"logo":     BaseUrl + "storage/img/BTC.png",
		},
		"main_net": map[string]string{
			"chain_id": "56",
			"token":    "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
			"symbol":   "BTC",
			"logo":     BaseUrl + "storage/img/BTC.png",
		},
	},
	"BTCB": map[string]map[string]string{
		"test_net": map[string]string{
			"chain_id": "97",
			"token":    "0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658",
			"symbol":   "BTC",
			"logo":     BaseUrl + "storage/img/BTC.png",
		},
		"main_net": map[string]string{
			"chain_id": "56",
			"token":    "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
			"symbol":   "BTC",
			"logo":     BaseUrl + "storage/img/BTC.png",
		},
	},
	"BUSD": map[string]map[string]string{
		"test_net": map[string]string{
			"chain_id": "97",
			"token":    "0xE676Dcd74f44023b95E0E2C6436C97991A7497DA",
			"symbol":   "BUSD",
			"logo":     BaseUrl + "storage/img/BUSD.png",
		},
		"main_net": map[string]string{
			"chain_id": "56",
			"token":    "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
			"symbol":   "BUSD",
			"logo":     BaseUrl + "storage/img/BUSD.png",
		},
	},
	"DAI": map[string]map[string]string{
		"test_net": map[string]string{
			"chain_id": "97",
			"token":    "0x490BC3FCc845d37C1686044Cd2d6589585DE9B8B",
			"symbol":   "DAI",
			"logo":     BaseUrl + "storage/img/DAI.png",
		},
		"main_net": map[string]string{
			"chain_id": "56",
			"token":    "0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3",
			"symbol":   "DAI",
			"logo":     BaseUrl + "storage/img/DAI.png",
		},
	},
	"ETH": map[string]map[string]string{
		"test_net": map[string]string{
			"chain_id": "97",
			"token":    "",
			"symbol":   "ETH",
			"logo":     BaseUrl + "storage/img/ETH.png",
		},
		"main_net": map[string]string{
			"chain_id": "56",
			"token":    "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
			"symbol":   "ETH",
			"logo":     BaseUrl + "storage/img/ETH.png",
		},
	},
	"USDT": map[string]map[string]string{
		"test_net": map[string]string{
			"chain_id": "97",
			"token":    "",
			"symbol":   "USDT",
			"logo":     BaseUrl + "storage/img/USDT.png",
		},
		"main_net": map[string]string{
			"chain_id": "56",
			"token":    "0x55d398326f99059ff775485246999027b3197955",
			"symbol":   "USDT",
			"logo":     BaseUrl + "storage/img/USDT.png",
		},
	},
	"CAKE": map[string]map[string]string{
		"test_net": map[string]string{
			"chain_id": "97",
			"token":    "0xEAEd08168a2D34Ae2B9ea1c1f920E0BC00F9fA67",
			"symbol":   "CAKE",
			"logo":     BaseUrl + "storage/img/CAKE.png",
		},
		"main_net": map[string]string{
			"chain_id": "56",
			"token":    "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
			"symbol":   "CAKE",
			"logo":     BaseUrl + "storage/img/CAKE.png",
		},
	},
	"PLGR": map[string]map[string]string{
		"test_net": map[string]string{
			"chain_id": "97",
			"token":    "",
			"symbol":   "PLGR",
			"logo":     BaseUrl + "storage/img/PLGR.png",
		},
		"main_net": map[string]string{
			"chain_id": "56",
			"token":    "0x6Aa91CbfE045f9D154050226fCc830ddbA886CED",
			"symbol":   "PLGR",
			"logo":     BaseUrl + "storage/img/PLGR.png",
		},
	},
}
