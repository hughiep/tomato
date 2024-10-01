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

	Database struct {
		MysqlDbHost         string `yaml:"mysql_host"`
		MysqlDbPort         string `yaml:"mysql_port"`
		MysqlDbName         string `yaml:"mysql_name"`
		MySqlDbUser         string `yaml:"mysql_user"`
		MySqlDbPassword     string `yaml:"mysql_password"`
		MySqlDbRootPassword string `yaml:"mysql_root_password"`
	} `yaml:"database"`
}

func (c *Config) Bind() {
	c.App.Port = os.Getenv("PORT")

	c.Database.MysqlDbName = getEnv("DATABASE_MYSQL_NAME", c.Database.MysqlDbName)
	c.Database.MySqlDbUser = getEnv("DATABASE_MYSQL_USER", c.Database.MySqlDbUser)
	c.Database.MySqlDbPassword = getEnv("DATABASE_MYSQL_PASSWORD", c.Database.MySqlDbPassword)
	c.Database.MySqlDbRootPassword = getEnv("DATABASE_MYSQL_ROOT_PASSWORD", c.Database.MySqlDbRootPassword)
	c.Database.MysqlDbHost = getEnv("DATABASE_MYSQL_HOST", c.Database.MysqlDbHost)
	c.Database.MysqlDbPort = getEnv("DATABASE_MYSQL_PORT", c.Database.MysqlDbPort)
}

// Check if os.Getenv has returned value, or using default value from config file
func getEnv(key string, initial string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return initial
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
