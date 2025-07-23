package main

import (
	"book-ddd/config"
	router "book-ddd/interfaces/http"
)

func main() {
	config.ConnectDB()

	app := router.InitRouter(config.DB)
	app.Listen(":8080")
}
