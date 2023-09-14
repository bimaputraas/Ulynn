package repository

import "gorm.io/gorm"

// Repository contract
type User struct {
	DB *gorm.DB
}

type VideoGame struct {
	DB *gorm.DB
}

type Histories struct {
	DB *gorm.DB
}

type Verification struct {
	DB *gorm.DB
}