package grpc_server

import (
	"ya_url_shortener/internal/app"
	"ya_url_shortener/pkg/api"
)

type server struct {
	shortener *app.Shortener
	api.UnimplementedURLShortenerServer
}

func New(shortener *app.Shortener) server {
	return server{shortener: shortener}
}
