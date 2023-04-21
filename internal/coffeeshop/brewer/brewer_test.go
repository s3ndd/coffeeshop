package brewer

import (
	"testing"
	"time"

	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestBrewerCreation(t *testing.T) {
	brewer := NewBrewer("testBrewer", 5)
	assert.NotNil(t, brewer, "Brewer should not be nil")
}

func TestBrewerStart(t *testing.T) {
	brewer := NewBrewer("testBrewer", 5)
	brewer.Start()

	coffee := types.NewCoffee(types.CoffeeType{
		Name:              "TestCoffee",
		BeansToWaterRatio: utils.FloatToDecimal(0.5),
		Price:             utils.FloatToDecimal(1.0),
		SizeInOunces:      12,
	}, types.Standard, []string{})

	brewer.Brew(coffee)

	time.Sleep(2 * time.Second)

	isWaterReady := <-coffee.WaterReady()
	assert.True(t, isWaterReady, "The coffee should be brewed")
}
