package dao

import (
	"github.com/lucasmilotich/checkout-golang/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var db *BasketDB

func init() {
	db = NewBasketDB()
}
func Test_discounts(t *testing.T) {
	t.Run("Get no existing record and get an error", func(t *testing.T) {

		_, err := db.Get("NOT_EXISTING")
		assert.Error(t, err)

	})

	t.Run("Create a record and get it", func(t *testing.T) {
		basket, err := db.Create(model.Basket{
			ID:          "",
			Products:    nil,
			TotalAmount: 0,
		})
		assert.Nil(t, err)
		ID := basket.ID
		basket, err = db.Get(ID)
		assert.Nil(t, err)
		assert.Equal(t, ID, basket.ID)

	})

	t.Run("Create a record, update and  and get it", func(t *testing.T) {

		basket := model.Basket{
			ID:          "",
			Products:    nil,
			TotalAmount: 0,
		}
		basket, err := db.Create(basket)

		assert.Nil(t, err)
		ID := basket.ID

		basket, err = db.Get(ID)

		assert.Nil(t, err)
		assert.Equal(t, ID, basket.ID)

		newTotal := 1.0
		basket.TotalAmount = 1.0
		db.Update(basket)

		basket, _ = db.Get(ID)
		assert.Equal(t, ID, basket.ID)
		assert.Equal(t, newTotal, basket.TotalAmount)

	})
}
