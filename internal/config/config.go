package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Destination DestinationConfig `yaml:"destination"`
	Sync        SyncConfig        `yaml:"sync"`
	Setup       CommandConfig     `yaml:"setup"`
	Command     interface{}       `yaml:"command"`
	Output      []string          `yaml:"output"`
}

type DestinationConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Port     int    `yaml:"port"`
}

type SyncConfig struct {
	Include []string `yaml:"include"`
	Exclude string `yaml:"exclude"`
}

type CommandConfig struct {
	Command interface{} `yaml:"command"`
}

func ParseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	// Convert single string command to array
	config.Setup.Command = normalizeCommand(config.Setup.Command)
	config.Command = normalizeCommand(config.Command)

	return &config, nil
}

func normalizeCommand(cmd interface{}) []string {
	switch v := cmd.(type) {
	case string:
		return []string{v}
	case []interface{}:
		result := make([]string, len(v))
		for i, item := range v {
			result[i] = fmt.Sprint(item)
		}
		return result
	default:
		return nil
	}
}
