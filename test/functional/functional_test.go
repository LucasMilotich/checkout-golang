package functional

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/lucasmilotich/checkout-golang/internal/model"
	checkoutServer "github.com/lucasmilotich/checkout-golang/internal/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server *httptest.Server
var httpClient *http.Client

var checkoutEndpoint string

func init() {
	r := chi.NewRouter()

	checkoutServer.BindEndpoints(r)
	server = httptest.NewServer(r)

	httpClient = &http.Client{}

	checkoutEndpoint = "/checkouts/"
}

func Test_checkout(t *testing.T) {

	body := `{	"product_ids" : ["PEN", "PEN", "PEN"] }`

	t.Run("Create checkout and get it", func(t *testing.T) {
		response, _ := httpClient.Post(server.URL+checkoutEndpoint, "application/json", bytes.NewBufferString(body))

		createdBasket := model.Basket{}
		_ = json.NewDecoder(response.Body).Decode(&createdBasket)

		assert.NotNil(t, createdBasket)

		response, _ = httpClient.Get(server.URL + checkoutEndpoint + createdBasket.ID)

		basket := model.Basket{}
		_ = json.NewDecoder(response.Body).Decode(&basket)

		assert.Equal(t, createdBasket, basket)

	})

	t.Run("Create checkout, add a product, and get it", func(t *testing.T) {
		response, _ := httpClient.Post(server.URL+checkoutEndpoint, "application/json", bytes.NewBufferString(body))

		createdBasket := model.Basket{}
		_ = json.NewDecoder(response.Body).Decode(&createdBasket)

		assert.NotNil(t, createdBasket)

		response, _ = httpClient.Get(server.URL + checkoutEndpoint + createdBasket.ID)

		basket := model.Basket{}
		_ = json.NewDecoder(response.Body).Decode(&basket)

		assert.Equal(t, createdBasket, basket)

	})

}

func init() {

}
