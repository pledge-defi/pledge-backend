package models

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"pledge-backend/db"
)

type TokenInfo struct {
	Id           int    `gorm:"column:id;primaryKey"`
	Logo         string `json:"logo" gorm:"column:logo"`
	Token        string `json:"token" gorm:"column:token"`
	Symbol       string `json:"symbol" gorm:"column:symbol"`
	ChainId      string `json:"chain_id" gorm:"column:chain_id"`
	Price        string `json:"price" gorm:"column:price"`
	AbiFileExist int    `json:"abi_file_exist" gorm:"column:abi_file_exist"`
	CreatedAt    string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    string `json:"updated_at" gorm:"column:updated_at"`
}

func NewTokenInfo() *TokenInfo {
	return &TokenInfo{}
}

// GetTokenInfo Get token information by token name
func (t *TokenInfo) GetTokenInfo(token, chainId string) (error, TokenInfo) {

	redisTokenInfoBytes, err := db.RedisGet("token_info:" + token + "_" + chainId)
	if err != nil {
		tokenInfo := TokenInfo{}
		err := db.Mysql.Table("token_info").Where("token=? and chain_id=?", token, chainId).First(&tokenInfo).Debug().Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, tokenInfo
			} else {
				return errors.New("record select err " + err.Error()), tokenInfo
			}
		}
		return nil, tokenInfo
	} else {
		redisTokenInfo := RedisTokenInfo{}
		err = json.Unmarshal(redisTokenInfoBytes, &redisTokenInfo)
		return nil, TokenInfo{
			Logo:    redisTokenInfo.Logo,
			Token:   token,
			Symbol:  redisTokenInfo.Symbol,
			ChainId: chainId,
			Price:   redisTokenInfo.Price,
		}
	}

}
