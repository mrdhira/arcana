package letters_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mrdhira/arcana/pkg/helpers"
	"github.com/mrdhira/arcana/pkg/logger"
	"github.com/mrdhira/arcana/src/domains/models"
	"go.uber.org/zap"
)

// CreateLetters controllers
func (c *LettersControllers) CreateLetters(ctx *gin.Context) {
	logger.Info("deliveries/http/public/letters_controllers.go/CreateLetters")

	var ReqBody *models.ReqCreateLetters
	if err := ctx.ShouldBindWith(&ReqBody, binding.JSON); err != nil {
		logger.Error("error when try to validate request body", zap.Error(err))
		helpers.SetHTTPResponse(ctx, http.StatusBadRequest, http.StatusBadRequest, err.Error(), "something when wrong", nil)
	}

	// ReqBodyByte, err := ioutil.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	logger.Error("error when try to read request body", zap.Error(err))
	// 	helpers.SetHTTPResponse(ctx, http.StatusBadRequest, http.StatusBadRequest, err.Error(), "something when wrong", nil)
	// }

	// var ReqBody *models.ReqCreateLetters
	// err = json.Unmarshal(ReqBodyByte, &ReqBody)
	// if err != nil {
	// 	logger.Error("error when try to unmarshal request body", zap.Error(err))
	// 	helpers.SetHTTPResponse(ctx, http.StatusBadRequest, http.StatusBadRequest, err.Error(), "something when wrong", nil)
	// }

	Letters, err := c.lettersUsecases.CreateLetters(ctx.Request.Context(), ReqBody)
	if err != nil {
		logger.Error("error when process create letters", zap.Error(err))
		helpers.SetHTTPResponse(ctx, http.StatusBadRequest, http.StatusBadRequest, err.Error(), "something when wrong", nil)
	}

	helpers.SetHTTPResponse(ctx, http.StatusOK, http.StatusOK, "", "Success", Letters)
}
