package json_models

// swagger:parameters getAuth
type ReqUserJSON struct {
	//in:body
	Body UserJSON
}

type UserJSON struct {
	// Required:true
	// Example:test@email.com
	Email    string `validate:"required, email" json:"email" binding:"required"`
	// Required:true
	// Example:123456
	Password string `validate:"required" json:"password" binding:"required"`
}

// authentication passed response
// swagger:response authSuccess
type RespAuthSuccess struct {
	// in:body
	Body AuthSuccess
}

type AuthSuccess struct {
	// response status
	// Example : 200
	Code int `json:"code"`
	// response message
	// Example: ok
	Msg string `json:"message"`
	// data contains jwt token
	Data AuthSuccessData `json:"data"`
}

type AuthSuccessData struct {
	// jwt token
	Token string `json:"token"`
}

// swagger:parameters register
type ReqRegInfo struct {
	//in:body
	Body RegInfo
}


type RegInfo struct {
	// registration information
	// Required:true
	// Example:test@email.com
	Email    string `validate:"required, email" json:"email" binding:"required"`
	// Required:true
	// Example:123456
	Password string `validate:"required" json:"password" binding:"required"`
	// Required:true
	// Example:778899
	VeriCode string `validate:"required" json:"code" binding:"required"`
}

// swagger:parameters sendVeriCode
type ReqTargetEmail struct {
	//in:body
	Body TargetEmail
}

type TargetEmail struct {
	// Required:true
	// Example:test@email.com
	Email string `validate:"required, email" json:"email" binding:"required"`
}

