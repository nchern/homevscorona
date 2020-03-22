package srv

import (
	"net/http"
)

func status(r *http.Request) (interface{}, error) {
	return okResponse, nil
}
