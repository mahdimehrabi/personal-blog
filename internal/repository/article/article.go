package article

import (
	"errors"
	"github.com/mahdimehrabi/personal-blog/internal/entity/article"
)

var (
	ErrArticleNotFound     = errors.New("article not found")
	ErrArticleAlreadyExist = errors.New("article already exist")
)

type Repository interface {
	Create(article *article.Article) error
	Update(article *article.Article) error
	Detail() *article.Article
	Delete() error
	List() []*article.Article
}
