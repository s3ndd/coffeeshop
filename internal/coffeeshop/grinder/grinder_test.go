package grinder

import (
	"testing"
	"time"

	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGrinderStart(t *testing.T) {
	grinder := NewGrinder("testGrinder", 100)
	grinder.Start()

	coffee := types.NewCoffee(types.CoffeeType{
		Name:              "TestCoffee",
		BeansToWaterRatio: utils.FloatToDecimal(0.5),
		Price:             utils.FloatToDecimal(1.0),
		SizeInOunces:      12,
	}, types.Standard, []string{})

	grinder.Grind(coffee)

	// Give some time for grinding
	time.Sleep(1 * time.Second)

	isBeansReady := <-coffee.BeansReady()
	assert.True(t, isBeansReady, "The coffee beans should be ground")
}

func TestGrinderGrindingProcess(t *testing.T) {
	// Create a grinder
	grinder := NewGrinder("TestGrinder", 100)

	// Start the grinder
	grinder.Start()

	coffee := types.NewCoffee(types.CoffeeType{
		Name:              "TestCoffee",
		BeansToWaterRatio: utils.FloatToDecimal(0.5),
		Price:             utils.FloatToDecimal(1.0),
		SizeInOunces:      12,
	}, types.Standard, []string{})

	// Grind the coffee
	grinder.Grind(coffee)

	// Wait for the grinding process to complete

	time.Sleep(time.Duration(int(coffee.BeansNeeded().Round(0).IntPart())/grinder.gramsPerSecond) * time.Second)

	// Check if the coffee beans are ground
	isBeansReady := <-coffee.BeansReady()
	assert.True(t, isBeansReady, "The coffee beans should be ground")

	// Check if the grind time is correct
	expectedGrindTime := time.Duration(int(coffee.BeansNeeded().Round(0).IntPart())/grinder.gramsPerSecond) * time.Second
	assert.Equal(t, expectedGrindTime, coffee.GrindTime(), "The grind time should match the expected grind time")
}

func TestGrindingTime(t *testing.T) {
	grinder := NewGrinder("testGrinder", 100)
	grinder.Start()

	coffee := types.NewCoffee(types.CoffeeType{
		Name:              "TestCoffee",
		BeansToWaterRatio: utils.FloatToDecimal(0.5),
		Price:             utils.FloatToDecimal(1.0),
		SizeInOunces:      12,
	}, types.Standard, []string{})

	grinder.Grind(coffee)

	// Give some time for grinding
	time.Sleep(1 * time.Second)

	isBeansReady := <-coffee.BeansReady()
	assert.True(t, isBeansReady, "The coffee beans should be ground")

	expectedGrindingTime := time.Duration(int(coffee.BeansNeeded().Round(1).IntPart())/grinder.gramsPerSecond) * time.Second
	assert.Equal(t, expectedGrindingTime, coffee.GrindTime(), "The actual grinding time should match the expected grinding time")
}

func TestGrindingMultipleCoffees(t *testing.T) {
	grinder := NewGrinder("testGrinder", 100)
	grinder.Start()

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

	grinder.Grind(coffee1)
	isBeansReady1 := <-coffee1.BeansReady()
	assert.True(t, isBeansReady1, "The coffee beans for coffee1 should be ground")

	grinder.Grind(coffee2)

	// Give some time for grinding
	time.Sleep(2 * time.Second)

	isBeansReady2 := <-coffee2.BeansReady()
	assert.True(t, isBeansReady2, "The coffee beans for coffee2 should be ground")
}
