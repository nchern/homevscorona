package restapi

import (
	"github.com/nchern/homevscorona/backend/api/pkg/model"
	"github.com/nchern/homevscorona/backend/api/pkg/store"
	"github.com/nchern/homevscorona/backend/api/pkg/store/pgstore"
)

var (
	users UserStore = store.NewInMemUserStore()
)

func init() {
	users = pgstore.New()
}

type UserStore interface {
	Create(email string, u *model.User) error
	GetByEmail(email string) (*model.User, error)
	GetByID(id string) (*model.User, error)
	SaveEvent(userID string, event *model.Event) error
	GetEvents(userID string) ([]*model.Event, error)
}
