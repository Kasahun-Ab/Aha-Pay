package main

import (
	"context"
	"fmt"
	"go_ecommerce/internal/config"
	"go_ecommerce/internal/handlers"
	customMiddleware "go_ecommerce/internal/middleware"
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
)

func main() {

	cfg, err := config.LoadConfig("config.yaml")

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dbConn, err := config.InitDB(cfg)

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

	wallerRepo := repositories.NewWalletRepository(dbConn)

	userRepo := repositories.NewUserRepository(dbConn)

	authService := services.NewAuthService(dbConn, "secretKey", *userRepo, *wallerRepo)
	authHandler := handlers.NewAuthHandler(authService)

	resetService := services.NewResetService(*userRepo)
	resetHandler := handlers.NewRestHandler(resetService)

	userAccountService := services.NewUserAccountService(*userRepo)
	userAccountHandler := handlers.NewUserAccountHandler(userAccountService)

	walletService := services.NewWalletService(*wallerRepo)
	walletHandler := handlers.NewWalletHandler(walletService)

	go func() {
		if err := e.Start(":8080"); err != nil {
			log.Printf("Error starting server: %v", err)
		}
	}()

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)
	e.POST("/forgot-password", resetHandler.ForgotPassword)
	e.POST("/reset-password", resetHandler.ResetPassword)

	userRoutes := e.Group("/user")
	userRoutes.Use(customMiddleware.AuthMiddleware)
	userRoutes.GET("", userAccountHandler.GetUser)
	userRoutes.POST("/email", userAccountHandler.GetUserByEmail)
	userRoutes.PUT("", userAccountHandler.UpdateUser)
	userRoutes.PUT("/email", userAccountHandler.UpdateUserByEmail)
	userRoutes.DELETE("", userAccountHandler.DeleteUser)
	userRoutes.POST("/wallet", walletHandler.CreateWallet)

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
