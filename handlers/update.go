package handlers

import (
	"net/http"
	"strconv"

	"github.com/alvarotolentino/product-api/data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products updateProduct
// Update an existing product
// responses:
//  200: noContent
//  400: validationError
//  404: productNotFound

// UpdateProducts updates a product by the ID
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}
