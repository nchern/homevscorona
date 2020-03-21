Endpoint to query a users recent locations
optional endpoint

### Request:

```http
GET /locations

## Required:
UserId=[integer]
AuthToken? (to be discussed)

Content-Type: application/json

{ "locations": [{"locationID": "...", "name": "...", 
"lastVisit": "..."}]}
```

### Responses

#### Success

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
