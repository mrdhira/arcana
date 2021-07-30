package postgresql

import "github.com/jackc/pgx/v4/pgxpool"

// PostgreSQL struct
type PostgreSQL struct {
	client *pgxpool.Pool
}

func New(postgreSQLClient *pgxpool.Pool) *PostgreSQL {
	return &PostgreSQL{
		client: postgreSQLClient,
	}
}
