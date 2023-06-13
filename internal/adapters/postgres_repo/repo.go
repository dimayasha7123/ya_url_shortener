package postgres_repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresRepo struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (postgresRepo, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return postgresRepo{}, fmt.Errorf("can't create pool: %v", err)
	}

	return postgresRepo{pool: pool}, nil
}
