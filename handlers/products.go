package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/alvarotolentino/product-api/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}


type KeyProduct struct{}

// MiddlewareProductValidation validates all request
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}
		
		err = prod.Validate()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error validationg product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
