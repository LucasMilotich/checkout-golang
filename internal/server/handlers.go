package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lucasmilotich/checkout-golang/internal/controller"
	"github.com/lucasmilotich/checkout-golang/internal/dao"
	"github.com/lucasmilotich/checkout-golang/internal/service"
	log "github.com/sirupsen/logrus"
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

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.Info("Server started")
}
