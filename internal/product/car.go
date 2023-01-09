package product

import "strings"

// Car Models the core Car struct which satisfies the interface "StoreProduct".
type Car struct {
	ProductType     string
	BrandName       string
	Model           string
	Description     string
	Features        []string
	Color           string
	FuelType        string
	ManufactureYear int
}

// Name returns the name of a Car. Name is a combination of BrandName and Model
// separated by a whitespace.
func (c *Car) Name() string {
	return strings.TrimSpace(c.GetBrandName() + " " + c.GetModel())
}

// GetDescription returns the Description of a Car.
func (c *Car) GetDescription() string {
	return c.Description
}

// GetModel returns the Model of a Car.
func (c *Car) GetModel() string {
	return c.Model
}

// GetBrandName returns the BrandName of a Car
func (c *Car) GetBrandName() string {
	return c.BrandName
}

// GetColor returns the Color of a Car.
func (c *Car) GetColor() string {
	return c.Color
}

// GetFuelType returns the fuel-type of a Car.
func (c *Car) GetFuelType() string {
	return c.FuelType
}

// GetProductType returns the product-type of a Car.
func (c *Car) GetProductType() string {
	return c.ProductType
}

// GetFeatures returns all features of a Car
func (c *Car) GetFeatures() []string {
	return c.Features
}

// GetManufactureDate returns all features of a Car
func (c *Car) GetManufactureYear() int {
	return c.ManufactureYear
}
