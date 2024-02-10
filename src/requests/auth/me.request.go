package authRequest

import (
	"api-artha/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(context *gin.Context) *string {
	id := context.Param("id")

	if id == ":id" {
		helpers.Response("Param 'id' is required", http.StatusBadRequest, context, nil)
		return nil
	}

	return &id
}
