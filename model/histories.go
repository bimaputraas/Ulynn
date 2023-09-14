package model

import "time"

type Histories struct {
	ID           	int
	UserID       	int       	`json:"user_id"`
	VideoGameID 	int       	`json:"video_game_id"`
	StartDate		time.Time 	`json:"start_date"`
	DueDate			time.Time 	`json:"due_date"`
	Status 			string      `json:"status"`
}