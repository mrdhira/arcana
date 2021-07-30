package repositories

import (
	"context"

	"github.com/mrdhira/arcana/src/domains/models"
)

// IPostgreSQL interface
type IPostgreSQL interface {
	// Letters
	CreateLetters(ctx context.Context, Letters *models.Letters) (*models.Letters, error)
}
