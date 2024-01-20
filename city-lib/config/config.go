package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	User       string `json:"user"`
	Password   string `json:"password"`
	DbName     string `json:"db_name"`
	DbHost     string `json:"db_host"`
	City       string `json:"city"`
	ServerHost string `json:"server_host"`
}

func (c *Config) Load(f *os.File) error {

	if f == nil {
		return ErrInvalidFilename
	}

	content, err := io.ReadAll(f)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, c); err != nil {
		return err
	}

	return nil
}
