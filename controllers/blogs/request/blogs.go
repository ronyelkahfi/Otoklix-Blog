package request

import (
	_blogDomain "otoklix-blog/business/blogs"
	"time"
)

type BlogRequest struct {
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Published_at time.Time `json:"published_at"`
}

func RequestToDomain(data BlogRequest) _blogDomain.Domain {
	return _blogDomain.Domain{
		Title:        data.Title,
		Content:      data.Content,
		Published_at: data.Published_at,
	}
}
