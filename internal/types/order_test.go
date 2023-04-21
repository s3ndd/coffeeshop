package types

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	customer := &Customer{name: "Alice"}
	coffeeType := CoffeeType{Name: "Latte", Price: decimal.NewFromFloat(3.5)}
	coffeeSize := Standard
	extras := []string{"milk"}

	order := NewOrder(customer, coffeeType, coffeeSize, extras)

	assert.Equal(t, customer, order.Customer())
	assert.Equal(t, coffeeType, order.Coffee().CoffeeType())
	assert.Equal(t, coffeeSize, order.Coffee().Size())
	assert.Equal(t, extras, order.Coffee().Extras())
	assert.Equal(t, calculatePrice(coffeeType, coffeeSize, extras), order.price)
}

func TestCompleteOrder(t *testing.T) {
	customer := &Customer{name: "Bob"}
	coffeeType := CoffeeType{Name: "Cappuccino", Price: decimal.NewFromFloat(4.0)}
	coffeeSize := Large
	extras := []string{"sugar"}

	order := NewOrder(customer, coffeeType, coffeeSize, extras)
	order.Complete()

	assert.NotNil(t, order.ServedTime())
	assert.NotNil(t, order.Customer().LeaveTime())
	assert.True(t, order.ProcessingTime() > 0)
}
