package mocks

import (
	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/stretchr/testify/mock"
)

type MockBarista struct {
	mock.Mock
}

func (m *MockBarista) MarkAvailable() {
	m.Called()
}

func (m *MockBarista) MarkBusy() {
	m.Called()
}

func (m *MockBarista) ProcessOrder(order *types.Order) {
	m.Called(order)
}
