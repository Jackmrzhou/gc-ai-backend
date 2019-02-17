package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/api-codes"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/json-models"
	"github.com/jackmrzhou/gc-ai-backend/utils"
	"net/http"
	"time"
)

func JwtValidation() gin.HandlerFunc{
	return func(c *gin.Context) {
		code := api_codes.SUCCESS
		token := c.GetHeader("Authorization")
		if token == ""{
			// no token
			code = api_codes.UnAuth
		}else{
			claim, err := utils.ParseToken(token)
			if err != nil{
				// invalid token
				code = api_codes.InvaildToken
			}else if time.Now().Unix() > claim.ExpiresAt{
				// token is expired
				code = api_codes.AuthTimeOut
			}
		}

		if code != api_codes.SUCCESS{
			// auth failed
			c.JSON(http.StatusBadRequest, json_models.Status{
				Code:code,
				Msg:api_codes.GetMsg(code),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
