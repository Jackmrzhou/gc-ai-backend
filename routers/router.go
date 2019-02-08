package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai/conf"
	"github.com/jackmrzhou/gc-ai/middleware"
	"github.com/jackmrzhou/gc-ai/routers/api"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(conf.RunMode)

	router.POST("/auth", api.GetAuth)
	router.POST("/register", api.Register)

	apis := router.Group("/api/v1")
	apis.Use(middleware.JwtValidation())
	{

	}
	return router
}
