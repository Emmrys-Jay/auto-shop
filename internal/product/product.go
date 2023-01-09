package product

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

// StoreProduct represents the core attributes of every product that could be in the store.
// It is not an exhaustive list of attributes for every product, some products may have more
// attributes, but they will implement this interface as far as they have these methods.
type StoreProduct interface {
	Name() string

	GetDescription() string

	GetModel() string

	GetBrandName() string

	GetProductType() string

	GetFeatures() []string
}

// IsValid is a set which stores all valid product types
var IsValid = map[string]bool{
	"accessory": true,
	"car":       true,
}

// Product models a product which can be placed in the store.
type Product struct {
	StoreProduct
	Quantity int
	Price    float64
}

// DisplayProduct prints all available products an io.Writer
func (c *Product) DisplayProduct(w io.Writer) {
	json, err := json.MarshalIndent(c.StoreProduct, "", " ")
	if err != nil {
		log.Println("could not display products: %w", err)
	}

	fmt.Fprintln(w, string(json))
}

// InStock checks if a product is in stock. It prints a response if product
// is in stock or otherwise.
func (c *Product) InStock() bool {
	return c.Quantity > 1
}

// Sell alters the Quantity of an Accessory based on the Quantity parameter passed in.
func (c *Product) Sell(Quantity int) {
	c.Quantity = c.Quantity - Quantity
}
