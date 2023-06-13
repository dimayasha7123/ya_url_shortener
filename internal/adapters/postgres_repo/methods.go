package postgres_repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func (r postgresRepo) GetShortenByOrigIfExists(ctx context.Context, orig string) (string, error) {
	query := `
		select short from urls where orig = $1;
	`

	var ret string
	err := r.pool.QueryRow(ctx, query, orig).Scan(&ret)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("can't query shorten by orig: %v", err)
	}

	return ret, nil
}

func (r postgresRepo) SaveNewURL(ctx context.Context, id uint64, short, orig string) error {
	query := `
		insert into urls (id, short, orig)
		values ($1, $2, $3);
	`

	_, err := r.pool.Query(ctx, query, id, short, orig)
	if err != nil {
		return fmt.Errorf("can't query saving new url: %v", err)
	}

	return nil
}

func (r postgresRepo) GetOrigByID(ctx context.Context, id uint64) (string, error) {
	query := `
		select orig from urls
		where id = $1;
	`

	var ret string
	err := r.pool.QueryRow(ctx, query, id).Scan(&ret)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return "", fmt.Errorf("id = %v not found", id)
		}
		return "", fmt.Errorf("can't query orig by id: %v", err)
	}

	return ret, nil
}
