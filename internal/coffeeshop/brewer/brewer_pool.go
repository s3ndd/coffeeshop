package brewer

import (
	"sync"

	"github.com/s3ndd/coffeeshop/pkg/utils"
)

// BrewerPool is a pool of brewers
type BrewerPool chan *Brewer

// NewBrewerPool creates a new brewer pool
func NewBrewerPool(size int) BrewerPool {
	return make(BrewerPool, size)
}

// AddBrewer adds a brewer to the pool
func (bp BrewerPool) AddBrewer(brewer *Brewer) {
	bp <- brewer
}

// Start starts all brewers in the pool
func (bp BrewerPool) Start() {
	wg := &sync.WaitGroup{}
	for i := 0; i < len(bp); i++ {
		brewer := <-bp
		wg.Add(1)
		go func(b *Brewer) {
			defer wg.Done()
			b.Start()
		}(brewer)
		bp <- brewer
	}

	// Wait for all brewers to be started
	wg.Wait()

	utils.Logger().Info("All brewers are started")
}
