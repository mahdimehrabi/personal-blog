package article

import "strings"

type Article struct {
	Title   string
	Slug    string
	Content string
	Tags    string
}

func NewArticle(title, slug, content string, tags []string) *Article {
	a := &Article{
		Title: title, Slug: slug, Content: content,
	}
	a.SetTags(tags)
	return a
}

func (a *Article) SetTags(tags []string) {
	a.Tags = strings.Join(tags, ",")
}
