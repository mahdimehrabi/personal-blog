package article

import (
	"errors"
	"github.com/mahdimehrabi/personal-blog/internal/entity/article"
	articleRepository "github.com/mahdimehrabi/personal-blog/internal/repository/article"
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
	articleMock.On("Create", articleEnt).Return(1, nil)

	articleService := NewArticle(articleMock, logger)
	id, err := articleService.Create(articleEnt)
	assert.Nil(t, err)
	assert.Equal(t, id, int64(1))
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
	articleMock.On("Create", articleEnt).Return(0, err)

	articleService := NewArticle(articleMock, logger)
	id, err := articleService.Create(articleEnt)
	assert.Equal(t, id, int64(0))
	assert.Error(t, err)
	articleMock.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestArticle_Update_Success(t *testing.T) {
	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	articleEnt := article.NewArticle("article1", "article-1",
		"dsggdsgds lkadsjgdajk dlsagjlgd jalgjdka lkadgs",
		[]string{"tag1", "tag2", "tag3"})
	articleEnt.ID = 1
	articleMock.On("Update", articleEnt).Return(nil)

	articleService := NewArticle(articleMock, logger)
	err := articleService.Update(articleEnt)
	assert.Nil(t, err)
	articleMock.AssertExpectations(t)
}

func TestArticle_Update_DBErr(t *testing.T) {
	expectedErr := errors.New("fake db error")
	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	logger.On("Error", mock.Anything).Return()
	articleEnt := article.NewArticle("article1", "article-1",
		"dsggdsgds lkadsjgdajk dlsagjlgd jalgjdka lkadgs",
		[]string{"tag1", "tag2", "tag3"})
	articleEnt.ID = 1
	articleMock.On("Update", articleEnt).Return(expectedErr)

	articleService := NewArticle(articleMock, logger)
	err := articleService.Update(articleEnt)
	assert.ErrorIs(t, err, expectedErr)
	articleMock.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestArticle_Update_NotFound(t *testing.T) {
	expectedErr := articleRepository.ErrArticleNotFound

	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	articleEnt := &article.Article{
		ID: 1,
	}
	articleEnt.ID = 1
	articleMock.On("Update", articleEnt).Return(expectedErr)

	articleService := NewArticle(articleMock, logger)
	err := articleService.Update(articleEnt)
	assert.ErrorIs(t, err, expectedErr)
	articleMock.AssertExpectations(t)
	logger.AssertNotCalled(t, "Error")
}

func TestArticle_Delete_Success(t *testing.T) {
	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	articleEnt := &article.Article{
		ID: 1,
	}
	articleMock.On("Delete", articleEnt).Return(nil)

	articleService := NewArticle(articleMock, logger)
	err := articleService.Delete(articleEnt)
	assert.Nil(t, err)
	articleMock.AssertExpectations(t)
}

func TestArticle_Delete_DBErr(t *testing.T) {
	expectedErr := errors.New("fake db error")
	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	logger.On("Error", mock.Anything).Return()
	articleEnt := &article.Article{
		ID: 1,
	}
	articleEnt.ID = 1
	articleMock.On("Delete", articleEnt).Return(expectedErr)

	articleService := NewArticle(articleMock, logger)
	err := articleService.Delete(articleEnt)
	assert.ErrorIs(t, err, expectedErr)
	articleMock.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestArticle_Delete_NotFound(t *testing.T) {
	expectedErr := articleRepository.ErrArticleNotFound

	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	articleEnt := &article.Article{
		ID: 1,
	}
	articleEnt.ID = 1
	articleMock.On("Delete", articleEnt).Return(expectedErr)

	articleService := NewArticle(articleMock, logger)
	err := articleService.Delete(articleEnt)
	assert.ErrorIs(t, err, expectedErr)
	articleMock.AssertExpectations(t)
	logger.AssertNotCalled(t, "Error")
}

func TestArticle_Detail_Success(t *testing.T) {
	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	articleEnt := &article.Article{
		ID: 1,
	}
	articleMock.On("Detail", articleEnt).Return(articleEnt, nil)

	articleService := NewArticle(articleMock, logger)
	returnedArticleEnt, err := articleService.Detail(articleEnt)
	assert.Nil(t, err)
	assert.Equal(t, returnedArticleEnt, articleEnt)
	articleMock.AssertExpectations(t)
}

func TestArticle_Detail_DBErr(t *testing.T) {
	expectedErr := errors.New("fake db error")
	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	logger.On("Error", mock.Anything).Return()
	articleEnt := &article.Article{
		ID: 1,
	}
	articleEnt.ID = 1
	articleMock.On("Detail", articleEnt).Return(nil, expectedErr)

	articleService := NewArticle(articleMock, logger)
	returnedArticleEnt, err := articleService.Detail(articleEnt)
	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, returnedArticleEnt)
	articleMock.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestArticle_Detail_NotFound(t *testing.T) {
	expectedErr := articleRepository.ErrArticleNotFound

	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	articleEnt := &article.Article{
		ID: 1,
	}
	articleEnt.ID = 1
	articleMock.On("Detail", articleEnt).Return(nil, expectedErr)

	articleService := NewArticle(articleMock, logger)
	returnedArticleEnt, err := articleService.Detail(articleEnt)
	assert.ErrorIs(t, err, expectedErr)
	assert.Nil(t, returnedArticleEnt)
	articleMock.AssertExpectations(t)
	logger.AssertNotCalled(t, "Error")
}
