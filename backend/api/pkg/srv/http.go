package srv

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nchern/homevscorona/backend/api/pkg/model"
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

type Context struct {
	Token             *AuthToken
	AuthenticatedUser *model.User
	Request           *http.Request
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
			status, resp = handleError(r, err)
		}

		w.WriteHeader(status)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logError(r, err)
		}
	}
}

type authenticatedHandler func(ctx *Context) (interface{}, error)

func authenticated(fn authenticatedHandler) handler {
	return func(r *http.Request) (interface{}, error) {
		token, err := authenticate(r.Header)
		if err != nil {
			return nil, err
		}

		user, err := users.GetByEmail(token.Email)
		if err != nil {
			return nil, err
		}

		ctx := &Context{
			Request:           r,
			AuthenticatedUser: user,
			Token:             token,
		}
		return fn(ctx)
	}
}

func handleError(r *http.Request, err error) (status int, resp interface{}) {
	logError(r, err)

	status = http.StatusInternalServerError

	if err == errAuthFailed {
		status = http.StatusUnauthorized
	}
	if err == io.EOF {
		status = http.StatusBadRequest
	}

	resp = errorResponse{Status: fmt.Sprintf("%d", status), Detail: err.Error()}
	return
}

func logError(r *http.Request, err error) {
	log.Printf("ERROR %s %s", r.URL, err)
}
