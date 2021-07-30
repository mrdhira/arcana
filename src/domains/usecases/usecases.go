package usecases

import (
	"context"

	"github.com/mrdhira/arcana/src/domains/models"
)

// ILetters interface
type ILetters interface {
	CreateLetters(ctx context.Context, ReqBody *models.ReqCreateLetters) (ResBody *models.ResCreateLetters, err error)
}
