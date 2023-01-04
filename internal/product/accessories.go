package product

import (
	"fmt"
	"strings"
)

// Accessory Models Accessory struct which satisfies the interface "StoreProduct".
type Accessory struct {
	ProductType string
	BrandName   string
	Model       string
	Description string
	Features    []string
	Color       string
	Quantity    int
	Price       float64
	// Expire date is an optional time string field of the format (2022-10)
	ExpiryDate string
}

// Name returns the name of an Accessory. Name is a combination of BrandName and Model
// separated by a whitespace.
func (c *Accessory) Name() string {
	return strings.TrimSpace(c.GetBrandName() + " " + c.GetModel())
}

// GetQuantity returns the Quantity of an Accessory.
func (c *Accessory) GetQuantity() int {
	return c.Quantity
}

// GetPrice returns the Price of an Accessory.
func (c *Accessory) GetPrice() string {
	return fmt.Sprintf("%.2f", c.Price)
}

// GetDescription returns the Description of an Accessory.
func (c *Accessory) GetDescription() string {
	return c.Description
}

// GetModel returns the Model of an Accessory.
func (c *Accessory) GetModel() string {
	return c.Model
}

// GetBrandName returns the BrandName of an Accessory
func (c *Accessory) GetBrandName() string {
	return c.BrandName
}

// GetColor returns the Color of an Accessory.
func (c *Accessory) GetColor() string {
	return c.Color
}

// GetProductType returns the product-type of an Accessory.
func (c *Accessory) GetProductType() string {
	return c.ProductType
}

// GetSell alters the Quantity of an Accessory based on the Quantity parameter passed in.
func (c *Accessory) Sell(Quantity int) {
	c.Quantity = c.Quantity - Quantity
}

// GetFeatures returns all features of an Accessory
func (c *Accessory) GetFeatures() []string {
	return c.Features
}

// GetExpiryDate returns all features of an Accessory
func (c *Accessory) GetExpiryDate() string {
	return c.ExpiryDate
}
