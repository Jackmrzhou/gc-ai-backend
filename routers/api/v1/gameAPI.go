package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/api-codes"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/json-models"
	"net/http"
)

func NewGame(c *gin.Context) {
	// swagger:route POST /api/v1/game newGame
	//
	// create a new game
	// only administrator is allowed
	//
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
	var json json_models.NewGameReq
	if err := c.ShouldBindJSON(&json); err != nil{
		// invalid parameters
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code: api_codes.INVALID,
			Msg:  api_codes.GetMsg(api_codes.INVALID),
		})
		return
	}

	game := models.Game{
		Name:json.Name,
		Introduction:json.Introduction,
	}
	if err := models.CreateGame(&game); err != nil{
		// create game failed
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code: api_codes.NewGameFailed,
			Msg:  api_codes.GetMsg(api_codes.NewGameFailed),
		})
	}else{
		// create game succeeded
		c.JSON(http.StatusOK, json_models.Status{
			Code: api_codes.SUCCESS,
			Msg:  api_codes.GetMsg(api_codes.SUCCESS),
		})
	}
}

func AllGames(c *gin.Context) {
	// swagger:route GET /api/v1/games/all getAllGame
	//
	// get all the games
	//
	//     Consumes:
	//     - application/x-www-form-urlencoded
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: allGame
	games := models.QueryAllGames()
	// return all the games queried
	c.JSON(http.StatusOK, json_models.AllGamesResp{
		Code: api_codes.SUCCESS,
		Msg:  api_codes.GetMsg(api_codes.SUCCESS),
		Data: games,
	})
}

func GetGameRank(c *gin.Context) {
	// swagger:route GET /api/v1/rank/game getGameRank
	//
	// get the ranking from a specific game
	// the API will return http status 400 if
	// no ranks found
	//
	//     Consumes:
	//     - application/x-www-form-urlencoded
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: gameRank
	//       400: statusResponse
	var query json_models.GetGameRankReq
	if c.ShouldBindQuery(&query) != nil{
		// invalid parameters
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code: api_codes.INVALID,
			Msg:  api_codes.GetMsg(api_codes.INVALID),
		})
		return
	}

	if ranks, err := models.QueryRankByGameID(query.ID); err != nil || len(ranks) == 0{
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code: api_codes.RankNotFound,
			Msg:  api_codes.GetMsg(api_codes.RankNotFound),
		})
	}else{
		var gameRanks []json_models.GameRank
		for _, rank := range ranks{
			profile, _ := models.QueryProfile(rank.UserID)
			gameRanks = append(gameRanks, json_models.GameRank{
				Score:    rank.Score,
				UserInGameRank: json_models.UserInGameRank{
					ID:rank.UserID,
					Nickname:profile.Nickname,
					LastUpdate:rank.UpdatedAt,
				},
			})
		}
		c.JSON(http.StatusOK, json_models.GameRankResp{
			Code: api_codes.SUCCESS,
			Msg:  api_codes.GetMsg(api_codes.SUCCESS),
			Data: gameRanks,
		})
	}
}

func GetUserRank(c *gin.Context) {
	// swagger:route GET /api/v1/rank/user getUserRank
	//
	// get the ranking from a specific user.
	// the API will return http status 400 if
	// no ranks found
	//
	//     Consumes:
	//     - application/x-www-form-urlencoded
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: userRank
	//       400: statusResponse
	var query json_models.GetUserRankReq
	if c.ShouldBindQuery(&query) != nil{
		// invalid parameters
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code: api_codes.INVALID,
			Msg:  api_codes.GetMsg(api_codes.INVALID),
		})
		return
	}

	if ranks, err := models.QueryRankByUserID(query.ID); err != nil || len(ranks) == 0{
		// no rank queried
		c.JSON(http.StatusBadRequest, json_models.APIError{
			Code: api_codes.RankNotFound,
			Msg:  api_codes.GetMsg(api_codes.RankNotFound),
		})
	}else{
		// succeeded
		var userRanks []json_models.UserRank
		for _, rank := range ranks{
			game, _ := models.QueryGameByID(rank.GameID)
			userRanks = append(userRanks, json_models.UserRank{
				Score:rank.Score,
				GameInUserRank:json_models.GameInUserRank{
					ID:rank.GameID,
					Name:game.Name,
					LastUpdate:rank.UpdatedAt,
				},
			})
		}
		c.JSON(http.StatusOK, json_models.UserRankResp{
			Code: api_codes.SUCCESS,
			Msg:  api_codes.GetMsg(api_codes.SUCCESS),
			Data: userRanks,
		})
	}
}
