package grinder

import (
	"sync"

	"github.com/s3ndd/coffeeshop/pkg/utils"
)

// GrinderPool is a pool of grinders
type GrinderPool chan *Grinder

// NewGrinderPool creates a new grinder pool
func NewGrinderPool(size int) GrinderPool {
	return make(GrinderPool, size)
}

// AddGrinder adds a grinder to the pool
func (gp GrinderPool) AddGrinder(grinder *Grinder) {
	gp <- grinder
}

// Start starts all grinders in the pool
// It waits for all grinders to be started
func (gp GrinderPool) Start() {
	wg := &sync.WaitGroup{}
	for i := 0; i < len(gp); i++ {
		grinder := <-gp
		wg.Add(1)
		go func(g *Grinder) {
			defer wg.Done()
			g.Start()
		}(grinder)
		gp <- grinder
	}
	// Wait for all grinders to be started
	wg.Wait()

	utils.Logger().Info("All grinders are started")
}
