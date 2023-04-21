package mocks

import (
	. "github.com/s3ndd/coffeeshop/internal/coffeeshop/grinder"
	"github.com/stretchr/testify/mock"
)

type MockGrinderPool struct {
	mock.Mock
}

func (m *MockGrinderPool) GrinderPool() chan *Grinder {
	args := m.Called()
	return args.Get(0).(chan *Grinder)
}
