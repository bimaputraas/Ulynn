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

type History struct {
	Repository repository.History
}

type Rent struct {
	Repository repository.Rent
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

func InitHistoryHandler(db *gorm.DB) *History{
	return &History{Repository : repository.History{
		DB : db,
	}}
}

func InitRentHandler(db *gorm.DB) *Rent{
	return &Rent{Repository : repository.Rent{
		DB : db,
	}}
}

