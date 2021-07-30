package http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mrdhira/arcana/pkg/logger"
	publicLettersControllers "github.com/mrdhira/arcana/src/deliveries/http/publics/letters_controllers"
)

// NewRoutes func
func NewRoutes(
	postgreSQLClient *pgxpool.Pool,
) *gin.Engine {
	logger.Info("deliveries/http/routes.go/NewRoutes")

	// Initiate Controllers
	PublicLettersControllers := publicLettersControllers.NewLettersControllers(context.Background(), postgreSQLClient)

	router := gin.Default()

	// Public Routes
	publicRoutes := router.Group("/public")
	{
		// Letters
		publicRoutes.POST("/letters", PublicLettersControllers.CreateLetters)
	}

	return router
}
