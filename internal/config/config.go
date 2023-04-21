package config

import (
	"os"
	"sync"

	"github.com/s3ndd/coffeeshop/pkg/utils"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
)

type Configurer interface {
	CoffeeTypes() []*CoffeeType
}

// Config is a struct that contains the settings for the coffee shop.
const configFile = "coffeeshop.yaml"

var (
	config *Config
	once   sync.Once
)

// BrewerSettings is a struct that contains the settings for a coffee brewer.
type BrewerSettings struct {
	Tag                  string `yaml:"tag"`
	OuncesWaterPerSecond int    `yaml:"ouncesWaterPerSecond"`
}

// GrinderSettings is a struct that contains the settings for a coffee grinder.
type GrinderSettings struct {
	Tag            string `yaml:"tag"`
	GramsPerSecond int    `yaml:"gramsPerSecond"`
}

// CoffeeType represents the type of a coffee
type CoffeeType struct {
	Name              string          `yaml:"name"`
	BeansToWaterRatio decimal.Decimal `yaml:"beansToWaterRatio"`
	Price             decimal.Decimal `yaml:"price"`
	SizeInOunces      int             `yaml:"sizeInOunces"`
}

// CoffeeShopSettings is a struct that contains the settings for a coffee shop.
type CoffeeShopSettings struct {
	NumberOfBaristas int               `yaml:"numberOfBaristas"`
	NumberOfCashiers int               `yaml:"numberOfCashiers"`
	NumberOfGreeters int               `yaml:"numberOfGreeters"`
	CashierQueueSize int               `yaml:"cashierQueueSize"`
	OrderQueueSize   int               `yaml:"orderQueueSize"`
	CoffeeTypes      []CoffeeType      `yaml:"coffeeTypes"`
	GrinderSettings  []GrinderSettings `yaml:"grinders"`
	BrewerSettings   []BrewerSettings  `yaml:"brewers"`
}

type Config struct {
	CoffeeShopSettings CoffeeShopSettings `yaml:"coffeeShop"`
}

func (c *Config) CoffeeTypes() []*CoffeeType {
	var coffeeTypes []*CoffeeType
	for _, ct := range c.CoffeeShopSettings.CoffeeTypes {
		coffeeTypes = append(coffeeTypes, &ct)
	}
	return coffeeTypes
}

func (c *Config) GrinderSettings() []*GrinderSettings {
	var grinderSettings []*GrinderSettings
	for _, gs := range c.CoffeeShopSettings.GrinderSettings {
		grinderSettings = append(grinderSettings, &gs)
	}
	return grinderSettings
}

func (c *Config) BrewerSettings() []*BrewerSettings {
	var brewerSettings []*BrewerSettings
	for _, bs := range c.CoffeeShopSettings.BrewerSettings {
		brewerSettings = append(brewerSettings, &bs)
	}
	return brewerSettings
}

func (c *Config) CoffeeShop() *CoffeeShopSettings {
	return &c.CoffeeShopSettings
}

func LoadConfig() *Config {
	once.Do(func() {
		logger := utils.Logger()
		logger.Info("Reading config file")

		data, err := os.ReadFile(configFile)
		if err != nil {
			logger.WithError(err).Fatal("Error reading config file")
		}
		cfg := &Config{}
		if err = yaml.Unmarshal(data, &cfg); err != nil {
			logger.WithError(err).Fatal("Error unmarshalling config file")
		}
		config = cfg
	})

	return config
}
