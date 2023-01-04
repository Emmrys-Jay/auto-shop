package product

import (
	"encoding/json"
	"io"
	"log"
)

// StoreProduct represents the core attributes of every product that could be in the store.
// It is not an exhaustive list of attributes for every product, some products may have more
// attributes, but they will implement this interface as far as they have these methods.
type StoreProduct interface {
	Name() string

	GetQuantity() int

	GetPrice() string

	GetDescription() string

	GetModel() string

	GetBrandName() string

	GetProductType() string

	GetFeatures() []string

	Sell(quantity int)
}

// Product models a product which can be placed in the store.
type Product struct {
	ID string
	StoreProduct
}

// DisplayProduct prints all available products to the command line
func (c *Product) DisplayProduct(w io.Writer) {
	if _, err := json.MarshalIndent(c.StoreProduct, "", " "); err != nil {
		log.Println("could not display products: %w", err)
	}
}

// InStock checks if a product is in stock. It prints a response if product
// is in stock or otherwise.
func (c *Product) InStock() bool {
	return c.GetQuantity() > 1
}
