package authController

import (
	"api-artha/src/helpers"
	authRequest "api-artha/src/requests/auth"
	authResult "api-artha/src/results/auth"
	authService "api-artha/src/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	loginRequest := authRequest.Login(context)
	if loginRequest == nil {
		return
	}

	user := authService.Login(context, loginRequest)
	if user == nil {
		return
	}

	helpers.Response("Login success", http.StatusOK, context, authResult.Login(user))
}
