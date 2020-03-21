Checking in a new event

### Request:

```http
POST /event

## Required:
UserId=[integer]
AuthToken? (to be discussed)

Content-Type: application/json
The content of "details" depends on the type of event

{ "type": "location/person" , "date": "datetime",
"details": {"companions_users":[{"userID":"..."],
            "companions_non_users":[{"name":"..."}]}}
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
