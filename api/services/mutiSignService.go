package services

import (
	"encoding/json"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
)

type MutiSignService struct{}

func NewMutiSign() *MutiSignService {
	return &MutiSignService{}
}

// Set Multi-Sign
func (c *MutiSignService) SetMultiSign(mutiSign *request.SetMultiSign) (int, error) {
	//db set
	err := models.NewMultiSign().Set(mutiSign)
	if err != nil {
		return statecode.COMMON_ERR_SERVER_ERR, err
	}
	return statecode.COMMON_SUCCESS, nil
}

// Get Multi-Sign
func (c *MutiSignService) GetMultiSign(mutiSign *response.MultiSign, chainId int) (int, error) {
	//db get
	multiSignModel := models.NewMultiSign()
	err := multiSignModel.Get(chainId)
	if err != nil {
		return statecode.COMMON_ERR_SERVER_ERR, err
	}
	var multiSignAccount []string
	json.Unmarshal([]byte(multiSignModel.MultiSignAccount), &multiSignAccount)

	mutiSign.SpName = multiSignModel.SpName
	mutiSign.SpToken = multiSignModel.SpToken
	mutiSign.JpName = multiSignModel.JpName
	mutiSign.JpToken = multiSignModel.JpToken
	mutiSign.SpAddress = multiSignModel.SpAddress
	mutiSign.JpAddress = multiSignModel.JpAddress
	mutiSign.SpHash = multiSignModel.SpHash
	mutiSign.JpHash = multiSignModel.JpHash
	mutiSign.MultiSignAccount = multiSignAccount
	return statecode.COMMON_SUCCESS, nil
}
