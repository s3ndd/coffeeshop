package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/s3ndd/coffeeshop/internal/coffeeshop"
	"github.com/s3ndd/coffeeshop/internal/config"
	"github.com/s3ndd/coffeeshop/internal/monitor"
	"github.com/s3ndd/coffeeshop/internal/types"
	"github.com/s3ndd/coffeeshop/pkg/utils"
)

// main is the entry point of the application
// It creates a new coffee shop and serves 100 customers
// It then prints the metrics summary and closes the coffee shop
func main() {
	logger := utils.Logger()
	logger.Info("Starting coffee shop")

	// Load the config file
	cfg := config.LoadConfig()
	// Log the config file
	logger.WithField("config", cfg).Info("Config file read successfully")

	// ordersWg is used to wait for all orders to be completed
	ordersWg := &sync.WaitGroup{}

	// Create a new event system and start the event listener
	eventSystem := monitor.NewEventSystem()
	go eventSystem.StartEventListener()

	// Create a new coffee shop
	coffeeShop := coffeeshop.NewCoffeeShop(cfg.CoffeeShop(), ordersWg, eventSystem)
	// Open the coffee shop
	coffeeShop.Open()

	// For simulation purposes, we will serve 20 customers
	for i := 0; i < 20; i++ {
		customer := types.NewCustomer(strconv.Itoa(i), cfg)
		coffeeShop.ServeCustomer(customer)
		// simulate a random delay
		time.Sleep(utils.RandomDelaySeconds())
	}

	// Wait for all orders to be completed
	ordersWg.Wait()

	// Print the metrics summary
	eventSystem.Stop()
	eventSystem.PrintMetricsSummary()

	// Close the coffee shop
	coffeeShop.Close()

	logger.Info("Coffee shop closed")
}
