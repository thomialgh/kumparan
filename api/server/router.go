package server

import (
	"kumparan/api/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func handler() http.Handler {
	c := echo.New()
	c.Use(middleware.Logger())
	c.POST("/", controllers.CreateArticle)
	c.GET("/", controllers.GetArticles, controllers.CacheMiddlerwere)
	return c
}
