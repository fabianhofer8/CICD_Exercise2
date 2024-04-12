// main.go

package main

import (
	"os"
)

var app App

func main() {
	println("Started main")
	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	app.Run(":8010")

}
