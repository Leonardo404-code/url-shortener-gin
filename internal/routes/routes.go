package routes

import (
	"url-shortener/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello")
	})

	r.POST("/create_short_url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/get_short_url", func(c *gin.Context) {
		handler.ShortUrlRedirect(c)
	})

	return r
}
