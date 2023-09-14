package handler

import (
	"videogames_rent_api/helper"
	"videogames_rent_api/model"

	"github.com/labstack/echo/v4"
)

func (h User) Register(c echo.Context) error {
	// bind
	var reqBody model.ReqBodyUserRegister
	err := c.Bind(&reqBody)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// validate
	err = helper.Validate(reqBody)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// hash
	hash,err := helper.HashPassword(reqBody.Password)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}
	reqBody.Password = hash

	// create
	user,err := h.Repository.Create(reqBody)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// send notif by email
	err = helper.SendMail(user.Email)
	if err != nil {
		return helper.ErrorResponse(500, err.Error())
	}

	// success
	helper.WriteResponseWithData(c,201,"Registration successful! A notification has been successfully sent to the email address "+user.Email,user)
	return nil
}

func (h User) Login(c echo.Context) error {
	// bind
	var reqBody model.ReqBodyUserLogin
	err := c.Bind(&reqBody)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// validate
	err = helper.Validate(reqBody)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// find user
	user,err := h.Repository.FindByEmail(reqBody.Email)
	if err != nil {
		return helper.ErrorResponse(400, "Wrong email or password")
	}

	// compare hash
	if !helper.CheckPasswordHash(reqBody.Password,user.Password){
		return helper.ErrorResponse(400, "Wrong email or password")
	}

	// generate token
	tokenString,err := helper.GenerateToken(user.ID)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}
	
	// update user jwt token
	user,err = h.Repository.UpdateJwtToken(tokenString,*user)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// success
	helper.WriteResponseWithData(c,201,"Success login",user)
	return nil
}

func (h User) GetInfo(c echo.Context) error {
	// get logged in user
	user := c.Get("user").(model.Users)

	// success
	helper.WriteResponseWithData(c,200,"Success get logged in user information",user)
	return nil
}

func (h User) TopUp(c echo.Context) error {
	// bind
	var reqBody model.ReqBodyUserTopUp
	err := c.Bind(&reqBody)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// validate
	err = helper.Validate(reqBody)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// get logged in user
	user := c.Get("user").(model.Users)

	// update
	updatedUser,err := h.Repository.UpdateDepositAmount(reqBody.TopUpAmount,user)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// success
	helper.WriteResponseWithData(c,201,"Success top-up",updatedUser)
	return nil
}

