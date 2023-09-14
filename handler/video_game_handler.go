package handler

import (
	"strconv"
	"videogames_rent_api/helper"

	"github.com/labstack/echo/v4"
)

// ViewAll godoc
// @Summary View all video games
// @Description View all video games with optional availability filter
// @Tags VideoGame
// @Accept json
// @Produce json
// @Param availability query string false "Filter by availability (true or false)"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Failure 401 {object} helper.ErrResponse
// @Router /video_games [get]
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

// ViewById godoc
// @Summary View a video game by ID
// @Description View a video game by its unique ID
// @Tags VideoGame
// @Accept json
// @Produce json
// @Param id path string true "Video Game ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.ErrResponse
// @Failure 401 {object} helper.ErrResponse
// @Router /video_games/{id} [get]
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
