package mocks

import (
	. "github.com/s3ndd/coffeeshop/internal/coffeeshop/brewer"
	"github.com/stretchr/testify/mock"
)

type MockBrewerPool struct {
	mock.Mock
}

func (m *MockBrewerPool) BrewerPool() chan *Brewer {
	args := m.Called()
	return args.Get(0).(chan *Brewer)
}
