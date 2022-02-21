package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
)

type PoolDataInfo struct{}

func NewPoolDataInfo() *PoolDataInfo {
	return &PoolDataInfo{}
}

func (v *PoolDataInfo) PoolDataInfo(c *gin.Context, req *request.PoolDataInfo) int {
	err := c.ShouldBind(req)
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

	if req.ChainId != 97 && req.ChainId != 56 {
		return statecode.CHAINID_ERR
	}

	return statecode.COMMON_SUCCESS
}
