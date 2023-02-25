package usecases

import (
	"MyBlog/constants"
	"MyBlog/models"
	"MyBlog/repositories/interfaces"
	interfaces2 "MyBlog/usecases/interfaces"
	"github.com/labstack/echo/v4"
	"log"
)

type ArticleUsecase struct {
	userRepository    interfaces.UserRepository
	articleRepository interfaces.ArticleRepository
}

func (a ArticleUsecase) Fetch(limit uint, orderColumn string, orderType string, where string) (*[]models.Article, *echo.HTTPError) {
	articles, err := a.articleRepository.Fetch(limit, orderColumn, orderType, where)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorInternalServer
	}
	if articles == nil || len(*articles) == 0 {
		return nil, constants.ErrorNotFound
	}

	return articles, nil
}

func (a ArticleUsecase) GetById(id uint) (*models.Article, *echo.HTTPError) {
	article, err := a.articleRepository.GetById(id)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorInternalServer
	}
	if article == nil {
		return nil, constants.ErrorNotFound
	}
	return article, nil
}

func (a ArticleUsecase) Add(userId uint, title string, content string) (*models.Article, *echo.HTTPError) {
	article, err := a.articleRepository.Add(userId, title, content)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorInternalServer
	}
	return article, nil
}

func (a ArticleUsecase) Update(userId uint, articleId uint, title string, content string) (*models.Article, *echo.HTTPError) {
	article, err := a.articleRepository.GetById(articleId)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorInternalServer
	}
	if article == nil {
		return nil, constants.ErrorNotFound
	}

	if article.AuthorId != userId {
		return nil, constants.ErrorUnauthorized
	}

	if title != "" {
		article.Title = title
	}

	if content != "" {
		article.Content = content
	}

	err = a.articleRepository.Update(article)
	if err != nil {
		return nil, constants.ErrorInternalServer
	}

	return article, nil
}

func (a ArticleUsecase) Delete(userId uint, articleId uint) *echo.HTTPError {
	article, err := a.articleRepository.GetById(articleId)
	if err != nil {
		log.Println(err)
		return constants.ErrorInternalServer
	}
	if article == nil {
		return constants.ErrorNotFound
	}

	if article.AuthorId != userId {
		return constants.ErrorUnauthorized
	}

	err = a.articleRepository.Delete(articleId)
	if err != nil {
		log.Println(err)
		return constants.ErrorInternalServer
	}
	return nil
}

func NewArticleUsecase(userRepository interfaces.UserRepository, articleRepository interfaces.ArticleRepository) interfaces2.ArticleUsecase {
	return &ArticleUsecase{
		userRepository:    userRepository,
		articleRepository: articleRepository,
	}
}
