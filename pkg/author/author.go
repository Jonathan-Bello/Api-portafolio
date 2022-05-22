package author

import "github.com/Jonathan-Bello/Api-portafolio/pkg"

type Model struct {
	pkg.GlobalModel `gorm:"embedded"`
	FirstName       string  `gorm:"type:varchar(50);not null" json:"first_name"`
	LastName        string  `gorm:"type:varchar(50);not null" json:"last_name"`
	Email           string  `gorm:"type:varchar(50);not null" json:"email"`
	Image           *string `gorm:"type:varchar(256);not null" json:"image"`
}

type Authors []Model

func (m *Model) TableName() string {
	return "authors"
}
