package blogs

import (
	"context"
	"time"
)

type BlogUsecase struct {
	ContextTimeout time.Duration
	Repository     Repository
}

func NewUseCase(repo Repository, timeout time.Duration) Usecase {
	return &BlogUsecase{
		ContextTimeout: timeout,
		Repository:     repo,
	}
}

func (uc *BlogUsecase) GetPosts(ctx context.Context) ([]Domain, error) {
	return uc.Repository.GetData(ctx)
}
func (uc *BlogUsecase) ShowPost(ctx context.Context, id uint) (Domain, error) {
	return uc.Repository.GetById(ctx, id)
}
func (uc *BlogUsecase) CreatePost(ctx context.Context, data Domain) (Domain, error) {
	return uc.Repository.Create(ctx, data)
}

func (uc *BlogUsecase) UpdatePost(ctx context.Context, data Domain) (Domain, error) {
	return uc.Repository.Update(ctx, data)
}

func (uc *BlogUsecase) DeletePost(ctx context.Context, id uint) (Domain, error) {
	return uc.Repository.Delete(ctx, id)
}
