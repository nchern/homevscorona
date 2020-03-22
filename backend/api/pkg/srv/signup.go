package srv

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

func signup(r *http.Request) (interface{}, error) {
	token, err := authenticate(r.Header)
	if err != nil {
		return nil, err
	}

	user := &model.User{Id: uuid.New(), Email: token.Email, Name: "John Doe"}
	err = users.Create(token.Email, user)
	if err != nil {
		return nil, err
	}

	return okResponse, nil
}
