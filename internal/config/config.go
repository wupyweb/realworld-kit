package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig   `yaml:"server"`
	DB     DatabaseConfig `yaml:"db"`
	// redis
	// k8s
	// docker
}

type ServerConfig struct {
	ENV     string `yaml:"env"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func LoadConfig(cfg_path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(cfg_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
