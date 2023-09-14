package handler

import (
	"strconv"
	"videogames_rent_api/helper"
	"videogames_rent_api/model"

	"github.com/labstack/echo/v4"
)

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
	user := c.Get("user").(model.Users)

	// update user,video_game, and create a new history
	resBody, err := h.Repository.UpdateCreate(int(reqBody.VideoGamesID), user)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// success
	helper.WriteResponseWithData(c, 201, "Success add a new rent", resBody)
	return nil
}

func (h Histories) UpdateRent(c echo.Context) error {
	// get logged in user
	user := c.Get("user").(model.Users)

	// bind
	var reqBody model.ReqBodyHistoryUpdate
	err := c.Bind(&reqBody)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// invalid bind status, bad request
	if reqBody.Status != "done" && reqBody.Status != "DONE"{
		return helper.ErrorResponse(400, "invalid status bind")
	}

	// validate
	err = helper.Validate(reqBody)
	if err != nil {
		return helper.ErrorResponse(400, err.Error())
	}

	// get id from path param
	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	history,err := h.Repository.Update(id,user.ID,reqBody.Status)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	// success
	helper.WriteResponseWithData(c,200,"Success updated rent",history)
	return nil
}

func (h Histories) ViewAll(c echo.Context) error {
	// get logged in user
	user := c.Get("user").(model.Users)

	// view all
	histories,err := h.Repository.FindAll(user.ID)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	// success
	helper.WriteResponseWithData(c,200,"Success view all histories",histories)
	return nil
}

func (h Histories) ViewById(c echo.Context) error {
	// get logged in user
	user := c.Get("user").(model.Users)

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