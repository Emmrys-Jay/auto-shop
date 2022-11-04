package product

import (
	"fmt"
	"io"
)

// StoreProduct represents the core attributes of every product that could be in the store.
// It is not an exhaustive list of attributes for every product, some products may have more
// attributes, but they will impliment this interface as far as they have these methods.
type StoreProduct interface {
	Name() string

	Quantity() int

	Price() string

	Description() string

	Model() string

	BrandName() string

	ProductType() string

	Sell(quantity int)
}

// Product models a product which can be placed in the store.
type Product struct {
	ID string
	StoreProduct
}

// DisplayProduct prints all available products to the command line
func (c *Product) DisplayProduct(w io.Writer) {
	fmt.Fprintln(w, "ID: ", c.ID)
	fmt.Fprintln(w, "Product Type: ", c.ProductType())
	fmt.Fprintln(w, "Brand: ", c.BrandName())
	fmt.Fprintln(w, "Model: ", c.Model())
	fmt.Fprintln(w, "Product Name: ", c.Name())
	fmt.Fprintln(w, "Quantity: ", c.Quantity())
	fmt.Fprintln(w, "Price: ", c.Price())

	switch p := c.StoreProduct.(type) {

	// After adding a new product, store manager should add a new 'switch Case' for it
	// and call its unique methods with a "fmt.Fprintln(w, )" funtion.

	case *Car:
		fmt.Fprintln(w, "Color: ", p.Color())
		fmt.Fprintln(w, "Fuel Type: ", p.FuelType())
	}

	fmt.Println()
}

// InStock checks if a product is in stock. It prints a response if product
// is in stock or otherwise.
func (c *Product) InStock() bool {
	return c.Quantity() > 1
}
