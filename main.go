package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sentinal-contactform/mailing"
	"sentinal-contactform/routes"

	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		os.Exit(0)
	}

	err = mailing.Init()
	if err != nil {
		log.Println("Server startup failed! Mailing Credentials not found.")
		os.Exit(0)
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Method=${method}, Uri=${uri}, Status=${status}\n",
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	routes.HealthRoutes(e)
	contact_api := e.Group("/contact")
	routes.ConactRoutes(contact_api)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
