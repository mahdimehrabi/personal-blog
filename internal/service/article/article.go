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

func (a Article) Create(artEnt *entityArticle.Article) (int64, error) {
	id, err := a.articleRepository.Create(artEnt)
	if err != nil {
		if errors.Is(err, article.ErrArticleAlreadyExist) {
			return id, err
		}
		a.logger.Error(err.Error())
		return id, err
	}
	return id, nil
}

func (a Article) Update(artEnt *entityArticle.Article) error {
	if err := a.articleRepository.Update(artEnt); err != nil {
		if errors.Is(err, article.ErrArticleAlreadyExist) {
			return err
		}
		if errors.Is(err, article.ErrArticleNotFound) {
			return err
		}
		a.logger.Error(err.Error())
		return err
	}
	return nil
}

func (a Article) Delete(artEnt *entityArticle.Article) error {
	if err := a.articleRepository.Delete(artEnt); err != nil {
		if errors.Is(err, article.ErrArticleNotFound) {
			return err
		}
		a.logger.Error(err.Error())
		return err
	}
	return nil
}

func (a Article) Detail(artEnt *entityArticle.Article) (*entityArticle.Article, error) {
	artEnt, err := a.articleRepository.Detail(artEnt)
	if err != nil {
		if errors.Is(err, article.ErrArticleNotFound) {
			return nil, err
		}
		a.logger.Error(err.Error())
		return nil, err
	}
	return artEnt, nil
}
