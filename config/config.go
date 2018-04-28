package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Module ModuleConfig
	Server ServerConfig
}

type ModuleConfig struct {
	Lib string `toml:"lib"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func existsFile(path string) error {
	_, err := os.Stat(path)
	return err
}

func NewConfig(path string) (*Config, error) {
	if err := existsFile(path); err != nil {
		return nil, err
	}

	var config *Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	if err := config.Validate(); err != nil {
		return config, err
	}

	return config, nil
}

func (c Config) Validate() error {
	if err := c.Module.Validate(); err != nil {
		return err
	}
	err := c.Server.Validate()
	return err
}

func (c ModuleConfig) Validate() error {
	err := existsFile(c.Lib)
	return err
}

func (c ServerConfig) Validate() error {
	return nil
}
