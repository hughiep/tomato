package db

import (
	"fmt"
	"time"

	"tomato/internal/config"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func Connect(c *config.Config) *gorm.DB {
	// Connect to the database
	dsn := dbURL(c)
	dbConnection := mysql.Open(dsn)
	db, err := gorm.Open(dbConnection, &gorm.Config{})
	if err != nil {
		zap.L().Error("Failed to connect to database", zap.Error(err))
		panic(err)
	}

	rawDB, err := db.DB()
	if err != nil {
		zap.L().Error("Failed to get raw database")
		panic(err)
	}

	// Set the maximum number of open connections
	rawDB.SetMaxOpenConns(10)

	// Set the maximum number of idle connections
	rawDB.SetMaxIdleConns(10)

	// Set the maximum lifetime of a connection
	rawDB.SetConnMaxLifetime(time.Hour)

	return db
}
