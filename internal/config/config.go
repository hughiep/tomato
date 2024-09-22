package config

import (
	"embed"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed config.yml
var Files embed.FS

type Config struct {
	App struct {
		Port string `yaml:"port"`
	} `yaml:"app"`
}

func (c *Config) Bind() {
	c.App.Port = os.Getenv("PORT")
}

func Load() *Config {
	c := &Config{}

	// Read default values from config file
	in, err := Files.ReadFile("config.yml")
	if err != nil {
		log.Panic(err)
	}
	// Parse config file
	if err := yaml.Unmarshal(in, c); err != nil {
		log.Fatalf("error: %v", err)
	}

	c.Bind()
	return c
}
