package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("scrape", scrape)
		}
	}

	router.Run("localhost:8000")
}

func scrape(ctx *gin.Context) {
	url := ctx.Request.URL.Query().Get("url")

	if url == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Please specify a url in the query parameters",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data": url,
	})
}
