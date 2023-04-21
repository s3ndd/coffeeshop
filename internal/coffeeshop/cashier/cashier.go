package cashier

import (
	"time"

	"github.com/s3ndd/coffeeshop/internal/monitor"
	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
)

// Cashier represents a cashier in the coffee shop
// It has a customer queue and a shared order queue
type Cashier struct {
	id            int
	customerQueue chan *types.Customer
	orderQueue    types.OrderQueueer
	eventSystem   monitor.EventSystemer
}

// NewCashier creates a new cashier
func NewCashier(id int, maximumCustomers int, orderQueue types.OrderQueueer, eventSystem monitor.EventSystemer) *Cashier {
	cashier := &Cashier{
		id:            id,
		customerQueue: make(chan *types.Customer, maximumCustomers),
		orderQueue:    orderQueue,
		eventSystem:   eventSystem,
	}

	return cashier
}

// Start starts the cashier and listens for customers
func (c *Cashier) Start() {
	logger := utils.Logger().WithField("cashier", c.id)
	go func() {
		for customer := range c.customerQueue {
			logger.WithField("customer", customer.Name()).Info("Customer is placing order")
			order := customer.PlaceOrder()
			// add random delay to Simulate the customer placing the order
			time.Sleep(utils.RandomDelaySeconds())
			c.orderQueue.Publish(order)
			// send event to monitor
			c.eventSystem.SendEvent(monitor.Event{Type: monitor.OrderReceived, Data: order})
			logger.WithField("customer", customer.Name()).Info("Customer is done placing order")
		}
	}()
}

// ID returns the cashier's ID
func (c *Cashier) ID() int {
	return c.id
}

// CustomerQueueSize returns the current size of the customer queue
func (c *Cashier) CustomerQueueSize() int {
	return len(c.customerQueue)
}

// OrderQueueSize returns the current size of the order queue
func (c *Cashier) OrderQueueSize() int {
	return c.orderQueue.Size()
}

// ServeCustomer adds the customer to the cashier's customer queue
func (c *Cashier) ServeCustomer(customer *types.Customer) {
	c.customerQueue <- customer
}
