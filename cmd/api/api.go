package main

import (
	"fmt"
	"go_ecommerce/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo with Logger and Recover!")
}

func postHandler(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, "Hello, "+name)
}

func greetHandler(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name)
}

func setupRoutes(e *echo.Echo) {
	e.GET("/", helloHandler)
	e.GET("/greet/:name", greetHandler)
	e.POST("/post", postHandler)
}

func initDB(cfg *config.Config) (*gorm.DB, error) {
	dbConn, err := config.ConnectDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to the database: %v", err)
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		return nil, fmt.Errorf("Error retrieving SQL DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("Error pinging the database: %v", err)
	}

	fmt.Println("Successfully connected to the MySQL database!")
	return dbConn, nil
}

func main() {

	cfg, err := config.LoadConfig("config.yaml")

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dbConn, err := initDB(cfg)

	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer func() {
		sqlDB, err := dbConn.DB()
		if err != nil {
			log.Printf("Error retrieving SQL DB: %v", err)
		}
		sqlDB.Close()
	}()

	e := echo.New()

	// Middleware for logging and recovering from panics
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	// Set up all routes
	setupRoutes(e)

	
	go func() {
		if err := e.Start(":8080"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()


	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Initiating server shutdown
	fmt.Println("Shutting down gracefully...")

	if err := e.Shutdown(nil);
	      err != nil {

		log.Fatalf("Server Shutdown failed: %v", err)
	}

	fmt.Println("Server stopped successfully.")
}
