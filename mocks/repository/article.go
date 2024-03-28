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
	errVal := args.Get(1)
	var err error
	if errVal != nil {
		err = errVal.(error)
	}
	return int64(args.Int(0)), err
}

func (m ArticleMock) Update(article *articleEntity.Article) error {
	args := m.Called(article)
	errVal := args.Get(0)
	var err error
	if errVal != nil {
		err = errVal.(error)
	}
	return err
}

func (m ArticleMock) Detail(article *articleEntity.Article) (*articleEntity.Article, error) {
	args := m.Called(article)
	articleVal := args.Get(0)
	var articleEntVal *articleEntity.Article
	if articleVal != nil {
		article = articleVal.(*articleEntity.Article)
	}

	errVal := args.Get(1)
	var err error
	if errVal != nil {
		err = errVal.(error)
	}
	return articleEntVal, err
}

func (m ArticleMock) Delete(article *articleEntity.Article) error {
	args := m.Called(article)
	errVal := args.Get(0)
	var err error
	if errVal != nil {
		err = errVal.(error)
	}
	return err
}

func (m ArticleMock) List() []*articleEntity.Article {
	args := m.Called()
	return args.Get(0).([]*articleEntity.Article)
}
