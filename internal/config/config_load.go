package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

func LoadConfig(path string, config interface{}) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("failed to stat config file: %w", err)
	}

	godotenv.Load() //не обрабатываем

	if err := cleanenv.ReadConfig(path, config); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	return nil
}
