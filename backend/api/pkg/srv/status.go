package srv

import (
	"net/http"
)

func status(r *http.Request) (interface{}, error) {
	return &struct {
		Status string `json:"status"`
	}{"ok"}, nil
}
