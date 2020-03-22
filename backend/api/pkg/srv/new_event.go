package srv

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

type newEventRequest struct {
	Type    string      `json:"type"`
	Date    int64       `json:"date"`
	Details *eventUsers `json:"details"`
}

type eventUsers struct {
	Users []*model.User
}

func newEvent(r *http.Request) (interface{}, error) {
	token, err := authenticate(r.Header)
	if err != nil {
		return nil, err
	}

	var req newEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	if req.Details == nil {
		// TODO: proper error, should be 400
		return nil, errors.New("empty details")
	}
	// TODO: add various validations

	user, err := users.GetByEmail(token.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("%s not found", token.Email)
	}

	event := &model.Event{ID: uuid.New().String(), Timestamp: time.Unix(req.Date, 0)}
	for _, u := range req.Details.Users {
		event.Users = append(event.Users, u)
	}
	if err := users.SaveEvent(user.ID, event); err != nil {
		return nil, err
	}

	return okResponse, nil
}
