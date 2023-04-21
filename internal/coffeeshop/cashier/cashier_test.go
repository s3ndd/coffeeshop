package cashier

import (
	"testing"
	"time"

	"github.com/s3ndd/coffeeshop/internal/coffeeshop/mocks"
	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCashier(t *testing.T) {
	mockOrderQueue := new(mocks.MockOrderQueue)
	mockEventSystem := new(mocks.MockEventSystem)

	cashier := NewCashier(1, 5, mockOrderQueue, mockEventSystem)
	assert.Equal(t, 1, cashier.ID(), "Cashier ID should be 1")

	cashier.ServeCustomer(types.NewCustomer("Shelly Shi", mocks.CreateMockConfig()))
	assert.Equal(t, 1, cashier.CustomerQueueSize(), "Cashier should have 1 customer in the queue")

	mockOrderQueue.On("Publish", mock.Anything)
	mockEventSystem.On("SendEvent", mock.Anything)

	cashier.Start()
	time.Sleep(5 * time.Second) // Give some time for the cashier to process the customer's order

	mockOrderQueue.AssertExpectations(t)
	mockEventSystem.AssertExpectations(t)
}
