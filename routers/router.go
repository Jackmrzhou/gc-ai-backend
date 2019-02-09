package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jackmrzhou/gc-ai-backend/middleware"
	"github.com/jackmrzhou/gc-ai-backend/routers/api"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @host 47.102.147.41

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(conf.RunMode)

	if conf.Swagger {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.POST("/auth", api.GetAuth)
	router.POST("/register", api.Register)
	router.POST("/vericode", api.SendVeriCode)

	apis := router.Group("/api/v1")
	apis.Use(middleware.JwtValidation())
	{

	}
	return router
}
