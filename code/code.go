package code

const (
	SUCCESS    = 200
	ERROR      = 500
	INVALID    = 400
	AuthFailed = 2001
	RegFailed = 2002
	Banned	= 2003
	VeriFailed = 2004
	SendMailFailed = 2005
)

var Msgs = map[int] string{
	SUCCESS :   "ok",
	ERROR :     "internal error",
	INVALID :   "invalid parameters",
	AuthFailed: "invalid email or password",
	RegFailed: "registration failed",
	Banned:"User is banned",
	VeriFailed:"Verification failed",
	SendMailFailed:"sending mail failed",
}

func GetMsg(code int) string {
	msg, ok := Msgs[code]
	if ok{
		return msg
	}
	return Msgs[ERROR]
}