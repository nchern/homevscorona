package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID `json:"id,omitempty"`

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`
}

type Event struct {
	Id        uuid.UUID `json:"id"`
	Timestamp time.Time `json:"timestamp"`

	Location *Location `json:"location"`

	Users []*User `json:"users"`
}

type Location struct {
}
