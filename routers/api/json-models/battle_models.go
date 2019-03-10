package json_models

import "github.com/jackmrzhou/gc-ai-backend/models"

// swagger:parameters startBattle
type swaggerStartBattleReq struct {
	// in:body
	Body StartBattleReq
}

type StartBattleReq struct {
	// Required:true
	// Example: 11
	GameID uint `validate:"required" json:"game_id" binding:"required"`
	// Required:true
	// Example: 15
	UserID uint	`validate:"required" json:"user_id" binding:"required"`
}

// operation succeeded, the info of the battle will be returned
// swagger:response startBattleResp
type swaggerStartBattleResp struct {
	// in:body
	Body StartBattleResp
}

type StartBattleResp struct {
	// response code
	// Example:200
	Code int `json:"code"`
	// response status
	// Example:ok
	Msg string `json:"message"`
	// contains the information of the battle created
	Data StartBattleData `json:"data"`
}

type StartBattleData struct {
	// battle id
	// Example:13
	BattleID uint `json:"battle_id"`
	// Current Battle Status: 1 == suspending, 2 == judging, 3 == finished
	// Example: 2
	Status uint `json:"status"`
}

// swagger:parameters queryProcess
type QueryProcessReq struct {
	BattleID uint `validate:"required" json:"battle_id" binding:"required"`
}

// query succeeded, the details of a battle will be returned
// swagger:response queryProcessResp
type swaggerQueryProcessResp struct {
	// in:body
	Body QueryProcessResp
}

type QueryProcessResp struct {
	// response status
	// Example:200
	Code int `json:"code"`
	// response message
	// Example:ok
	Msg string `json:"message"`
	// details of a battle
	// Example:{"status":2, "attacker_id":13, "defender_id": 14, "game_id":15, "detail":"string", "winner_id":0, "reward_score":0, "penalty_score":0}
	Data models.Battle `json:"data"`
}

// contains all battles of a user
// swagger:response getUserBattlesResp
type swaggerGetUserBattlesResp struct {
	// in:body
	Body GetUserBattlesResp
}

type GetUserBattlesResp struct {
	// response status
	// Example:200
	Code int `json:"code"`
	// response message
	// Example:ok
	Msg string `json:"message"`
	// battles
	// Example:[{"attacker_id":13, "defender_id": 14, "game_id":15, "battle_id":22}]
	Data []BriefBattleInfo `json:"data"`
}

type BriefBattleInfo struct {
	BattleID uint `json:"battle_id"`
	GameID uint `json:"game_id"`
	AttackerID uint `json:"attacker_id"`
	DefenderID uint `json:"defender_id"`
}