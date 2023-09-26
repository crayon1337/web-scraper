package helper

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

// Scrape a URL using playwright
func ScrapeUrl(url string) Resturant {
	pw, err := playwright.Run()
	handleError("Could not start playwright", err)

	browser, err := pw.Chromium.Launch()
	handleError("Could not launch browser", err)

	page, err := browser.NewPage()
	handleError("Could not create page", err)

	if _, err = page.Goto(url); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}

	// Change timeout to 1 second instead of 30.
	page.SetDefaultTimeout(1000)

	// Populate the struct
	var resturant Resturant
	resturant.Name = getInnerTextBySelector(page, ".resturant-name .title")
	resturant.Rate = getInnerTextBySelector(page, ".vue-star-rating-rating-text")
	resturant.Address = getInnerTextBySelector(page, ".resturant-info-container .info-item .info-value")
	resturant.ReviewCount = getInnerTextBySelector(page, ".reviews")

	// Populate Menu struct
	categories, err := page.Locator(".cat-section").All()
	handleError("Could not find categories", err)

	for _, category := range categories {
		name, err := category.Locator(".section-title").InnerText()
		handleError("Could not locate category name", err)

		categoryStruct := Menu{}
		categoryStruct.Category = name

		// Populate Menu Item(s)
		items, err := category.Locator(".section-body .menu-item").All()
		handleError("Could not locate menu item(s)", err)

		for _, item := range items {
			itemStruct := Item{}

			// Item Name
			itemName, err := item.Locator(".item-header .title").InnerText()
			handleError("Could not locate menu item name", err)
			itemStruct.Name = itemName

			// Item Description
			itemDesc, err := item.Locator("p").First().InnerText()
			handleError("Could not locate item description", err)
			itemStruct.Description = itemDesc

			// Item Old Price
			itemOldPrice, err := item.Locator(".old-price").InnerText()

			if err == nil {
				itemStruct.OldPrice = itemOldPrice
			}

			itemPrice, err := item.Locator(".item-footer .price .bold").InnerText()

			if err == nil {
				itemStruct.Price = itemPrice
			}

			itemCurrency, err := item.Locator(".item-footer .currency").First().TextContent()

			if err == nil {
				itemStruct.Currency = itemCurrency
			}

			// Append item to category struct
			categoryStruct.Items = append(categoryStruct.Items, itemStruct)
		}

		// Append menu to resturant struct
		resturant.Menu = append(resturant.Menu, categoryStruct)
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("Could not close browser: %v", err)
	}

	if err = pw.Stop(); err != nil {
		log.Fatalf("Could not stop Playwright: %v", err)
	}

	return resturant
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
