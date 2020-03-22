package srv

import (
	"net/http"
)

func getEvents(r *http.Request) (interface{}, error) {
	return &eventsResponseMock{}, nil
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
