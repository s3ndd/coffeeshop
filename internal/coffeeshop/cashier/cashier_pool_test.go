package cashier

import (
	"testing"
	"time"

	"github.com/s3ndd/coffeeshop/internal/coffeeshop/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCashierPool(t *testing.T) {
	mockOrderQueue := &mocks.MockOrderQueue{}
	mockEventSystem := &mocks.MockEventSystem{}

	// Create cashiers
	cashier1 := NewCashier(1, 10, mockOrderQueue, mockEventSystem)
	cashier2 := NewCashier(2, 10, mockOrderQueue, mockEventSystem)

	// Create a cashier pool
	cashierPool := NewCashierPool(2)

	// Test adding cashiers to the cashier pool
	cashierPool.AddCashier(cashier1)
	cashierPool.AddCashier(cashier2)
	assert.Equal(t, 2, cashierPool.Len())

	// Test starting cashiers in the cashier pool
	cashierPool.Start()
	time.Sleep(1 * time.Second) // Give some time for cashiers to start

	// Test customer queue length comparison
	assert.False(t, cashierPool.Less(0, 1))
	assert.False(t, cashierPool.Less(1, 0))

	// Test swapping cashiers in the cashier pool
	cashierPool.Swap(0, 1)
	assert.Equal(t, cashier2, cashierPool[0])
	assert.Equal(t, cashier1, cashierPool[1])

	// Test pushing and popping cashiers in the cashier pool
	cashier3 := NewCashier(3, 10, mockOrderQueue, mockEventSystem)
	cashierPool.Push(cashier3)
	assert.Equal(t, 3, cashierPool.Len())
	assert.Equal(t, cashier3, cashierPool[2])

	poppedCashier := cashierPool.Pop().(*Cashier)
	assert.Equal(t, cashier3, poppedCashier)
	assert.Equal(t, 2, cashierPool.Len())
}
