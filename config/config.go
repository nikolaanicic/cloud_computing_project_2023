package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	User              string `json:"user"`
	Password          string `json:"password"`
	CityDbName        string `json:"db_name"`
	CityDbHost        string `json:"db_host"`
	CentralDbHost     string `json:"central_db_host"`
	CentralDbName     string `json:"central_db_name"`
	City              string `json:"city"`
	CityServer        string `json:"city_server"`
	CentralServerHost string `json:"central_server"`
}

func (c *Config) Load(f *os.File) error {

	if f == nil {
		return fmt.Errorf("invalid file pointer")
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

func (c *Config) GetCityServerHost() string {
	return c.CityServer
}

func (c *Config) GetCentralServerHost() string {
	return c.CentralServerHost
}

func (c *Config) GetCityDbHost() string {
	return c.CityDbHost
}

func (c *Config) GetCentralDbHost() string {
	return c.CentralDbHost
}
