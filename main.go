package main

import (
	"fmt"

	"github.com/Emmrys-Jay/auto-shop/product"
	"github.com/Emmrys-Jay/auto-shop/store"
	"github.com/Emmrys-Jay/auto-shop/util"
)

func main() {
	// DEMO
	util.Message("DEMO")

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

	util.Message("ADD PRODUCTS")

	ids := availableProducts.AddProduct(&p)
	for i := 0; i < len(ids); i++ {
		fmt.Printf("You added product %v with id %v\n", p[i].BrandName+" "+p[i].Model, ids[i])
	}

	util.Message("LIST PRODUCTS")

	availableProducts.ListProducts()

	util.Message("SELL PRODUCT")

	availableProducts.SellProduct(ids[0], 10, &soldProducts)

	fmt.Println()

	availableProducts.SellProduct(ids[1], 20, &soldProducts)

	util.Message("LIST SOLD ITEMS")

	availableProducts.ListSoldItems(&soldProducts)

	util.Message("END OF DEMO")
}
