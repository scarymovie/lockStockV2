package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

// LoadConfig reads the YAML file, substitutes environment variables, and unmarshals into the Config struct.
func LoadConfig(path string) (*Config, error) {
	// Read the YAML file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Expand environment variables
	dataWithEnv := []byte(os.ExpandEnv(string(data)))

	// Unmarshal into Config struct
	var cfg Config
	if err := yaml.Unmarshal(dataWithEnv, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
