package response

import (
	"github.com/gin-gonic/gin"
)

func HandleResponse(ctx *gin.Context, status Status, err error, data interface{}) {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	ctx.AbortWithStatusJSON(status.Code, ResponseModel{
		Status:       status.Status,
		Description:  status.Description,
		ErrorMessage: errorMessage,
		SnapData:     data,
	})
}

// ResponseModel ...
type ResponseModel struct {
	Status       string      `json:"status"`
	Description  string      `json:"description"`
	SnapData     interface{} `json:"snapData"`
	ErrorMessage string      `json:"error"`
}
