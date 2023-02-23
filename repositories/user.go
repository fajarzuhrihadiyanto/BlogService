package repositories

import (
	"MyBlog/models"
	"MyBlog/repositories/interfaces"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u UserRepository) GetById(id uint) (*models.User, error) {
	// Get user by id
	var user models.User
	result := u.DB.First(&user, id)

	// Check if there is an error
	if result.Error != nil {

		// Check if record not found
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Println(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (u UserRepository) GetByEmail(email string) (*models.User, error) {
	// Get user by email
	var user models.User
	result := u.DB.Where("email = ?", email).First(&user)

	// Check if there is an error
	if result.Error != nil {

		// Check if record not found
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Println(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (u UserRepository) Add(email string, password string, name string) (*models.User, error) {
	// Create user
	user := models.User{Email: email, Password: password, Name: name, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := u.DB.Create(&user)

	// Check if there is an error
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (u UserRepository) Update(user *models.User) error {
	// Update user
	result := u.DB.Save(user)

	return result.Error
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepository{db}
}
