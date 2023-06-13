package app

import "context"

type Repository interface {
	GetShortenByOrigIfExists(ctx context.Context, orig string) (string, error)
	SaveNewURL(ctx context.Context, id uint64, short, orig string) error
	GetOrigByID(ctx context.Context, id uint64) (string, error)
}
