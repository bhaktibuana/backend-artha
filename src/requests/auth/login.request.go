package authRequest

import (
	"api-artha/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Encrypted bool   `json:"encrypted"`
}

func Login(context *gin.Context) *LoginRequest {
	var loginRequest LoginRequest

	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		helpers.Response(err.Error(), http.StatusBadRequest, context, nil)
		return nil
	}

	return &loginRequest
}
