package interfaces

import "MyBlog/models"

type UserRepository interface {
	GetById(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Add(email string, password string, name string) (*models.User, error)
	Update(user *models.User) error
}
