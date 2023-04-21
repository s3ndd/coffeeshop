package monitor

import (
	"sync"
	"time"

	"github.com/s3ndd/coffeeshop/pkg/utils"
)

// Metrics represents the metrics of the coffee shop
type Metrics struct {
	receivedOrders   int
	processedOrders  int
	completedOrders  int
	totalProcessTime time.Duration
	totalGrindTime   time.Duration
	totalBrewTime    time.Duration
	totalWaitTime    time.Duration
	metricsMutex     sync.Mutex
}

// NewMetrics creates a new metrics object
func NewMetrics() *Metrics {
	return &Metrics{
		receivedOrders:   0,
		processedOrders:  0,
		completedOrders:  0,
		totalProcessTime: 0,
		totalGrindTime:   0,
		totalBrewTime:    0,
		totalWaitTime:    0,
		metricsMutex:     sync.Mutex{},
	}
}

// IncrementReceivedOrders increments the number of received orders
func (m *Metrics) IncrementReceivedOrders() {
	m.metricsMutex.Lock()
	m.receivedOrders++
	m.metricsMutex.Unlock()
}

// IncrementProcessedOrders increments the number of processed orders
func (m *Metrics) IncrementProcessedOrders() {
	m.metricsMutex.Lock()
	m.processedOrders++
	m.metricsMutex.Unlock()
}

// IncrementCompletedOrders increments the number of completed orders
func (m *Metrics) IncrementCompletedOrders() {
	m.metricsMutex.Lock()
	m.completedOrders++
	m.metricsMutex.Unlock()
}

// AddProcessTime adds the given duration to the total process time
func (m *Metrics) AddProcessTime(duration time.Duration) {
	m.metricsMutex.Lock()
	m.totalProcessTime += duration
	m.metricsMutex.Unlock()
}

// AddGrindTime adds the given duration to the total grind time
func (m *Metrics) AddGrindTime(duration time.Duration) {
	m.metricsMutex.Lock()
	m.totalGrindTime += duration
	m.metricsMutex.Unlock()
}

// AddBrewTime adds the given duration to the total brew time
func (m *Metrics) AddBrewTime(duration time.Duration) {
	m.metricsMutex.Lock()
	m.totalBrewTime += duration
	m.metricsMutex.Unlock()
}

// AddWaitTime adds the given duration to the total wait time
func (m *Metrics) AddWaitTime(duration time.Duration) {
	m.metricsMutex.Lock()
	m.totalWaitTime += duration
	m.metricsMutex.Unlock()
}

// PrintSummary prints the metrics summary
// This function is not thread safe, it should be called only after all the events are processed
// and the event listener is stopped
func (m *Metrics) PrintSummary() {
	m.metricsMutex.Lock()
	defer m.metricsMutex.Unlock()

	logger := utils.Logger().WithFields(utils.LogFields{
		"received_orders":  m.receivedOrders,
		"processed_orders": m.processedOrders,
		"completed_orders": m.completedOrders,
	})
	if m.completedOrders > 0 {
		// Calculate average times for each order
		// avoid division by zero
		logger = logger.WithFields(utils.LogFields{
			"average_grinding_time": (m.totalGrindTime / time.Duration(m.completedOrders)).Seconds(),
			"average_brewing_time":  (m.totalBrewTime / time.Duration(m.completedOrders)).Seconds(),
			"average_waiting_time":  (m.totalWaitTime / time.Duration(m.completedOrders)).Seconds(),
			"average_process_time":  (m.totalProcessTime / time.Duration(m.completedOrders)).Seconds(),
		})
	}

	logger.Info("Metrics summary")
}
