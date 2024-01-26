package article

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestArticle_SetTags(t *testing.T) {
	tags := []string{"tag1", "tag2", "tag3"}
	article := NewArticle("title", "sdgds", "dsgdsg", []string{})
	article.SetTags(tags)

	assert.Equal(t, "tag1,tag2,tag3", article.Tags)
}

func TestNewArticle(t *testing.T) {
	tags := []string{"tag1", "tag2", "tag3"}
	article := NewArticle("title", "sdgds", "dsgdsg", tags)
	assert.Equal(t, "tag1,tag2,tag3", article.Tags)
	assert.Equal(t, time.Now().Unix(), article.CreatedAt)
}
