package postgresql

import (
	"context"

	"github.com/mrdhira/arcana/pkg/logger"
	"github.com/mrdhira/arcana/src/domains/models"
	"go.uber.org/zap"
)

// CreateLetters func
func (pql *PostgreSQL) CreateLetters(ctx context.Context, Letters *models.Letters) (*models.Letters, error) {
	logger.Info("domains/repositories/postgresql/create_letters.go/CreateLetters")

	err := pql.client.
		QueryRow(
			ctx,
			`
			INSERT INTO letters (to, text, created_at, updated_at)
			VALUES($1. $2, $3, $4)
			RETURNING id
			`,
			Letters.To,
			Letters.Text,
			Letters.CreatedAt,
			Letters.UpdatedAt,
		).Scan(&Letters.ID)
	if err != nil {
		logger.Error("error when try to create letters", zap.Error(err))
		return Letters, err
	}

	return Letters, err
}
