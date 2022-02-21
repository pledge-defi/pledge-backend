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

func (c *TokenList) GetTokenList(req *request.TokenList) (int, []models.TokenInfo) {
	err, res := models.NewTokenInfo().GetTokenInfo(req)
	if err != nil {
		return statecode.COMMON_ERR_SERVER_ERR, nil
	}
	return statecode.COMMON_SUCCESS, res

}
