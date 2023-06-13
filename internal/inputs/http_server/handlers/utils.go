package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeErrorAndLog(w http.ResponseWriter, errMsg string, errObj error) {
	log.Printf("ERROR %s: %v", errMsg, errObj)

	body := errResp{Error: errMsg}
	bytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("can't marshall body to json: %v", body)
		return
	}

	writeRespAndLogIfCant(w, bytes)
}

func writeRespAndLogIfCant(w http.ResponseWriter, bytes []byte) {
	bytes = append(bytes, '\n')
	_, err := w.Write(bytes)
	if err != nil {
		log.Printf("can't write %q to response writer: %v", string(bytes), err)
	}
}
