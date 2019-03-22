package server

import (
	"github.com/ShawnRong/bento/controllers"
	"github.com/ShawnRong/bento/middlewares"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// NewRouter gin router
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware())

	// Serve frontend views folder
	router.Use(static.Serve("/", static.LocalFile("./app/build/", true)))

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.POST("/", user.Create)
			userGroup.GET("/:id", user.Retrieve)
			userGroup.DELETE("/:id", user.Delete)
			userGroup.PATCH("/:id", user.Update)
		}
	}

	return router
}
