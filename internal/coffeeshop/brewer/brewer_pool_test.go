package brewer

import (
	"testing"
	"time"

	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestBrewerPoolCreation(t *testing.T) {
	brewerPool := NewBrewerPool(2)
	assert.Equal(t, 2, cap(brewerPool), "Brewer pool should have a capacity of 2")
}

func TestBrewerPoolAddBrewer(t *testing.T) {
	brewerPool := NewBrewerPool(1)
	brewer := NewBrewer("testBrewer", 5)

	brewerPool.AddBrewer(brewer)

	assert.Equal(t, 1, len(brewerPool), "Brewer pool should have 1 brewer")
}

func TestBrewerPoolStart(t *testing.T) {
	brewerPool := NewBrewerPool(2)
	brewer1 := NewBrewer("testBrewer1", 5)
	brewer2 := NewBrewer("testBrewer2", 7)

	brewerPool.AddBrewer(brewer1)
	brewerPool.AddBrewer(brewer2)

	brewerPool.Start()

	// Test grinding coffee with both grinders
	coffee1 := types.NewCoffee(types.CoffeeType{
		Name:              "TestCoffee",
		BeansToWaterRatio: utils.FloatToDecimal(0.5),
		Price:             utils.FloatToDecimal(1.0),
		SizeInOunces:      12,
	}, types.Standard, []string{})
	coffee2 := types.NewCoffee(types.CoffeeType{
		Name:              "TestCoffee",
		BeansToWaterRatio: utils.FloatToDecimal(0.5),
		Price:             utils.FloatToDecimal(1.0),
		SizeInOunces:      12,
	}, types.Standard, []string{})

	brewer1ToTest := <-brewerPool
	brewer1ToTest.Brew(coffee1)
	brewerPool <- brewer1ToTest

	brewer2ToTest := <-brewerPool
	brewer2ToTest.Brew(coffee2)
	brewerPool <- brewer2ToTest

	time.Sleep(3 * time.Second)

	isWaterReady1 := <-coffee1.WaterReady()
	assert.True(t, isWaterReady1, "The coffee1 should be brewed")

	isWaterReady2 := <-coffee2.WaterReady()
	assert.True(t, isWaterReady2, "The coffee2 should be brewed")
}
