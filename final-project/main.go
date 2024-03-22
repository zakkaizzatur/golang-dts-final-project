package main

import (
	"final-project/database"
	"final-project/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}