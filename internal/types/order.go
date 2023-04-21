package types

import (
	"time"

	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/shopspring/decimal"
)

// Order represents an order
type Order struct {
	customer   *Customer
	coffee     *Coffee
	orderTime  time.Time
	servedTime *time.Time
	price      decimal.Decimal
}

// NewOrder creates a new order
func NewOrder(customer *Customer, coffeeType CoffeeType, coffeeSize CoffeeSize, extras []string) *Order {
	return &Order{
		customer:  customer,
		orderTime: time.Now(),
		coffee:    NewCoffee(coffeeType, coffeeSize, extras),
		price:     calculatePrice(coffeeType, coffeeSize, extras),
	}
}

// Customer returns the order's customer
func (o *Order) Customer() *Customer {
	return o.customer
}

// Coffee returns the order's coffee
func (o *Order) Coffee() *Coffee {
	return o.coffee
}

// Complete completes the order
func (o *Order) Complete() {
	now := time.Now()
	o.servedTime = &now
	o.Customer().SetLeaveTime(now)
	utils.Logger().WithField("order", o.customer.Name()).Info("Order completed")
}

func (o *Order) OrderTime() time.Time {
	return o.orderTime
}

func (o *Order) ServedTime() *time.Time {
	return o.servedTime
}

func (o *Order) ProcessingTime() time.Duration {
	if o.servedTime == nil {
		return 0
	}
	return o.servedTime.Sub(o.orderTime)
}

// calculatePrice calculates the price of an order
func calculatePrice(coffeeType CoffeeType, size CoffeeSize, extras []string) decimal.Decimal {
	// TODO: 0.5 and 0.25 could be configurable
	sizePrice := decimal.NewFromInt(int64(size)).Mul(utils.FloatToDecimal(0.5))
	extrasPrice := decimal.NewFromInt(int64(len(extras))).Mul(utils.FloatToDecimal(0.25))

	return coffeeType.Price.Add(sizePrice).Add(extrasPrice)
}
