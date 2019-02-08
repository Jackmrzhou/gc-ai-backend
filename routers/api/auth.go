package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai/code"
	"github.com/jackmrzhou/gc-ai/models"
	"github.com/jackmrzhou/gc-ai/routers/api/json-models"
	"github.com/jackmrzhou/gc-ai/utils"
	"github.com/jackmrzhou/gc-ai/verification"
	"net/http"
)

func GetAuth(c *gin.Context) {
	var json json_models.UserJSON
	if err := c.ShouldBindJSON(&json); err != nil{
		// invalid parameters
		c.JSON(http.StatusBadRequest, JSONTemplate{
			Code:code.INVALID,
			Msg: code.GetMsg(code.INVALID),
		})
		return
	}

	if user,err := models.QueryUser(json.Email, json.Password); err != nil{
		// authentication failed
		c.JSON(http.StatusBadRequest, JSONTemplate{
			Code:code.AuthFailed,
			Msg: code.GetMsg(code.AuthFailed),
		})
	}else if models.IsBanned(&user){
		// the user is banned
		c.JSON(http.StatusBadRequest, JSONTemplate{
			Code:code.Banned,
			Msg:code.GetMsg(code.Banned),
		})
	}else {
		// return the token
		c.JSON(http.StatusOK, JSONTemplate{
			Code: code.SUCCESS,
			Msg:  code.GetMsg(code.SUCCESS),
			JSONData: struct {
				Token string
			}{utils.GenerateToken(user.ID, user.Email)},
		})
	}
}

func Register(c *gin.Context) {
	var json json_models.RegInfo
	if err := c.ShouldBindJSON(&json); err != nil{
		// invalid parameters
		c.JSON(http.StatusBadRequest, JSONTemplate{
			Code:code.INVALID,
			Msg: code.GetMsg(code.INVALID),
		})
		return
	}
	if !verification.CheckAndDelCode(json.Email, json.VeriCode){
		// verification failed
		c.JSON(http.StatusBadRequest, JSONTemplate{
			Code:code.VeriFailed,
			Msg:code.GetMsg(code.VeriFailed),
		})
		return
	}
	if user, err := models.CreateUser(json.Email, json.Password); err != nil{
		// registration failed
		c.JSON(http.StatusBadRequest, JSONTemplate{
			Code : code.RegFailed,
			Msg:code.GetMsg(code.RegFailed),
		})
	}else {
		// registration succeeded, return user_id and token
		c.JSON(http.StatusOK, JSONTemplate{
			Code: code.SUCCESS,
			Msg:  code.GetMsg(code.SUCCESS),
			JSONData: struct {
				Token string
			}{
				utils.GenerateToken(user.ID, user.Email),
			},
		})
	}
}

func SendVeriCode(c *gin.Context) {
	var json json_models.SendCodeRecv
	if err := c.ShouldBindJSON(&json); err != nil{
		// invalid parameters
		c.JSON(http.StatusBadRequest, JSONTemplate{
			Code:code.INVALID,
			Msg: code.GetMsg(code.INVALID),
		})
		return
	}
	if _, err := verification.SendCode(json.Email); err != nil{
		// send mail failed
		c.JSON(http.StatusBadRequest, JSONTemplate{
			Code:code.SendMailFailed,
			Msg: code.GetMsg(code.SendMailFailed),
		})
	}else{
		// send mail succeeded
		c.JSON(http.StatusBadRequest, JSONTemplate{
			Code:code.SUCCESS,
			Msg: code.GetMsg(code.SUCCESS),
		})
	}
}