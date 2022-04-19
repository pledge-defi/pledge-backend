package models

type AbiJson struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type Outputs struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}
type AbiData struct {
	Inputs          []interface{} `json:"inputs"`
	Payable         bool          `json:"payable,omitempty"`
	StateMutability string        `json:"stateMutability,omitempty"`
	Type            string        `json:"type"`
	Anonymous       bool          `json:"anonymous,omitempty"`
	Name            string        `json:"name,omitempty"`
	Constant        bool          `json:"constant,omitempty"`
	Outputs         []Outputs     `json:"outputs,omitempty"`
}
