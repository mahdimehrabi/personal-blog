package article

import (
	"strings"
	"time"
)

type Article struct {
	ID        int64
	Title     string
	Slug      string
	Content   string
	Tags      string
	CreatedAt int64
	UpdateAt  int64
	DeletedAt int64
}

func NewArticle(title, slug, content string, tags []string) *Article {
	a := &Article{
		Title:     title,
		Slug:      slug,
		Content:   content,
		CreatedAt: time.Now().Unix(),
	}
	a.SetTags(tags)
	return a
}

func (a *Article) SetTags(tags []string) {
	a.Tags = strings.Join(tags, ",")
}
