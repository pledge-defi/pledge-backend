package services

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/log"
)

type poolService struct{}

func NewPool() *poolService {
	return &poolService{}
}

func (s *poolService) PoolBaseInfo(chainId int, result *[]models.PoolBaseInfoRes) int {

	err := models.NewPoolBases().PoolBaseInfo(chainId, result)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr
	}
	return statecode.CommonSuccess
}

func (s *poolService) PoolDataInfo(chainId int, result *[]models.PoolDataInfoRes) int {

	err := models.NewPoolData().PoolDataInfo(chainId, result)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr
	}
	return statecode.CommonSuccess
}
