// Package config is for loading environment variables
package config

import (
	"github.com/bareish/captnHook/pkg/services"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Service is the environment implementation
type Service struct{}

// Load will load the configuration variables
func (c *Service) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("did not load env vars from .env")
	}
}

// Get will return the default config
func (c *Service) Get() services.Config {
	return services.Config{
		General: getGeneralConfig(),
		Alpaca:  getAlpacaConfig(),
		Binance: getBinanceConfig(),
		Coinbase: getCoinbaseConfig(),
	}
}

// return the server general configuration
func getGeneralConfig() services.GeneralConfig {
	// default to dev settings
	config := services.GeneralConfig{
		Port:     os.Getenv("PORT"),
		CertPath: os.Getenv("CERT_PATH"),
		KeyPath:  os.Getenv("KEY_PATH"),
		AppEnv:   os.Getenv("MODE"),
		BaseURL:  os.Getenv("BASE_URL"),
	}
	// change urls if we are in production
	if config.AppEnv == "PROD" {
		config.BaseURL = "https://0.0.0.0"
	}

	return config
}

// return the Alpaca configuration
func getAlpacaConfig() services.AlpacaConfig {
	// setup config values
	config := services.AlpacaConfig{
		// client id
		ClientID: os.Getenv("ALPACA_CLIENT_ID"),
		// client secret
		ClientSecret: os.Getenv("ALPACA_CLIENT_SECRET"),
		// account type
		AccountType: os.Getenv("ALPACA_ACCOUNT_TYPE"),
		// base - we default to paper url
		BaseURL: "https://paper-api.alpaca.markets",
		// websocket url
		WebSocketURL: os.Getenv("ALPACA_WEBSOCKET_URL"),
	}
	// check if we are using the live url
	if config.AccountType == "live" {
		config.BaseURL = "https://api.alpaca.markets"
	}

	return config
}

// return the Binance configuration
func getBinanceConfig() services.BinanceConfig {
	// setup config values
	config := services.BinanceConfig{
		// client secret
		ClientSecret: os.Getenv("BINANCE_SECRET_KEY"),
		// client id
		ClientId: os.Getenv("BINANCE_CLIENT_ID"),
	}

	return config
}

// return the Coinbase configuration
func getCoinbaseConfig() services.CoinbaseConfig {
	return services.CoinbaseConfig{
		ClientID: os.Getenv("COINBASE_CLIENT_ID"),
		ClientSecret: os.Getenv("COINBASE_SECRET_KEY"),
	}
}
