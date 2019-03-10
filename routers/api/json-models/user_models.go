package json_models

import "github.com/jackmrzhou/gc-ai-backend/models"

type UpdateProfileReq struct {
	Nickname string	`json:"nickname" validate:"required" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
}

type GetProfileResp struct {
	// response code
	// Example:200
	Code int `json:"code"`
	// response status
	// Example:ok
	Msg string `json:"message"`
	// contains profile
	Data *models.Profile `json:"data"`
}

type TestNicknameReq struct {
	Nickname string `json:"nickname" validate:"required" binding:"required"`
}

type TestNicknameResp struct {
	// response code
	// Example:200
	Code int `json:"code"`
	// response status
	// Example:ok
	Msg string `json:"message"`
	//contains result
	Data TestNicknameData `json:"data"`
}

type TestNicknameData struct {
	Exists bool `json:"exists"`
}