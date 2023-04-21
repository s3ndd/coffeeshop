package coffeeshop

import (
	"sync"

	"github.com/s3ndd/coffeeshop/internal/coffeeshop/barista"
	brewer1 "github.com/s3ndd/coffeeshop/internal/coffeeshop/brewer"
	cashier2 "github.com/s3ndd/coffeeshop/internal/coffeeshop/cashier"
	greeter2 "github.com/s3ndd/coffeeshop/internal/coffeeshop/greeter"
	grinder2 "github.com/s3ndd/coffeeshop/internal/coffeeshop/grinder"
	"github.com/s3ndd/coffeeshop/internal/config"
	"github.com/s3ndd/coffeeshop/internal/monitor"
	"github.com/s3ndd/coffeeshop/internal/types"
)

type CoffeeShop struct {
	grinderPool grinder2.GrinderPool
	brewerPool  brewer1.BrewerPool
	greeterPool greeter2.GreeterPool
	cashierPool cashier2.CashierPool
	baristaPool *barista.BaristaPool
	orderQueue  *types.OrderQueue
	ordersWg    *sync.WaitGroup
}

func NewCoffeeShop(coffeeShop *config.CoffeeShopSettings, ordersWg *sync.WaitGroup, eventSystem monitor.EventSystemer) *CoffeeShop {
	// create an order queue
	orderQueue := types.NewOrderQueue(coffeeShop.OrderQueueSize)

	// create cashiers
	cashierPool := cashier2.NewCashierPool(coffeeShop.NumberOfCashiers)
	for i := 0; i < coffeeShop.NumberOfCashiers; i++ {
		cashier := cashier2.NewCashier(i, coffeeShop.CashierQueueSize, orderQueue, eventSystem)
		cashierPool.AddCashier(cashier)
	}

	// create greeters
	greeterPool := greeter2.NewGreeterPool(coffeeShop.NumberOfGreeters)
	for i := 0; i < coffeeShop.NumberOfGreeters; i++ {
		greeter := greeter2.NewGreeter(i, &cashierPool)
		greeterPool.AddGreeter(greeter)
	}

	// create grinder pool
	grinderPool := grinder2.NewGrinderPool(len(coffeeShop.GrinderSettings))
	for _, settings := range coffeeShop.GrinderSettings {
		grinder := grinder2.NewGrinder(settings.Tag, settings.GramsPerSecond)
		grinderPool.AddGrinder(grinder)
	}

	// create brewer pool
	brewerPool := brewer1.NewBrewerPool(len(coffeeShop.BrewerSettings))
	for _, settings := range coffeeShop.BrewerSettings {
		brewer := brewer1.NewBrewer(settings.Tag, settings.OuncesWaterPerSecond)
		brewerPool.AddBrewer(brewer)
	}

	// create barista pool
	baristas := make([]barista.Baristaer, coffeeShop.NumberOfBaristas)
	for i := 0; i < coffeeShop.NumberOfBaristas; i++ {
		barista := barista.NewBarista(i, grinderPool, brewerPool, ordersWg, eventSystem)
		baristas[i] = barista
	}

	baristaPool := barista.NewBaristaPool(orderQueue, baristas)

	return &CoffeeShop{
		grinderPool: grinderPool,
		brewerPool:  brewerPool,
		greeterPool: greeterPool,
		cashierPool: cashierPool,
		baristaPool: baristaPool,
		orderQueue:  orderQueue,
		ordersWg:    ordersWg,
	}
}

// Open opens the coffee shop
func (cs *CoffeeShop) Open() {
	cs.cashierPool.Start()
	cs.baristaPool.Start()
	cs.grinderPool.Start()
	cs.brewerPool.Start()
}

// Close closes the coffee shop
func (cs *CoffeeShop) Close() {
	// TODO: implement
}

// ServeCustomer serves a customer
func (cs *CoffeeShop) ServeCustomer(customer *types.Customer) {
	cs.ordersWg.Add(1)
	cs.greeterPool.AssignCustomer(customer)
}
