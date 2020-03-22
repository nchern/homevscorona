package srv

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
)

type errorResponse struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type handler func(*http.Request) (interface{}, error)

func handle(fn handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		w.Header()["Content-Type"] = []string{"application/json"}

		log.Printf("%s Authorization: '%s'", r.URL, r.Header.Get(headerAuthorization))

		resp, err := fn(r)

		if err != nil {
			log.Printf("%s %s", r.URL, err)
			resp = errorResponse{Status: "500", Detail: err.Error()}
			status = http.StatusInternalServerError
		}
		w.WriteHeader(status)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("%s %s", r.URL, err)
		}
	}
}
