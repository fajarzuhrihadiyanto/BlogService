package interfaces

import "MyBlog/models"

type ArticleRepository interface {
	Fetch(limit uint, orderColumn string, orderType string, where string) (*[]models.Article, error)
	GetById(id uint) (*models.Article, error)
	Add(author *models.User, title string, content string) (*models.Article, error)
	Update(article *models.Article) error
	Delete(id uint) error
}
