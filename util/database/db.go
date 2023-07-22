package database

import (
	"database/sql"
	"fmt"
	"github.com/Anupam-dagar/baileys/configuration"
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
	MaxIdleConnections int
	MaxOpenConnections int
}

const (
	DsnStringFormat           = "postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Kolkata"
	ConfigKeyDatabasePort     = "database.port"
	ConfigKeyDatabaseHost     = "database.host"
	ConfigKeyDatabaseUserName = "database.username"
	ConfigKeyDatabasePassword = "database.password"
	ConfigKeyDatabaseName     = "database.name"
)

func (c Config) buildDSN() string {
	return fmt.Sprintf(DsnStringFormat, c.Username, c.Password, c.Host, c.Port, c.DbName)
}

// InitDatabaseWithConfig initializes postgres connection with provided configuration.
func InitDatabaseWithConfig(config Config) {
	var err error
	db, err = gorm.Open(postgres.Open(config.buildDSN()), &gorm.Config{
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

// InitDatabase initializes postgres connection with default yaml configuration keys.
func InitDatabase() {
	config := Config{
		Host:     configuration.GetStringConfig(ConfigKeyDatabaseHost),
		Port:     configuration.GetStringConfig(ConfigKeyDatabasePort),
		Username: configuration.GetStringConfig(ConfigKeyDatabaseUserName),
		Password: configuration.GetStringConfig(ConfigKeyDatabasePassword),
		DbName:   configuration.GetStringConfig(ConfigKeyDatabaseName),
	}
	InitDatabaseWithConfig(config)
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

func GetTableName(db *gorm.DB, entityStruct interface{}) (string, error) {
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(entityStruct)
	if err != nil {
		return "", err
	}
	return stmt.Schema.Table, nil
}
