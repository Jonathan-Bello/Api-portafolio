package storage

import (
	"log"

	"github.com/Jonathan-Bello/Api-portafolio/pkg/author"
	"github.com/Jonathan-Bello/Api-portafolio/pkg/blog"
	"github.com/Jonathan-Bello/Api-portafolio/pkg/tech"
)

// Migration creates the database tables
func Migration() error {
	err := db.AutoMigrate(&author.Model{}, &tech.Model{}, &blog.Model{})
	if err != nil {
		log.Printf("Error en la migracion: %v\n", err)
		return err
	}
	return nil
}
