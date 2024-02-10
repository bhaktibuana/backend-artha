package authRequest

import (
	"api-artha/src/helpers"
	"api-artha/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Encrypted bool   `json:"encrypted"`
}

func Register(context *gin.Context) *RegisterRequest {
	var request RegisterRequest
	var user models.Users

	if err := context.ShouldBindJSON(&request); err != nil {
		helpers.Response(err.Error(), http.StatusBadRequest, context, nil)
		return nil
	}

	if err := models.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &request
		}
	}

	helpers.Response("Email already exist", http.StatusConflict, context, nil)
	return nil
}
