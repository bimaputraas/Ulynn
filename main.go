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

// // simpan dana topup ke db golang
// _, err = db.Exec("INSERT INTO topups (xendit_id, external_id, amount, bank_code, account_name) VALUES (?, ?, ?, ?, ?)",
// topUp.ID, topUp.ExternalID, topUp.Amount, topUp.BankCode, topUp.AccountName)
// if err != nil {
// log.Fatal("Error inserting top-up data into the database:", err)
// }