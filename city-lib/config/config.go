package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	DbHost   string `json:"db_host"`
	City     string `json:"city"`
}

func LoadConfig(f *os.File) (*Config, error) {

	if f == nil {
		return nil, ErrInvalidFilename
	}

	content, err := io.ReadAll(f)

	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := json.Unmarshal(content, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
