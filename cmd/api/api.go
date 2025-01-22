package main

import (
	"context"
	"fmt"
	"go_ecommerce/internal/config"
	"go_ecommerce/internal/handlers"
	"go_ecommerce/internal/models"
	"go_ecommerce/internal/repositories"
	"go_ecommerce/internal/services"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func initDB(cfg *config.Config) (*gorm.DB, error) {
	dbConn, err := config.ConnectDB(cfg)
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
		if sqlDB, err := dbConn.DB(); err != nil {
			log.Printf("Error retrieving SQL DB: %v", err)
		} else {
			sqlDB.Close()
		}
	}()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err = dbConn.AutoMigrate(
		&models.User{},
		&models.AuditLog{},
		&models.CardDetail{},
		&models.Log{},
		&models.Notification{},
		&models.PaymentMethod{},
		&models.RecurringTransaction{},
		&models.RewardPoint{},
		&models.SecurityLog{},
		&models.Transaction{},
		&models.Transfer{},
		&models.UserSession{},
		&models.UserVerification{},
		&models.Wallet{},
	)

	if err != nil {
		log.Fatal("Error auto-migrating: ", err)
	}

	fmt.Println("Auto migration completed")
	userRepo := repositories.NewUserRepository(dbConn)
	authService := services.NewAuthService(dbConn, "secretKey", *userRepo)

	authHandler := handlers.NewAuthHandler(authService)

	restService := services.NewResetService(*userRepo)

	userHandler := handlers.NewRestHandler(restService)

	go func() {
		if err := e.Start(":8080"); err != nil {
			log.Printf("Error starting server: %v", err)
		}
	}()

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)
	e.POST("/forgot-password", userHandler.ForgotPassword)
	e.POST("/reset-password", userHandler.ResetPassword)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown failed: %v", err)
	}

	fmt.Println("Server stopped successfully.")
}
