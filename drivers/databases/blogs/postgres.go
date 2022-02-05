package blogs

import (
	"context"
	_blogDomain "otoklix-blog/business/blogs"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(database *gorm.DB) _blogDomain.Repository {
	return &BlogRepository{
		db: database,
	}
}

func (repo *BlogRepository) GetData(ctx context.Context) ([]_blogDomain.Domain, error) {
	blog := []Blogs{}
	result := repo.db.Find(&blog)
	if result.Error != nil {
		return []_blogDomain.Domain{}, result.Error
	}
	return ToListDomain(blog), nil
}
func (repo *BlogRepository) GetById(ctx context.Context, id uint) (_blogDomain.Domain, error) {
	blog := Blogs{}
	result := repo.db.Find(&blog, id)
	if result.Error != nil {
		return _blogDomain.Domain{}, result.Error
	}
	return ToDomain(blog), nil
}
func (repo *BlogRepository) Create(ctx context.Context, data _blogDomain.Domain) (_blogDomain.Domain, error) {
	blog := Blogs{
		Title:       data.Title,
		Content:     data.Content,
		PublishedAt: data.Published_at,
	}
	result := repo.db.Create(&blog)
	if result.Error != nil {
		return _blogDomain.Domain{}, result.Error
	}
	return ToDomain(blog), nil
}
func (repo *BlogRepository) Update(ctx context.Context, data _blogDomain.Domain) (_blogDomain.Domain, error) {
	blog := Blogs{}

	// result := repo.db.Model(&blog).Update(Blog{Title: data.Title, Content: data.Content, PublishedAt: data.Published_at})
	// if result.Error != nil {
	// 	return _blogDomain.Domain{}, result.Error
	// }
	repo.db.Find(&blog, data.Id)
	returndata := blog
	blog.Title = data.Title
	blog.Content = data.Content
	blog.PublishedAt = data.Published_at
	repo.db.Save(&blog)

	return ToDomain(returndata), nil
}
func (repo *BlogRepository) Delete(ctx context.Context, id uint) (_blogDomain.Domain, error) {
	var blog Blogs
	repo.db.Find(&blog, id)
	result := repo.db.Delete(&blog, id)
	if result.Error != nil {
		return _blogDomain.Domain{}, result.Error
	}
	return ToDomain(blog), nil
}
