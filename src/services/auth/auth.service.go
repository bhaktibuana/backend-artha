package authService

import (
	"api-artha/src/helpers"
	"api-artha/src/models"
	authRequest "api-artha/src/requests/auth"
	"database/sql"
	"net/http"
	"time"

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

func Register(context *gin.Context, request *authRequest.RegisterRequest) *models.Users {
	var user models.Users
	var role models.Roles
	var username, tagLine string

	if !request.Encrypted {
		request.Password = helpers.HashPassword(request.Password)
	}

	if err := models.DB.First(&role, "name = ?", "Common User").Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helpers.Response("No role found with the specified name", http.StatusNotFound, context, nil)
			return nil
		default:
			helpers.Response(err.Error(), http.StatusInternalServerError, context, nil)
			return nil
		}
	}

	for true {
		username, tagLine, _ = helpers.GenerateRandomAccountId()

		if err := models.DB.
			First(&user, "username = ? AND tag_line = ?", username, tagLine).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				break
			}
		}
	}

	user = models.Users{
		Name:      request.Name,
		Username:  username,
		TagLine:   tagLine,
		Email:     request.Email,
		Password:  request.Password,
		BirthDate: sql.NullTime{Time: time.Time{}, Valid: false},
		GenderId:  sql.NullInt64{Int64: 0, Valid: false},
		RoleId:    role.Id,
		Status:    models.USER_STATUS_UNVERIFIED,
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{Time: time.Time{}, Valid: false},
		DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
	}

	if err := models.DB.Create(&user).Error; err != nil {
		helpers.Response("Register failed", http.StatusBadRequest, context, nil)
		return nil
	}

	return &user
}
