package srv

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

type signupRequest struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

func signup(ctx *Context) (interface{}, error) {
	var req signupRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		return nil, err
	}

	email := ctx.Token.Email
	user := &model.User{ID: uuid.New().String(), Email: email, Name: req.Name}

	err := users.Create(email, user)
	if err != nil {
		return nil, err
	}

	return okResponse, nil
}
