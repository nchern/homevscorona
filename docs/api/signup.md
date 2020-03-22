Sign up endpoint

### Request:

```http
POST /api/signup

Authorization: Bearer <token>
Content-Type: application/json

{ "provider": "google", "name": "John Doe" }
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
