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
}

func (app *App)Initialize(ConfigFile string) error {
	if err := conf.LoadConf(ConfigFile); err != nil{
		return err
	}
	app.Router = routers.InitRouter()
	if err := models.OpenDB(); err != nil{
		return err
	}
	return nil
}

func (app *App)Run() {
	app.Router.Run()
}