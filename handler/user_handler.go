package handler

import (
	"strconv"
	"videogames_rent_api/helper"
	"videogames_rent_api/model"

	"github.com/labstack/echo/v4"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.ReqBodyUserRegister true "User registration details"
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Failure 500 {object} helper.ErrResponse
// @Router /users/register [post]
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

// StatusVerification godoc
// @Summary Verify user's account
// @Description Verify user's account by providing a verification code
// @Tags User
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param code path string true "Verification code"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Router /users/verify/{userId}/{code} [get]
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

		// update status user
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

// Login godoc
// @Summary User login
// @Description Log in a user with the provided email and password
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.ReqBodyUserLogin true "User login details"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 403 {object} helper.Response
// @Router /users/login [post]
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
		return helper.ErrorResponse(400, "wrong email or password")
	}

	// compare hash
	if !helper.CheckPasswordHash(reqBody.Password,user.Password){
		return helper.ErrorResponse(400, "wrong email or password")
	}
	c.Set("user",user)
	return nil
}

// GetInfo godoc
// @Summary Get logged-in user information
// @Description Get information about the currently logged-in user
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} helper.Response
// @Failure 401 {object} helper.ErrResponse
// @Router /users/info [get]
func (h User) GetInfo(c echo.Context) error {
	// get logged in user
	user := c.Get("user").(*model.Users)

	// success
	helper.WriteResponseWithData(c,200,"Success get logged in user information",user)
	return nil
}

// TopUp godoc
// @Summary Top up user's deposit
// @Description Top up the deposit of the logged-in user
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param user body model.ReqBodyUserTopUp true "Top-up details"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Failure 401 {object} helper.ErrResponse
// @Router /users/topup [put]
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

