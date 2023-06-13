package mem_repo

import (
	"context"
	"fmt"
)

func (r memRepo) GetShortenByOrigIfExists(ctx context.Context, orig string) (string, error) {
	short, ok := r.shortens[orig]
	if !ok {
		return "", nil
	}
	return short, nil
}

func (r memRepo) SaveNewURL(ctx context.Context, id uint64, short, orig string) error {
	_, ok := r.shortens[orig]
	if ok {
		return fmt.Errorf("item with orig = %v already exists", orig)
	}
	_, ok = r.originals[id]
	if ok {
		return fmt.Errorf("item with id = %v already exists", id)
	}

	r.shortens[orig] = short
	r.originals[id] = orig

	return nil
}

func (r memRepo) GetOrigByID(ctx context.Context, id uint64) (string, error) {
	orig, ok := r.originals[id]
	if !ok {
		return "", fmt.Errorf("no item with id = %v", id)
	}

	return orig, nil
}
