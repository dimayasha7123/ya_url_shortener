package grpc_server

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ya_url_shortener/internal/app"
	"ya_url_shortener/pkg/api"
)

func (s server) Shorten(ctx context.Context, req *api.OriginalURL) (*api.ShortURL, error) {
	short, err := s.shortener.Shorten(req.URL)
	if err != nil {
		if errors.As(err, &app.ValidationError{}) {
			return nil, status.Error(codes.InvalidArgument, "url not valid")
		}
		return nil, status.Errorf(codes.Internal, "can't shorten url")
	}

	return &api.ShortURL{URL: short}, nil
}
