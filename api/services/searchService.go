package services

import (
	"fmt"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/log"
)

type SearchService struct{}

func NewSearch() *SearchService {
	return &SearchService{}
}

func (c *SearchService) Search(req *request.Search) (int, int64, []models.Pool) {

	whereCondition := fmt.Sprintf(`chain_id='%v'`, req.ChainID)
	if req.LendTokenSymbol != "" {
		whereCondition += fmt.Sprintf(` and lend_token_symbol='%v'`, req.LendTokenSymbol)
	}
	if req.State != "" {
		whereCondition += fmt.Sprintf(` and state='%v'`, req.State)
	}
	err, total, data := models.NewPool().Pagination(req, whereCondition)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr, 0, nil
	}
	return 0, total, data
}
