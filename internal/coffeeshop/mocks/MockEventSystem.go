package mocks

import (
	"github.com/s3ndd/coffeeshop/internal/monitor"
	"github.com/stretchr/testify/mock"
)

type MockEventSystem struct {
	mock.Mock
}

func NewMockEventSystem() *MockEventSystem {
	return &MockEventSystem{}
}

func (m *MockEventSystem) SendEvent(event monitor.Event) {
	m.Called(event)
}

func (m *MockEventSystem) StartEventListener() {
	m.Called()
}
