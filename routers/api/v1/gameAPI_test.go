package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/json-models"
	"github.com/jackmrzhou/gc-ai-backend/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var games = []json_models.NewGameReq{
	{
		Name:"game 1",
		Introduction:"introduction 1",
	},
	{
		Name:"game 1",
		Introduction:"introduction 2",
	},
}

func TestNewGame(t *testing.T) {
	r.POST("/game", NewGame)
	for i, game := range games{
		body, _ := json.Marshal(game)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/game", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if i == 0{
			assert.Equal(t, 200, w.Code)
		}else {
			assert.Equal(t, 400, w.Code)
		}
	}
}

func TestAllGames(t *testing.T) {
	r.GET("/games/all", AllGames)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/games/all", nil)
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)
}

func TestGetGameRank(t *testing.T) {
	r.GET("/rank/game", GetGameRank)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rank/game", nil)
	q := req.URL.Query()
	q.Add("game_id", fmt.Sprint(initGame.ID))
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/rank/game", nil)
	q = req.URL.Query()
	q.Add("game_id", fmt.Sprint(1))
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	fmt.Println(w.Body)
}

func TestGetUserRank(t *testing.T) {
	r.GET("/rank/user", GetUserRank)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rank/user", nil)
	q := req.URL.Query()
	q.Add("user_id", fmt.Sprint(initUser.ID))
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/rank/user", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	fmt.Println(w.Body)
}

var initGame models.Game
var initUser models.User
var initRank models.Rank

func initRun() {
	user , _ := models.CreateUser("1234@a.b", "123456")
	initUser = user
	initGame = models.Game{
		Name:"game 0",
		Introduction:"introduction 0",
	}
	_ = models.CreateGame(&initGame)
	rank, _ := models.CreateRank(&user, &initGame, 10)
	initRank = *rank
	fmt.Println(*rank)
}

func clear() {
	models.DeleteUserDebug(initUser.Email)
	models.DeleteGameByIDDebug(initGame.ID)
	models.DeleteRankByIDDebug(initRank.ID)
	for _, game:= range games{
		models.DeleteGameByNameDebug(game.Name)
	}
}

var r *gin.Engine

func TestMain(m *testing.M) {
	if err := conf.LoadConf("../../../conf/app.conf"); err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	r = utils.TestRouter()
	if err := models.OpenDB(); err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	initRun()
	m.Run()
	clear()
}