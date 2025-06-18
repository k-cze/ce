package config

import (
	"ce/internal/csvloader"
	"log"
	"os"
)

type Config interface {
	GetPort() string
	GetOpenExchangeAPIKey() string
	GetCSVDrivers() csvloader.Drivers
}

type config struct {
	port       string
	apiKey     string
	csvDrivers csvloader.Drivers
}

func New() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default if not provided
	}

	apiKey := os.Getenv("OPENEXCHANGERATES_API_KEY")
	if apiKey == "" {
		log.Fatal("missing required env var: OPENEXCHANGE_API_KEY")
	}

	drivers, err := csvloader.InitAllDrivers()
	if err != nil {
		log.Fatal("failed to initialize drivers:", err)
	}

	return &config{
		port:       port,
		apiKey:     apiKey,
		csvDrivers: drivers,
	}
}

func (c *config) GetPort() string {
	return c.port
}

func (c *config) GetOpenExchangeAPIKey() string {
	return c.apiKey
}

func (e *config) GetCSVDrivers() csvloader.Drivers {
	return e.csvDrivers
}
