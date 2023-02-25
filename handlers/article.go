package handlers

import (
	"MyBlog/constants"
	"MyBlog/models"
	_struct "MyBlog/struct"
	"MyBlog/usecases/interfaces"
	"MyBlog/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ArticleHandler struct {
	ArticleUsecase interfaces.ArticleUsecase
}

func (handler ArticleHandler) Get(ctx echo.Context) error {

	var limit uint = 10
	var orderBy = "id"
	var orderType = "asc"
	var where = ""

	errs := echo.QueryParamsBinder(ctx).
		Uint("limit", &limit).
		String("order_by", &orderBy).
		String("order", &orderType).
		String("where", &where).
		BindErrors()

	if errs != nil {
		var errors []_struct.HTTPFieldError
		for _, err := range errs {
			log.Println(err)
			errors = append(errors, _struct.HTTPFieldError{
				Message: err.Error(),
			})
		}
		return ctx.JSON(constants.ErrorInternalServer.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorInternalServer.Message),
			Errors:  errors,
		})
	}

	articles, err := handler.ArticleUsecase.Fetch(limit, orderBy, orderType, where)
	if err != nil {
		return ctx.JSON(err.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(err.Message),
		})
	}
	if articles == nil {
		return ctx.JSON(http.StatusNotFound, &_struct.HTTPBodyResponse{
			Message: "Articles not found",
		})
	}

	return ctx.JSON(http.StatusOK, &_struct.HTTPBodyResponse{
		Message: "Articles found",
		Data:    articles,
	})
}

func (handler ArticleHandler) GetById(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Println(err)
		return ctx.JSON(constants.ErrorInternalServer.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorInternalServer.Message),
		})
	}

	article, httpError := handler.ArticleUsecase.GetById(uint(id))
	if httpError != nil {
		return ctx.JSON(httpError.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(httpError.Message),
		})
	}
	if article == nil {
		return ctx.JSON(http.StatusNotFound, &_struct.HTTPBodyResponse{
			Message: "Article not found",
		})
	}

	return ctx.JSON(http.StatusOK, &_struct.HTTPBodyResponse{
		Message: "Article found",
		Data:    article,
	})
}

func (handler ArticleHandler) Add(ctx echo.Context) error {
	body := new(models.Article)
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

	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		// Return unauthorized error
		return ctx.JSON(constants.ErrorUnauthorized.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorUnauthorized.Message),
		})
	}
	authToken := strings.Split(authHeader, " ")[1]

	claims, _, err := utils.VerifyToken(authToken)
	if err != nil {
		log.Println(err)
		return ctx.JSON(constants.ErrorInternalServer.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorInternalServer.Error()),
		})
	}

	userId, _ := strconv.ParseUint(fmt.Sprint(claims["user_id"]), 10, 32)

	newArticle, httpError := handler.ArticleUsecase.Add(uint(userId), body.Title, body.Content)
	if httpError != nil {
		return ctx.JSON(httpError.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(httpError.Message),
		})
	}

	return ctx.JSON(http.StatusCreated, &_struct.HTTPBodyResponse{
		Message: "Article Created",
		Data:    newArticle,
	})
}

func (handler ArticleHandler) Update(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Println(err)
		return ctx.JSON(constants.ErrorInternalServer.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorInternalServer.Message),
		})
	}

	body := new(models.ArticleUpdate)
	err = ctx.Bind(body)
	if err != nil {
		log.Println(err)
		return ctx.JSON(constants.ErrorInternalServer.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorInternalServer.Error()),
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

	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		// Return unauthorized error
		return ctx.JSON(constants.ErrorUnauthorized.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorUnauthorized.Message),
		})
	}
	authToken := strings.Split(authHeader, " ")[1]

	claims, _, err := utils.VerifyToken(authToken)
	if err != nil {
		log.Println(err)
		return ctx.JSON(constants.ErrorInternalServer.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorInternalServer.Error()),
		})
	}
	userId, _ := strconv.ParseUint(fmt.Sprint(claims["user_id"]), 10, 32)

	article, httpError := handler.ArticleUsecase.Update(uint(userId), uint(id), body.Title, body.Content)
	if httpError != nil {
		return ctx.JSON(httpError.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(httpError.Message),
		})
	}

	return ctx.JSON(http.StatusOK, &_struct.HTTPBodyResponse{
		Message: "Article updated",
		Data:    article,
	})
}

func (handler ArticleHandler) Delete(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Println(err)
		return ctx.JSON(constants.ErrorInternalServer.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorInternalServer.Message),
		})
	}

	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		// Return unauthorized error
		return ctx.JSON(constants.ErrorUnauthorized.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorUnauthorized.Message),
		})
	}
	authToken := strings.Split(authHeader, " ")[1]

	claims, _, err := utils.VerifyToken(authToken)
	if err != nil {
		log.Println(err)
		return ctx.JSON(constants.ErrorInternalServer.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(constants.ErrorInternalServer.Error()),
		})
	}
	userId, _ := strconv.ParseUint(fmt.Sprint(claims["user_id"]), 10, 32)

	httpError := handler.ArticleUsecase.Delete(uint(userId), uint(id))
	if httpError != nil {
		return ctx.JSON(httpError.Code, &_struct.HTTPBodyResponse{
			Message: fmt.Sprint(httpError.Message),
		})
	}

	return ctx.JSON(http.StatusOK, &_struct.HTTPBodyResponse{
		Message: "Article deleted",
	})
}

func NewArticleHandler(e *echo.Echo, articleUsecase interfaces.ArticleUsecase) {
	handler := &ArticleHandler{
		ArticleUsecase: articleUsecase,
	}

	e.GET("/article", handler.Get)
	e.GET("/article/:id", handler.GetById)
	e.POST("/article", handler.Add)
	e.PUT("/article/:id", handler.Update)
	e.DELETE("/article/:id", handler.Delete)
}
