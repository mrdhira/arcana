package helpers

import "github.com/gin-gonic/gin"

// HTTPResponse
type HTTPResponse struct {
	Code    int         `json:"code"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SetHTTPResponse func
func SetHTTPResponse(ctx *gin.Context, httpStatus int, code int, error string, message string, data interface{}) {
	ctx.JSON(
		httpStatus,
		&HTTPResponse{
			Code:    code,
			Error:   error,
			Message: message,
			Data:    data,
		},
	)
}
