Endpoint to query a users recent contacts
optional endpoint

### Request:

```http
GET /contacts

## Required:
UserId=[integer]
AuthToken? (to be discussed)

Content-Type: application/json

{"companions_users":[{"userID":"...", "userName":"...", "no_visits":"..."}],
 "companions_non_users":[{"name":"...", "no_visits":"..."}]
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
