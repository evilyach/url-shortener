package main

import (
	"fmt"

	"github.com/evilyach/url-shortener/handlers"
	"github.com/evilyach/url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API.",
		})
	})

	r.POST("/create", func(c *gin.Context) {
		handlers.CreateShortURL(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handlers.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web server: '%v'", err))
	}
}
