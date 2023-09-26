package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
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

type Resturant struct {
	Name         string
	ReviewsCount string
	Address      string
	Menu         []Menu
}

type Menu struct {
	Category string
	Items    []Item
}

type Item struct {
	Name        string
	Description string
	OldPrice    string
	Price       string
	Currency    string
}

func scrape(ctx *gin.Context) {
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

	collyCollector := colly.NewCollector(
		colly.AllowedDomains("www.elmenus.com"),
	)

	resturant := Resturant{}

	// Resturant Information will be visible in the resturant-info-container div.
	collyCollector.OnHTML(".resturant-info-container", func(element *colly.HTMLElement) {
		// Split the string into an array by using \n as delimiter then pick the first element which is the wanted string.
		resturant.Name = strings.Split(element.ChildText("h1[class=title]"), "\n")[0]
		resturant.Address = strings.Split(element.ChildText("p[class=info-value]"), "\n")[0]
	})

	collyCollector.OnHTML(".rest-rate", func(element *colly.HTMLElement) {
		resturant.ReviewsCount = element.ChildText("a")
	})

	// TODO: Append Menu & Items to resturant's Menu struct.

	collyCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitng", r.URL)
	})

	collyCollector.Visit(url)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data": resturant,
	})
}
