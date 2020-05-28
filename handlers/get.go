package handlers

import (
	"net/http"
	"strconv"

	"github.com/alvarotolentino/product-api/data"
	"github.com/gorilla/mux"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := data.ToJSON(lp, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products getProduct
// Returns a product
// responses:
//  200: productResponse
//  404: productNotFound

// GetProduct returns a product from the data store
func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to conver id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle GET Product", id)
	prod, err := data.GetProduct(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		http.Error(w, "Unable to read product", http.StatusInternalServerError)
		return
	}

}
