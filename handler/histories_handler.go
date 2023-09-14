package handler

import (
	"strconv"
	"videogames_rent_api/helper"
	"videogames_rent_api/model"

	"github.com/labstack/echo/v4"
)

// AddRent godoc
// @Summary Add a new rental
// @Description Add a new rental with the provided information
// @Tags Histories
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header with token"
// @Param user body model.ReqBodyVideoGameRent true "Rental details"
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Failure 401 {object} helper.ErrResponse
// @Router /histories/rent [post]
func (h Histories) AddRent(c echo.Context) error {
	// bind
	var reqBody model.ReqBodyVideoGameRent
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

	// update user,video_game, and create a new history
	resBody, err := h.Repository.UpdateCreate(int(reqBody.VideoGamesID), *user)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// success
	helper.WriteResponseWithData(c, 201, "Success add a new rent", resBody)
	return nil
}
// UpdateRent godoc
// @Summary Update a rental status
// @Description Update the status of a rental by ID
// @Tags Histories
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header with token"
// @Param id path string true "Rental ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Failure 401 {object} helper.ErrResponse
// @Router /histories/rent/{id} [put]
func (h Histories) UpdateRent(c echo.Context) error {
	// get logged in user
	user := c.Get("user").(*model.Users)

	// get id from path param
	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	history,err := h.Repository.Update(id,user.ID)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	// success
	helper.WriteResponseWithData(c,200,"Success updated rent",history)
	return nil
}

// ViewAll godoc
// @Summary View all rental histories
// @Description View all rental histories for the logged-in user
// @Tags Histories
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header with token"
// @Success 200 {array} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Failure 401 {object} helper.ErrResponse
// @Router /histories [get]
func (h Histories) ViewAll(c echo.Context) error {
	// get logged in user
	user := c.Get("user").(*model.Users)

	// with query param
	status := c.QueryParam("status")

	// available
	if status == "in-progress" || status == "In-Progress" || status == "IN-PROGRESS"{
		ok := false
		videoGames, err := h.Repository.FindAllWithStatus(user.ID,ok)
		if err != nil {
			helper.ErrorResponse(400,err.Error())
		}

		helper.WriteResponseWithData(c,200,"Success view all available video games",videoGames)
		return nil
	} 
	
	// not available
	if status == "done" || status == "Done" || status == "DONE"{
		ok := true
		videoGames, err := h.Repository.FindAllWithStatus(user.ID,ok)
		if err != nil {
			helper.ErrorResponse(400,err.Error())
		}

		helper.WriteResponseWithData(c,200,"Success view all unavailable video games",videoGames)
		return nil
	} 

	// view all
	histories,err := h.Repository.FindAll(user.ID)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	// success
	helper.WriteResponseWithData(c,200,"Success view all histories",histories)
	return nil
}

// ViewById godoc
// @Summary View a rental history by ID
// @Description View a rental history by its unique ID
// @Tags Histories
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header with token"
// @Param id path string true "Rental History ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Failure 401 {object} helper.ErrResponse
// @Router /histories/{id} [get]
func (h Histories) ViewById(c echo.Context) error {
	// get logged in user
	user := c.Get("user").(*model.Users)

	// get id from path param
	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	history,err := h.Repository.FindById(id,user.ID)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	// success
	helper.WriteResponseWithData(c,200,"Success view history by id",history)
	return nil
}

// DeleteById godoc
// @Summary Delete a rental history by ID
// @Description Delete a rental history by its unique ID
// @Tags Histories
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header with token"
// @Param id path string true "Rental History ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Failure 401 {object} helper.ErrResponse
// @Router /histories/{id} [get]
func (h Histories) Delete(c echo.Context) error {
	// get id from path param
	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	// get user
	user := c.Get("user").(*model.Users)

	// get history
	history,err := h.Repository.FindById(id,user.ID)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	// if status in-progress(not finish)
	if history.Status == "In-Progress"{
		return helper.ErrorResponse(400, "unable to delete history in progress status, please complete your rental process first.")
	}

	// delete history
	deletedHistory,err := h.Repository.Delete(*history)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}
	
	// success
	helper.WriteResponseWithData(c,200,"Your history has been successfully deleted",deletedHistory)
	return nil
}