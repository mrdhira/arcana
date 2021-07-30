package publics_controllers

import "github.com/gin-gonic/gin"

// ILettersControllers interface
type ILettersControllers interface {
	CreateLetters(g *gin.Context)
}
