package grpc_server

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ya_url_shortener/internal/app"
	"ya_url_shortener/pkg/api"
)

func (s server) GetOriginal(ctx context.Context, req *api.ShortURL) (*api.OriginalURL, error) {
	orig, err := s.shortener.GetOrig(req.URL)
	if err != nil {
		if errors.As(err, &app.ValidationError{}) || errors.As(err, &app.DecodingError{}) {
			return nil, status.Error(codes.InvalidArgument, "url not valid")
		}
		return nil, status.Error(codes.Internal, "can't get orig url")
	}

	return &api.OriginalURL{URL: orig}, nil
}
