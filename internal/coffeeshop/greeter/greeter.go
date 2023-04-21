package greeter

import (
	"container/heap"

	cashier2 "github.com/s3ndd/coffeeshop/internal/coffeeshop/cashier"
	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
)

// Greeter represents a greeter
type Greeter struct {
	id          int
	cashierPool *cashier2.CashierPool
}

// NewGreeter creates a new greeter
// cashierPool is a shared resource between all greeters
func NewGreeter(id int, cashierPool *cashier2.CashierPool) *Greeter {
	return &Greeter{
		id:          id,
		cashierPool: cashierPool,
	}
}

// Greet assigns the customer to the cashier with the shortest queue and logs the assignment
func (g *Greeter) Greet(customer *types.Customer) {
	// Assign the customer to the cashier with the shortest queue
	cashier := heap.Pop(g.cashierPool).(*cashier2.Cashier)
	cashier.ServeCustomer(customer)
	// Return the cashier to the pool
	heap.Push(g.cashierPool, cashier)

	utils.Logger().WithFields(utils.LogFields{
		"greeter":   g.id,
		"customer":  customer.Name(),
		"cashier":   cashier.ID(),
		"queueSize": cashier.CustomerQueueSize(),
	}).Info("Greeter assigned customer to cashier")
}

// ID returns the greeter's ID
func (g *Greeter) ID() int {
	return g.id
}
