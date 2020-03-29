package srv

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

type signupRequest struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

func signup(r *http.Request) (interface{}, error) {
	token, err := authenticate(r.Header)
	if err != nil {
		return nil, err
	}
	var req signupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	user := &model.User{ID: uuid.New().String(), Email: token.Email, Name: req.Name}
	err = users.Create(token.Email, user)
	if err != nil {
		return nil, err
	}

	return okResponse, nil
}
