package response

type TokenList struct {
	Name    string  `json:"name"`
	LogoURI string  `json:"logoURI"`
	Tokens  []Token `json:"tokens"`
}

type Token struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	ChainID int    `json:"chainId"`
	LogoURI string `json:"logoURI,omitempty"`
}
