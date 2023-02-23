package usecases

import (
	"MyBlog/constants"
	"MyBlog/models"
	"MyBlog/repositories/interfaces"
	interfaces2 "MyBlog/usecases/interfaces"
	"MyBlog/utils"
	"github.com/labstack/echo/v4"
	"log"
)

type AuthUsecase struct {
	userRepository interfaces.UserRepository
}

func (a AuthUsecase) Register(email string, password string, password2 string, name string) (*models.User, *echo.HTTPError) {
	// Check if password is the same as password2
	if password != password2 {
		return nil, constants.ErrorPasswordConfirmation
	}

	// Check if email is already used
	user, err := a.userRepository.GetByEmail(email)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorInternalServer
	}
	if user != nil {
		return nil, constants.ErrorEmailUsed
	}

	// Hash password before stored to database
	hashedPassword, err := utils.Hash(password)
	if err != nil {
		return nil, constants.ErrorInternalServer
	}

	// Create new user with a given data
	newUser, err := a.userRepository.Add(email, hashedPassword, name)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorInternalServer
	}

	return newUser, nil
}

func (a AuthUsecase) Login(email string, password string) (*models.User, *echo.HTTPError) {
	// Check if user exist
	user, err := a.userRepository.GetByEmail(email)
	if err != nil {
		return nil, constants.ErrorInternalServer
	}
	if user == nil {
		return nil, constants.ErrorUnauthorized
	}

	// Check if password is correct
	if !utils.CheckHash(password, user.Password) {
		return nil, constants.ErrorUnauthorized
	}

	return user, nil
}

func NewAuthUsecase(userRepository interfaces.UserRepository) interfaces2.AuthUsecase {
	return &AuthUsecase{
		userRepository: userRepository,
	}
}
