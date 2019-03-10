package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/api-codes"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/json-models"
	"github.com/jackmrzhou/gc-ai-backend/utils"
	"log"
	"net/http"
)

func UpdateProfile(c *gin.Context) {
	var json json_models.UpdateProfileReq
	if err := c.ShouldBindJSON(&json); err != nil{
		utils.ErrorResponse(c, api_codes.INVALID)
		return
	}

	userID, _ := utils.GetCurrentUserID(c)

	if err := models.UpdateProfile(&models.Profile{
		UserID:userID,
		Nickname:json.Nickname,
		Avatar:json.Avatar,
		Introduction:json.Introduction,
	}); err != nil{
		log.Println(err)
		utils.ErrorResponse(c, api_codes.UpdateProfileFailed)
	}else {
		utils.SuccessResponse(c, api_codes.SUCCESS)
	}
}

func TestNickname(c *gin.Context) {
	var query json_models.TestNicknameReq
	if c.ShouldBindQuery(&query) != nil{
		utils.ErrorResponse(c, api_codes.INVALID)
		return
	}

	var IsExists bool

	if _, err := models.QueryProfileByNickname(query.Nickname); err != nil{
		IsExists = true
	}else {
		IsExists = false
	}

	c.JSON(http.StatusOK, json_models.TestNicknameResp{
		Code:api_codes.SUCCESS,
		Msg:api_codes.GetMsg(api_codes.SUCCESS),
		Data:json_models.TestNicknameData{
			Exists:IsExists,
		},
	})
}

func GetProfile(c *gin.Context) {
	userId, _ := utils.GetCurrentUserID(c)

	if profile, err := models.QueryProfileByUserID(userId); err != nil{
		log.Println(err)
		utils.ErrorResponse(c, api_codes.ObjectNotFound)
	}else {
		c.JSON(http.StatusOK, json_models.GetProfileResp{
			Code:api_codes.SUCCESS,
			Msg:api_codes.GetMsg(api_codes.SUCCESS),
			Data:profile,
		})
	}
}
