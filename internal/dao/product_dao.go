package dao

import (
	"errors"
	"sync"

	"github.com/lucasmilotich/coreapi/internal/model"
)

type ProductDB struct {
	elements map[string]model.Product
	mux      *sync.Mutex
}

func NewProductDB() *ProductDB {
	return &ProductDB{
		elements: map[string]model.Product{},
		mux:      &sync.Mutex{},
	}
}

func (db *ProductDB) Get(id string) (model.Product, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	if id == "PEN" {
		return model.Product{Code: id, Name: "Lana Pen", Price: 5.0}, nil
	}

	if id == "TSHIRT" {
		return model.Product{Code: id, Name: "Lana T-Shirt", Price: 20.0}, nil
	}

	if id == "MUG" {
		return model.Product{Code: id, Name: "Lana Coffee Mug", Price: 7.50}, nil
	}

	entity, ok := db.elements[id]
	if !ok {
		return model.Product{}, errors.New("not found")
	}

	return entity, nil
}

func (db *ProductDB) GetAllByIDs(ids []string) ([]model.Product, error) {

	products := make([]model.Product, 0)

	for _, b := range ids {
		product, err := db.Get(b)
		if err != nil {
			return nil, errors.New("product not found")
		}
		products = append(products, product)
	}

	return products, nil

}
