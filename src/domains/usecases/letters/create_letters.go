package letters

import (
	"context"
	"time"

	"github.com/mrdhira/arcana/pkg/logger"
	"github.com/mrdhira/arcana/src/domains/models"
	"go.uber.org/zap"
)

// CreateLetters func
func (l *Letters) CreateLetters(ctx context.Context, ReqBody *models.ReqCreateLetters) (ResBody *models.ResCreateLetters, err error) {
	logger.Info("domains/usecases/letters/create_letters.go/CreateLetters")

	Letters := &models.Letters{
		To:        ReqBody.To,
		Text:      ReqBody.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	Letters, err = l.postgreSQL.CreateLetters(ctx, Letters)
	if err != nil {
		logger.Error("error on create letters", zap.Error(err))
		return
	}

	return &models.ResCreateLetters{
		ID:        Letters.ID,
		To:        Letters.To,
		Text:      Letters.Text,
		CreatedAt: Letters.CreatedAt,
	}, nil
}
