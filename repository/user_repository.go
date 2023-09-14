package repository

import (
	"videogames_rent_api/model"
)

func (r User) Create(reqBody model.ReqBodyUserRegister) (*model.Users,error){ 
	user := model.Users{
		Email: reqBody.Email,
		Password: reqBody.Password,
		Name: reqBody.Name,
	}
	result := r.DB.Create(&user)
	if result.Error != nil {
		return nil,result.Error
	}

	return &user,nil
}

func (r User) FindByEmail(email string) (*model.Users,error){
	var user model.Users
	result := r.DB.Where("email = ?",email).First(&user)
	if result.Error != nil {
		return nil,result.Error
	}

	return &user,nil
}

func (r User) UpdateJwtToken(tokenString string,user model.Users) (*model.Users,error){
	user.JwtToken = tokenString
	result := r.DB.Save(&user)
	if result.Error != nil {
		return nil,result.Error
	}

	return &user,nil
}

func (r User) UpdateDepositAmount(reqAmount float64,user model.Users) (*model.Users,error){
	user.DepositAmount += reqAmount
	result := r.DB.Save(&user)
	if result.Error != nil {
		return nil,result.Error
	}

	return &user,nil
}

