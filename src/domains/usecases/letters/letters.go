package letters

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mrdhira/arcana/pkg/logger"
	"github.com/mrdhira/arcana/src/domains/repositories"
	"github.com/mrdhira/arcana/src/domains/repositories/postgresql"
)

// Letters struct
type Letters struct {
	postgreSQL repositories.IPostgreSQL
}

func New(
	ctx context.Context,
	postgreSQLClient *pgxpool.Pool,
) *Letters {
	logger.Info("domains/usecases/letters/letters.go")

	// Init Repositories
	postgreSQL := postgresql.New(postgreSQLClient)

	return &Letters{
		postgreSQL: postgreSQL,
	}
}
