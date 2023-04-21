package barista

import (
	"github.com/s3ndd/coffeeshop/internal/types"
)

// BaristaPool represents a pool of baristas
type BaristaPool struct {
	orderQueue types.OrderQueueer
	baristas   []Baristaer
}

// NewBaristaPool creates a new barista pool
// The barista pool will subscribe to the order queue which is a shared resource
func NewBaristaPool(orderQueue types.OrderQueueer, baristas []Baristaer) *BaristaPool {
	return &BaristaPool{
		orderQueue: orderQueue,
		baristas:   baristas,
	}
}

// Start starts the barista pool
func (bp *BaristaPool) Start() {
	for _, barista := range bp.baristas {
		go func(b Baristaer) {
			b.MarkAvailable()
			for order := range bp.orderQueue.Subscribe() {
				// mark the barista as busy
				b.MarkBusy()
				b.ProcessOrder(order)
				// mark the barista as available again
				b.MarkAvailable()
			}
		}(barista)
	}
}
