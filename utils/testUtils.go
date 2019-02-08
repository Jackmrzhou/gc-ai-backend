package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jackmrzhou/gc-ai-backend/middleware"
)

func TestRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(conf.RunMode)

	apis := router.Group("/api/v1")
	apis.Use(middleware.JwtValidation())
	{

	}
	return router
}