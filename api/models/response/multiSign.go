package response

// multi-sign signature
type MultiSign struct {
	SpName           string   `json:"sp_name"`
	SpToken          string   `json:"_spToken"`
	JpName           string   `json:"jp_name"`
	JpToken          string   `json:"_jpToken"`
	SpAddress        string   `json:"sp_address"`
	JpAddress        string   `json:"jp_address"`
	SpHash           string   `json:"spHash"`
	JpHash           string   `json:"jpHash"`
	MultiSignAccount []string `json:"multi_sign_account"`
}
