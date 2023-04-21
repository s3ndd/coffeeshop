package mocks

import (
	. "github.com/s3ndd/coffeeshop/internal/types"
	"github.com/stretchr/testify/mock"
)

type MockOrderQueue struct {
	mock.Mock
}

func (m *MockOrderQueue) Publish(order *Order) {
	m.Called(order)
}

func (m *MockOrderQueue) Subscribe() <-chan *Order {
	args := m.Called()
	return args.Get(0).(chan *Order)
}

func (m *MockOrderQueue) Size() int {
	args := m.Called()
	return args.Int(0)
}
