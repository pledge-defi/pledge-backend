package request

type PoolBaseInfo struct {
	ChainId int `form:"chainId" binding:"required"`
}
