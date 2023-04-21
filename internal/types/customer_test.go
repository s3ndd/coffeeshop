package types

import (
	"testing"
	"time"

	"github.com/s3ndd/coffeeshop/internal/config"
	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/stretchr/testify/assert"
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

// Create a mock object that implements the Configurer interface
func createMockConfig() *MockConfig {
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

func TestNewCustomer(t *testing.T) {
	customer := NewCustomer("Shelly Shi", createMockConfig())
	assert.NotNil(t, customer, "Customer should not be nil")
	assert.Equal(t, "Shelly Shi", customer.Name(), "Customer name should be 'Shelly Shi'")
	assert.True(t, time.Since(customer.ArrivedTime()) < time.Second, "Customer arrived time should be set")
	assert.Nil(t, customer.LeaveTime(), "Customer leave time should be nil")
}

func TestCustomerSetLeaveTime(t *testing.T) {
	customer := NewCustomer("Shelly Shi", createMockConfig())
	leaveTime := time.Now()
	customer.SetLeaveTime(leaveTime)

	assert.NotNil(t, customer.LeaveTime(), "Customer leave time should not be nil")
	assert.Equal(t, leaveTime, *customer.LeaveTime(), "Customer leave time should be set to the given value")
}

func TestCustomerWaitTime(t *testing.T) {
	customer := NewCustomer("Shelly Shi", createMockConfig())
	leaveTime := time.Now().Add(5 * time.Minute)
	customer.SetLeaveTime(leaveTime)

	waitTime := customer.WaitTime()

	assert.True(t, waitTime >= 5*time.Minute && waitTime < 6*time.Minute, "Customer wait time should be approximately 5 minutes")
}

func TestCustomerPlaceOrder(t *testing.T) {
	// Create a mock configuration object
	mockConfig := createMockConfig()

	// Create a new customer with the mock configuration object
	customer := NewCustomer("Shelly", mockConfig)

	// Call the PlaceOrder method and check the result
	order := customer.PlaceOrder()
	assert.NotNil(t, order, "Order should not be nil")
	assert.Equal(t, customer, order.Customer(), "Order customer should be the customer that placed the order")
	assert.Equal(t, CoffeeType(*mockConfig.CoffeeTypes()[0]), order.Coffee().CoffeeType(), "Order coffee type should be the first coffee type in the configuration")
}
