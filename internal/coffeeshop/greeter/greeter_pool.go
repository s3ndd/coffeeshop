package greeter

import (
	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
)

// GreeterPool is a pool of greeters
type GreeterPool chan *Greeter

// NewGreeterPool creates a new greeter pool
func NewGreeterPool(numOfGreeters int) GreeterPool {
	return make(GreeterPool, numOfGreeters)
}

// AddGreeter adds a greeter to the pool
func (gp GreeterPool) AddGreeter(greeter *Greeter) {
	gp <- greeter
}

// AssignCustomer assigns a customer to a greeter, and logs the assignment
func (gp GreeterPool) AssignCustomer(customer *types.Customer) {
	// get an available greeter from the pool
	greeter := <-gp

	logger := utils.Logger().WithFields(utils.LogFields{
		"greeter":  greeter.ID(),
		"customer": customer.Name(),
	})

	logger.Info("Greeter is greeting customer")

	greeter.Greet(customer)

	// return the greeter to the pool
	gp <- greeter

	logger.Info("Greeter is done greeting customer")
}
