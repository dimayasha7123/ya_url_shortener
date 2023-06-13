package http_server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"ya_url_shortener/internal/app"
	"ya_url_shortener/internal/inputs/http_server/handlers"
)

func New(shortener *app.Shortener, addr string) *http.Server {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// api/v1/data/shorten
	subrouter.Handle("/data/shorten", handlers.NewShortenHandler(shortener)).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")
	// api/v1/{shortURL}
	subrouter.Handle("/{shortUrl:[A-Za-z0-9_]+}", handlers.NewRedirectHandler(shortener)).
		Methods(http.MethodGet)

	router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 10,
	}

	return srv
}
