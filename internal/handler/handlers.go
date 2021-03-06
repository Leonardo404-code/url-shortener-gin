package handler

import (
	"fmt"
	"net/http"
	"url-shortener/pkg/database"
	"url-shortener/pkg/shortener"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Request struct {
	OriginalUrl string `json:"original_url" biding:"required"`
	UserId      string `json:"user_id" biding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest Request

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.OriginalUrl, creationRequest.UserId)

	database.SaveUrl(shortUrl, creationRequest.OriginalUrl, creationRequest.UserId)

	host := fmt.Sprintf("http://localhost:%s/get_short_url?shortUrl=", viper.GetString("PORT"))

	c.JSON(200, gin.H{
		"short_url": host + shortUrl,
	})
}

func ShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Query("shortUrl")

	initialUrl := database.RetrieveOriginalUrl(shortUrl)

	c.Redirect(302, initialUrl)
}
