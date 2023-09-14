package handler

import (
	"strconv"
	"videogames_rent_api/helper"
	"videogames_rent_api/model"

	"github.com/labstack/echo/v4"
)

func (h History) ViewAll(c echo.Context) error {
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

func (h History) ViewById(c echo.Context) error {
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



