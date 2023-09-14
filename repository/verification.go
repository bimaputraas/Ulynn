package repository

import (
	"videogames_rent_api/model"
)


func (r Verification) UpdateJwtToken(tokenString string,user model.Users) (*model.Users,error){
	user.JwtToken = tokenString
	result := r.DB.Save(&user)
	if result.Error != nil {
		return nil,result.Error
	}

	return &user,nil
}

func (r Verification) FindUserById(id interface{}) (*model.Users,error){
	var user model.Users
	result := r.DB.First(&user,id)
	if result.Error != nil {
		return nil,result.Error
	}

	return &user,nil

}