package model

type Basket struct {
	ID          string    `json:"id"`
	Products    []Product `json:"products"`
	TotalAmount float64   `json:"total_amount"`
}
