package main

import (
	"go-web-scraper/api"
	"go-web-scraper/database"
	"go-web-scraper/scraper"
)

func main() {
	api.HandleRequest()
	database.ConnectDB()
	scraper.StartScraper()
}
