package model

import "time"

type Histories struct {
	ID           	int
	UserID       	int       	`json:"user_id"`
	VideoGameID 	int       	`json:"video_game_id"`
	VideoGame 		VideoGames  `json:"video_game" gorm:"foreignKey:VideoGameID"`
	StartDate		time.Time 	`json:"start_date"`
	DueDate			time.Time 	`json:"due_date"`
	TotalRentalCost float64		`json:"totan_rental_cost"`
	Status 			string      `json:"status"`
}