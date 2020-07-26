package dao

import (
	"github.com/lucasmilotich/coreapi/internal/model"
	"sync"
)

type DiscountDB struct {
	elements map[string]model.Discount
	mux      *sync.Mutex
}

func NewDiscountDB() *DiscountDB {
	return &DiscountDB{
		elements: map[string]model.Discount{},
		mux:      &sync.Mutex{},
	}
}

func (db *DiscountDB) GetAllDiscounts() ([]model.Discount, error) {

	discounts := make([]model.Discount, 0)

	discounts = append(discounts, model.Discount{
		ProductCode:                "PEN",
		MinProductsInBasket:        3,
		PackagePromotion:           true,
		DiscountPerUnitToBeApplied: 0.333,
	}, model.Discount{
		ProductCode:                "TSHIRT",
		MinProductsInBasket:        3,
		PackagePromotion:           false,
		DiscountPerUnitToBeApplied: 0.250,
	})

	return discounts, nil
}
