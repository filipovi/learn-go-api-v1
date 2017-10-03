package config

import (
	"encoding/json"
	"os"
)

// Config is a struct loaded from the config.json file.
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// New returns a Config struct filled with the json file
func New(path string) (Config, error) {
	var cfg Config

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return cfg, err
	}

	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, err
}
