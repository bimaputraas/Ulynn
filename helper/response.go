package helper

import "github.com/labstack/echo/v4"

type Response struct {
	Message string
	Data    interface{}
}

func WriteResponse(c echo.Context,code int,msg string){
	c.JSON(code,Response{
		Message: msg,
		Data: "-",
	})
}

func WriteResponseWithData(c echo.Context,code int,msg string,data interface{}){
	c.JSON(code,Response{
		Message: msg,
		Data: data,
	})
}

func ErrorResponse(code int,msg string) *echo.HTTPError{
	err := echo.NewHTTPError(code,Response{
		Message: msg,
		Data: "-",
	})

	return err
}