package gorm

import (
	"github.com/mahdimehrabi/personal-blog/internal/entity/article"
	"gorm.io/gorm"
)

type Article struct {
	db *gorm.DB
}

func NewArticle(db *gorm.DB) *Article {
	return &Article{db: db}
}

func (a Article) Create(article *article.Article) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a Article) Update(article *article.Article) error {
	//TODO implement me
	panic("implement me")
}

func (a Article) Detail(article *article.Article) *article.Article {
	//TODO implement me
	panic("implement me")
}

func (a Article) Delete(article *article.Article) error {
	//TODO implement me
	panic("implement me")
}

func (a Article) List() []*article.Article {
	//TODO implement me
	panic("implement me")
}
