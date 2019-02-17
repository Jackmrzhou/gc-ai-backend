package json_models

import (
	"github.com/jackmrzhou/gc-ai-backend/models"
	"time"
)

// contains all games
// swagger:response allGame
type swaggerAllGamesResp struct {
	// in:body
	Body AllGamesResp
}

// swagger:parameters newGame
type swaggerNewGameReq struct {
	//in:body
	Body NewGameReq
}

type NewGameReq struct {
	// Required:true
	// Example:chess
	Name string `validate:"required" json:"name" binding:"required"`
	// Required:true
	// Example:chess game
	Introduction string `validate:"required" json:"introduction" binding:"required"`
}

type AllGamesResp struct {
	// response status
	// Example: 200
	Code int `json:"code"`
	// response message
	// Example: ok
	Msg string `json:"message"`
	// all the games
	// Example: [{"id":1,"name":"chess", "introduction":"chess game"}]
	Data []models.Game `json:"data"`
}


// swagger:parameters getGameRank
type GetGameRankReq struct {
	// game id
	//
	// in:query
	// Required:true
	ID uint `form:"game_id" binding:"required" validate:"required"`
}

type GameRank struct {
	// score
	// Example:1200
	Score uint `json:"score"`
	UserInGameRank
}

type UserInGameRank struct {
	// user id
	// Example:12
	ID uint `json:"user_id"`
	// user's nickname
	// Example:jack
	Nickname string `json:"nickname"`
	// the time of last change of score
	// Example:2019-02-13T22:17:05+08:00
	LastUpdate time.Time `json:"last_update"`
}

// contains ranks of a game, these ranks are not sorted
// swagger:response gameRank
type swaggerGameRankResp struct {
	//in:body
	Body GameRankResp
}

type GameRankResp struct {
	// response status
	// Example:200
	Code int `json:"code"`
	// response message
	// Example:ok
	Msg string `json:"message"`
	// ranks of a game
	// Example:[{"score":10,"user_id":34,"nickname":"jack","last_update":"2019-02-13T22:17:05+08:00"}]
	Data []GameRank `json:"data"`
}

// swagger:parameters getUserRank
type GetUserRankReq struct {
	// user id
	//
	// in:query
	// Required:true
	ID uint `form:"user_id" binding:"required" validate:"required"`
}

type UserRank struct {
	// score
	// Example:1200
	Score uint `json:"score"`
	GameInUserRank
}

type GameInUserRank struct {
	// game id
	// Example:13
	ID uint `json:"game_id"`
	// game's name
	// Example:chess
	Name string `json:"game_name"`
	// the time of last change of score
	// Example:2019-02-13T22:17:05+08:00
	LastUpdate time.Time `json:"last_update"`
}

type UserRankResp struct {
	// response status
	// Example: 200
	Code int `json:"code"`
	// response message
	// Example:ok
	Msg string `json:"message"`
	// ranks of a user
	// Example:[{"score":10,"game_id":19,"game_name":"game 0","last_update":"2019-02-13T22:17:05+08:00"}]
	Data []UserRank `json:"data"`
}

//contains ranks of a user
// swagger:response userRank
type swaggerUserRankResp struct {
	//in:body
	Body UserRankResp
}