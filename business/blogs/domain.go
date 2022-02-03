package blogs

import (
	"context"
	"time"
)

type Domain struct {
	Id           uint
	Title        string
	Content      string
	Published_at time.Time
}

type Usecase interface {
	GetPosts(ctx context.Context) ([]Domain, error)
	ShowPost(ctx context.Context, id uint) (Domain, error)
	CreatePost(ctx context.Context, data Domain) (Domain, error)
	UpdatePost(ctx context.Context, data Domain) (Domain, error)
	DeletePost(ctx context.Context, id uint) (Domain, error)
}

type Repository interface {
	GetData(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, data Domain) (Domain, error)
	Update(ctx context.Context, data Domain) (Domain, error)
	Delete(ctx context.Context, id uint) (Domain, error)
}
