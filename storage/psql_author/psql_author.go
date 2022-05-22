package psqlauthor

import (
	"errors"
	"fmt"

	"github.com/Jonathan-Bello/Api-portafolio/models"
	"github.com/Jonathan-Bello/Api-portafolio/storage"
	"gorm.io/gorm"
)

// Create create a new author in database
func Create(author models.Author) error {
	res := storage.DB().Create(&author)

	if res.Error != nil {
		return fmt.Errorf("can't create author in database: %v", res.Error)
	}

	return nil
}

// GetAll get all authors from database
func GetAll() (models.Authors, error) {
	var authors models.Authors

	res := storage.DB().Find(&authors)

	if res.Error != nil {
		return nil, fmt.Errorf("can't get authors from database: %v", res.Error)
	}

	return authors, nil
}

// GetByID get author by id from database
func GetByID(id uint) (models.Author, error) {
	author := models.Author{}
	res := storage.DB().First(&author, id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return author, res.Error
	}

	if res.Error != nil {
		return author, fmt.Errorf("can't get author from database: %v", res.Error)
	}

	return author, nil
}

// Update update author in database
func Update(a models.Author) error {
	authorCondition := models.Author{}
	authorCondition.ID = a.ID
	res := storage.DB().Model(&authorCondition).Updates(a)

	if res.Error != nil {
		return fmt.Errorf("can't update author in database: %v", res.Error)
	}

	if res.RowsAffected == 0 {
		return storage.ErrNotFoundUpdate
	}

	return nil
}

// Delete delete author from database
func Delete(id uint) error {
	res := storage.DB().Delete(models.Author{}, id)

	if res.Error != nil {
		return fmt.Errorf("can't delete author from database: %v", res.Error)
	}

	if res.RowsAffected == 0 {
		return storage.ErrNotFoundDelete
	}

	return nil
}
