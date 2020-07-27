package controller

import (
	"encoding/json"
	"github.com/lucasmilotich/checkout-golang/internal/dto"
	"github.com/lucasmilotich/checkout-golang/internal/service"
	"net/http"

	"github.com/go-chi/chi"
)

type CheckoutController struct {
	CheckoutService service.CheckoutService
}

func (c *CheckoutController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	entity, apiError := c.CheckoutService.Delete(id)
	if apiError != nil {
		w.WriteHeader(apiError.Code)
		errorData, _ := json.Marshal(apiError)
		_, _ = w.Write(errorData)
		return
	}

	data, err := json.Marshal(entity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("{\"error\": failed to marshal}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (c *CheckoutController) Modify(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	checkoutDTO := dto.CreationCheckoutDTO{}
	err := json.NewDecoder(r.Body).Decode(&checkoutDTO)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("{\"error\": error unmarshalling}"))
		return
	}

	entity, apiError := c.CheckoutService.ModifyBasket(id, checkoutDTO)
	if apiError != nil {
		w.WriteHeader(apiError.Code)
		errorData, _ := json.Marshal(apiError)
		_, _ = w.Write(errorData)
		return
	}

	data, err := json.Marshal(entity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("{\"error\": failed to marshal}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (c *CheckoutController) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	entity, apiError := c.CheckoutService.GetById(id)
	if apiError != nil {
		w.WriteHeader(apiError.Code)
		errorData, _ := json.Marshal(apiError)
		_, _ = w.Write(errorData)
		return
	}

	data, err := json.Marshal(entity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("{\"error\": failed to marshal}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (c *CheckoutController) Create(w http.ResponseWriter, r *http.Request) {

	checkoutDTO := dto.CreationCheckoutDTO{}
	err := json.NewDecoder(r.Body).Decode(&checkoutDTO)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("{\"error\": error unmarshalling}"))
		return
	}

	entity, apiError := c.CheckoutService.CreateCheckout(checkoutDTO.ProductIds)
	if apiError != nil {
		w.WriteHeader(apiError.Code)
		errorData, _ := json.Marshal(apiError)
		_, _ = w.Write(errorData)
		return
	}

	data, err := json.Marshal(entity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("{\"error\": failed to marshal}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
