package handler

import (
	"strconv"
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

	// set verification code
	codeString := helper.GenerateCode()
	err = h.Repository.CreateVerificationCode(user.ID,codeString)
	if err != nil {
		return helper.ErrorResponse(500, err.Error())
	}

	// send notif and code verification by email 
	err = helper.SendMail(user.Email,user.ID,codeString)
	if err != nil {
		return helper.ErrorResponse(500, err.Error())
	}

	// success
	helper.WriteResponseWithData(c,201,"Registration successful! Your verify link has been sent to your email",user)
	return nil
}

func (h User) StatusVerification(c echo.Context) error {
	// param path
	idStr := c.Param("userId")
	codeStr := c.Param("code")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		return helper.ErrorResponse(400, "failed verification")
	}

	// get user verification code
	userVerif,err := h.Repository.FindUserVerificationByUserId(id)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// if match
	if userVerif.VerifyCode == codeStr {
		// get user code from db
		user,err := h.Repository.FindById(id)
		if err != nil {
			return helper.ErrorResponse(400, err.Error())
		}

		// update status
		updatedUser, err := h.Repository.UpdateStatusActivated(*user)
		if err != nil {
			return helper.ErrorResponse(400, err.Error())
		}

		// success
		helper.WriteResponseWithData(c,200,"Your account has been successfully verified. You can now access our services.",updatedUser)
		return nil
	}

	return helper.ErrorResponse(400, "failed verification")
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
	c.Set("user",user)
	return nil
}

func (h User) GetInfo(c echo.Context) error {
	// get logged in user
	user := c.Get("user").(*model.Users)

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
	user := c.Get("user").(*model.Users)

	// update
	updatedUser,err := h.Repository.UpdateDepositAmount(reqBody.TopUpAmount,*user)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// success
	helper.WriteResponseWithData(c,201,"Success top-up",updatedUser)
	return nil
}

