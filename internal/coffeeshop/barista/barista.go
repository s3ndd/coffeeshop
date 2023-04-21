package barista

import (
	"sync"

	"github.com/s3ndd/coffeeshop/internal/coffeeshop/brewer"
	"github.com/s3ndd/coffeeshop/internal/coffeeshop/grinder"
	"github.com/s3ndd/coffeeshop/internal/monitor"
	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
)

type Baristaer interface {
	MarkAvailable()
	MarkBusy()
	ProcessOrder(order *types.Order)
}

// Barista is a worker that processes orders
type Barista struct {
	ID          int
	grinderPool chan *grinder.Grinder
	brewerPool  chan *brewer.Brewer
	available   chan struct{}
	ordersWg    *sync.WaitGroup
	eventSystem monitor.EventSystemer
}

// NewBarista creates a new barista
// the grinderPool and brewerPool are used to get available grinders and brewers
// the ordersWg is used to wait for all orders to be processed
func NewBarista(id int, grinderPool chan *grinder.Grinder, brewerPool chan *brewer.Brewer, ordersWg *sync.WaitGroup, eventSystem monitor.EventSystemer) *Barista {
	return &Barista{
		ID:          id,
		grinderPool: grinderPool,
		brewerPool:  brewerPool,
		available:   make(chan struct{}, 1),
		ordersWg:    ordersWg,
		eventSystem: eventSystem,
	}
}

// MarkAvailable marks the barista as available
func (b *Barista) MarkAvailable() {
	b.available <- struct{}{}
}

// MarkBusy marks the barista as busy
func (b *Barista) MarkBusy() {
	<-b.available
}

// ProcessOrder processes an order
// It gets an available grinder and brewer from the pool, processes the order, and returns the grinder and brewer to the pool
// The processing workflow is
// 1. Get an available grinder from the pool
// 2. Grind coffee
// 3. Return the grinder to the pool
// 4. Get an available brewer from the pool
// 5. Brew coffee
// 6. Return the brewer to the pool
// 7. Complete order and notify the customer
func (b *Barista) ProcessOrder(order *types.Order) {
	logger := utils.Logger().WithFields(utils.LogFields{
		"barista":  b.ID,
		"customer": order.Customer().Name(),
	})

	// send event to monitor
	b.eventSystem.SendEvent(monitor.Event{Type: monitor.OrderProcessed, Data: order})

	logger.Info("Barista is processing order")
	// Get an available grinder from the pool
	grinder := <-b.grinderPool
	// Grind coffee
	grinder.Grind(order.Coffee())
	<-order.Coffee().BeansReady()
	// Return the grinder to the pool
	b.grinderPool <- grinder

	// Get an available brewer from the pool
	brewer := <-b.brewerPool
	// Brew coffee
	brewer.Brew(order.Coffee())
	<-order.Coffee().WaterReady()
	// Return the brewer to the pool
	b.brewerPool <- brewer

	// Complete order and notify the customer
	order.Complete()
	b.ordersWg.Done()
	// send event to monitor
	b.eventSystem.SendEvent(monitor.Event{Type: monitor.OrderCompleted, Data: order})
	logger.Info("Barista is done processing order")
}
