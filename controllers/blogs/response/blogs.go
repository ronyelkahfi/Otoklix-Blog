package response

import (
	_blogDomain "otoklix-blog/business/blogs"
	"time"
)

type BlogResponse struct {
	Id           uint      `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Published_at time.Time `json:"published_at"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

func DomainToResponse(data _blogDomain.Domain) BlogResponse {
	return BlogResponse{
		Id:           data.Id,
		Title:        data.Title,
		Content:      data.Content,
		Published_at: data.Published_at,
		Created_at:   data.Created_at,
		Updated_at:   data.Updated_at,
	}
}
func ListDomainToResponse(data []_blogDomain.Domain) []BlogResponse {
	var output []BlogResponse
	for _, row := range data {
		output = append(output, DomainToResponse(row))
	}
	return output
}
