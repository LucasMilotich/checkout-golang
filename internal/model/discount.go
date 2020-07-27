package model

type Discount struct {
	ProductCode         string
	MinProductsInBasket int32
	// package promotion is true when you have to get a discount buying groups of X number
	PackagePromotion           bool
	DiscountPerUnitToBeApplied float64
}
