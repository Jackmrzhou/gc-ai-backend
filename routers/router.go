package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jackmrzhou/gc-ai-backend/middleware"
	"github.com/jackmrzhou/gc-ai-backend/routers/api"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/v1"
)


func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(conf.RunMode)

	if conf.Swagger {
		router.Static("/swagger/", "swaggerui/")
	}
	router.POST("/auth", api.GetAuth)
	router.POST("/register", api.Register)
	router.POST("/vericode", api.SendVeriCode)

	apis := router.Group("/api/v1")
	apis.Use(middleware.JwtValidation())
	{
		apis.POST("/game", v1.NewGame)
		apis.GET("/games/all", v1.AllGames)
		apis.GET("/rank/game", v1.GetGameRank)
		apis.GET("/rank/user", v1.GetUserRank)

		apis.POST("/sourcecode", v1.UploadSourceCode)
		apis.GET("/user/sourcecode", v1.GetSourceCodesByUser)
		apis.GET("/sourcecode", v1.GetSrcByUserAndGame)

		apis.POST("/battle", v1.StartBattle)
		apis.GET("/battle", v1.QueryProcess)
	}
	return router
}
