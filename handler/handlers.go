package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/solarcreature/GoURL/algo"
	"github.com/solarcreature/GoURL/store"
	"net/http"
)

type URLRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}
func CreateShortURL(c *gin.Context) {
	var request URLRequest
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		
		return
	}

	shortURL := algo.GenerateURL(request.LongURL,request.UserID)
	store.StoreMap(shortURL,request.LongURL,request.UserID)

	host := "http://localhost:9808/"

	c.JSON(200,gin.H{
		"message":  	"URL shortened successfully",
		"short_url":	 host + shortURL,
	})
}

func RedirectShortURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	initialURL := store.RetrieveOriginalURL(shortURL)
	c.Redirect(302,initialURL)

}