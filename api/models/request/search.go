package request

type Search struct {
	ChainID         int    `form:"chainID" json:"chainID" binding:"required"`
	LendTokenSymbol string `form:"lend_token_symbol" json:"lend_token_symbol" binding:"omitempty"`
	State           string `form:"state" json:"state" binding:"omitempty"`
	Page            int    `form:"page" json:"page" `
	PageSize        int    `form:"pageSize" json:"pageSize" `
}
