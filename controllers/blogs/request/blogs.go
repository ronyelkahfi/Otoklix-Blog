package request

type BlogRequest struct {
	Id           uint   `json:"id"`
	Title        string `json:"title"`
	Content      string
	Published_at string
}
