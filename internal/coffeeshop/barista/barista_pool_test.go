package barista

import (
	"sync"
	"testing"
	"time"

	. "github.com/s3ndd/coffeeshop/internal/coffeeshop/mocks"
	. "github.com/s3ndd/coffeeshop/internal/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBaristaPool(t *testing.T) {
	// Create a mock order queue
	mockOrderQueue := new(MockOrderQueue)
	mockEventSystem := new(MockEventSystem)

	// Create a channel to simulate order queue subscription
	orderChan := make(chan *Order)
	mockOrderQueue.On("Subscribe").Return(orderChan)

	mockEventSystem.On("SendEvent", mock.Anything)

	// Publish two orders
	order1 := NewOrder(NewCustomer("Alice", CreateMockConfig()), CoffeeType{Name: "Cappuccino", Price: decimal.NewFromFloat(4.0)}, Standard, []string{})
	order2 := NewOrder(NewCustomer("Bob", CreateMockConfig()), CoffeeType{Name: "Latte", Price: decimal.NewFromFloat(3.5)}, Large, []string{"milk"})

	// Create a wait group and mock event system
	ordersWg := &sync.WaitGroup{}

	mockBarista1 := new(MockBarista)
	mockBarista2 := new(MockBarista)

	// Set expectations for the mock baristas
	mockBarista1.On("MarkAvailable").Twice()
	mockBarista1.On("MarkBusy").Once()
	mockBarista1.On("ProcessOrder", mock.Anything).Once().Run(func(args mock.Arguments) {
		order1.Complete()
		ordersWg.Done()
	})

	mockBarista2.On("MarkAvailable").Twice()
	mockBarista2.On("MarkBusy").Once()
	mockBarista2.On("ProcessOrder", mock.Anything).Once().Run(func(args mock.Arguments) {
		order2.Complete()
		ordersWg.Done()
	})

	// Create a BaristaPool with the mock order queue and baristas
	baristaPool := NewBaristaPool(mockOrderQueue, []Baristaer{mockBarista1, mockBarista2})

	// Start the BaristaPool
	baristaPool.Start()

	ordersWg.Add(2)

	orderChan <- order1
	orderChan <- order2
	done := make(chan struct{})

	go func() {
		ordersWg.Wait()
		close(done)
	}()

	// Set a timeout for the test
	select {
	case <-done:
		// Check if both orders are processed
		assert.NotNil(t, order1.ServedTime())
		assert.NotNil(t, order2.ServedTime())
	case <-time.After(15 * time.Second):
		t.Fatal("Test timed out")
	}
}
