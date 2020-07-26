package service

import (
	"errors"
	"github.com/lucasmilotich/checkout-golang/internal/dao"
	"github.com/lucasmilotich/checkout-golang/internal/dto"
	"github.com/lucasmilotich/checkout-golang/internal/model"
	"math"
)

type CheckoutService struct {
	ProductService ProductService
	DB             *dao.BasketDB
	DiscountDB     *dao.DiscountDB
}

type basketAux struct {
	quantity int32
	price    float64
}

func (s CheckoutService) GetById(id string) (model.Basket, *model.ApiError) {
	basket, err := s.DB.Get(id)
	if err != nil {
		return model.Basket{}, &model.ApiError{
			Code:    404,
			Message: "basket not found",
			Error:   err,
		}
	}
	basket.TotalAmount = s.calculateTotalAmount(basket.Products)

	return basket, nil
}

func (s CheckoutService) Delete(id string) (model.Basket, *model.ApiError) {
	basket, err := s.DB.Get(id)
	if err != nil {
		return model.Basket{}, &model.ApiError{
			Code:    404,
			Message: "basket not found",
			Error:   err,
		}
	}
	s.DB.Delete(basket.ID)
	return basket, nil
}

func (s CheckoutService) CreateCheckout(productsIDs []string) (model.Basket, *model.ApiError) {

	products, err := s.ProductService.GetAllByIDs(productsIDs)
	if err != nil {
		println("no product")
		return model.Basket{}, &model.ApiError{
			Code:    400,
			Message: "product does not exist",
			Error:   errors.New("error getting product"),
		}
	}

	basket, err := s.DB.Create(model.Basket{Products: products})

	if err != nil {
		return model.Basket{}, &model.ApiError{
			Code:    500,
			Message: "error saving basket",
			Error:   errors.New("error saving basket"),
		}
	}

	basket.TotalAmount = s.calculateTotalAmount(basket.Products)

	return basket, nil
}

func (s CheckoutService) ModifyBasket(id string, dto dto.CreationCheckoutDTO) (model.Basket, *model.ApiError) {

	basket, err := s.DB.Get(id)
	if err != nil {
		return model.Basket{}, &model.ApiError{
			Code:    404,
			Message: "basket did not found",
			Error:   err,
		}
	}

	products, err := s.ProductService.GetAllByIDs(dto.ProductIds)

	if err != nil {
		return model.Basket{}, &model.ApiError{
			Code:    400,
			Message: "some of the products does not exist",
			Error:   err,
		}
	}

	basket.Products = products
	return s.DB.Update(basket), nil

}

func (s CheckoutService) calculateTotalAmount(products []model.Product) float64 {
	totalAmount := 0.0

	quantityPerProduct := make(map[string]basketAux)

	for _, product := range products {
		totalAmount += product.Price
		aux := quantityPerProduct[product.Code]
		aux.quantity += 1
		aux.price = product.Price
		quantityPerProduct[product.Code] = aux
	}

	return totalAmount - s.getCalculatedDiscounts(quantityPerProduct)
}

func (s CheckoutService) getCalculatedDiscounts(quantityPerProduct map[string]basketAux) float64 {

	totalDiscount := 0.0

	discounts, _ := s.DiscountDB.GetAllDiscounts()

	for _, discount := range discounts {
		aux := quantityPerProduct[discount.ProductCode]

		if aux.quantity >= discount.MinProductsInBasket {
			totalDiscount += s.getDiscount(discount, aux)
		}
	}

	return totalDiscount
}

func (s CheckoutService) getDiscount(discount model.Discount, aux basketAux) float64 {
	if discount.PackagePromotion {
		quantityOfPromotionsToBeApplied := aux.quantity / discount.MinProductsInBasket
		return math.Round(
			float64(quantityOfPromotionsToBeApplied) *
				float64(discount.MinProductsInBasket) *
				discount.DiscountPerUnitToBeApplied *
				aux.price)
	} else {
		return discount.DiscountPerUnitToBeApplied * aux.price * float64(aux.quantity)
	}
}
