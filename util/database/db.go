package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type Config struct {
	Host               string
	Port               string
	Username           string
	Password           string
	DbName             string
	DSN                string
	MaxIdleConnections int
	MaxOpenConnections int
}

const DsnStringFormat = "postgres://%s:%s@%s:%s/%s?sslmode=disabled&TimeZone=Asia/Kolkata"

func (c *Config) buildDSN() string {
	return fmt.Sprintf(DsnStringFormat, c.Username, c.Password, c.Host, c.Port, c.DbName)
}

// InitDatabase initializes postgres connection with provided configuration.
func InitDatabase(config *Config) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		panic(fmt.Sprintf("Connection : unable to establish connection with database: %s ", err.Error()))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("SQLDB : unable to establish connection with database: %s", err.Error()))
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Ping failed : unable to establish connection with database: %s ", err.Error()))
	}

	SetMaxIdleConnections(sqlDB, config.MaxIdleConnections)
	SetMaxOpenConnections(sqlDB, config.MaxOpenConnections)
}

func SetMaxIdleConnections(sqlDB *sql.DB, connections int) {
	if connections <= 0 {
		return
	}

	sqlDB.SetMaxIdleConns(connections)
}

func SetMaxOpenConnections(sqlDB *sql.DB, connections int) {
	if connections <= 0 {
		return
	}

	sqlDB.SetMaxOpenConns(connections)
}

// DisconnectDatabase closes the connection to database.
func DisconnectDatabase() {
	fmt.Println(`Disconnecting from database...`)
	dbInstance, err := GetDatabase().DB()
	if err != nil {
		fmt.Printf(`Error retrieving sql DB: %v`, err.Error())

		return
	}

	err = dbInstance.Close()
	if err != nil {
		fmt.Printf(`Error disconnecting from db: %v`, err.Error())

		return
	}
}

// GetDatabase returns the database instance
func GetDatabase() *gorm.DB {
	return db
}
