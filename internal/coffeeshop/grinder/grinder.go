package grinder

import (
	"time"

	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/shopspring/decimal"
)

// Grinder represents a coffee grinder
// tag is the name of the grinder
// gramsPerSecond is the number of grams that can be ground per second
// grindingChannel is the channel that receives the coffee to be ground
// the grinderChannel is a unbuffered channel, so the grinder will block until the coffee is ground
type Grinder struct {
	tag             string
	gramsPerSecond  int
	grindingChannel chan *types.Coffee
}

// NewGrinder creates a new coffee grinder
func NewGrinder(tag string, gramsPerSecond int) *Grinder {
	return &Grinder{
		tag:             tag,
		gramsPerSecond:  gramsPerSecond,
		grindingChannel: make(chan *types.Coffee),
	}
}

// Start starts the grinder
func (g *Grinder) Start() {
	logger := utils.Logger().WithField("grinder", g.tag)
	go func() {
		for coffee := range g.grindingChannel {
			logger = logger.WithFields(utils.LogFields{
				"coffee": coffee.CoffeeType().Name,
				"size":   coffee.Size(),
				"beans":  coffee.BeansNeeded(),
			})
			logger.Info("Grinding coffee beans")

			grindingTime := time.Duration(coffee.BeansNeeded().Div(decimal.NewFromInt(int64(g.gramsPerSecond))).IntPart()) * time.Second
			// Simulate the grinding process
			time.Sleep(grindingTime)
			coffee.SetGrindTime(grindingTime)
			coffee.SetBeansReady(true)
			logger.Info("Coffee beans are ground")
		}
	}()
}

// Grind adds the coffee to the grinder's grinding channel
func (g *Grinder) Grind(coffee *types.Coffee) {
	g.grindingChannel <- coffee
}
