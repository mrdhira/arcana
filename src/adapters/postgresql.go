package adapters

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mrdhira/arcana/config"
	"github.com/mrdhira/arcana/pkg/logger"
	"go.uber.org/zap"
)

// NewPostgreSQLAdapter func
func NewPostgreSQLAdapter(ctx context.Context, Config *config.PostgresqlConfig) *pgxpool.Pool {
	logger.Info("adapters/postgresql.go/NewPostgreSQLAdapter")

	pool, err := pgxpool.Connect(ctx, Config.ConnectionURL)
	if err != nil {
		logger.Error("error when try to connect postgreSQL pool connection", zap.Error(err))
		panic(err)
	}

	logger.Info("postgreSQL connected")

	return pool
}
