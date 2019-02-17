package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/api-codes"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/json-models"
	"github.com/jackmrzhou/gc-ai-backend/utils"
	"github.com/jackmrzhou/gc-ai-backend/verification"
	"net/http"
)

func GetAuth(c *gin.Context) {
	// swagger:route POST /auth getAuth
	//
	// get jwt token from server
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: authSuccess
	//       400: statusResponse
	var json json_models.UserJSON
	if err := c.ShouldBindJSON(&json); err != nil{
		// invalid parameters
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code: api_codes.INVALID,
			Msg:  api_codes.GetMsg(api_codes.INVALID),
		})
		return
	}

	if user,err := models.QueryUser(json.Email, json.Password); err != nil{
		// authentication failed
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code: api_codes.AuthFailed,
			Msg:  api_codes.GetMsg(api_codes.AuthFailed),
		})
	}else if models.IsBanned(&user){
		// the user is banned
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code: api_codes.Banned,
			Msg:  api_codes.GetMsg(api_codes.Banned),
		})
	}else {
		// return the token
		c.JSON(http.StatusOK, json_models.AuthSuccess{
			Code: api_codes.SUCCESS,
			Msg:  api_codes.GetMsg(api_codes.SUCCESS),
			Data: json_models.AuthSuccessData{
				Token: utils.GenerateToken(user.ID, user.Email),
			},
		})
	}
}

func Register(c *gin.Context) {
	// swagger:route POST /register register
	//
	// register a new account
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: authSuccess
	//       400: statusResponse
	var json json_models.RegInfo
	if err := c.ShouldBindJSON(&json); err != nil{
		// invalid parameters
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code:     api_codes.INVALID,
			Msg:      api_codes.GetMsg(api_codes.INVALID),
			JSONData: struct {}{},
		})
		return
	}
	if !verification.CheckAndDelCode(json.Email, json.VeriCode){
		// verification failed
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code:     api_codes.VeriFailed,
			Msg:      api_codes.GetMsg(api_codes.VeriFailed),
			JSONData: struct {}{},
		})
		return
	}
	if user, err := models.CreateUser(json.Email, json.Password); err != nil{
		// registration failed
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code :    api_codes.RegFailed,
			Msg:      api_codes.GetMsg(api_codes.RegFailed),
			JSONData: struct {}{},
		})
	}else {
		// registration succeeded, return user_id and token
		c.JSON(http.StatusOK, json_models.AuthSuccess{
			Code: api_codes.SUCCESS,
			Msg:  api_codes.GetMsg(api_codes.SUCCESS),
			Data: json_models.AuthSuccessData{
				Token: utils.GenerateToken(user.ID, user.Email),
			},
		})
	}
}

func SendVeriCode(c *gin.Context) {
	// swagger:route POST /vericode sendVeriCode
	//
	// request to server to send a verification code
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: statusResponse
	//       400: statusResponse
	//		 500: statusResponse
	var json json_models.TargetEmail
	if err := c.ShouldBindJSON(&json); err != nil{
		// invalid parameters
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code:     api_codes.INVALID,
			Msg:      api_codes.GetMsg(api_codes.INVALID),
			JSONData: struct {}{},
		})
		return
	}
	if _, err := verification.SendCode(json.Email); err != nil{
		// send mail failed
		c.JSON(http.StatusInternalServerError, json_models.APIError{
			Code:     api_codes.SendMailFailed,
			Msg:      api_codes.GetMsg(api_codes.SendMailFailed),
			JSONData: struct {}{},
		})
	}else{
		// send mail succeeded
		c.JSON(http.StatusBadRequest, json_models.Status{
			Code:     api_codes.SUCCESS,
			Msg:      api_codes.GetMsg(api_codes.SUCCESS),
			JSONData: struct {}{},
		})
	}
}