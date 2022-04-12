package services

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
)

type TokenList struct{}

func NewTokenList() *TokenList {
	return &TokenList{}
}

func (c *TokenList) DebtTokenList(req *request.TokenList) (int, []models.TokenInfo) {
	err, res := models.NewTokenInfo().GetTokenInfo(req)
	if err != nil {
		return statecode.CommonErrServerErr, nil
	}
	return statecode.CommonSuccess, res

}

func (c *TokenList) GetTokenList(req *request.TokenList) (int, []models.TokenList) {
	err, tokenList := models.NewTokenInfo().GetTokenList(req)
	if err != nil {
		return statecode.CommonErrServerErr, nil
	}
	return statecode.CommonSuccess, tokenList

}
