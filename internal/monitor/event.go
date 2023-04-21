package monitor

// EventType is the type of event
type EventType int

const (
	// OrderReceived is the event type for when an order is received
	OrderReceived EventType = iota
	// OrderProcessed is the event type for when an order is processed
	OrderProcessed
	// OrderCompleted is the event type for when an order is completed
	OrderCompleted
)

// Event is the event struct
type Event struct {
	Type EventType
	Data interface{}
}
