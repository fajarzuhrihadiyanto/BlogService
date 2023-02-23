package handlers

import (
	"MyBlog/constants"
	"MyBlog/models"
	_struct "MyBlog/struct"
	"MyBlog/usecases/interfaces"
	"MyBlog/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type AuthHandler struct {
	AuthUseCase interfaces.AuthUsecase
}

func (handler AuthHandler) Login(ctx echo.Context) error {
	// Parse body request
	body := new(models.UserLogin)
	err := ctx.Bind(body)
	if err != nil {
		return constants.ErrorDataValidation
	}

	// Body validation
	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		log.Println(err)
		var errors []_struct.HTTPFieldError

		// Iterate through field errors
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, _struct.HTTPFieldError{
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}

		// Return validation error
		return ctx.JSON(constants.ErrorDataValidation.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorDataValidation.Message),
			Errors:  errors,
		})
	}

	// Do login
	user, httpError := handler.AuthUseCase.Login(body.Email, body.Password)
	if httpError != nil {
		log.Println(httpError)

		// Return another error
		return ctx.JSON(httpError.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(httpError.Message),
		})
	}

	// Create JWT
	tokenStr, err := utils.CreateToken(jwt.MapClaims{
		"user_id": strconv.Itoa(int(user.ID)),
	})

	type Data struct {
		Token string `json:"token"`
	}

	// Return data
	return ctx.JSON(http.StatusOK, &_struct.HTTPBodyResponse{
		Message: "User logged in successfully",
		Data: &Data{
			Token: tokenStr,
		},
	})
}

func (handler AuthHandler) Register(ctx echo.Context) error {
	// Parse body request
	body := new(struct {
		models.User
		Password2 string `json:"password2" validate:"required"`
	})
	err := ctx.Bind(body)
	if err != nil {
		log.Println(err)
		return ctx.JSON(500, &_struct.HTTPBodyResponse{
			Message: constants.ErrorInternalServer.Error(),
		})
	}

	// Body Validation
	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		log.Println(err)
		var errors []_struct.HTTPFieldError

		// Iterate through field errors
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, _struct.HTTPFieldError{
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}

		// Return validation error
		return ctx.JSON(constants.ErrorDataValidation.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorDataValidation.Message),
			Errors:  errors,
		})
	}

	// Register user
	newUser, httpError := handler.AuthUseCase.Register(body.Email, body.Password, body.Password2, body.Name)
	if httpError != nil {
		log.Println(httpError)

		// Return another error
		return ctx.JSON(httpError.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(httpError.Message),
		})
	}

	// Return data
	return ctx.JSON(http.StatusOK, &_struct.HTTPBodyResponse{
		Message: "User registered successfully",
		Data:    newUser,
	})
}

func NewAuthHandler(e *echo.Echo, usecase interfaces.AuthUsecase) {
	handler := &AuthHandler{
		AuthUseCase: usecase,
	}
	e.POST("auth/login", handler.Login)
	e.POST("auth/register", handler.Register)
}
