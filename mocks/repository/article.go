package repository

import (
	articleEntity "github.com/mahdimehrabi/personal-blog/internal/entity/article"
	"github.com/stretchr/testify/mock"
)

type ArticleMock struct {
	mock.Mock
}

func (m ArticleMock) Create(article *articleEntity.Article) (int64, error) {
	args := m.Called(article)
	return int64(args.Int(0)), args.Error(1)
}

func (m ArticleMock) Update(article *articleEntity.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m ArticleMock) Detail(article *articleEntity.Article) *articleEntity.Article {
	args := m.Called(article)
	return args.Get(0).(*articleEntity.Article)
}

func (m ArticleMock) Delete(article *articleEntity.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m ArticleMock) List() []*articleEntity.Article {
	args := m.Called()
	return args.Get(0).([]*articleEntity.Article)
}
