package models

import (
	"errors"
	"pledge-backend/api/models/request"
	"pledge-backend/db"
)

type TokenInfo struct {
	Id      int32  `json:"-" gorm:"column:id;primaryKey"`
	Symbol  string `json:"symbol" gorm:"column:symbol"`
	Token   string `json:"token" gorm:"column:token"`
	ChainId int    `json:"chain_id" gorm:"column:chain_id"`
}

type TokenList struct {
	Id       int32  `json:"-" gorm:"column:id;primaryKey"`
	Symbol   string `json:"symbol" gorm:"column:symbol"`
	Decimals int    `json:"decimals" gorm:"column:decimals"`
	Token    string `json:"token" gorm:"column:token"`
	Logo     string `json:"logo" gorm:"column:logo"`
	ChainId  int    `json:"chain_id" gorm:"column:chain_id"`
}

func NewTokenInfo() *TokenInfo {
	return &TokenInfo{}
}

func (m *TokenInfo) GetTokenInfo(req *request.TokenList) (error, []TokenInfo) {
	var tokenInfo = make([]TokenInfo, 0)
	err := db.Mysql.Table("token_info").Where("chain_id", req.ChainId).Find(&tokenInfo).Debug().Error
	if err != nil {
		return errors.New("record select err " + err.Error()), nil
	}
	return nil, tokenInfo
}

func (m *TokenInfo) GetTokenList(req *request.TokenList) (error, []TokenList) {
	var tokenList = make([]TokenList, 0)
	err := db.Mysql.Table("token_info").Where("chain_id", req.ChainId).Find(&tokenList).Debug().Error
	if err != nil {
		return errors.New("record select err " + err.Error()), nil
	}
	return nil, tokenList
}
