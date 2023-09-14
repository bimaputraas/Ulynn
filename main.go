package main

import (
	"videogames_rent_api/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// init db
	db := config.InitDb()
	
	// init echo instance
	e := config.InitEchoInstance(db)
	
	// run echo on localhost port 8080
	e.Logger.Fatal(e.Start(":8080"))
}