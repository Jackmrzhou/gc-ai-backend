package api_codes

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID        = 400
	AuthFailed     = 2001
	RegFailed      = 2002
	Banned         = 2003
	VeriFailed     = 2004
	SendMailFailed = 2005
	NewGameFailed  = 2006
	RankNotFound   = 2007
	UnAuth = 2008
	AuthTimeOut = 2009
	InvaildToken = 2100
	NewSrcFailed = 2010
	SourceCodesNotFound = 2011
	ObjectUnavailable = 2012
	ObjectNotFound	= 2013
	WaitForFinish = 2014
	StartBattleFailed = 2015
)

var Msgs = map[int] string{
	SUCCESS :       "ok",
	ERROR :         "internal error",
	INVALID :       "invalid parameters",
	AuthFailed:     "invalid email or password",
	RegFailed:      "registration failed",
	Banned:         "User is banned",
	VeriFailed:     "Verification failed",
	SendMailFailed: "sending mail failed",
	NewGameFailed:  "create new game failed",
	RankNotFound:   "rank not found",
	UnAuth:"Unauthorized",
	AuthTimeOut :"Authorization is expired",
	InvaildToken : "Token is invalid",
	NewSrcFailed : "Uploading source code failed",
	SourceCodesNotFound:"no source codes found",
	ObjectUnavailable:"object is unavailable",
	ObjectNotFound :"object not found",
	WaitForFinish : "you still have programs in execution",
	StartBattleFailed:"start battle failed",
}

func GetMsg(code int) string {
	msg, ok := Msgs[code]
	if ok{
		return msg
	}
	return Msgs[ERROR]
}