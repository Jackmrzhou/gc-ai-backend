// Package classification gc-ai API.
//
// gc-ai API specification
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: 47.102.147.41:8080
//     BasePath: /
//     Version: 0.1.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: 594780735@qq.com
//
// swagger:meta
package ab_app

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"github.com/jackmrzhou/gc-ai-backend/routers"
)

type AppInterface interface {
	Initialize(ConfigFile string) error
	Run()
}

type App struct {
	Router *gin.Engine
	// db *gorm.DB
	// actually db will be maintained in db.go
	// to follow the design pattern, we will store one global variable
	models.DBMaintainer
}

func (app *App)Initialize(ConfigFile string) error {
	if err := conf.LoadConf(ConfigFile); err != nil{
		return err
	}
	app.Router = routers.InitRouter()
	if err := models.OpenDB(); err != nil{
		return err
	}
	app.DBMaintainer.AddFunc(models.MaintainVeriCode, conf.CodeActiveTime)
	return nil
}

func (app *App)Run() {
	app.Router.Run()
}