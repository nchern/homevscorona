package srv

import (
	"log"
	"net/http"

	"github.com/nchern/homevscorona/backend/api/pkg/restapi"
)

var listenAddr = ":8080"

var (
	handle        = restapi.Handle
	authenticated = restapi.Authenticated
)

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
	http.HandleFunc("/api/status", handle(restapi.Status))
	http.HandleFunc("/api/signup", handle(authenticated(restapi.Signup)))
	http.HandleFunc("/api/new_event", handle(authenticated(restapi.NewEvent)))
	http.HandleFunc("/api/get_events", handle(authenticated(restapi.GetEvents)))
}
