package main

import (
	"MyBlog/handlers"
	"MyBlog/repositories"
	"MyBlog/usecases"
	"MyBlog/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := utils.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("Get index")
		return nil
	})

	userRepository := repositories.NewUserRepository(db)
	articleRepository := repositories.NewArticleRepository(db)

	authUsecase := usecases.NewAuthUsecase(userRepository)
	articleUsecase := usecases.NewArticleUsecase(userRepository, articleRepository)

	handlers.NewAuthHandler(e, authUsecase)
	handlers.NewArticleHandler(e, articleUsecase)

	err = e.Start(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
