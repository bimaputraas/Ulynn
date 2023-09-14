package handler

import (
	"strconv"
	"videogames_rent_api/helper"

	"github.com/labstack/echo/v4"
)

func (h VideoGame) ViewAll(c echo.Context) error {
	// with query param
	isTrueStr := c.QueryParam("availability")

	// available
	if isTrueStr == "true"{
		ok := true
		videoGames, err := h.Repository.FindAllWithAvailability(ok)
		if err != nil {
			helper.ErrorResponse(400,err.Error())
		}

		helper.WriteResponseWithData(c,200,"Success view all available video games",videoGames)
		return nil
	} 
	
	// not available
	if isTrueStr == "false"{
		ok := false
		videoGames, err := h.Repository.FindAllWithAvailability(ok)
		if err != nil {
			helper.ErrorResponse(400,err.Error())
		}

		helper.WriteResponseWithData(c,200,"Success view all unavailable video games",videoGames)
		return nil
	} 

	// without query param, view all
	videoGames,err := h.Repository.FindAll()
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	helper.WriteResponseWithData(c,200,"Success view all video games",videoGames)
	return nil
}

func (h VideoGame) ViewById(c echo.Context) error {
	// get id from path param
	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	// find by id
	videoGame,err := h.Repository.FindById(id)
	if err != nil {
		return helper.ErrorResponse(400,err.Error())
	}

	// success
	helper.WriteResponseWithData(c,200,"Success view video game by id",videoGame)
	return nil
}
