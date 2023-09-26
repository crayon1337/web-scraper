package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/scrape", scrape)

	router.Run("localhost:8000")
}

func scrape(ctx *gin.Context) {
	url := ctx.Request.URL.Query().Get("url")

	if url == "" {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Please enter a url"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": url})
}
