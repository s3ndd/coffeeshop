package monitor

import (
	"sync"

	"github.com/s3ndd/coffeeshop/internal/types"
)

type EventSystemer interface {
	SendEvent(event Event)
}

// EventSystem is the event system that will be used to monitor the coffee shop
type EventSystem struct {
	eventChannel chan Event
	metrics      *Metrics
	wg           sync.WaitGroup
}

// NewEventSystem creates a new EventSystem
func NewEventSystem() *EventSystem {
	return &EventSystem{
		eventChannel: make(chan Event),
		metrics:      NewMetrics(),
		wg:           sync.WaitGroup{},
	}
}

// SendEvent sends an event to the event system
func (es *EventSystem) SendEvent(event Event) {
	es.wg.Add(1)
	es.eventChannel <- event
}

// StartEventListener starts the event listener
// It will listen to the event channel and update the metrics based on the event type
func (es *EventSystem) StartEventListener() {
	for event := range es.eventChannel {
		switch event.Type {
		case OrderReceived:
			es.metrics.IncrementReceivedOrders()
		case OrderProcessed:
			es.metrics.IncrementProcessedOrders()
		case OrderCompleted:
			es.metrics.IncrementCompletedOrders()
			order := event.Data.(*types.Order)
			es.metrics.AddGrindTime(order.Coffee().GrindTime())
			es.metrics.AddBrewTime(order.Coffee().BrewTime())
			es.metrics.AddWaitTime(order.Customer().WaitTime())
			es.metrics.AddProcessTime(event.Data.(*types.Order).ProcessingTime())
		}
		es.wg.Done()
	}
}

// PrintMetricsSummary prints the metrics summary
func (es *EventSystem) PrintMetricsSummary() {
	es.metrics.PrintSummary()
}

// Stop stops the event system
// It closes the event channel and waits for all events to be processed
func (es *EventSystem) Stop() {
	close(es.eventChannel)
	es.wg.Wait()
}
