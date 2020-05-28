// Package handlers Provides a Product Rest API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import "github.com/alvarotolentino/product-api/data"

// swagger:response productNotFound
type productNotFoundError struct {
	// The error message
	// in: body
	// Example: Product not found
	Message string
}

// swagger:response validationError
type validationErrorWrapper struct {
	// The error message
	// in:body
	Body struct{
		Key string
		Error string
	}

}

// A product record
// swagger:response productResponse
type productResponseWrapper struct {
	// A product in the system
	// in: body
	Body data.Product
}

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productsNoContect struct {
	
}

// swagger:parameters deleteProduct getProduct updateProduct
type productIDParameterWrapper struct {
	// The id of the product
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters addProduct updateProduct
type productAddParameterWrapper struct {
	// Product data structure to Update or Create.
	// in: body
	// required: true
	Body data.Product
}