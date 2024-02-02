package article

import (
	"errors"
	entityArticle "github.com/mahdimehrabi/personal-blog/internal/entity/article"
	loggerInfra "github.com/mahdimehrabi/personal-blog/internal/infrastracture/logger"
	"github.com/mahdimehrabi/personal-blog/internal/repository/article"
)

type Article struct {
	logger            loggerInfra.Logger
	articleRepository article.Repository
}

func NewArticle(articleRepository article.Repository,
	logger loggerInfra.Logger) *Article {
	return &Article{
		articleRepository: articleRepository,
		logger:            logger,
	}
}

func (a Article) Create(artEnt *entityArticle.Article) error {
	if err := a.articleRepository.Create(artEnt); err != nil {
		if errors.Is(err, article.ErrArticleAlreadyExist) {
			return err
		}
		a.logger.Error(err.Error())
		return err
	}
	return nil
}
