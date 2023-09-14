package model

type VideoGames struct {
	ID           int
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Availability bool    `json:"availability"`
	RentalCost   float64 `json:"rental_cost"`
}
