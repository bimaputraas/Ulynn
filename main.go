package main

import (
	"videogames_rent_api/config"
	"videogames_rent_api/router"

	_ "github.com/joho/godotenv/autoload"
)

// @title Documentation Video Game Rent API
// @version 1.0
// @description Documentation Video Game Rent API using Swaggo/echo-swagger.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host http://localhost:8080
func main() {
	// init db
	db := config.InitDb()
	
	// init echo instance
	e := router.InitEchoInstance(db)
	
	// run echo on localhost port 8080
	e.Logger.Fatal(e.Start(":8080"))
}

// // simpan dana topup ke db golang
// _, err = db.Exec("INSERT INTO topups (xendit_id, external_id, amount, bank_code, account_name) VALUES (?, ?, ?, ?, ?)",
// topUp.ID, topUp.ExternalID, topUp.Amount, topUp.BankCode, topUp.AccountName)
// if err != nil {
// log.Fatal("Error inserting top-up data into the database:", err)
// }