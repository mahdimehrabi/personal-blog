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
	Create(article *article.Article) (int64, error)
	Update(article *article.Article) error
	Detail(article *article.Article) (*article.Article, error)
	Delete(article *article.Article) error
	List() []*article.Article
}
