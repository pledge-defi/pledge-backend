package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
)

type MutiSign struct{}

func NewMutiSign() *MutiSign {
	return &MutiSign{}
}

func (v *MutiSign) SetMultiSign(c *gin.Context, req *request.SetMultiSign) int {

	err := c.ShouldBind(req)
	if req.ChainId != 97 && req.ChainId != 56 {
		return statecode.CHAINID_ERR
	}
	if err == io.EOF {
		return statecode.PARAMETER_EMPTY_ERR
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "SpName" && e.Tag() == "required" {
				return statecode.P_NAME_EMPTY
			}
		}
		return statecode.COMMON_ERR_SERVER_ERR
	}

	return statecode.COMMON_SUCCESS
}

func (v *MutiSign) GetMultiSign(c *gin.Context, req *request.GetMultiSign) int {

	err := c.ShouldBind(req)
	if req.ChainId != 97 && req.ChainId != 56 {
		return statecode.CHAINID_ERR
	}
	if err == io.EOF {
		return statecode.PARAMETER_EMPTY_ERR
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "ChainId" && e.Tag() == "required" {
				return statecode.CHAINID_EMPTY
			}
		}
		return statecode.COMMON_ERR_SERVER_ERR
	}

	return statecode.COMMON_SUCCESS
}
