package model

import (
	"time"
)

type User struct {
	ID string `json:"id,omitempty"`

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`
}

type Event struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`

	Location *Location `json:"location"`

	Users []*User `json:"users"`
}

type Location struct {
}
