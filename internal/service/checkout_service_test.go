package service

import (
	"github.com/lucasmilotich/checkout-golang/internal/dao"
	"github.com/stretchr/testify/assert"
	"testing"
)

var checkoutService CheckoutService

func init() {
	checkoutService = CheckoutService{
		ProductService: ProductService{},
		// with other type of db, we can mock it using golang/mock or creating a mock struct
		DB:         dao.NewBasketDB(),
		DiscountDB: dao.NewDiscountDB(),
	}
}

func Test_discounts(t *testing.T) {
	t.Run("Get total discount for 10 pens( buy two get one free)", func(t *testing.T) {

		aux := basketAux{price: 5, quantity: 10}
		products := make(map[string]basketAux)
		products["PEN"] = aux

		assert.Equal(t, 15.0, checkoutService.getCalculatedDiscounts(products))
	})

	t.Run("Get total discount for 10 t-shirts(more than 3, receives 25% discount each)", func(t *testing.T) {

		aux := basketAux{price: 10, quantity: 10}
		products := make(map[string]basketAux)
		products["TSHIRT"] = aux

		assert.Equal(t, 25.0, checkoutService.getCalculatedDiscounts(products))
	})

	t.Run("Get total discount for 10 t-shirts(more than 3, receives 25% discount each) and 10 pens", func(t *testing.T) {

		aux := basketAux{price: 10, quantity: 10}
		products := make(map[string]basketAux)
		products["TSHIRT"] = aux

		aux2 := basketAux{price: 5, quantity: 10}
		products["PEN"] = aux2

		assert.Equal(t, 40.0, checkoutService.getCalculatedDiscounts(products))
	})
}
