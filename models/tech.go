package models

type Tech struct {
	GlobalModel `gorm:"embedded"`
	Name        string `gorm:"type:varchar(20);not null" json:"name"`
	Image       string `gorm:"type:varchar(256);not null" json:"image"`
}

type Techs []Tech

func (m *Tech) TableName() string {
	return "techs"
}
