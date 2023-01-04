package store

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Emmrys-Jay/auto-shop/internal/product"
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

// ProductTypes is a set which stores all valid product types
var ProductTypes = map[string]bool{
	"accessory":  true,
	"vehicle":    true,
	"van":        true,
	"car":        true,
	"suv":        true,
	"ambulance":  true,
	"bus":        true,
	"motorcycle": true,
	"bike":       true,
	"scooter":    true,
}

// AddProductType adds a product type to the list of valid product types
func AddProductType(w io.Writer, productType string) {
	ProductTypes[productType] = true

	fmt.Fprintf(w, "Successfully added a product of type %s", productType)
}

// NoOfProductsForSale calculates the number of products available in the store
func (s *Store) NoOfProductsForSale(w io.Writer) {
	quantity := 0

	for _, v := range *s {
		if v.InStock() {
			quantity++
		}
	}

	fmt.Fprintf(w, "There are %v products up for sale\n", quantity)
}

// AddProduct adds new products to the store
func (s *Store) AddProducts(w io.Writer, req []product.StoreProduct) []string {
	if len(req) == 0 {
		fmt.Fprintf(w, "No products specified")
	}

	fmt.Fprintln(w, "Adding products to store...")
	ids := make([]string, 0)

	var count int
	var invalidProducts string
	for _, v := range req {
		if !ProductTypes[strings.ToLower(v.GetProductType())] {
			invalidProducts += " " + v.GetProductType()
			continue
		}
		count++
		id := uuid.New()

		var p = &product.Product{
			ID:           id.String(),
			StoreProduct: v,
		}

		(*s)[id.String()] = p
		ids = append(ids, id.String())
	}

	if invalidProducts == "" {
		fmt.Fprintf(w, "Added %v products successfully...\n", count)
	} else if count != 0 {
		fmt.Fprintf(w, "Added %v products successfully...\n", count)
		fmt.Fprintf(w, "Products with types %s are invalid and were not added", strings.ReplaceAll(strings.TrimSpace(invalidProducts), " ", ", "))
	} else {
		fmt.Fprintf(w, "All product types specified are invalid")
	}

	return ids
}

// ListProducts prints all products available in the store
func (s *Store) ListProducts(w io.Writer) {
	for _, v := range *s {
		if v.InStock() {
			v.DisplayProduct(w)
		}
	}
}

// SellProduct sells a product by updating its quantity, and adds the sold product to
// the store for products sold.
func (s *Store) SellProduct(w io.Writer, id string, quantity int, salesStore *Store) {
	if _, ok := (*s)[id]; ok {
		(*s)[id].Sell(quantity)
	} else {
		fmt.Fprintln(w, "Product you specified does not exist")
		return
	}

	if _, ok := (*salesStore)[id]; !ok {
		(*salesStore)[id] = (*s)[id]
		(*salesStore)[id].Sell((*s)[id].GetQuantity() - quantity)
	} else {
		(*salesStore)[id].Sell(-quantity)
	}

	fmt.Fprintf(w, "Congrats! You just sold %v %v, with id %v\n", quantity, (*s)[id].Name(), id)
}

// ListSoldItems displays all products that has been sold from the store.
func (s *Store) ListSoldItems(w io.Writer, salesStore *Store) {
	si := make([]SoldItems, 0)
	totalPrice := 0.0

	for _, v := range *salesStore {
		s := SoldItems{
			ID:           v.ID,
			Name:         v.Name(),
			QuantitySold: v.GetQuantity(),
			PriceSold:    v.GetPrice(),
		}

		si = append(si, s)

		price, _ := strconv.ParseFloat(v.GetPrice(), 64)
		totalPrice += price
	}

	siDisplay := SoldItemsDisplay{
		TotalPrice: totalPrice,
		SoldItems:  si,
	}

	jsonDisplay, err := json.MarshalIndent(siDisplay, "", " ")
	if err != nil {
		fmt.Fprintln(w, "error: Could not encode JSON")
		return
	}

	fmt.Fprintln(w, string(jsonDisplay))
}
