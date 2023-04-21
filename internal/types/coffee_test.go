package types

import (
	"testing"

	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBeansNeeded(t *testing.T) {
	coffeeType := CoffeeType{
		BeansToWaterRatio: utils.FloatToDecimal(0.05),
		SizeInOunces:      8,
	}
	size := Large

	expectedBeansNeeded := utils.FloatToDecimal(14.18)
	beansNeeded := calculateBeansNeeded(&coffeeType, size)

	assert.Equal(t, expectedBeansNeeded, beansNeeded, "Beans needed should match expected value")
}

func TestCalculateWaterNeeded(t *testing.T) {
	coffeeType := CoffeeType{
		SizeInOunces: 8,
	}
	size := Large

	expectedWaterNeeded := decimal.NewFromInt(10).Round(2)
	waterNeeded := calculateWaterNeeded(&coffeeType, size)

	assert.Equal(t, expectedWaterNeeded, waterNeeded, "Water needed should match expected value")
}
