package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"ya_url_shortener/internal/app"
)

type shortenHandler struct {
	shortener *app.Shortener
}

func NewShortenHandler(shortener *app.Shortener) shortenHandler {
	return shortenHandler{
		shortener: shortener,
	}
}

func (h shortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeErrorAndLog(w, "can't read request body", err)
		return
	}

	err = r.Body.Close()
	if err != nil {
		log.Printf("can't close request body: %v", err)
	}

	var data shortenBody
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeErrorAndLog(w, "can't unmarshall body", err)
		return
	}

	shorten, err := h.shortener.Shorten(data.URL)
	if err != nil {
		if errors.As(err, &app.ValidationError{}) {
			w.WriteHeader(http.StatusBadRequest)
			writeErrorAndLog(w, "url not valid", err)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		writeErrorAndLog(w, "can't shorten url", err)
		return
	}

	resp := shortenResp{URL: shorten}
	bytes, err = json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeErrorAndLog(w, "can't marshall body", err)
		return
	}

	writeRespAndLogIfCant(w, bytes)
}
