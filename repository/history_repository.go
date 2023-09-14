package repository

import (
	"videogames_rent_api/model"
)

// find all
func (r History) FindAll(userId int) (*[]model.Histories,error){
	var histories []model.Histories
	result := r.DB.Where("user_id = ?",userId).Find(&histories)
	if result.Error != nil {
		return nil,result.Error
	}

	return &histories,nil
}

// find by id
func (r History) FindById(historyId,userId int) (*model.Histories,error){
	var history model.Histories
	result := r.DB.Where("user_id = ?",userId).First(&history,historyId)
	if result.Error != nil {
		return nil,result.Error
	}

	return &history,nil
}
