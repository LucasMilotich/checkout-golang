package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lucasmilotich/checkout-golang/internal/server"
)

// go run main.go
// curl localhost:8080/777
func main() {
	r := chi.NewRouter()

	server.BindEndpoints(r)

	err := http.ListenAndServe(":8080", r)
	panic(err.Error())
}
