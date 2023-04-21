package brewer

import (
	"time"

	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/shopspring/decimal"
)

// Brewer represents a coffee brewer
// It has a brewing channel where it receives the coffee to brew
// The brewing channel is unbuffered, so the brewer will block until the coffee is brewed
// tag is the name of the brewer
// ouncesWaterPerSecond is the number of ounces of water that can be brewed per second
type Brewer struct {
	tag                  string
	ouncesWaterPerSecond int
	brewingChannel       chan *types.Coffee
}

// NewBrewer creates a new coffee brewer
func NewBrewer(tag string, ouncesWaterPerSecond int) *Brewer {
	return &Brewer{
		tag:                  tag,
		ouncesWaterPerSecond: ouncesWaterPerSecond,
		brewingChannel:       make(chan *types.Coffee),
	}
}

// Start starts the brewer
func (b *Brewer) Start() {
	logger := utils.Logger().WithField("brewer", b.tag)
	go func() {
		for coffee := range b.brewingChannel {
			logger = logger.WithFields(utils.LogFields{
				"coffee": coffee.CoffeeType().Name,
				"size":   coffee.Size(),
				"water":  coffee.WaterNeeded(),
			})
			logger.Info("Brewing coffee")
			brewingTime := time.Duration(coffee.WaterNeeded().Div(decimal.NewFromInt(int64(b.ouncesWaterPerSecond))).IntPart()) * time.Second
			time.Sleep(brewingTime)
			coffee.SetBrewTime(brewingTime)
			coffee.SetWaterReady(true)
			logger.Info("Coffee is brewed")
		}
	}()
}

// Brew adds the coffee to the brewer's brewing channel
func (b *Brewer) Brew(coffee *types.Coffee) {
	b.brewingChannel <- coffee
}
