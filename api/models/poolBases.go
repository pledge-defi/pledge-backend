package models

import (
	"encoding/json"
	"pledge-backend/db"
)

type PoolBaseInfo struct {
	PoolID                 int             `json:"pool_id"`
	AutoLiquidateThreshold string          `json:"autoLiquidateThreshold"`
	BorrowSupply           string          `json:"borrowSupply"`
	BorrowToken            string          `json:"borrowToken"`
	BorrowTokenInfo        BorrowTokenInfo `json:"borrowTokenInfo"`
	EndTime                string          `json:"endTime"`
	InterestRate           string          `json:"interestRate"`
	JpCoin                 string          `json:"jpCoin"`
	LendSupply             string          `json:"lendSupply"`
	LendToken              string          `json:"lendToken"`
	LendTokenInfo          LendTokenInfo   `json:"lendTokenInfo"`
	MartgageRate           string          `json:"martgageRate"`
	MaxSupply              string          `json:"maxSupply"`
	SettleTime             string          `json:"settleTime"`
	SpCoin                 string          `json:"spCoin"`
	State                  string          `json:"state"`
}

type PoolBases struct {
	PoolID                 int    `json:"pool_id"`
	AutoLiquidateThreshold string `json:"autoLiquidateThreshold"`
	BorrowSupply           string `json:"borrowSupply"`
	BorrowToken            string `json:"borrowToken"`
	BorrowTokenInfo        string `json:"borrowTokenInfo"`
	EndTime                string `json:"endTime"`
	InterestRate           string `json:"interestRate"`
	JpCoin                 string `json:"jpCoin"`
	LendSupply             string `json:"lendSupply"`
	LendToken              string `json:"lendToken"`
	LendTokenInfo          string `json:"lendTokenInfo"`
	MartgageRate           string `json:"martgageRate"`
	MaxSupply              string `json:"maxSupply"`
	SettleTime             string `json:"settleTime"`
	SpCoin                 string `json:"spCoin"`
	State                  string `json:"state"`
}

type BorrowTokenInfo struct {
	BorrowFee  string `json:"borrowFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}

type LendTokenInfo struct {
	LendFee    string `json:"lendFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}

type PoolBaseInfoRes struct {
	Index    int          `json:"index"`
	PoolData PoolBaseInfo `json:"pool_data"`
}

func NewPoolBases() *PoolBases {
	return &PoolBases{}
}

func (p *PoolBases) PoolBaseInfo(chainId int, res *[]PoolBaseInfoRes) error {
	var poolBases []PoolBases

	err := db.Mysql.Table("poolbases").Where("chain_id=?", chainId).Order("pool_id asc").Find(&poolBases).Debug().Error
	if err != nil {
		return err
	}

	for _, v := range poolBases {
		borrowTokenInfo := BorrowTokenInfo{}
		_ = json.Unmarshal([]byte(v.BorrowTokenInfo), &borrowTokenInfo)
		lendTokenInfo := LendTokenInfo{}
		_ = json.Unmarshal([]byte(v.LendTokenInfo), &lendTokenInfo)
		*res = append(*res, PoolBaseInfoRes{
			Index: v.PoolID - 1,
			PoolData: PoolBaseInfo{
				PoolID:                 v.PoolID,
				AutoLiquidateThreshold: v.AutoLiquidateThreshold,
				BorrowSupply:           v.BorrowSupply,
				BorrowToken:            v.BorrowToken,
				BorrowTokenInfo:        borrowTokenInfo,
				EndTime:                v.EndTime,
				InterestRate:           v.InterestRate,
				JpCoin:                 v.JpCoin,
				LendSupply:             v.LendSupply,
				LendToken:              v.LendToken,
				LendTokenInfo:          lendTokenInfo,
				MartgageRate:           v.MartgageRate,
				MaxSupply:              v.MaxSupply,
				SettleTime:             v.SettleTime,
				SpCoin:                 v.SpCoin,
				State:                  v.State,
			},
		})
	}
	return nil
}
