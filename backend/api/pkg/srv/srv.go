package srv

import (
	"log"
	"net/http"

	"github.com/nchern/homevscorona/backend/api/pkg/model"
	"github.com/nchern/homevscorona/backend/api/pkg/store"
)

var listenAddr = ":8080"

var (
	users UserStore = store.NewInMemUserStore()
)

type UserStore interface {
	Create(email string, u *model.User) error
	GetByEmail(email string) (*model.User, error)
	GetById(id string) (*model.User, error)
	SaveEvent(userID string, event *model.Event) error
	GetEvents(userID string) ([]*model.Event, error)
}

// Start runs the api server
func Start(name string) {
	log.Printf("%s srv is running", name)

	setRoutes()

	log.Fatal(http.ListenAndServe(listenAddr, nil))

}

// Stop performs all necessary operations for graceful service tear down.
func Stop() {
	log.Println("Stopping...")
}

func setRoutes() {
	http.HandleFunc("/status", handle(status))
	http.HandleFunc("/api/get_events", handle(getEvents))
	http.HandleFunc("/api/signup", handle(signup))
	http.HandleFunc("/api/new_event", handle(newEvent))
}
