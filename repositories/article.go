package repositories

import (
	"MyBlog/models"
	"MyBlog/repositories/interfaces"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func (a ArticleRepository) Fetch(limit uint, orderColumn string, orderType string, where string) (*[]models.Article, error) {
	var articles []models.Article
	result := a.DB

	query := "SELECT articles.*, users.name AS author_name FROM articles INNER JOIN users ON users.id = articles.author_id"

	if where != "" {
		query = fmt.Sprintf("%v WHERE %v", query, where)
	}

	query = fmt.Sprintf("%v ORDER BY %v %v LIMIT %v", query, orderColumn, orderType, limit)

	result = result.Raw(query).Scan(&articles)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		log.Println(result.Error)
		return nil, result.Error
	}

	return &articles, nil
}

func (a ArticleRepository) GetById(id uint) (*models.Article, error) {
	// Get article by id
	var article models.Article
	result := a.DB.Select("articles.*, users.name AS author_name").Joins("INNER JOIN users ON users.id = articles.author_id").First(&article, id).Scan(&article)

	// Check if there is an error
	if result.Error != nil {

		// Check if record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		log.Println(result.Error)
		return nil, result.Error
	}

	return &article, nil
}

func (a ArticleRepository) Add(authorId uint, title string, content string) (*models.Article, error) {
	article := models.Article{
		AuthorId:  authorId,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := a.DB.Create(&article)

	// Check if there is an error
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &article, nil
}

func (a ArticleRepository) Update(article *models.Article) error {
	// Update article
	result := a.DB.Save(article)

	return result.Error
}

func (a ArticleRepository) Delete(id uint) error {
	result := a.DB.Delete(&models.Article{}, id)

	return result.Error
}

func NewArticleRepository(db *gorm.DB) interfaces.ArticleRepository {
	return &ArticleRepository{db}
}
