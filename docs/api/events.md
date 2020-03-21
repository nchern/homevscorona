Endpoint for receiving all recent Events (maybe last 14days?) of a User

### Request:

```http
GET /events

## Required:
UserId=[integer]
AuthToken? (to be discussed)

##Optional:
(to be discussed)
startdate=[datetime]
enddate=[datetime]

Content-Type: application/problem+json

{"username":"...", "events": [
  {"type": "location", "time":"datetime",
  "details": ["locationID":"", "name":"...", "adress": "..."]},
  {"type": "person", "time":"datetime",
  "details": ["companions_users":["userID":"", "userName"],
              "companions_non_users":["name"]]},
  ]
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
