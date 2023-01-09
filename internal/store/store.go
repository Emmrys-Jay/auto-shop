package store

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/Emmrys-Jay/auto-shop/internal/product"
	"github.com/google/uuid"
)

type (
	// Store models the store for available products and sold products
	Store struct {
		AvailableProducts map[string]*product.Product
		SoldProducts      map[string]*SoldItems
		Orders            map[string]*OrderedItems
	}

	// SoldItems is a struct which models a sold item
	SoldItems struct {
		ProductID    string `json:"product_id"`
		ProductName  string `json:"product_name"`
		QuantitySold int    `json:"quantity_sold"`
	}

	// OrderedItems is a struct which models an ordered item
	OrderedItems struct {
		ProductID    string  `json:"id"`
		ProductName  string  `json:"product_name"`
		PriceSold    float64 `json:"price_sold"`
		QuantitySold int     `json:"quantity_sold"`
	}
)

// AddProductType adds a product type to the list of valid product types
func AddProductType(w io.Writer, productType string) {
	product.IsValid[productType] = true

	fmt.Fprintf(w, "Successfully added a product of type %s", productType)
}

// AddProduct adds new products to the store
func (s *Store) AddProducts(w io.Writer, req []product.Product) []string {
	if len(req) == 0 {
		fmt.Fprintf(w, "No products specified")
	}

	fmt.Fprintln(w, "Adding products to store...")
	ids := make([]string, 0)

	var count int
	var invalidProducts string
	for _, v := range req {
		if !product.IsValid[strings.ToLower(v.GetProductType())] {
			invalidProducts += " " + v.GetProductType()
			continue
		}
		count++
		id := uuid.New()

		var p = &product.Product{
			StoreProduct: v.StoreProduct,
			Quantity:     v.Quantity,
			Price:        v.Price,
		}

		s.AvailableProducts[id.String()] = p
		ids = append(ids, id.String())
	}

	if invalidProducts == "" {
		fmt.Fprintf(w, "Added %v products successfully...\n", count)
	} else if count != 0 {
		fmt.Fprintf(w, "Added %v products successfully...\n", count)
		fmt.Fprintf(w, "Products with type(s) %s are invalid and were not added", strings.ReplaceAll(strings.TrimSpace(invalidProducts), " ", ", "))
	} else {
		fmt.Fprintf(w, "All product types specified are invalid")
	}

	return ids
}

// ListProducts prints all products available in the store
func (s *Store) ListProducts(w io.Writer) {
	for _, v := range s.AvailableProducts {
		if v.InStock() {
			v.DisplayProduct(w)
		}
	}
}

// SellProduct sells a product by updating its quantity, and adds the sold product to
// the store for products sold.
func (s *Store) SellProduct(w io.Writer, id string, quantity int) {
	var product *product.Product
	if _, ok := s.AvailableProducts[id]; ok {
		product = s.AvailableProducts[id]
		product.Sell(quantity)
	} else {
		fmt.Fprintln(w, "Product you specified does not exist")
		return
	}

	if _, ok := s.SoldProducts[id]; !ok {
		s.SoldProducts[id] = &SoldItems{
			ProductID:    id,
			ProductName:  product.Name(),
			QuantitySold: quantity,
		}
	} else {
		s.SoldProducts[id].QuantitySold += quantity
	}

	orderID := uuid.New()
	s.Orders[orderID.String()] = &OrderedItems{
		ProductID:    id,
		ProductName:  product.Name(),
		PriceSold:    product.Price,
		QuantitySold: quantity,
	}

	fmt.Fprintf(w, "Congrats! You just sold %d %s(s), with id %v\n", quantity, product.Name(), id)
}

// NoOfProductsForSale writes the number of products available in the store
func (s *Store) NoOfProductsForSale(w io.Writer) {
	quantity := 0

	for _, v := range s.AvailableProducts {
		if v.InStock() {
			quantity++
		}
	}

	fmt.Fprintf(w, "There are %v products up for sale\n", quantity)
}

// ListSoldItems writes all products that has been sold from the store to w.
func (s *Store) ListSoldItems(w io.Writer) {
	for _, v := range s.SoldProducts {
		result, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			log.Fatalln("error: could not print all sold products: %w", err)
		}

		fmt.Fprintln(w, string(result))
	}
}

// ListOrderedItems writes all orders to w.
func (s *Store) ListOrderedItems(w io.Writer) {
	for _, v := range s.Orders {
		result, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			log.Fatalln("error: could not print all ordered products: %w", err)
		}

		fmt.Fprintln(w, string(result))
	}
}

// TotalPriceOfProductsLeft writes the total price of all products left to w.
func (s *Store) TotalPriceOfProductsLeft(w io.Writer) {
	var total float64
	for _, v := range s.AvailableProducts {
		if v.InStock() {
			total += (float64(v.Quantity) * v.Price)
		}
	}

	fmt.Fprintf(w, "The total price of products left in store is %.2f\n", total)
}

// TotalPriceOfProductsSold writes the total price of all products sold to w.
func (s *Store) TotalPriceOfProductsSold(w io.Writer) {
	var total float64
	for _, v := range s.Orders {
		total += (float64(v.QuantitySold) * v.PriceSold)
	}

	fmt.Fprintf(w, "The total price of products sold is %.2f\n", total)
}

// UpdateProductPrice updates the price of a product. It writes either an error or
// success statement to w.
func (s *Store) UpdateProductPrice(w io.Writer, id string, price float64) {
	var product *product.Product
	if _, ok := s.AvailableProducts[id]; !ok {
		fmt.Fprintln(w, "Product you specifed does not exist")
		return
	}

	product = s.AvailableProducts[id]
	if price > 0 {
		product.Price = price
	} else {
		fmt.Fprintln(w, "Invalid price")
		return
	}

	fmt.Fprintf(w, "You successfully updated the price of %s with id: %s\n", product.Name(), id)

}

// AddToProductQuantity updates the quantity of a product. It writes either an error or
// success statement to w.
func (s *Store) AddToProductQuantity(w io.Writer, id string, quantity int) {
	var product *product.Product
	if _, ok := s.AvailableProducts[id]; !ok {
		fmt.Fprintln(w, "Product you specifed does not exist")
		return
	}

	product = s.AvailableProducts[id]
	if quantity > 0 {
		product.Quantity += quantity
	} else {
		fmt.Fprintln(w, "Invalid Quantity: You can only increase product quantity through this method")
		return
	}

	fmt.Fprintf(w, "You successfully added %d %s(s) with id: %s\n", quantity, product.Name(), id)

}
