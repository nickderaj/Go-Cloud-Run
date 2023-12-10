package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()

	e.GET("/", HealthCheck)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port environment variable not found")
	}
	address := ":" + port
	e.Logger.Fatal(e.Start(address))
}
