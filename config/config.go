package config

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"sync"
)

type Config struct {
	Directory string
}

type config struct {
	config  Config
	created bool
	mu      sync.Mutex
}

var cfg config

func NewConfig(configPath string) error {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if cfg.created {
		return nil
	}

	err := godotenv.Load(configPath)
	if err != nil {
		return err
	}

	config := Config{}
	if err := env.Parse(&config); err != nil {
		return err
	}
	cfg.created = true

	cfg.config = config
	return nil
}

func GetCfg() (Config, error) {
	if !cfg.created {
		return Config{}, fmt.Errorf("Config haven't been init yet")
	}

	return cfg.config, nil
}
