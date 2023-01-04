package product

import "fmt"

// Vehicle Models the core Vehicle struct which satisfies the interface "StoreProduct".
type Vehicle struct {
	ProductType     string
	BrandName       string
	Model           string
	Description     string
	Features        []string
	Color           string
	FuelType        string
	Quantity        int
	Price           float64
	ManufactureYear int
}

// Name returns the name of a Vehicle. Name is a combination of BrandName and Model
// separated by a whitespace.
func (c *Vehicle) Name() string {
	return c.GetBrandName() + " " + c.GetModel()
}

// GetQuantity returns the Quantity of a Vehicle.
func (c *Vehicle) GetQuantity() int {
	return c.Quantity
}

// GetPrice returns the Price of a Vehicle.
func (c *Vehicle) GetPrice() string {
	return fmt.Sprintf("%.2f", c.Price)
}

// GetDescription returns the Description of a Vehicle.
func (c *Vehicle) GetDescription() string {
	return c.Description
}

// GetModel returns the Model of a Vehicle.
func (c *Vehicle) GetModel() string {
	return c.Model
}

// GetBrandName returns the BrandName of a Vehicle
func (c *Vehicle) GetBrandName() string {
	return c.BrandName
}

// GetColor returns the Color of a Vehicle.
func (c *Vehicle) GetColor() string {
	return c.Color
}

// GetFuelType returns the fuel-type of a Vehicle.
func (c *Vehicle) GetFuelType() string {
	return c.FuelType
}

// GetProductType returns the product-type of a Vehicle.
func (c *Vehicle) GetProductType() string {
	return c.ProductType
}

// GetSell alters the Quantity of a Vehicle based on the Quantity parameter passed in.
func (c *Vehicle) Sell(Quantity int) {
	c.Quantity = c.Quantity - Quantity
}

// GetFeatures returns all features of a vehicle
func (c *Vehicle) GetFeatures() []string {
	return c.Features
}

// GetManufactureDate returns all features of a vehicle
func (c *Vehicle) GetManufactureYear() int {
	return c.ManufactureYear
}
