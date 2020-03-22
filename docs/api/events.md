Endpoint for receiving all recent Events (maybe last 14days?) of a User

### Request:

```http
POST /api/get_events
Authorization: Bearer <token>

{"start_date": 12345678, "end_date": 12345678}  # Both fields are optional

HTTP 200 OK

Content-Type: application/json

{
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
            // no "user_id" for unregistered user
            "name": "Sarah ."
          }
        ]
      }
    }
  ]
}
```

### Error responses

#### Client error
```http
HTTP/1.1 400

{ "status": "400", "title": "You provided some wrong data" }
```

#### Server error
```http
HTTP/1.1 500

{ "status": "500", "title": "Something went wrong" }
```
