package service

import (
	"github.com/lucasmilotich/coreapi/internal/dao"
	"github.com/lucasmilotich/coreapi/internal/model"
)

type ProductService struct {
	DB *dao.ProductDB
}

func (s ProductService) GetById(id string) (model.Product, error) {
	return s.DB.Get(id)
}
func (s ProductService) GetAllByIDs(ids []string) ([]model.Product, error) {
	return s.DB.GetAllByIDs(ids)
}