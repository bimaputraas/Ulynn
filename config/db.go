package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb()*gorm.DB{
	dsn := os.Getenv("DB")
 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
  