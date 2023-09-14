package config

import (
	"videogames_rent_api/handler"
	"videogames_rent_api/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitEchoInstance(db *gorm.DB) *echo.Echo{
	// router and server
	e := echo.New()
	
	// auth
	authMiddleware := middleware.Auth{
		DB:db,
	}
	
	// user router
	userHandler := handler.InitUserHandler(db)
	user := e.Group("/users")
	{
		// POST http://localhost:8080/users/register - user register
		user.POST("/register", userHandler.Register)
		// POST http://localhost:8080/users/login - user login
		user.POST("/login", userHandler.Login)
		// POST http://localhost:8080/users/login - user login
		user.GET("", userHandler.GetInfo,authMiddleware.Authentication)
		// PUT http://localhost:8080/users/login - user top-up deposit amount
		user.PUT("/top_up",userHandler.TopUp,authMiddleware.Authentication)
		

		// User histories router
		historiesHandler := handler.InitHistoryHandler(db)
		history := user.Group("/histories",authMiddleware.Authentication)
		{
			// GET http://localhost:8080/users/histories - view all video games
			history.GET("",historiesHandler.ViewAll)
			// GET http://localhost:8080/users/histories/:id - view all video games
			history.GET("/:id",historiesHandler.ViewById)
			

		}

		// User rent
		rentHandler := handler.InitRentHandler(db)
		rent := user.Group("/rent",authMiddleware.Authentication)
		{
			// GET http://localhost:8080/users/histories/:id - user update rent
			rent.POST("",rentHandler.AddRent)
			// PUT http://localhost:8080/users/login - user create rent
			rent.PUT("/:id",rentHandler.UpdateRent)
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