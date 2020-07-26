package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lucasmilotich/checkout-golang/internal/server"
)

// go run main.go
// curl localhost:8080/777
func main() {
	err := startServer()
	panic(err.Error())
}

func startServer() error {
	r := chi.NewRouter()

	server.BindEndpoints(r)

	err := http.ListenAndServe(":8080", r)
	return err
}
