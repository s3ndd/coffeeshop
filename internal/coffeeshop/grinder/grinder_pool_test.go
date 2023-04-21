package grinder

import (
	"testing"
	"time"

	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGrinderPoolCreation(t *testing.T) {
	grinderPool := NewGrinderPool(2)
	assert.Equal(t, 2, cap(grinderPool), "Grinder pool should have a capacity of 2")
}

func TestGrinderPoolAddGrinder(t *testing.T) {
	grinderPool := NewGrinderPool(1)
	grinder := NewGrinder("testGrinder", 10)

	grinderPool.AddGrinder(grinder)

	assert.Equal(t, 1, len(grinderPool), "Grinder pool should have 1 grinder")
}

func TestGrinderPoolStart(t *testing.T) {
	grinderPool := NewGrinderPool(2)
	grinder1 := NewGrinder("testGrinder1", 10)
	grinder2 := NewGrinder("testGrinder2", 12)

	grinderPool.AddGrinder(grinder1)
	grinderPool.AddGrinder(grinder2)

	grinderPool.Start()

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

	grinder1ToTest := <-grinderPool
	grinder1ToTest.Grind(coffee1)
	grinderPool <- grinder1ToTest

	grinder2ToTest := <-grinderPool
	grinder2ToTest.Grind(coffee2)
	grinderPool <- grinder2ToTest

	// Give some time for grinding
	time.Sleep(2 * time.Second)

	isBeansReady1 := <-coffee1.BeansReady()
	assert.True(t, isBeansReady1, "The coffee beans for coffee1 should be ground")

	isBeansReady2 := <-coffee2.BeansReady()
	assert.True(t, isBeansReady2, "The coffee beans for coffee2 should be ground")
}
