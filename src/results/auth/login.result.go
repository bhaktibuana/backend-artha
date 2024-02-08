package authResult

import (
	"api-artha/src/helpers"
	"api-artha/src/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginResult struct {
	Id      int64  `json:"id"`
	ArthaId string `json:"artha_id"`
	Token   string `json:"token"`
}

func Login(user *models.Users) LoginResult {
	claims := jwt.MapClaims{
		"id":       user.Id,
		"email":    user.Email,
		"tag_line": user.TagLine,
		"username": user.Username,
		"name":     user.Name,
	}

	token, _ := helpers.GenerateJWT(claims, time.Hour*24*30)

	return LoginResult{
		Id:      user.Id,
		ArthaId: user.Username + "#" + user.TagLine,
		Token:   token,
	}
}
