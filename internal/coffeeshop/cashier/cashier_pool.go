package cashier

import (
	"sync"

	"github.com/s3ndd/coffeeshop/pkg/utils"
)

// CashierPool is a pool of cashiers using a priority queue
// It implements the heap interface so that we can get the cashier with the shortest queue
type CashierPool []*Cashier

// NewCashierPool creates a new cashier pool
func NewCashierPool(size int) CashierPool {
	return make(CashierPool, 0, size)
}

// AddCashier adds a cashier to the cashier pool
func (cq *CashierPool) AddCashier(cashier *Cashier) {
	cq.Push(cashier)
}

// Start starts all cashiers in the cashier pool
func (cq *CashierPool) Start() {
	wg := &sync.WaitGroup{}
	for _, cashier := range *cq {
		wg.Add(1)
		go func(cashier *Cashier) {
			defer wg.Done()
			cashier.Start()
		}(cashier)
	}
	// Wait for all cashiers to be started
	wg.Wait()

	utils.Logger().Info("All cashiers are started")
}

// Len returns the length of the cashier pool
func (cq CashierPool) Len() int {
	return len(cq)
}

// Less returns true if the length of the customer queue of the cashier at index i is less than
// the length of the customer queue of the cashier at index j
func (cq CashierPool) Less(i, j int) bool {
	return len(cq[i].customerQueue) < len(cq[j].customerQueue)
}

// Swap swaps the cashiers at index i and j
func (cq CashierPool) Swap(i, j int) {
	cq[i], cq[j] = cq[j], cq[i]
}

// Push pushes a cashier to the cashier pool
func (cq *CashierPool) Push(x interface{}) {
	cashier := x.(*Cashier)
	*cq = append(*cq, cashier)
}

// Pop pops a cashier from the cashier pool
func (cq *CashierPool) Pop() interface{} {
	old := *cq
	n := len(old)
	x := old[n-1]
	*cq = old[0 : n-1]
	return x
}
