package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lucasmilotich/coreapi/internal/controller"
	"github.com/lucasmilotich/coreapi/internal/dao"
	"github.com/lucasmilotich/coreapi/internal/service"
)

func setupMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
}

func BindEndpoints(r *chi.Mux) {

	setupMiddlewares(r)

	checkoutController := controller.CheckoutController{
		CheckoutService: service.CheckoutService{
			ProductService: service.ProductService{DB: dao.NewProductDB()},
			DB:             dao.NewBasketDB(),
			DiscountDB:     dao.NewDiscountDB(),
		},
	}

	r.Use(middleware.Logger)
	r.Get("/checkouts/{id}", checkoutController.Get)
	r.Put("/checkouts/{id}", checkoutController.Modify)
	r.Delete("/checkouts/{id}", checkoutController.Delete)
	r.Post("/checkouts", checkoutController.Create)

	println("server started")
}
