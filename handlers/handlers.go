package handlers

import (
	"net/http"

	"github.com/evilyach/url-shortener/shortener"
	"github.com/evilyach/url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortURL(c *gin.Context) {
	var creationRequest UrlCreationRequest

	err := c.ShouldBindJSON(&creationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortURL, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url cereated successfully",
		"short_url": host + shortURL,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortURL := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortURL)
	c.Redirect(302, initialUrl)
}
