package article

import (
	entityArticle "github.com/mahdimehrabi/personal-blog/internal/entity/article"
	"github.com/mahdimehrabi/personal-blog/internal/repository/article"
)

type Article struct {
	articleRepository article.Repository
}

func NewArticle(articleRepository article.Repository) *Article {
	return &Article{
		articleRepository: articleRepository,
	}
}

func (a Article) Create(artEnt *entityArticle.Article) error {
	//if err := a.articleRepository.Create(artEnt); err != nil {
	//	if errors.Is(err, article.ErrArticleAlreadyExist) {
	//		return err
	//	}
	//	log.Error().Err(err).Msg("error in creating article")
	//	return err
	//}
	return nil
}
