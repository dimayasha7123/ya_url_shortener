package http_server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		out := fmt.Sprintf("%s %s ", r.Method, r.URL)
		if r.Method == http.MethodPost {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				out += fmt.Sprintf("body unavailable: %v", err)
			} else {
				out += string(body)
			}
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		log.Println(out)

		next.ServeHTTP(w, r)
	})
}
