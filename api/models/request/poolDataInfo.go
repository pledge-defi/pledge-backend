package request

type PoolDataInfo struct {
	ChainId int `form:"chainId" binding:"required"`
}
