package psqlblog

import (
	"fmt"

	"github.com/Jonathan-Bello/Api-portafolio/models"
	"github.com/Jonathan-Bello/Api-portafolio/storage"
)

// Create create a new blog in database
func Create(blog models.Blog) error {
	res := storage.DB().Create(&blog)

	if res.Error != nil {
		return fmt.Errorf("can't create blog in database: %v", res.Error)
	}

	return nil
}

// GetAll get all blogs from database
func GetAll() (models.Blogs, error) {
	var blogs models.Blogs

	res := storage.DB().Preload("Techs").Preload("Authors").Find(&blogs)

	if res.Error != nil {
		return nil, fmt.Errorf("can't get blogs from database: %v", res.Error)
	}

	return blogs, nil
}
