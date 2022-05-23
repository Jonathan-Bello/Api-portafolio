package main

import (
	"log"
	"os"

	"github.com/Jonathan-Bello/Api-portafolio/routes"
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
	err := storage.Migration()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Migration success YEAH")
	}

	// Routes
	routes.Router(*e)

	err = e.Start(":" + port)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
