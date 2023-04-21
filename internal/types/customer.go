package types

import (
	"math/rand"
	"time"

	"github.com/s3ndd/coffeeshop/internal/config"
)

// extrasOptions is a list of extras options
// for simulation purposes, the different options are hard-coded
// different options will be affected the order's price
var extrasOptions = [][]string{
	{},
	{"milk"},
	{"sugar"},
	{"milk", "sugar"},
}

// Customer represents a customer
type Customer struct {
	name        string
	arrivedTime time.Time
	leaveTime   *time.Time
	config      config.Configurer
}

// NewCustomer creates a new customer
func NewCustomer(name string, config config.Configurer) *Customer {
	return &Customer{
		name:        name,
		arrivedTime: time.Now(),
		config:      config,
	}
}

// Name returns the customer's name
func (c *Customer) Name() string {
	return c.name
}

// ArrivedTime returns the time the customer arrived
func (c *Customer) ArrivedTime() time.Time {
	return c.arrivedTime
}

// LeaveTime returns the time the customer left
func (c *Customer) LeaveTime() *time.Time {
	return c.leaveTime
}

// SetLeaveTime sets the time the customer left
func (c *Customer) SetLeaveTime(leaveTime time.Time) {
	c.leaveTime = &leaveTime
}

// WaitTime returns the time the customer waited
func (c *Customer) WaitTime() time.Duration {
	if c.leaveTime == nil {
		return 0
	}
	return c.leaveTime.Sub(c.arrivedTime)
}

// PlaceOrder places an order
// for simulation purposes, we will randomly generate an order
func (c *Customer) PlaceOrder() *Order {
	coffeeTypes := c.config.CoffeeTypes()
	randomCoffeeType := coffeeTypes[rand.Intn(len(coffeeTypes))]
	randomCoffeeSize := CoffeeSize(rand.Intn(3)) // There are 3 coffee sizes: Standard, Large, and ExtraLarge
	randomExtras := extrasOptions[rand.Intn(len(extrasOptions))]
	return NewOrder(c, CoffeeType(*randomCoffeeType), randomCoffeeSize, randomExtras)
}
