package dao

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lucasmilotich/coreapi/internal/model"
	"sync"
)

type BasketDB struct {
	elements map[string]model.Basket
	mux      *sync.Mutex
}

func NewBasketDB() *BasketDB {
	return &BasketDB{
		elements: map[string]model.Basket{},
		mux:      &sync.Mutex{},
	}
}

func (db *BasketDB) Get(id string) (model.Basket, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	entity, ok := db.elements[id]
	if !ok {
		return model.Basket{}, errors.New("not found")
	}

	return entity, nil
}

func (db *BasketDB) Create(basket model.Basket) (model.Basket, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	basket.ID = uuid.New().String()
	db.elements[basket.ID] = basket

	return basket, nil
}

func (db *BasketDB) Update(basket model.Basket) model.Basket {
	db.mux.Lock()
	defer db.mux.Unlock()

	db.elements[basket.ID] = basket

	return basket
}

func (db *BasketDB) Delete(id string) {
	db.mux.Lock()
	defer db.mux.Unlock()

	delete(db.elements, id)

}
