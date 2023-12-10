package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nickderaj/goose/api"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", HealthCheck)
	e.GET("/users", api.GetAllUsers)
	e.POST("/users", api.CreateUser)
	e.GET("/users/:id", api.GetUser)
	e.PUT("/users/:id", api.UpdateUser)
	e.DELETE("/users/:id", api.DeleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port environment variable not found")
	}
	address := ":" + port
	e.Logger.Fatal(e.Start(address))
}
