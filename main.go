package main

import (
	"github.com/ravelinejunior/go_api_gin/database"
	"github.com/ravelinejunior/go_api_gin/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}
