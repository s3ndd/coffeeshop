package types

type OrderQueueer interface {
	Publish(order *Order)
	Subscribe() <-chan *Order
	Size() int
}

// OrderQueue is a queue of orders
type OrderQueue chan *Order

// NewOrderQueue creates a new order queue
func NewOrderQueue(orderQueueSize int) *OrderQueue {
	orderQueue := make(OrderQueue, orderQueueSize)
	return &orderQueue
}

// Publish publishes an order to the order queue
func (oq *OrderQueue) Publish(order *Order) {
	*oq <- order
}

// Subscribe subscribes to the order queue
func (oq *OrderQueue) Subscribe() <-chan *Order {
	return *oq
}

// Size returns the current size of the order queue
func (oq *OrderQueue) Size() int {
	return len(*oq)
}
