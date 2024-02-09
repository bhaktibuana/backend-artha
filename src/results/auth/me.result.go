package authResult

import (
	"api-artha/src/models"
	"fmt"
)

type MeResult struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	TagLine   string `json:"tag_line"`
	AccountId string `json:"account_id"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	RoleId    int64  `json:"role_id"`
	Role      string `json:"role"`
	ImageUrl  string `json:"image_url"`
}

func Me(user *models.Users) MeResult {
	return MeResult{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		TagLine:   user.TagLine,
		AccountId: fmt.Sprintf("%s#%s", user.Username, user.TagLine),
		Email:     user.Email,
		Status:    user.Status,
		RoleId:    user.RoleId,
		Role:      user.Role.Name,
		ImageUrl:  user.ImageUrl.String,
	}
}
