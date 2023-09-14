package middleware

import (
	"videogames_rent_api/repository"

	"gorm.io/gorm"
)

type Verification struct {
	Repository repository.Verification
}

func InitVerification(db *gorm.DB) *Verification{
	return &Verification{
		Repository: repository.Verification{
			DB: db,
		},
	}
}
