package handlers

import (
	"net/http"

	"github.com/alvarotolentino/product-api/data"
)

// AddProduct adds a new product to the data store
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}
