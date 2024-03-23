package main

import (
	"os"

	"github.com/zakkaizzatur/golang-dts-final-project/database"
	"github.com/zakkaizzatur/golang-dts-final-project/router"

	_ "github.com/zakkaizzatur/golang-dts-final-project/docs"
)

// @title MyGram API
// @version 1.0
// @description In this application you can save photos or make comments on other people's photos
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email zakka.izzatur@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /

func main() {

	var PORT = os.Getenv("PORT")

	database.StartDB()
	r := router.StartApp()
	r.Run(":" + PORT)
}