package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port      string `yaml:"port"`
	AuthToken string `yaml:"auth_token"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	return &cfg, err
}
