package request

type SetMultiSign struct {
	SpName           string   `json:"sp_name" binding:"required"`
	ChainId          int      `json:"chain_id"`
	SpToken          string   `json:"_spToken"`
	JpName           string   `json:"jp_name"`
	JpToken          string   `json:"_jpToken"`
	SpAddress        string   `json:"sp_address"`
	JpAddress        string   `json:"jp_address"`
	SpHash           string   `json:"spHash"`
	JpHash           string   `json:"jpHash"`
	MultiSignAccount []string `json:"multi_sign_account"`
}

type GetMultiSign struct {
	ChainId int `json:"chain_id"`
}
