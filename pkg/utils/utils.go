package utils

import (
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

func RandomDelaySeconds() time.Duration {
	return time.Duration(rand.Intn(5)) * time.Second
}

// OuncesToGrams converts ounces to grams
func OuncesToGrams(ounces decimal.Decimal) decimal.Decimal {
	return ounces.Mul(FloatToDecimal(28.3495))
}

// FloatToDecimal converts a float to a decimal
func FloatToDecimal(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f).Round(2)
}

// IntToDecimal converts an int to a decimal
func IntToDecimal(i int) decimal.Decimal {
	return decimal.NewFromInt(int64(i)).Round(2)
}
