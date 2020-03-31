package restapi

import (
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

func (er newEventRequest) GetTime() time.Time {
	if er.Time == 0 {
		return time.Now()
	}
	return time.Unix(er.Time, 0)
}

type eventDetails struct {
	Users []*model.User `json:"users"`
}

func NewEvent(ctx *Context) (interface{}, error) {
	if ctx.AuthenticatedUser == nil {
		return nil, errUnknownUserToken(ctx.Token)
	}

	var req newEventRequest
	if err := ctx.ParseJSONBody(&req); err != nil {
		return nil, err
	}

	event := &model.Event{ID: uuid.New().String(), Timestamp: req.GetTime()}

	switch req.Type {
	case eventPerson:
		if len(req.Details.Users) == 0 {
			return nil, fmt.Errorf("%w: details nil or empty", errValidationFailed)
		}
		for _, u := range req.Details.Users {
			event.Users = append(event.Users, u)
		}

	default:
		return nil, fmt.Errorf("%w: unknown event type: '%s'", errValidationFailed, req.Type)
	}

	// TODO: add various validations

	if err := users.SaveEvent(ctx.AuthenticatedUser.ID, event); err != nil {
		return nil, err
	}

	return okResponse, nil
}
