package helpers

import "github.com/gin-gonic/gin"

type TResponse struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func Response(message string, httpStatus int, context *gin.Context, data interface{}) {
	response := TResponse{
		Message: message,
		Status:  httpStatus >= 200 && httpStatus < 300,
		Data:    data,
	}

	context.JSON(httpStatus, response)
}
