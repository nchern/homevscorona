package srv

import (
	"log"
	"net/http"
)

var listenAddr = ":8080"

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
}
