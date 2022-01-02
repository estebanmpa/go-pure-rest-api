package endpoints

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Product struct {
	ID       int64
	Name     string
	LastName string
	Email    float32
}

type Products []Product

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/product", getProduct).Methods("GET")
	router.HandleFunc("/product", saveProduct).Methods("POST", "PUT")
}

func getProduct(w http.ResponseWriter, r *http.Request) {
}

func saveProduct(w http.ResponseWriter, r *http.Request) {
}
