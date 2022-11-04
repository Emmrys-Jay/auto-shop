package store

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/Emmrys-Jay/auto-shop/product"
	"github.com/google/uuid"
)

type (
	// Store models the store for available products and sold products
	Store map[string]*product.Product

	// SoldItemsDisplay models the structure products will be displayed from
	// the sold items store
	SoldItemsDisplay struct {
		TotalPrice float64     `json:"total_price"`
		SoldItems  []SoldItems `json:"products_sold"`
	}

	// SoldItems is a sub model of SoldItemsDisplay
	SoldItems struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		QuantitySold int    `json:"quantity_sold"`
		PriceSold    string `json:"price_sold"`
	}
)

// NoOfProductsForSale calculates the number of products available in the store
func (s *Store) NoOfProductForSale() {
	quantity := 0

	for _, v := range *s {
		if v.Quantity() > 0 {
			quantity++
		}
	}

	fmt.Printf("There are %v products up for sale\n", quantity)
}

// AddProduct adds new products to the store
func (s *Store) AddProduct(req *[]product.AddCarRequest) []string {
	fmt.Println("Adding products to store...")
	ids := make([]string, 0)

	for _, v := range *req {
		car := product.CarRequestToCar(&v)
		id := uuid.New()

		var p = &product.Product{
			ID:           id.String(),
			StoreProduct: car,
		}

		(*s)[id.String()] = p
		ids = append(ids, id.String())
	}

	fmt.Printf("Added %v products successfully...\n", len(*req))

	return ids
}

// ListProducts prints all products available in the store
func (s *Store) ListProducts() {
	for _, v := range *s {
		if v.Quantity() > 0 {
			v.DisplayProduct(os.Stdout)
		}
	}
}

// SellProduct sells a product by updating its quantity, and adds the sold product to
// the store for products sold.
func (s *Store) SellProduct(id string, quantity int, salesStore *Store) {
	if _, ok := (*s)[id]; ok {
		(*s)[id].Sell(quantity)
	} else {
		fmt.Println("Product you specified does not exist")
		return
	}

	if _, ok := (*salesStore)[id]; !ok {
		(*salesStore)[id] = (*s)[id]
		(*salesStore)[id].Sell(-(*s)[id].Quantity() + quantity)
	} else {
		(*salesStore)[id].Sell(-quantity)
	}

	fmt.Printf("Congrats! You just sold %v %v, with id %v\n", quantity, (*s)[id].Name(), id)
}

// ListSoldItems displays all products that has been sold from the store.
func (s *Store) ListSoldItems(salesStore *Store) {
	si := make([]SoldItems, 0)
	totalPrice := 0.0

	for _, v := range *salesStore {
		s := SoldItems{
			ID:           v.ID,
			Name:         v.Name(),
			QuantitySold: v.Quantity(),
			PriceSold:    v.Price(),
		}

		si = append(si, s)

		price, _ := strconv.ParseFloat(v.Price(), 64)
		totalPrice += price
	}

	siDisplay := SoldItemsDisplay{
		TotalPrice: totalPrice,
		SoldItems:  si,
	}

	jsonDisplay, err := json.MarshalIndent(siDisplay, "", " ")
	if err != nil {
		fmt.Println("error: Could not encode JSON")
		return
	}

	fmt.Println(string(jsonDisplay))
}
