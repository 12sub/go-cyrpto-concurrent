package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type FileTask struct {
	Mode  string   `yaml:"mode"`
	Files []string `yaml:"files"`
}

type Config struct {
	DefaultScheme   string   `yaml:"default_scheme"`
	Concurrent      bool     `yaml:"concurrent"`
	DefaultPassword string   `yaml:"default_password"`
	LogLevel        string   `yaml:"log_level"`
	FileTask        FileTask `yaml:"file_task"`
}

// more changes will be made for reading commands from configuration files 

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't read from config file: %w", err)
	}
	var AppConfig Config
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		return nil, fmt.Errorf("couldn't parse config file: %w", err)
	}
	return &AppConfig, nil
}
