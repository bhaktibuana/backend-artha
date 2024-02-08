package authService

import (
	"api-artha/src/helpers"
	"api-artha/src/models"
	authRequest "api-artha/src/requests/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(context *gin.Context, request *authRequest.LoginRequest) *models.Users {
	var user models.Users

	if !request.Encrypted {
		request.Password = helpers.HashPassword(request.Password)
	}

	if err := models.DB.
		Preload("Role").
		Joins("JOIN roles ON users.role_id = roles.id").
		Preload("Gender").
		First(&user, "email = ? AND password = ? AND status <> ?", request.Email, request.Password, "deleted").
		Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helpers.Response("Wrong email or password", http.StatusBadRequest, context, nil)
			return nil
		default:
			helpers.Response(err.Error(), http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	if user.Status == "unverified" {
		helpers.Response("Unverified email", http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}
