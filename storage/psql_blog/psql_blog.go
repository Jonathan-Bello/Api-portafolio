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

// GetByID get blog by id from database
func GetByID(id uint) (models.Blog, error) {
	var blog models.Blog

	res := storage.DB().Preload("Techs").Preload("Authors").First(&blog, id)

	if res.Error != nil {
		return models.Blog{}, fmt.Errorf("can't get blog from database: %v", res.Error)
	}

	return blog, nil
}

// Update update blog in database
func Update(blog models.Blog) error {
	// Update blogs and relations many2many with techs
	res := storage.DB().Model(&blog).Updates(blog)

	if res.Error != nil {
		return fmt.Errorf("can't update blog in database: %v", res.Error)
	}

	return nil
}

// Delete delete blog in database
func Delete(id uint) error {
	blog := models.Blog{}
	blog.ID = id
	res := storage.DB().Delete(&blog)

	if res.Error != nil {
		return fmt.Errorf("can't delete blog in database: %v", res.Error)
	}

	if res.RowsAffected == 0 {
		return storage.ErrNotFoundDelete
	}

	return nil
}
