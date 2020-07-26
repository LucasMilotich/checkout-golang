package model

type Product struct {
	Code  string  `json:"code,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}
