package interfaces

import (
	"MyBlog/models"
	"github.com/labstack/echo/v4"
)

type ArticleUsecase interface {
	Fetch(limit uint, orderColumn string, orderType string, where string) (*[]models.Article, *echo.HTTPError)
	GetById(id uint) (*models.Article, *echo.HTTPError)
	Add(userId uint, title string, content string) (*models.Article, *echo.HTTPError)
	Update(userId uint, articleId uint, title string, content string) (*models.Article, *echo.HTTPError)
	Delete(userId uint, articleId uint) *echo.HTTPError
}
