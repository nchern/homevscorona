package srv

import (
	"fmt"
	"net/http"

	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

func getEvents(r *http.Request) (interface{}, error) {
	token, err := authenticate(r.Header)
	if err != nil {
		return nil, err
	}
	user, err := users.GetByEmail(token.Email)
	if user == nil {
		return nil, fmt.Errorf("%s not found", token.Email)
	}

	return &eventsResponse{
		UserName: user.Name,
	}, nil
}

type eventsResponse struct {
	UserName string `json:"user_name"`

	Events []*model.Event `json:"events"`
}

type eventsResponseMock struct{}

func (r *eventsResponseMock) MarshalJSON() ([]byte, error) {
	return []byte(`{
		  "user_name": "John Doe",
		  "events": [
			{
			  "type": "location",
			  "time": 12345678,
			  "details": {
				"location_id": "location-id-1",
				"name": "Rewe",
				"adress": "Berlin ...."
			  }
			},
			{
			  "type": "person",
			  "time": 12345678,
			  "details": {
				"users": [
				  {
					"user_id": "user-id",
					"user_name": "User regitered name",
					"name": "..."
				  }, 
				  {
					"name": "Sarah ."
				  }
				]
			  }
			}
		  ]
}`), nil
}
