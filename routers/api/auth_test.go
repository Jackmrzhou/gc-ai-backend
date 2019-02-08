package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai/conf"
	"github.com/jackmrzhou/gc-ai/models"
	"github.com/jackmrzhou/gc-ai/routers/api/json-models"
	"github.com/jackmrzhou/gc-ai/utils"
	"github.com/jackmrzhou/gc-ai/verification"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var regs = []json_models.RegInfo{
	{
		Email:"594780735@qq.com",
		Password:"123456",
	},
	{
		Email:"test2@qq.com",
		Password:"1234567",
	},
}

var users = []json_models.UserJSON{
	{
		Email:"594780735@qq.com",
		Password:"123456",
	},
	{
		Email:"test1@qq.com",
		Password:"123",
	},
}

var r *gin.Engine

func TestRegister(t *testing.T) {
	r.POST("/register", Register)
	for i, reg := range regs {
		body, _ := json.Marshal(reg)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if i == 0{
			assert.Equal(t, 200, w.Code)
		}else {
			assert.Equal(t, 400, w.Code)
		}
		fmt.Println(w.Body)
	}
}

func TestGetAuth(t *testing.T) {
	r.POST("/auth", GetAuth)
	for i, user := range users{
		body, _ := json.Marshal(user)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if i == 0{
			assert.Equal(t, 200, w.Code)
		}else {
			assert.Equal(t, 400, w.Code)
		}
		fmt.Println(w.Body)
	}
}

func clear() {
	models.DeleteUser(regs[0].Email)
}

func TestMain(m *testing.M) {
	if err := conf.LoadConf("../../conf/app.conf"); err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	r = utils.TestRouter()
	if err := models.OpenDB(); err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	code, err := verification.SendCode(regs[0].Email)
	if err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	regs[0].VeriCode = code
	m.Run()
	clear()
}