package request

type TokenList struct {
	ChainId int `form:"chainId" binding:"required"`
}
