package models

type Author struct {
	GlobalModel `gorm:"embedded"`
	FirstName       string  `gorm:"type:varchar(50);not null" json:"first_name"`
	LastName        string  `gorm:"type:varchar(50);not null" json:"last_name"`
	Email           string  `gorm:"type:varchar(50);not null" json:"email"`
	Image           string `gorm:"type:varchar(256);not null" json:"image"`
}

type Authors []Author

func (m *Author) TableName() string {
	return "authors"
}