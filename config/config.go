package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig() (*Config, error) {
	data, err := os.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
