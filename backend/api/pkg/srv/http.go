package srv

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
)

var (
	okResponse = responseBase{Status: "200"}
)

type responseBase struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Status string `json:"status"`

	Code   string `json:"code,omitempty"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type handler func(*http.Request) (interface{}, error)

func handle(fn handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		log.Printf("INFO %s Authorization: '%s'", r.URL, r.Header.Get(headerAuthorization))

		resp, err := fn(r)

		if err != nil {
			log.Printf("ERROR %s %s", r.URL, err)

			status = http.StatusInternalServerError

			if err == errAuthFailed {
				status = http.StatusUnauthorized
			}
			if err == io.EOF {
				status = http.StatusBadRequest
			}

			resp = errorResponse{Status: fmt.Sprintf("%d", status), Detail: err.Error()}
		}

		w.WriteHeader(status)

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("ERROR %s %s", r.URL, err)
		}
	}
}
