package db

import (
	"fmt"
	"time"

	"tomato/internal/config"
	"tomato/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func dbURL(c *config.Config) string {
	// Get the database URL
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Database.MySqlDbUser,
		c.Database.MySqlDbPassword,
		c.Database.MysqlDbHost,
		c.Database.MysqlDbPort,
		c.Database.MysqlDbName,
	)
}

func Connect(c *config.Config) {
	// Connect to the database
	dsn := dbURL(c)
	connection := mysql.Open(dsn)
	mysql, err := gorm.Open(connection, &gorm.Config{})
	if err != nil {
		logger.Logger.Error(dsn)
		panic("Failed to connect to database" + dsn)
	}

	rawDB, err := mysql.DB()
	if err != nil {
		panic("Failed to get raw database")
	}

	// Set the maximum number of open connections
	rawDB.SetMaxOpenConns(10)

	// Set the maximum number of idle connections
	rawDB.SetMaxIdleConns(10)

	// Set the maximum lifetime of a connection
	rawDB.SetConnMaxLifetime(time.Hour)

	db = mysql
}
