package storage

import (
	"log"

	"github.com/Jonathan-Bello/Api-portafolio/models"
)

// Migration creates the database tables
func Migration() error {
	err := db.AutoMigrate(&models.Author{}, &models.Tech{}, &models.Blog{})
	if err != nil {
		log.Printf("Error en la migracion: %v\n", err)
		return err
	}
	return nil
}
