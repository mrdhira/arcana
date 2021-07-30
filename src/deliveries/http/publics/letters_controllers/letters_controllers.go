package letters_controllers

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mrdhira/arcana/pkg/logger"
	"github.com/mrdhira/arcana/src/domains/usecases"
	"github.com/mrdhira/arcana/src/domains/usecases/letters"
)

// LettersControllers struct
type LettersControllers struct {
	lettersUsecases usecases.ILetters
}

func NewLettersControllers(
	ctx context.Context,
	postgreSQLClient *pgxpool.Pool,
) *LettersControllers {
	logger.Info("deliveries/http/public/letters_controllers/letters_controllers.go/NewLettersControllers")

	// Initiate Usecases
	lettersUsecases := letters.New(ctx, postgreSQLClient)

	return &LettersControllers{
		lettersUsecases: lettersUsecases,
	}
}
