package response

type TokenList struct {
	Id      int32  `json:"-" gorm:"column:id;primaryKey"`
	Symbol  string `json:"symbol" gorm:"column:symbol"`
	Token   string `json:"token" gorm:"column:token"`
	Logo    string `json:"logo" gorm:"column:logo"`
	ChainId int    `json:"chain_id" gorm:"column:chain_id"`
}
