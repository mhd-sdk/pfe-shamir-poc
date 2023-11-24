package config

import (
	"encoding/json"
	"os"

	"github.com/pfe-manager/pkg/models"
)

type Config struct {
	ShamirSplitNumber int              `json:"shamirSplitNumber"`
	ShamirThreshold   int              `json:"shamirThreshold"`
	Services          []models.Service `json:"services"`
	Mode              string           `json:"mode"` // can be "dev" or "prod"
}

var cfg Config

func GetConfig() Config {
	return cfg
}

func GetServices() []models.Service {
	return cfg.Services
}

func GetServiceByName(name string) models.Service {
	for _, service := range cfg.Services {
		if service.Name == name {
			return service
		}
	}
	return models.Service{}
}

func GetNumberOfServices() int {
	return len(cfg.Services)
}

func GetNumberOfUpServices() int {
	up := 0
	for _, service := range cfg.Services {
		if service.Status == models.ServiceUp {
			up++
		}
	}
	return up
}

func GetNumberOfDownServices() int {
	down := 0
	for _, service := range cfg.Services {
		if service.Status == models.ServiceDown {
			down++
		}
	}
	return down
}

func LoadConfig() error {
	file, error := os.Open("./config/config.json")
	if error != nil {
		return error
	}
	defer file.Close()

	decode := json.NewDecoder(file)
	configReadResult := Config{}
	err := decode.Decode(&configReadResult)
	if err != nil {
		return err
	}
	cfg = configReadResult
	return nil
}
