package interfaces

import (
	"MyBlog/models"
	"github.com/labstack/echo/v4"
)

type AuthUsecase interface {
	Register(email string, password string, password2 string, name string) (*models.User, *echo.HTTPError)
	Login(email string, password string) (*models.User, *echo.HTTPError)
}
