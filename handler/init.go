package handler

import (
	"videogames_rent_api/repository"

	"gorm.io/gorm"
)

// handler contract
type User struct {
	Repository repository.User
}

type VideoGame struct {
	Repository repository.VideoGame
}

type Histories struct {
	Repository repository.Histories
}

// init handler
func InitUserHandler(db *gorm.DB) *User{
	return &User{Repository : repository.User{
		DB : db,
	}}
}

func InitVideoGameHandler(db *gorm.DB) *VideoGame{
	return &VideoGame{Repository : repository.VideoGame{
		DB : db,
	} }
}

func InitHistoriesHandler(db *gorm.DB) *Histories{
	return &Histories{Repository : repository.Histories{
		DB : db,
	}}
}

