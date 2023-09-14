package middleware

import (
	"videogames_rent_api/helper"
	"videogames_rent_api/model"

	"github.com/labstack/echo/v4"
)


func(h Verification) AuthorizeUserStatus(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// handler login
		// next(c)
		err := next(c)
		if err != nil {
			return err
		}

		// get user
		user := c.Get("user").(*model.Users)

		// if status pending verification
		if user.Status != "Activated"{
		return helper.ErrorResponse(403, "Your account activation is pending, please check your email for activate an account.")
		}

		// generate token
		tokenString,err := helper.GenerateToken(user.ID)
		if err != nil {
			return helper.ErrorResponse(400, err.Error())
		}

		// update user jwt token
		loggedUser,err := h.Repository.UpdateJwtToken(tokenString,*user)
		if err != nil {
			return helper.ErrorResponse(400, err.Error())
		}

		// success
		helper.WriteResponseWithData(c,201,"Success login",loggedUser)
		return nil
	}
}


func(h Verification) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
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
		user,err := h.Repository.FindUserById(id)
		if err != nil {
			return helper.ErrorResponse(401,"Undefined user")
		}

		// success
		c.Set("user",user)
		return next(c)
	}
}
