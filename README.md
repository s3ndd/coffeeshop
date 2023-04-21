# CoffeeShop
CoffeeShop is a simulation of a coffee shop operation, demonstrating the use of concurrency patterns and data structures in Golang. The coffee shop has greeters, cashiers, baristas, grinders, and brewers working together to serve customers efficiently.

## Design
**Greeters as Load Balancers:** Greeters act as load balancers to assign customers to cashiers. They assess the customer queue length of each cashier and assign customers to the cashier with the shortest queue, ensuring efficient customer service.

**Cashiers' Customer Queue:** The cashiers' customer queue is implemented using a priority queue, making it easier to maintain the shortest customer queue and optimize cashier assignment.

**Barista Pool:** All baristas are in a pool, working concurrently to process orders.

**Order Queue:** An order queue is set up between cashiers and baristas. When a customer places an order, the cashier publishes the order to the order queue. Available baristas subscribe to the order queue and pick up the orders as they come in.

**Grinder and Brewer Pools:** All grinders and brewers are part of their respective pools. Baristas can choose an available grinder and brewer from these pools, optimizing resource utilization.

**Workflow:** The workflow of the CoffeeShop simulation is as follows:
- Customer arrives.
- Greeter assigns the customer to a cashier.
- Cashier serves the customer.
- Customer places an order.
- Cashier publishes the order to the order queue.
- Barista picks up the order from the queue.
- Barista chooses an available grinder to grind the coffee beans.
- After grinding, the barista chooses an available brewer to brew the coffee.
- Once all steps are completed, the order is ready for the customer.

By implementing these design principles, CoffeeShop efficiently simulates the workings of a real coffee shop and demonstrates the power of concurrency and data structures in Golang.

## Code Structure
The code is structured in the following way:
- `ci`: Contains shell scripts for building, linting, and testing the project.
- `cmd`: Contains the entry point of the application.
    - `main.go`: The main function that initializes and runs the simulation.
- `coffeeshop.yaml`: The configuration file for the CoffeeShop simulation.
- `internal`: Contains the main packages and components of the application.
    - `coffeeshop`: The core package containing the coffee shop components.
        - `barista`: Contains the Barista struct and related methods, as well as the BaristaPool and related methods.
        - `brewer`: Contains the Brewer struct and related methods, as well as the BrewerPool and related methods.
        - `cashier`: Contains the Cashier struct and related methods, as well as the CashierPool and related methods.
        - `coffeeshop.go`: The main CoffeeShop struct and its methods.
        - `greeter`: Contains the Greeter struct and related methods, as well as the GreeterPool and related methods.
        - `grinder`: Contains the Grinder struct and related methods, as well as the GrinderPool and related methods.
        - `mocks`: Contains mock implementations of various interfaces for testing purposes.
    - `config`: Contains the Config struct and related methods for loading the configuration file.
    - `monitor`: Contains the EventSystemer interface and related implementations for monitoring and event handling, as well as the Metrics struct and related methods.
    - `types`: A package containing common types and interfaces used across the application, such as Coffee, Customer, Order, and OrderQueue.
- `pkg`: Contains utility packages.
- `utils`: Contains utility functions and logging setup.
- `go.mod` and go.sum: Go module configuration files.

## How to Run
1. List itemClone the repository to your local machine.
```
git clone https://github.com/s3ndd/coffeeshop.git
```

2. Navigate to the project directory.
```
cd coffeeshop
```

3. Run the simulation.
```
go run cmd/main.go 
```

After running the simulation, you will see a metrics summary printed in the console. This summary will include information about order processing times.

```json
{
  "level": "info",
  "ts": "2023-04-21T16:03:56.775Z",
  "caller": "monitor/metrics.go:108",
  "msg": "Metrics summary",
  "completed_orders": 20,
  "processed_orders": 20,
  "received_orders": 20,
  "average_brewing_time": 4.05,
  "average_grinding_time": 4.7,
  "average_process_time": 11.400986343,
  "average_waiting_time": 11.501073568
}
```

## Configuration
The CoffeeShop simulation can be customized by modifying the coffeeshop.yaml configuration file. This file allows you to adjust various settings, such as the number of greeters, cashiers, and baristas, as well as configure the equipment and coffee types available in the shop.

To customize the simulation, edit the coffeeshop.yaml file to reflect your desired settings, and then re-run the simulation using go run cmd/main.go. The simulation will adapt to the new configuration, and the metrics summary will reflect the changes made.

## Continuous Integration
This project uses GitHub Actions for continuous integration. You can find the build and test results under the "Actions" tab in the GitHub repository. It ensures that the code is working correctly and helps maintain code quality.

