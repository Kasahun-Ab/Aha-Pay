// db.go
package config

import (
	"fmt"
	"io/ioutil"

	// Importing the MySQL driver for GORM
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3" // For parsing YAML configuration
	"gorm.io/driver/mysql" // MySQL driver for GORM
	"gorm.io/gorm" // GORM library for ORM functionality
)

// Config struct holds the configuration for MySQL from the YAML file
type Config struct {
	MySQL struct {
		User     string `yaml:"user"`     // MySQL username
		Password string `yaml:"password"` // MySQL password
		Dbname   string `yaml:"dbname"`   // Database name
		Host     string `yaml:"host"`     // Database host (e.g., "localhost")
		Port     int    `yaml:"port"`     // Port on which MySQL is running (e.g., 3306)
		SSLMode  string `yaml:"sslmode"`  // SSL mode for MySQL (optional)
	} `yaml:"mysql"`
}

// LoadConfig reads and parses the configuration file (config.yaml)
// It returns a Config object containing MySQL connection details.
func LoadConfig(filename string) (*Config, error) {
	// Read the YAML configuration file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err // If there's an error reading the file, return the error
	}

	var config Config

	// Parse the YAML data into the Config struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err // If parsing fails, return the error
	}

	// Return the populated Config struct
	return &config, nil
}

func ConnectDB(config *Config) (*gorm.DB, error) {
	// Build the Data Source Name (DSN) string to connect to the MySQL database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySQL.User,       // MySQL username
		config.MySQL.Password,   // MySQL password
		config.MySQL.Host,       // MySQL host (e.g., "localhost")
		config.MySQL.Port,       // MySQL port (e.g., 3306)
		config.MySQL.Dbname)     // MySQL database name

	// Open a connection to the MySQL database using the DSN string and GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err // If there's an error connecting to the database, return the error
	}

	// Get the underlying SQL DB connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err // If there's an error retrieving the SQL DB instance, return the error
	}

	// Set the maximum number of idle and open connections
	sqlDB.SetMaxIdleConns(10)  // Set the max number of idle connections to 10
	sqlDB.SetMaxOpenConns(100) // Set the max number of open connections to 100

	// Return the GORM DB object
	return db, nil
}

func InitDB(cfg *Config) (*gorm.DB, error) {
	dbConn, err := ConnectDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to the database: %w", err)
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		return nil, fmt.Errorf("Error retrieving SQL DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("Error pinging the database: %w", err)
	}

	fmt.Println("Successfully connected to the MySQL database!")
	return dbConn, nil
}