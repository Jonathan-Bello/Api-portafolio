package main

import (
	"log"
	"os"

	"github.com/Jonathan-Bello/Api-portafolio/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Database
	driver := storage.Postgres
	storage.New(driver)
	storage.Migration()

	err := e.Start(":" + port)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
