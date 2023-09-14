package middleware

import (
	"videogames_rent_api/helper"
	"videogames_rent_api/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

func(h Auth) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get user id
		tokenString := c.Request().Header.Get("Authorization")

		// no token
		if tokenString == ""{
			return helper.ErrorResponse(401, "No token")
		}

		// parse token
		claims,err := helper.ParseToken(tokenString)
		if err != nil {
			return helper.ErrorResponse(401, err.Error())
		}

		// check and get user from db
		id := claims["id"]
		var user model.Users
		result := h.DB.First(&user,id)
		if result.Error != nil {
			return helper.ErrorResponse(401, "Unauthorized user")
		}

		// success
		c.Set("user",user)
		return next(c)
	}
}
