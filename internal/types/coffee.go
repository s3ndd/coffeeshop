package types

import (
	"time"

	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/shopspring/decimal"
)

// CoffeeSize represents the size of a coffee
type CoffeeSize int

// Coffee sizes
const (
	// Standard represents a standard coffee
	Standard CoffeeSize = iota
	// Large represents a large coffee
	Large
	// ExtraLarge represents an extra large coffee
	ExtraLarge
)

const (
	// StandardSizeInOunces represents the standard size in ounces
	// different coffee sizes will affect the amount of water and beans needed
	sizeWaterRatio = 0.25
)

// CoffeeType represents the type of a coffee
type CoffeeType struct {
	Name              string
	BeansToWaterRatio decimal.Decimal
	Price             decimal.Decimal
	SizeInOunces      int
}

// Coffee represents a coffee
type Coffee struct {
	coffeeType  CoffeeType
	size        CoffeeSize
	extras      []string
	beansNeeded decimal.Decimal
	waterNeeded decimal.Decimal
	beansReady  chan bool
	waterReady  chan bool
	grindTime   *time.Duration
	brewTime    *time.Duration
}

// NewCoffee creates a new coffee
func NewCoffee(coffeeType CoffeeType, size CoffeeSize, extras []string) *Coffee {
	return &Coffee{
		coffeeType:  coffeeType,
		size:        size,
		extras:      extras,
		waterNeeded: calculateWaterNeeded(&coffeeType, size),
		beansNeeded: calculateBeansNeeded(&coffeeType, size),
		beansReady:  make(chan bool),
		waterReady:  make(chan bool),
	}
}

// CoffeeType returns the coffee's type
func (c *Coffee) CoffeeType() CoffeeType {
	return c.coffeeType
}

// Size returns the coffee's size
func (c *Coffee) Size() CoffeeSize {
	return c.size
}

// Extras returns the coffee's extras
func (c *Coffee) Extras() []string {
	return c.extras
}

// BeansNeeded returns the amount of beans needed for the coffee
func (c *Coffee) BeansNeeded() decimal.Decimal {
	return c.beansNeeded
}

// WaterNeeded returns the amount of water needed for the coffee
func (c *Coffee) WaterNeeded() decimal.Decimal {
	return c.waterNeeded
}

// BeansReady returns a channel that will be notified when the beans are ready
func (c *Coffee) BeansReady() chan bool {
	return c.beansReady
}

// WaterReady returns a channel that will be notified when the water is ready
func (c *Coffee) WaterReady() chan bool {
	return c.waterReady
}

// SetGrindTime sets the time it took to grind the beans
func (c *Coffee) SetGrindTime(grindTime time.Duration) {
	c.grindTime = &grindTime
}

// GrindTime returns the time it took to grind the beans
func (c *Coffee) GrindTime() time.Duration {
	return *c.grindTime
}

// SetBrewTime sets the time it took to brew the coffee
func (c *Coffee) SetBrewTime(brewTime time.Duration) {
	c.brewTime = &brewTime
}

// BrewTime returns the time it took to brew the coffee
func (c *Coffee) BrewTime() time.Duration {
	return *c.brewTime
}

// SetBeansReady sets the beans ready flag
func (c *Coffee) SetBeansReady(ready bool) {
	c.beansReady <- ready
}

// SetWaterReady sets the water ready flag
func (c *Coffee) SetWaterReady(ready bool) {
	c.waterReady <- ready
}

// CalculateBeansNeeded calculates the amount of beans needed for a coffee
// The unit of measure is grams
func calculateBeansNeeded(coffeeType *CoffeeType, size CoffeeSize) decimal.Decimal {
	return coffeeType.BeansToWaterRatio.Mul(utils.OuncesToGrams(calculateWaterNeeded(coffeeType, size))).Round(2)
}

// calculateWaterNeeded calculates the amount of water needed for a coffee based on the coffee type and size
// The unit of measure is ounces
func calculateWaterNeeded(coffeeType *CoffeeType, size CoffeeSize) decimal.Decimal {
	return utils.FloatToDecimal(1 + sizeWaterRatio*float64(size)).Mul(decimal.NewFromInt(int64(coffeeType.SizeInOunces))).Round(2)
}
