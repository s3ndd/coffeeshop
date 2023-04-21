package barista

import (
	"sync"
	"testing"

	"github.com/s3ndd/coffeeshop/internal/coffeeshop/brewer"
	"github.com/s3ndd/coffeeshop/internal/coffeeshop/grinder"
	"github.com/s3ndd/coffeeshop/internal/coffeeshop/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBaristaProcessOrder(t *testing.T) {
	// create mock objects for the grinder and brewer pools
	grinderPool := make(chan *grinder.Grinder, 1)
	brewerPool := make(chan *brewer.Brewer, 1)

	mockGrinderPool := new(mocks.MockGrinderPool)
	mockGrinderPool.On("GrinderPool").Return(grinderPool)

	mockBrewerPool := new(mocks.MockBrewerPool)
	mockBrewerPool.On("BrewerPool").Return(brewerPool)

	// create a mock event system
	eventSystem := mocks.NewMockEventSystem()

	// create a barista with the mock objects
	barista := NewBarista(1, mockGrinderPool.GrinderPool(), mockBrewerPool.BrewerPool(), &sync.WaitGroup{}, eventSystem)

	// test MarkAvailable and MarkBusy
	barista.MarkAvailable()
	assert.Equal(t, len(barista.available), 1)
	barista.MarkBusy()
	assert.Equal(t, len(barista.available), 0)
	barista.MarkAvailable()
	assert.Equal(t, len(barista.available), 1)
}
