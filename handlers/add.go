package handlers

import (
	"net/http"

	"github.com/alvarotolentino/product-api/data"
)

// swagger:route POST /products products addProduct
// Create a new product
// responses:
//  200: noContent
//  400: validationError

// AddProduct adds a new product to the data store
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(prod)
}
