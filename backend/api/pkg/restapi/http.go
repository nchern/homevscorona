package restapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

const (
	headerAuthorization = "Authorization"
)

var (
	errNotFound         = errors.New("not found")
	errValidationFailed = errors.New("invalid params")

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
	Token             Token
	AuthenticatedUser *model.User
	Request           *http.Request
}

func (c *Context) ParseJSONBody(req interface{}) error {
	if err := json.NewDecoder(c.Request.Body).Decode(req); err != nil {
		return fmt.Errorf("%w: json: %s", errValidationFailed, err)
	}
	return nil
}

type Handler func(*http.Request) (interface{}, error)

func Handle(fn Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.WriteHeader(status)

			return
		}

		log.Printf("INFO %s Authorization: len=%d", r.URL, len(r.Header.Get(headerAuthorization)))

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

type AuthenticatedHandler func(ctx *Context) (interface{}, error)

func Authenticated(fn AuthenticatedHandler) Handler {
	return func(r *http.Request) (interface{}, error) {
		token, err := authenticate(r.Header)
		if err != nil {
			return nil, err
		}

		user, err := users.GetByEmail(token.GetEmail())
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

	if errors.Is(err, errAuthFailed) {
		status = http.StatusUnauthorized
	}
	if errors.Is(err, errValidationFailed) ||
		errors.Is(err, errNotFound) {
		status = http.StatusBadRequest
	}

	resp = errorResponse{Status: fmt.Sprintf("%d", status), Detail: err.Error()}
	return
}

func logError(r *http.Request, err error) {
	log.Printf("ERROR %s %s", r.URL, err)
}
