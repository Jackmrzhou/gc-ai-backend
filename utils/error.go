package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jackmrzhou/gc-ai-backend/api-codes"
	"github.com/jackmrzhou/gc-ai-backend/routers/api/json-models"
	"net/http"
)

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, json_models.APIError{
		Code: code,
		Msg:  api_codes.GetMsg(code),
	})
}
