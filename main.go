package main

import (
	"api"
	"database"
)

func main() {
	api.HandleRequest()
	database.ConnectDB()
}

