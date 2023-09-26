package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/webscraper/helper"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("scrape", handleScrapeEndpoint)
		}
	}

	router.Run("localhost:8000")
}

func handleScrapeEndpoint(ctx *gin.Context) {
	url := ctx.Request.URL.Query().Get("url")

	if url == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Please specify a url in the query parameters",
		})
		return
	}

	if !strings.HasPrefix(url, "https://www.elmenus.com") {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "The only data provider allowed right now is elmenus.com",
		})
		return
	}

	resturant := helper.ScrapeUrl(url)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data": resturant,
	})
}
