package router

import (
	_ "videogames_rent_api/docs"
	"videogames_rent_api/handler"
	"videogames_rent_api/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func InitEchoInstance(db *gorm.DB) *echo.Echo{
	// router and server
	e := echo.New()

	// swaggo
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// auth
	authMiddleware := middleware.InitVerification(db)
	
	// user router
	userHandler := handler.InitUserHandler(db)
	user := e.Group("/users")
	{
		// POST http://localhost:8080/users/register - user register
		user.POST("/register", userHandler.Register)
		// PUT http://localhost:8080/users/statusverification/:userId/:code - user verification
		user.GET("/statusverification/:userId/:code", userHandler.StatusVerification)
		// POST http://localhost:8080/users/login - user login
		user.POST("/login", userHandler.Login,authMiddleware.AuthorizeUserStatus)
		// POST http://localhost:8080/users - get info user
		user.GET("", userHandler.GetInfo,authMiddleware.Authentication)
		// PUT http://localhost:8080/users/top_up - user top-up deposit amount
		user.PUT("/top_up",userHandler.TopUp,authMiddleware.Authentication)
		

		// User histories
		historiesHandler := handler.InitHistoriesHandler(db)
		histories := user.Group("/rent/histories",authMiddleware.Authentication)
		{
			// GET http://localhost:8080/users/rent - user update rent
			histories.POST("",historiesHandler.AddRent)
			// PUT http://localhost:8080/users/rent/:id - user create rent
			histories.PUT("/:id",historiesHandler.UpdateRent)
			// GET http://localhost:8080/users/rent/histories - view all video games
			histories.GET("",historiesHandler.ViewAll)
			// GET http://localhost:8080/users/rent/histories/:id - view all video games
			histories.GET("/:id",historiesHandler.ViewById)
			// PUT http://localhost:8080/users/rent/delete - delete user
			histories.DELETE("/:id",historiesHandler.Delete)
		}

	}

	// video_games router
	vgHandler := handler.InitVideoGameHandler(db)
	vg := e.Group("/video_games",authMiddleware.Authentication)
	{
		// GET http://localhost:8080/video_games - view all video games
		vg.GET("",vgHandler.ViewAll)
		// GET http://localhost:8080/video_games/:id - view by id video game
		vg.GET("/:id",vgHandler.ViewById)
		
	}
	
	
	return e
}