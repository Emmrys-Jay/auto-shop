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

	myStore := &store.Store{
		AvailableProducts: map[string]*product.Product{},
		SoldProducts:      map[string]*store.SoldItems{},
		Orders:            map[string]*store.OrderedItems{},
	}

	p := []product.Product{
		product.Product{
			Price:    2500990.00,
			Quantity: 50,
			StoreProduct: &product.Car{
				ProductType:     "car",
				BrandName:       "Mercedes",
				Model:           "Sedan",
				Description:     "Your dream personal ride",
				Color:           "Black",
				FuelType:        "electric",
				ManufactureYear: 2010,
			},
		},
		product.Product{
			Price:    2500990.00,
			Quantity: 50,
			StoreProduct: &product.Car{
				ProductType:     "Car",
				BrandName:       "Lexus",
				Model:           "RX 350",
				Description:     "Be your own boss while seated",
				Color:           "White",
				FuelType:        "PMS",
				ManufactureYear: 2011,
			},
		},
		product.Product{
			Price:    200.00,
			Quantity: 1000,
			StoreProduct: &product.Accessory{
				ProductType: "accessory",
				BrandName:   "Wipers",
				Model:       "",
				Description: "Black car wipers of length 3ft",
				Features:    []string{},
				Color:       "black",
			},
		},
	}

	util.Message(w, "ADD PRODUCTS")

	// Add the three products
	ids := myStore.AddProducts(w, p)

	util.Message(w, "LIST PRODUCTS")

	// List the number of products added
	myStore.ListProducts(w)

	util.Message(w, "SELL PRODUCT")

	// Sell one of the products
	myStore.SellProduct(w, ids[0], 10)

	fmt.Fprintln(w)

	// Sell another one of the products
	myStore.SellProduct(w, ids[2], 20)

	fmt.Fprintln(w)

	// Sell the first product again
	myStore.SellProduct(w, ids[0], 10)

	util.Message(w, "LIST SOLD ITEMS")

	// List all sold items
	myStore.ListSoldItems(w)

	util.Message(w, "LIST ALL ORDERS")

	// List all ordered items
	myStore.ListOrderedItems(w)

	util.Message(w, "LIST TOTAL PRICE OF ALL AVAILABLE PRODUCTS")

	// List total price of all products available
	myStore.TotalPriceOfProductsLeft(w)

	util.Message(w, "LIST TOTAL PRICE OF ALL SOLD PRODUCTS")

	// List total price of all products available
	myStore.TotalPriceOfProductsSold(w)

	util.Message(w, "END OF DEMO")
}
