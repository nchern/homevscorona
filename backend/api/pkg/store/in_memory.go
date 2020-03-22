package store

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

type InMemUserStore struct {
	emailToUser StringUserPtrMap
}

func NewInMemUserStore() *InMemUserStore {
	return &InMemUserStore{
		emailToUser: NewStringUserPtrMapSyncronized(),
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

func (u *InMemUserStore) GetById(id uuid.UUID) (*model.User, error) {
	return nil, errors.New("GetById not implemented")
}
