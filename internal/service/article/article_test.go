package article

import (
	"errors"
	"github.com/mahdimehrabi/personal-blog/internal/entity/article"
	"github.com/mahdimehrabi/personal-blog/mocks/infrastructure"
	"github.com/mahdimehrabi/personal-blog/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestArticle_Create_Success(t *testing.T) {
	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	articleEnt := article.NewArticle("article1", "article-1",
		"dsggdsgds lkadsjgdajk dlsagjlgd jalgjdka lkadgs",
		[]string{"tag1", "tag2", "tag3"})
	articleMock.On("Create", articleEnt).Return(nil)

	articleService := NewArticle(articleMock, logger)
	err := articleService.Create(articleEnt)
	assert.Nil(t, err)
	articleMock.AssertExpectations(t)
}

func TestArticle_Create_DBErr(t *testing.T) {
	err := errors.New("error")
	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	logger.On("Error", mock.Anything).Return()

	articleEnt := article.NewArticle("article1", "article-1",
		"dsggdsgds lkadsjgdajk dlsagjlgd jalgjdka lkadgs",
		[]string{"tag1", "tag2", "tag3"})
	articleMock.On("Create", articleEnt).Return(err)

	articleService := NewArticle(articleMock, logger)
	err = articleService.Create(articleEnt)
	assert.Error(t, err)
	articleMock.AssertExpectations(t)
	logger.AssertExpectations(t)
}
