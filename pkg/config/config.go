package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env  string `yaml:"env"`
	Http struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"http"`
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}

func Load(fileData []byte) (*Config, error) {
	cfg := &Config{}
	err := yaml.Unmarshal(fileData, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return cfg, nil
}
