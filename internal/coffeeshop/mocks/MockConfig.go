package mocks

import (
	"github.com/s3ndd/coffeeshop/internal/config"
	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/stretchr/testify/mock"
)

// MockConfig is a mock implementation of the Configurer interface
type MockConfig struct {
	mock.Mock
}

func (c *MockConfig) CoffeeTypes() []*config.CoffeeType {
	args := c.Called()
	return args.Get(0).([]*config.CoffeeType)
}

// CreateMockConfig a mock object that implements the Configurer interface
func CreateMockConfig() *MockConfig {
	mockConfig := new(MockConfig)

	// Set up a mock implementation for the CoffeeTypes method
	mockConfig.On("CoffeeTypes").Return([]*config.CoffeeType{
		{
			Name:              "Espresso",
			BeansToWaterRatio: utils.FloatToDecimal(0.05),
			Price:             utils.FloatToDecimal(2.99),
			SizeInOunces:      2,
		},
	})

	return mockConfig
}
