package srv

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	headerAuthorization = "Authorization"
)

var (
	errAuthFailed = errors.New("auth failed")

	okResponse = &baseResponse{Status: "ok"}
)

type baseResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type AuthToken struct {
	Email string `json:"email"`
}

func authenticate(headers http.Header) (*AuthToken, error) {
	val := headers.Get(headerAuthorization)
	// check bearer
	tokens := strings.Split(val, " ")
	if len(tokens) != 2 {
		return nil, errAuthFailed
	}
	token := strings.TrimSpace(tokens[1])
	if token == "" {
		return nil, errAuthFailed
	}
	return &AuthToken{
		Email: "test@localhost.io",
	}, nil
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
			status = http.StatusInternalServerError
			log.Printf("ERROR %s %s", r.URL, err)

			if err == errAuthFailed {
				status = http.StatusUnauthorized
			}

			resp = errorResponse{Status: fmt.Sprintf("%d", status), Detail: err.Error()}
		}
		w.WriteHeader(status)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("%s %s", r.URL, err)
		}
	}
}
