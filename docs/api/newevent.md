Checking in a new event

### Request:

```http
POST /api/new_event

Authorization: Bearer <token>
Content-Type: application/json

{
  "type": "location/person",
  "date": 12345678,
  "details": { /*  The content of "details" depends on the type of event */ }
}
```

#### Details for users
```
{
    "users": [
      {
        "user_id": "user-id",
        "name": "John Doe"
      },
      {
            "name": "John Doe"
      }
    ]
}
```

#### Details for location
```
{
     "location_id": "location-id-1",
      "name": "Rewe",
      "adress": "Berlin ...."
}
```


### Responses

#### Sucess

```http
HTTP/1.1 200 OK

{ "status": "200" }
```



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
