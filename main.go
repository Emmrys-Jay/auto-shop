package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Emmrys-Jay/auto-shop/product"
	"github.com/Emmrys-Jay/auto-shop/store"
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

	var p = []product.AddCarRequest{
		{
			ProductType: "Car",
			BrandName:   "Mercedes",
			Model:       "Sedan",
			Description: "Your dream personal ride",
			Color:       "Black",
			FuelType:    "electric",
			Quantity:    22,
			Price:       5898990.00,
		},
		{
			ProductType: "Car",
			BrandName:   "Lexus",
			Model:       "RX 350",
			Description: "Be your own boss while seated",
			Color:       "White",
			FuelType:    "PMS",
			Quantity:    50,
			Price:       2500990.00,
		},
	}

	util.Message(w, "ADD PRODUCTS")

	ids := availableProducts.AddProduct(w, &p)
	for i := 0; i < len(ids); i++ {
		fmt.Printf("You added product %v with id %v\n", p[i].BrandName+" "+p[i].Model, ids[i])
	}

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
