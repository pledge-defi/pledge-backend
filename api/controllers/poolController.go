package controllers

import (
	"github.com/gin-gonic/gin"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
	"pledge-backend/api/services"
	"pledge-backend/api/validate"
)

type PoolController struct {
}

func (c *PoolController) PoolBaseInfo(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.PoolBaseInfo{}
	var result []models.PoolBaseInfoRes

	errCode := validate.NewPoolBaseInfo().PoolBaseInfo(ctx, &req)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode = services.NewPool().PoolBaseInfo(req.ChainId, &result)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.COMMON_SUCCESS, result)
	return
}

func (c *PoolController) PoolDataInfo(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.PoolDataInfo{}
	var result []models.PoolDataInfoRes

	errCode := validate.NewPoolDataInfo().PoolDataInfo(ctx, &req)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode = services.NewPool().PoolDataInfo(req.ChainId, &result)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.COMMON_SUCCESS, result)
	return
}

func (c *PoolController) TokenList(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.TokenList{}
	//result := make([]response.TokenList, 0)

	errCode := validate.NewTokenList().TokenList(ctx, &req)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode, result := services.NewTokenList().GetTokenList(&req)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.COMMON_SUCCESS, result)
	return
}

func (c *PoolController) Search(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.Search{}
	result := response.Search{}

	errCode := validate.NewSearch().Search(ctx, &req)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode, count, pools := services.NewSearch().Search(&req)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	result.Rows = pools
	result.Count = count
	res.Response(ctx, statecode.COMMON_SUCCESS, result)
	return
}

func (c *PoolController) DebtTokenList(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.TokenList{}

	errCode := validate.NewTokenList().TokenList(ctx, &req)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode, result := services.NewTokenList().DebtTokenList(&req)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.COMMON_SUCCESS, result)
	return
}
