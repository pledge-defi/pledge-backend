package sv22

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"pledge-backend/config"
	abifile "pledge-backend/contract/abi"
	"pledge-backend/db"
	"pledge-backend/log"
	"pledge-backend/schedule/models"
	"pledge-backend/utils"
	"strings"
	"time"
)

type TokenSymbol struct{}

func NewTokenSymbol() *TokenSymbol {
	return &TokenSymbol{}
}

// UpdateContractSymbol get contract symbol
func (s *TokenSymbol) UpdateContractSymbol() {
	var tokens []models.TokenInfo
	nowDateTime := utils.GetCurDateTimeFormat()
	db.Mysql.Table("token_info").Find(&tokens)
	for _, t := range tokens {
		if t.Token == "" {
			log.Logger.Sugar().Error("UpdateContractSymbol token empty", t.Symbol, t.ChainId)
			continue
		}
		err := errors.New("")
		symbol := ""
		if t.ChainId == "97" {
			err, symbol = s.GetContractSymbolOnTestNet(t.Token, config.Config.TestNet.NetUrl)
		} else if t.ChainId == "56" {
			if t.AbiFileExist == 0 {
				err = s.GetRemoteAbiFileByToken(t.Token, t.ChainId)
				if err != nil {
					log.Logger.Sugar().Error("UpdateContractSymbol GetRemoteAbiFileByToken err ", t.Symbol, t.ChainId, err)
					continue
				}
			}
			err, symbol = s.GetContractSymbolOnMainNet(t.Token, config.Config.MainNet.NetUrl)
		} else {
			log.Logger.Sugar().Error("UpdateContractSymbol chain_id err", t.Symbol, t.ChainId)
			continue
		}
		if err != nil {
			log.Logger.Sugar().Error("UpdateContractSymbol err", t.Symbol, t.ChainId, err)
			continue
		}
		err = db.Mysql.Table("token_info").Where("id=?", t.Id).Updates(map[string]interface{}{
			"symbol":     symbol,
			"updated_at": nowDateTime,
		}).Debug().Error
		if err != nil {
			log.Logger.Sugar().Error("UpdateContractSymbol err", t.Symbol, t.ChainId, err)
			continue
		}

		// Avoid blockchain servers treating programs as crawlers
		time.Sleep(time.Second)
	}
}

// GetRemoteAbiFileByToken get and save remote abi file on main net
func (s *TokenSymbol) GetRemoteAbiFileByToken(token, chainId string) error {
	url := "https://api.bscscan.com/api?module=contract&action=getabi&apikey=HJ3WS4N88QJ6S7PQ8D89BD49IZIFP1JFER&address=" + token
	res, err := utils.HttpGet(url)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	resStr := s.FormatAbiJsonStr(string(res))

	abiJson := models.AbiJson{}
	err = json.Unmarshal([]byte(resStr), &abiJson)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	if abiJson.Status != "1" {
		log.Logger.Error("get remote abi file failed: status 0 ")
		return errors.New("get remote abi file failed: status 0")
	}

	abiJsonBytes, err := json.MarshalIndent(abiJson.Result, "", "\t")
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	newAbi := abifile.GetCurrentAbPathByCaller() + "/v" + config.Config.Env.Version + "/" + token + ".abi"

	err = os.WriteFile(newAbi, abiJsonBytes, 0777)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	err = db.Mysql.Table("token_info").Where("token=? and chain_id=?", token, chainId).Updates(map[string]interface{}{
		"abi_file_exist": 1,
	}).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

// FormatAbiJsonStr format the abi string
func (s *TokenSymbol) FormatAbiJsonStr(result string) string {
	resStr := strings.Replace(result, `\`, "", -1)
	resStr = strings.Replace(resStr, `result\":\"[`, `result":[`, -1)
	resStr = strings.Replace(resStr, `result":"[`, `result":[`, -1)
	resStr = strings.Replace(resStr, `]\"}`, `]}`, -1)
	resStr = strings.Replace(resStr, `]"}`, `]}`, -1)
	return resStr
}

// GetContractSymbolOnMainNet get contract symbol on main net
func (s *TokenSymbol) GetContractSymbolOnMainNet(token, network string) (error, string) {
	ethereumConn, err := ethclient.Dial(network)
	if nil != err {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", token, err)
		return err, ""
	}
	abiStr, err := abifile.GetAbiByToken(token)
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", token, err)
		return err, ""
	}
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", token, err)
		return err, ""
	}
	contract, err := bind.NewBoundContract(common.HexToAddress(token), parsed, ethereumConn, ethereumConn, ethereumConn), nil
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", token, err)
		return err, ""
	}

	res := make([]interface{}, 0)
	err = contract.Call(nil, &res, "symbol")
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", err)
		return err, ""
	}

	return nil, res[0].(string)
}

// GetContractSymbolOnTestNet get contract symbol on test net
func (s *TokenSymbol) GetContractSymbolOnTestNet(token, network string) (error, string) {
	ethereumConn, err := ethclient.Dial(network)
	if nil != err {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", token, err)
		return err, ""
	}
	abiStr, err := abifile.GetAbiByToken("erc20")
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", token, err)
		return err, ""
	}
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", token, err)
		return err, ""
	}
	contract, err := bind.NewBoundContract(common.HexToAddress(token), parsed, ethereumConn, ethereumConn, ethereumConn), nil
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", token, err)
		return err, ""
	}

	res := make([]interface{}, 0)
	err = contract.Call(nil, &res, "symbol")
	if err != nil {
		log.Logger.Sugar().Error("GetContractSymbolOnMainNet err", token, err)
		return err, ""
	}

	return nil, res[0].(string)
}
