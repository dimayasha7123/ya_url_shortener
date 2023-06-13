package app

import (
	"fmt"
	"ya_url_shortener/internal/domain/base63"
	"ya_url_shortener/internal/domain/id_gen"
)

type Shortener struct {
	repo    Repository
	idGen   id_gen.IDGenerator
	encoder base63.Encoder
}

func New(nodeNum int64, repo Repository) (*Shortener, error) {
	gen, err := id_gen.New(nodeNum)
	if err != nil {
		return nil, fmt.Errorf("can't create new IDGenerator: %v", err)
	}
	return &Shortener{
		repo:    repo,
		idGen:   gen,
		encoder: base63.New(),
	}, nil
}
