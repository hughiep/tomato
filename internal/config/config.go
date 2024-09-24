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
		MysqlDbHost         string `yaml:"mysqlhost"`
		MysqlDbPort         string `yaml:"mysqlport"`
		MysqlDbName         string `yaml:"mysqlname"`
		MySqlDbUser         string `yaml:"mysqluser"`
		MySqlDbPassword     string `yaml:"mysqlpassword"`
		MySqlDbRootPassword string `yaml:"mysqlroot_password"`
	} `yaml:"database"`
}

func (c *Config) Bind() {
	c.App.Port = os.Getenv("PORT")

	c.Database.MysqlDbName = os.Getenv("DATABASE_MYSQLNAME")
	c.Database.MySqlDbUser = os.Getenv("DATABASE_MYSQLUSER")
	c.Database.MySqlDbPassword = os.Getenv("DATABASE_MYSQLPASSWORD")
	c.Database.MySqlDbRootPassword = os.Getenv("DATABASE_MYSQL_ROOT_PASSWORD")
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
