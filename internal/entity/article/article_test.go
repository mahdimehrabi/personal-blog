package article

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArticle_SetTags(t *testing.T) {
	tags := []string{"tag1", "tag2", "tag3"}
	article := NewArticle()
	article.SetTags(tags)

	assert.Equal(t, "tag1,tag2,tag3", article.Tags)
}
