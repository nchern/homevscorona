package store

import (
	"errors"

	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

type eventEntry struct {
	Events []*model.Event
}

type InMemUserStore struct {
	emailToUser    StringUserPtrMap
	userIDToEvents StringEventEntryPtrMap
}

func NewInMemUserStore() *InMemUserStore {
	return &InMemUserStore{
		emailToUser:    NewStringUserPtrMapSyncronized(),
		userIDToEvents: NewStringEventEntryPtrMapSyncronized(),
	}

}

func (u *InMemUserStore) Create(email string, user *model.User) error {
	// Race codition
	found, _ := u.GetByEmail(email)
	if found != nil {
		return nil
	}
	u.emailToUser.Set(email, user)
	return nil
}

func (u *InMemUserStore) GetByEmail(email string) (*model.User, error) {
	user, found := u.emailToUser.Get(email)
	if found {
		return user, nil
	}
	return nil, nil
}

func (u *InMemUserStore) GetByID(id string) (*model.User, error) {
	return nil, errors.New("GetById not implemented")
}

func (u *InMemUserStore) SaveEvent(userID string, event *model.Event) error {
	key := userID
	entry, found := u.userIDToEvents.Get(key)
	if !found {
		entry := &eventEntry{Events: []*model.Event{event}}
		u.userIDToEvents.Set(key, entry)
		return nil
	}
	entry.Events = append(entry.Events, event)
	return nil
}

func (u *InMemUserStore) GetEvents(userID string) ([]*model.Event, error) {
	key := userID
	entry, found := u.userIDToEvents.Get(key)
	if !found {
		return []*model.Event{}, nil
	}
	return entry.Events, nil
}
