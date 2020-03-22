Endpoint for receiving all recent Events (maybe last 14days?) of a User

### Request:

```http
POST /get_events
Authorization: Bearer <token>

{"start_date": <datetime>, "end_date": <datetime>}  # Both fields are optional

HTTP 200 OK

Content-Type: application/json

{
  "user_name": "John Doe",
  "events": [
    {
      "type": "location",
      "time": "datetime",
      "details": {
        "location_id": "location-id-1",
        "name": "Rewe",
        "adress": "Berlin ...."
      }
    },
    {
      "type": "person",
      "time": "datetime",
      "details": {
        "companions_users": [
          {
            "user_id": "user-id",
            "user_name": "User regitered name"
          }
        ],
        "companions_non_users": [
          {
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
