package blogs

import (
	_blogDomain "otoklix-blog/business/blogs"
	"time"
)

type Blog struct {
	Id          uint   `gorm:"primarykey;autoIncrement"`
	Title       string `gorm:"size:100"`
	Content     string `gorm:"type:text"`
	PublishedAt time.Time
	CreateAt    time.Time
	UpdatedAt   time.Time
}

func ToDomain(data Blog) _blogDomain.Domain {
	return _blogDomain.Domain{
		Id:           data.Id,
		Title:        data.Title,
		Content:      data.Content,
		Published_at: data.PublishedAt,
	}
}

func ToListDomain(data []Blog) []_blogDomain.Domain {
	var output []_blogDomain.Domain
	for _, row := range data {
		output = append(output, ToDomain(row))
	}
	return output
}
