package dao

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	Title   string `gorm:"index"`
	Slug    string `gorm:"index"`
	Content string
	Tags    string
}
