package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Emmrys-Jay/auto-shop/internal/product"
	"github.com/Emmrys-Jay/auto-shop/internal/store"
	"github.com/Emmrys-Jay/auto-shop/util"
)

func main() {

	// Set the output writer where the DEMO results will be printed to.
	// This variable allows the store manager to specify an output stream.
	// In this DEMO, the output stream is os.Stdout (the OS console)
	var w io.Writer = os.Stdout

	// DEMO
	util.Message(w, "DEMO")

	availableProducts := make(store.Store)
	soldProducts := make(store.Store)

	var p = []product.StoreProduct{
		&product.Vehicle{
			ProductType:     "car",
			BrandName:       "Mercedes",
			Model:           "Sedan",
			Description:     "Your dream personal ride",
			Color:           "Black",
			FuelType:        "electric",
			Quantity:        22,
			Price:           5898990.00,
			ManufactureYear: 2010,
		},
		&product.Vehicle{
			ProductType:     "Car",
			BrandName:       "Lexus",
			Model:           "RX 350",
			Description:     "Be your own boss while seated",
			Color:           "White",
			FuelType:        "PMS",
			Quantity:        50,
			Price:           2500990.00,
			ManufactureYear: 2011,
		},
		&product.Accessory{
			ProductType: "wipers",
			BrandName:   "",
			Model:       "",
			Description: "Black car wipers of length 3ft",
			Features:    []string{},
			Color:       "black",
			Quantity:    1000,
			Price:       200.00,
			ExpiryDate:  "2022-10",
		},
	}

	util.Message(w, "ADD PRODUCTS")

	ids := availableProducts.AddProducts(w, p)

	util.Message(w, "LIST PRODUCTS")

	availableProducts.ListProducts(w)

	util.Message(w, "SELL PRODUCT")

	availableProducts.SellProduct(w, ids[0], 10, &soldProducts)

	fmt.Fprintln(w)

	availableProducts.SellProduct(w, ids[1], 20, &soldProducts)

	util.Message(w, "LIST SOLD ITEMS")

	availableProducts.ListSoldItems(w, &soldProducts)

	util.Message(w, "END OF DEMO")
}
