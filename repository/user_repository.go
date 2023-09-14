package repository

import (
	"videogames_rent_api/model"
)

func (r User) Create(reqBody model.ReqBodyUserRegister) (*model.Users,error){ 
	user := model.Users{
		Email: reqBody.Email,
		Password: reqBody.Password,
		Name: reqBody.Name,
		Status: "Pending Activation",
		DepositAmount: 0,
	}
	result := r.DB.Create(&user)
	if result.Error != nil {
		return nil,result.Error
	}

	return &user,nil
}

func (r User) CreateVerificationCode(userId int,codeString string) error{ 
	userVerification := model.UserVerifications{
		UserID: userId,
		VerifyCode: codeString,
	}
	result := r.DB.Create(&userVerification)
	if result.Error != nil {
		return nil
	}

	return nil
}

func (r User) FindByEmail(email string) (*model.Users,error){
	var user model.Users
	result := r.DB.Where("email = ?",email).First(&user)
	if result.Error != nil {
		return nil,result.Error
	}

	return &user,nil
}

func (r User) FindById(id int) (*model.Users,error){
	var user model.Users
	result := r.DB.First(&user,id)
	if result.Error != nil {
		return nil,result.Error
	}

	return &user,nil
}

func (r User) FindUserVerificationByUserId(userId int) (*model.UserVerifications,error){
	var userVerif model.UserVerifications
	result := r.DB.Where("user_id = ?",userId).First(&userVerif)
	if result.Error != nil {
		return nil,result.Error
	}

	return &userVerif,nil
}


func (r User) UpdateStatusActivated(user model.Users) (*model.Users,error){
	user.Status = "Activated"
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

