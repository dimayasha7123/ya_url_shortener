package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"ya_url_shortener/internal/app"
)

type redirectHandler struct {
	shortener *app.Shortener
}

func NewRedirectHandler(shortener *app.Shortener) redirectHandler {
	return redirectHandler{
		shortener: shortener,
	}
}

func (h redirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	shortUrl, ok := vars["shortUrl"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		writeErrorAndLog(w, "can't find shortUrl variable", fmt.Errorf("no shortUrl in url"))
		return
	}

	orig, err := h.shortener.GetOrig(shortUrl)
	if err != nil {
		if errors.As(err, &app.ValidationError{}) || errors.As(err, &app.DecodingError{}) {
			w.WriteHeader(http.StatusBadRequest)
			writeErrorAndLog(w, "url not valid", err)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		writeErrorAndLog(w, "can't get orig url", err)
		return
	}

	resp := redirectResp{URL: orig}
	bytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeErrorAndLog(w, "can't marshall body", err)
		return
	}

	writeRespAndLogIfCant(w, bytes)
}
