package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/solarcreature/GoURL/store"
	"github.com/solarcreature/GoURL/handler"
)

func main() {
	r := gin.Default()
	r.GET("/",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to GoURL!",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context){
		handler.CreateShortURL(c)
	})

	r.GET("/:shortURL", func(c *gin.Context){
		handler.RedirectShortURL(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web server - Error: %v", err))
	}
}