package main

import (
	"github.com/aleroxac/alura-golang-gin/database"
	"github.com/aleroxac/alura-golang-gin/routes"
)

func main() {
	database.InitDatabaseConnection()
	routes.HandleRequests()
}
