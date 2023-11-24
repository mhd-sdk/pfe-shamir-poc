package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Name string `json:"name"`
	Host string `json:"host"`
}

var cfg Config

func GetConfig() Config {
	return cfg
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

func Override(newConfig Config) {
	cfg = newConfig
}
