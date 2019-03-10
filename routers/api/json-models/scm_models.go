package json_models

import "github.com/jackmrzhou/gc-ai-backend/models"

// swagger:parameters uploadSourceCode
type swaggerUploadSourceCodeReq struct {
	// in:body
	Body UploadSourceCodeReq
}

type UploadSourceCodeReq struct {
	// Required:true
	// Example:12
	GameID uint `validate:"required" json:"game_id" binding:"required"`
	// Required:true
	// Example:1
	CodeType int `json:"code_type"`
	// Required:true
	// Example:cpp
	Language string `validate:"required" json:"language" binding:"required"`
	// Required:true
	SourceCode string `validate:"required" json:"source_code" binding:"required"`
}

// operation succeeded, the source code's ID will returned
// swagger:response uploadSourceCodeResp
type swaggerUploadSourceCodeResp struct {
	// in:body
	Body UploadSourceCodeResp
}

type UploadSourceCodeResp struct {
	// response status
	// Example: 200
	Code int `json:"code"`
	// response message
	// Example: ok
	Msg string `json:"message"`
	// contains source code's ID
	Data UploadSourceCodeData `json:"data"`
}

type UploadSourceCodeData struct {
	// source code's ID
	// Example: 23
	SourceCodeID uint `json:"source_code_id"`
}

// contains all the source codes
// swagger:response getSourceCodesResp
type swaggerGetSourceCodes struct {
	// in:body
	Body GetSourceCodes
}

type GetSourceCodes struct {
	// response status
	// Example:200
	Code int `json:"code"`
	// response message
	// Example:ok
	Msg string `json:"message"`
	// source codes
	// Example:[{"user_id":1, "game_id":2, "code_type":1, "language":"cpp", "source_code":"string"}]
	Data []models.SourceCode `json:"data"`
}

// swagger:parameters getSrcByUserAndGame
type GetSrcByUserAndGameReq struct {
	// game id
	//
	// in:query
	// Required:true
	GameID uint `json:"game_id" validate:"required" binding:"required"`
}
