package models

type Blog struct {
	GlobalModel `gorm:"embedded"`
	Title       string   `gorm:"type:varchar(100);not null" json:"title"`
	Body        string   `gorm:"type:text;not null" json:"body"`
	Description string   `gorm:"type:varchar(100);not null" json:"description"`
	Url         *string  `gorm:"type:varchar(256)" json:"url"`
	Type        uint8    `gorm:"type:smallint;not null" json:"type"`
	Techs       []Tech   `gorm:"many2many:blog_techs" json:"techs"`
	Authors     []Author `gorm:"many2many:blog_authors" json:"authors"`
}

type Blogs []Blog

func (m *Blog) TableName() string {
	return "blogs"
}
