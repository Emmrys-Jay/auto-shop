package product

import "fmt"

// AddCarRequest is used to create a car in main.
type AddCarRequest struct {
	ProductType string
	BrandName   string
	Model       string
	Description string
	Color       string
	FuelType    string
	Quantity    int
	Price       float64
}

// CarRequestToCar maps an AddCarRequest to a Car.
func CarRequestToCar(p *AddCarRequest) *Car {
	var car = &Car{
		brandName:   p.BrandName,
		productType: p.ProductType,
		model:       p.Model,
		description: p.Description,
		color:       p.Color,
		fuelType:    p.FuelType,
		quantity:    p.Quantity,
		price:       p.Price,
	}

	return car
}

// Car models the core Car struct which satisfies the interface "StoreProduct".
type Car struct {
	productType string
	brandName   string
	model       string
	description string
	color       string
	fuelType    string
	quantity    int
	price       float64
}

// Name returns the name of a Car. Name is a combination of BrandName and Model
// separated by a whitespace.
func (c *Car) Name() string {
	return c.BrandName() + " " + c.Model()
}

// Quantity returns the quantity of a Car.
func (c *Car) Quantity() int {
	return c.quantity
}

// Price returns the price of a Car.
func (c *Car) Price() string {
	return fmt.Sprintf("%.2f", c.price)
}

// Description returns the description of a Car.
func (c *Car) Description() string {
	return c.description
}

// Model returns the model of a Car.
func (c *Car) Model() string {
	return c.model
}

// BrandName returns the BrandName of a car
func (c *Car) BrandName() string {
	return c.brandName
}

// Color retuns the color of a Car.
func (c *Car) Color() string {
	return c.color
}

// FuelType returns the fuel type of a Car.
func (c *Car) FuelType() string {
	return c.fuelType
}

// ProductType returns the product type of a Car.
func (c *Car) ProductType() string {
	return c.productType
}

// Sell alters the quantity of a Car based on the parameter passed in.
func (c *Car) Sell(quantity int) {
	c.quantity -= quantity
}
