package restapi

import (
	"github.com/google/uuid"
	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

type signupRequest struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

func Signup(ctx *Context) (interface{}, error) {
	var req signupRequest
	if err := ctx.ParseJSONBody(&req); err != nil {
		return nil, err
	}

	email := ctx.Token.GetEmail()
	user := &model.User{ID: uuid.New().String(), Email: email, Name: req.Name}

	err := users.Create(email, user)
	if err != nil {
		return nil, err
	}

	return okResponse, nil
}
