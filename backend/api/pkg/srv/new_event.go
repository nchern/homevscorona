package srv

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

const (
	eventPerson   = "person"
	eventLocation = "location"
)

type newEventRequest struct {
	Type    string       `json:"type"`
	Time    int64        `json:"time"`
	Details eventDetails `json:"details"`
}

type eventDetails struct {
	Users []*model.User `json:"users"`
}

func newEvent(ctx *Context) (interface{}, error) {
	if ctx.AuthenticatedUser == nil {
		return nil, fmt.Errorf("not found: %s", ctx.Token.GetEmail())
	}

	var req newEventRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		return nil, err
	}
	if req.Type != eventPerson {
		// TODO: proper error, should be 400
		return nil, errors.New("unknown event type: " + req.Type)
	}
	if len(req.Details.Users) == 0 {
		// TODO: proper error, should be 400
		return nil, errors.New("empty details")
	}
	// TODO: add various validations

	event := &model.Event{ID: uuid.New().String(), Timestamp: time.Unix(req.Time, 0)}
	for _, u := range req.Details.Users {
		event.Users = append(event.Users, u)
	}
	if err := users.SaveEvent(ctx.AuthenticatedUser.ID, event); err != nil {
		return nil, err
	}

	return okResponse, nil
}
