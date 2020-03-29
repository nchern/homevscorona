package srv

import (
	"fmt"

	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

func getEvents(ctx *Context) (interface{}, error) {
	if ctx.AuthenticatedUser == nil {
		return nil, fmt.Errorf("not found: %s", ctx.Token.Email)
	}

	user := ctx.AuthenticatedUser

	events, err := users.GetEvents(user.ID)
	if err != nil {
		return nil, err
	}
	return &eventsResponse{
		responseBase: okResponse,

		Events:   events,
		UserName: user.Name,
	}, nil
}

type eventsResponse struct {
	responseBase

	UserName string `json:"user_name"`

	Events []*model.Event `json:"events"`
}
