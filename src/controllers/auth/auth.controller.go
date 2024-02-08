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
	request := authRequest.Login(context)
	if request == nil {
		return
	}

	user := authService.Login(context, request)
	if user == nil {
		return
	}

	helpers.Response("Login success", http.StatusOK, context, authResult.Login(user))
}

func Register(context *gin.Context) {
	request := authRequest.Register(context)
	if request == nil {
		return
	}

	user := authService.Register(context, request)
	if user == nil {
		return
	}

	helpers.Response("Register success", http.StatusCreated, context, authResult.Register(user))
}
