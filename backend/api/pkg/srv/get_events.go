package srv

import (
	"fmt"
	"net/http"

	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

func getEvents(r *http.Request) (interface{}, error) {
	token, err := authenticate(r.Header)
	if err != nil {
		return nil, err
	}
	user, err := users.GetByEmail(token.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("%s not found", token.Email)
	}
	events, err := users.GetEvents(user.Id)
	if err != nil {
		return nil, err
	}
	return &eventsResponse{
		UserName: user.Name,
		Events:   events,
	}, nil
}

type eventsResponse struct {
	UserName string `json:"user_name"`

	Events []*model.Event `json:"events"`
}
