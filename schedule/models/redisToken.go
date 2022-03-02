package models

type RedisTokenInfo struct {
	Logo    string `json:"logo" gorm:"column:logo"`
	Token   string `json:"token" gorm:"column:token"`
	Symbol  string `json:"symbol" gorm:"column:symbol"`
	ChainId string `json:"chain_id" gorm:"column:chain_id"`
	Price   string `json:"price" gorm:"column:price"`
}
