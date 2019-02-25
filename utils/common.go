package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/api-codes"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/json-models"
	"log"
	"net/http"
)

func SuccessResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, json_models.Status{
		Code:code,
		Msg:api_codes.GetMsg(code),
	})
}

func LogIfNotNil(err error) {
	if err != nil{
		log.Fatal(err)
	}
}
