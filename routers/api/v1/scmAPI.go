package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/api-codes"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/json-models"
	"github.com/jackmrzhou/gc-ai-backend/utils"
	"net/http"
)

func UploadSourceCode(c *gin.Context) {
	// swagger:route POST /api/v1/sourcecode uploadSourceCode
	//
	// upload source code
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: uploadSourceCodeResp
	//       400: statusResponse
	var returnCode int

	var json json_models.UploadSourceCodeReq
	if err := c.ShouldBindJSON(&json); err != nil{
		utils.ErrorResponse(c, api_codes.INVALID)
		return
	}

	userID, _ := utils.GetCurrentUserID(c)

	var user *models.User
	var game *models.Game
	var err error

	if user, err = models.QueryUserByID(userID); err != nil{
		returnCode = api_codes.NewSrcFailed
	}

	if game, err = models.QueryGameByID(json.GameID); err != nil{
		returnCode = api_codes.NewSrcFailed
	}

	if returnCode == api_codes.NewSrcFailed{
		utils.ErrorResponse(c, returnCode)
		return
	}

	var sourceCode *models.SourceCode
	if sourceCode, err = models.CreateSourceCode(user, game, json.CodeType, json.Language, json.SourceCode); err != nil{
		returnCode = api_codes.NewSrcFailed
		utils.ErrorResponse(c, returnCode)
	}else{
		returnCode = api_codes.SUCCESS
		c.JSON(http.StatusOK, json_models.UploadSourceCodeResp{
			Code:returnCode,
			Msg:api_codes.GetMsg(returnCode),
			Data:json_models.UploadSourceCodeData{
				SourceCodeID:sourceCode.ID,
			},
		})
	}
}

func GetSourceCodesByUser(c *gin.Context) {
	// swagger:route GET /api/v1/user/sourcecode getSourceCodeByUser
	//
	// get the user's all source codes
	//
	//     Consumes:
	//     - application/x-www-form-urlencoded
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: getSourceCodesResp
	//		 400: statusResponse
	userID , _ := utils.GetCurrentUserID(c)

	if codes, err := models.QuerySourceCodeByUserID(userID); err != nil || len(codes) == 0{
		utils.ErrorResponse(c, api_codes.SourceCodesNotFound)
	}else{
		c.JSON(http.StatusOK, json_models.GetSourceCodes{
			Code:api_codes.SUCCESS,
			Msg:api_codes.GetMsg(api_codes.SUCCESS),
			Data:codes,
		})
	}
}

func GetSrcByUserAndGame(c *gin.Context) {
	// swagger:route GET /api/v1/sourcecode getSrcByUserAndGame
	//
	// get the user's all source codes
	//
	//     Consumes:
	//     - application/x-www-form-urlencoded
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: getSourceCodesResp
	//		 400: statusResponse
	var query json_models.GetSrcByUserAndGameReq
	if c.ShouldBindQuery(&query) != nil{
		utils.ErrorResponse(c, api_codes.INVALID)
		return
	}

	userID, _ := utils.GetCurrentUserID(c)
	if codes, err := models.QuerySourceCodeByUserGameIDs(userID, query.GameID); err != nil || len(codes) == 0{
		utils.ErrorResponse(c, api_codes.SourceCodesNotFound)
	}else{
		c.JSON(http.StatusOK, json_models.GetSourceCodes{
			Code:api_codes.SUCCESS,
			Msg:api_codes.GetMsg(api_codes.SUCCESS),
			Data:codes,
		})
	}
}