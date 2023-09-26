package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
	"github.com/webscrapper/helper"
)

func main() {
	pw, err := playwright.Run()
	handleError("Could not start playwright", err)

	browser, err := pw.Chromium.Launch()
	handleError("Could not launch browser", err)

	page, err := browser.NewPage()
	handleError("Could not create page", err)

	if _, err = page.Goto("https://www.elmenus.com/cairo/bazooka-myp8d"); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}

	// Populate the struct
	var resturant helper.Resturant
	resturant.Name = getInnerTextBySelector(page, ".resturant-name .title")
	resturant.Rate = getInnerTextBySelector(page, ".vue-star-rating-rating-text")
	resturant.Address = getInnerTextBySelector(page, ".resturant-info-container .info-item .info-value")
	resturant.ReviewCount = getInnerTextBySelector(page, ".reviews")

	if err = browser.Close(); err != nil {
		log.Fatalf("Could not close browser: %v", err)
	}

	if err = pw.Stop(); err != nil {
		log.Fatalf("Could not stop Playwright: %v", err)
	}

	fmt.Println(resturant)
}

func getInnerTextBySelector(page playwright.Page, selector string) string {
	text, err := page.Locator(selector).First().InnerText()
	handleError("Could not get text for element "+selector, err)

	return text
}

func handleError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
