package blog

import (
	"github.com/Jonathan-Bello/Api-portafolio/pkg"
	"github.com/Jonathan-Bello/Api-portafolio/pkg/author"
	"github.com/Jonathan-Bello/Api-portafolio/pkg/tech"
)

type Model struct {
	pkg.GlobalModel `gorm:"embedded"`
	Title           string         `gorm:"type:varchar(100);not null" json:"title"`
	Body            string         `gorm:"type:text;not null" json:"body"`
	Description     string         `gorm:"type:varchar(100);not null" json:"description"`
	Url             *string        `gorm:"type:varchar(256)" json:"url"`
	Type            uint8          `gorm:"type:varchar(256);not null" json:"type"`
	Techs           []tech.Model   `gorm:"many2many:blog_techs" json:"techs"`
	Authors         []author.Model `gorm:"many2many:blog_authors" json:"authors"`
}

func (m *Model) TableName() string {
	return "blogs"
}
