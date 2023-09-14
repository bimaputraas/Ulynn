package repository

import (
	"videogames_rent_api/model"
)

// find all
func (r VideoGame) FindAll() (*[]model.VideoGames,error){
	var videoGames []model.VideoGames
	result := r.DB.Find(&videoGames)
	if result.Error != nil {
		return nil,result.Error
	}

	return &videoGames,nil
}

// find all with availability 
func (r VideoGame) FindAllWithAvailability(ok bool) (*[]model.VideoGames,error){
	var videoGames []model.VideoGames
	result := r.DB.Where("availability = ?",ok).Find(&videoGames)
	if result.Error != nil {
		return nil,result.Error
	}

	return &videoGames,nil
}

// find by id
func (r VideoGame) FindById(id int) (*model.VideoGames,error){
	var videoGame model.VideoGames
	result := r.DB.First(&videoGame,id)
	if result.Error != nil {
		return nil,result.Error
	}

	return &videoGame,nil
}




