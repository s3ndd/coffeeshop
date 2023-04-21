package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderQueueCreation(t *testing.T) {
	orderQueue := NewOrderQueue(3)
	assert.NotNil(t, orderQueue, "Order queue should not be nil")
	assert.Equal(t, 0, orderQueue.Size(), "Order queue should be empty")
}
