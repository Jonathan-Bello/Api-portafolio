package psqltech

import (
	"errors"
	"fmt"

	"github.com/Jonathan-Bello/Api-portafolio/pkg/tech"
	"github.com/Jonathan-Bello/Api-portafolio/storage"
	"gorm.io/gorm"
)

// Create create a new tech in database
func Create() error {
	res := storage.DB().Create(&tech.Model{})

	if res.Error != nil {
		return fmt.Errorf("can't create tech in database: %v", res.Error)
	}

	return nil
}

// GetAll get all techs from database
func GetAll() (tech.Techs, error) {
	var techs tech.Techs

	res := storage.DB().Find(&techs)

	if res.Error != nil {
		return nil, fmt.Errorf("can't get techs from database: %v", res.Error)
	}

	return techs, nil
}

// GetByID get tech by id from database
func GetByID(id uint) (tech.Model, error) {
	tech := tech.Model{}
	res := storage.DB().First(&tech, id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return tech, res.Error
	}

	if res.Error != nil {
		return tech, fmt.Errorf("can't get tech from database: %v", res.Error)
	}

	return tech, nil
}

// Update update tech in database
func Update(t tech.Model) error {
	techCondition := tech.Model{}
	techCondition.ID = t.ID
	res := storage.DB().Model(&techCondition).Updates(t)

	if res.Error != nil {
		return fmt.Errorf("can't update tech in database: %v", res.Error)
	}

	if res.RowsAffected == 0 {
		return storage.ErrNotFoundUpdate
	}

	return nil
}

// Delete delete tech from database
func Delete(id uint) error {
	res := storage.DB().Delete(tech.Model{}, id)

	if res.Error != nil {
		return fmt.Errorf("can't delete tech in database: %v", res.Error)
	}

	if res.RowsAffected == 0 {
		return storage.ErrNotFoundDelete
	}

	return nil
}