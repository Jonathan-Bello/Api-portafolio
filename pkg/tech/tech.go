package tech

import "github.com/Jonathan-Bello/Api-portafolio/pkg"

type Model struct {
	pkg.GlobalModel `gorm:"embedded"`
	Name            string `gorm:"type:varchar(20);not null" json:"name"`
	Image           string `gorm:"type:varchar(256);not null" json:"image"`
}

type Techs []Model
