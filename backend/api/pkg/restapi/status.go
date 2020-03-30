package restapi

import (
	"net/http"
)

func Status(r *http.Request) (interface{}, error) {
	return okResponse, nil
}
