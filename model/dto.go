package model

import "time"

// users

type ReqBodyUserRegister struct {
	Name     string `json:"name" `
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ReqBodyUserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ReqBodyUserTopUp struct {
	TopUpAmount    float64 `json:"top_up_amount" validate:"required"`
}

type ReqBodyVideoGameRent struct {
	VideoGamesID    float64 `json:"video_game_id" validate:"required"`
}

type ResBodyHistoryVideoGame struct {
	Title string `json:"video_game_title"`
	StartDate time.Time `json:"start_date"`
	DueDate time.Time `json:"due_date"`
	Status string `json:"status"`
}

