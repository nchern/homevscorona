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

func newEvent(r *http.Request) (interface{}, error) {
	token, err := authenticate(r.Header)
	if err != nil {
		return nil, err
	}

	var req newEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	user, err := users.GetByEmail(token.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("%s not found", token.Email)
	}

	event := &model.Event{ID: uuid.New().String(), Timestamp: time.Unix(req.Time, 0)}
	for _, u := range req.Details.Users {
		event.Users = append(event.Users, u)
	}
	if err := users.SaveEvent(user.ID, event); err != nil {
		return nil, err
	}

	return okResponse, nil
}
