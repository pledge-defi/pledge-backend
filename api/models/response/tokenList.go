package response

import "time"

type TokenList struct {
	Name      string    `json:"name"`
	LogoURI   string    `json:"logoURI"`
	Tokens    []Token   `json:"tokens"`
	Version   Version   `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

type Token struct {
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	ChainID  int    `json:"chainId"`
	LogoURI  string `json:"logoURI,omitempty"`
}

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}
