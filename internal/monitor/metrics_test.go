package monitor

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	metrics := NewMetrics()

	// Test IncrementReceivedOrders
	metrics.IncrementReceivedOrders()
	assert.Equal(t, 1, metrics.receivedOrders)

	// Test IncrementProcessedOrders
	metrics.IncrementProcessedOrders()
	assert.Equal(t, 1, metrics.processedOrders)

	// Test IncrementCompletedOrders
	metrics.IncrementCompletedOrders()
	assert.Equal(t, 1, metrics.completedOrders)

	// Test AddProcessTime
	metrics.AddProcessTime(time.Second)
	assert.Equal(t, time.Second, metrics.totalProcessTime)

	// Test AddGrindTime
	metrics.AddGrindTime(time.Second)
	assert.Equal(t, time.Second, metrics.totalGrindTime)

	// Test AddBrewTime
	metrics.AddBrewTime(time.Second)
	assert.Equal(t, time.Second, metrics.totalBrewTime)

	// Test AddWaitTime
	metrics.AddWaitTime(time.Second)
	assert.Equal(t, time.Second, metrics.totalWaitTime)

	// Test PrintSummary
	metrics.PrintSummary() // Just test that it does not panic

	// Test concurrency with multiple goroutines
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			metrics.IncrementReceivedOrders()
			metrics.IncrementProcessedOrders()
			metrics.IncrementCompletedOrders()
			metrics.AddProcessTime(time.Second)
			metrics.AddGrindTime(time.Second)
			metrics.AddBrewTime(time.Second)
			metrics.AddWaitTime(time.Second)
		}()
	}
	wg.Wait()

	assert.Equal(t, 101, metrics.receivedOrders)
	assert.Equal(t, 101, metrics.processedOrders)
	assert.Equal(t, 101, metrics.completedOrders)
	assert.Equal(t, 101*time.Second, metrics.totalProcessTime)
	assert.Equal(t, 101*time.Second, metrics.totalGrindTime)
	assert.Equal(t, 101*time.Second, metrics.totalBrewTime)
	assert.Equal(t, 101*time.Second, metrics.totalWaitTime)

	// Test PrintSummary with completed orders
	metrics.PrintSummary() // Just test that it does not panic

	// Test PrintSummary without completed orders
	metrics2 := NewMetrics()
	metrics2.IncrementReceivedOrders()
	metrics2.PrintSummary() // Just test that it does not panic
}
