package authResult

import (
	"api-artha/src/models"
)

type RegisterResult struct {
	Id int64 `json:"id"`
}

func Register(user *models.Users) RegisterResult {
	return RegisterResult{
		Id: user.Id,
	}
}
