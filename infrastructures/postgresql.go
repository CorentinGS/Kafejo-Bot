package infrastructures

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/corentings/kafejo-bot/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBParams struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// DBConn is the database connection
var DBConn *gorm.DB

// GetDBConn returns the database connection object
func GetDBConn() *gorm.DB {
	return DBConn
}

// DBConfig returns the database configuration parameters
// from environment variables
//
// see: utils/env.go
func DBConfig() DBParams {
	dbConfig := new(DBParams) // Create a new DBParams struct

	// Get database configuration from environment variables
	if os.Getenv("APP_ENV") == "dev" {
		dbConfig.User = os.Getenv("DEBUG_DB_USER")         // Get DB_USER from env
		dbConfig.Password = os.Getenv("DEBUG_DB_PASSWORD") // Get DB_PASSWORD from env
		dbConfig.Host = os.Getenv("DEBUG_DB_HOST")         // Get DB_HOST from env
		dbConfig.DBName = os.Getenv("DEBUG_DB_DB")         // Get DB_DB (db name) from env
		dbConfig.Port = os.Getenv("DEBUG_DB_PORT")         // Get DB_PORT from env
	} else {
		dbConfig.User = os.Getenv("DB_USER")         // Get DB_USER from env
		dbConfig.Password = os.Getenv("DB_PASSWORD") // Get DB_PASSWORD from env
		dbConfig.Host = os.Getenv("DB_HOST")         // Get DB_HOST from env
		dbConfig.DBName = os.Getenv("DB_DB")         // Get DB_DB (db name) from env
		dbConfig.Port = os.Getenv("DB_PORT")         // Get DB_PORT from env
	}

	return *dbConfig
}

// Connect creates a connection to database
// dbConfig is the database configuration parameters
//
// see: utils/config.go and utils/env.go for more details
func (dbConfig *DBParams) Connect() error {
	// Convert port
	dbPort, err := strconv.Atoi(dbConfig.Port)
	if err != nil {
		return utils.ErrConversion
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)

	// Create postgres connection string
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.DBName, dbPort)
	// Open connection
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 newLogger, // Logger
		SkipDefaultTransaction: true,      // Skip default transaction
	})
	if err != nil {
		return utils.ErrOpenDB
	}

	sqlDB, err := DBConn.DB() // Get sql.DB object from gorm.DB
	if err != nil {
		return utils.ErrGetDB
	}
	sqlDB.SetMaxIdleConns(utils.SQLMaxIdleConns) // Set max idle connections
	sqlDB.SetMaxOpenConns(utils.SQLMaxOpenConns) // Set max open connections
	sqlDB.SetConnMaxLifetime(time.Hour)          // Set max connection lifetime

	return nil
}
