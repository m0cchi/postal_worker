package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Module ModuleConfig
	Server ServerConfig
}

type ModuleConfig struct {
	Dir string `toml:"modules_dir"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func isFile(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("%s is not FILE", path)
	}
	return nil
}

func isDir(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not FILE", path)
	}
	return nil
}

func NewConfig(path string) (*Config, error) {
	if err := isFile(path); err != nil {
		return nil, err
	}

	var config *Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
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
	err := isDir(c.Dir)
	return err
}

func (c ServerConfig) Validate() error {
	return nil
}
