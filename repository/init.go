package repository

import "gorm.io/gorm"

// Repository contract
type User struct {
	DB *gorm.DB
}

type VideoGame struct {
	DB *gorm.DB
}

type History struct {
	DB *gorm.DB
}

type Rent struct {
	DB *gorm.DB
}