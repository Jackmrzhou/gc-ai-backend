package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/api-codes"
	"github.com/jackmrzhou/gc-ai-backend/battle"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/json-models"
	"github.com/jackmrzhou/gc-ai-backend/utils"
	"log"
	"net/http"
)

func StartBattle(c *gin.Context) {
	// swagger:route POST /api/v1/battle startBattle
	//
	// start a battle
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: startBattleResp
	//       400: statusResponse

	var json json_models.StartBattleReq

	if c.ShouldBindJSON(&json) != nil{
		// invalid parameters
		utils.ErrorResponse(c, api_codes.INVALID)
		return
	}

	attackerID, _ := utils.GetCurrentUserID(c)
	defenderID := json.UserID


	returnCode := api_codes.SUCCESS
	var _battle *models.Battle

	if attackerSrc, err := models.QueryATKSrcByUserID(attackerID); err != nil{
		returnCode = api_codes.ObjectNotFound
	}else if defenderSrc, err := models.QueryDEFSrcByUserID(defenderID); err != nil{
		returnCode = api_codes.ObjectNotFound
	}else if game, err := models.QueryGameByID(json.GameID); err != nil{
		returnCode = api_codes.ObjectNotFound
	}else if models.IsEngagedInBattle(attackerID) {
		returnCode = api_codes.WaitForFinish
	}else{
		// all objects are available

		if _battle, err = models.CreateBattle(attackerSrc.UserID, defenderSrc.UserID, game); err != nil{
			log.Fatalf("Crate battle failed. %d, %d", attackerSrc.UserID, defenderSrc.UserID)
			returnCode = api_codes.StartBattleFailed
		}else {
			// judge
			go battle.Judge(attackerSrc, defenderSrc, game, _battle)
		}
	}

	if returnCode == api_codes.SUCCESS{
		c.JSON(http.StatusOK, json_models.StartBattleResp{
			Code:api_codes.SUCCESS,
			Msg:api_codes.GetMsg(api_codes.SUCCESS),
			Data:json_models.StartBattleData{
				BattleID: _battle.ID,
				Status:_battle.Status,
			},
		})
	}else{
		utils.ErrorResponse(c, returnCode)
	}
}

func QueryProcess(c *gin.Context) {
	// swagger:route GET /api/v1/battle queryProcess
	//
	// query the process of a battle
	//
	//     Consumes:
	//     - application/x-www-form-urlencoded
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: queryProcessResp
	//       400: statusResponse
	var query json_models.QueryProcessReq

	if c.ShouldBindJSON(&query) != nil{
		// invalid parameters
		utils.ErrorResponse(c, api_codes.INVALID)
		return
	}

	if b, err := models.QueryBattleByID(query.BattleID); err != nil{
		utils.ErrorResponse(c, api_codes.ObjectNotFound)
	}else{
		c.JSON(http.StatusOK, json_models.QueryProcessResp{
			Code:api_codes.SUCCESS,
			Msg:api_codes.GetMsg(api_codes.SUCCESS),
			Data:*b,
		})
	}
}
