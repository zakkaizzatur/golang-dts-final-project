package main

import (
	"os"

	"github.com/zakkaizzatur/golang-dts-final-project/database"
	"github.com/zakkaizzatur/golang-dts-final-project/router"
)

func main() {

	var PORT = os.Getenv("PORT")

	database.StartDB()
	r := router.StartApp()
	r.Run(":" + PORT)
}