# Blog Service

link development : https://blog-service-3h0p.onrender.com/

## Authentication

### Register

POST `/auth/register`

request field
```json
{
  "email": "example@mail.com",
  "password": "testPassword",
  "password2": "testPassword",
  "name": "John Doe"
}
```

Success response 200
```json
{
    "Message": "User registered successfully",
    "Data": {
        "id": 27,
        "email": "test17@mail.com",
        "password": "$2a$14$RqpcaBY4a18J0C8t/mqLCOzlHfTHh3GRQoGUPp1Rl8kL3ttb25iKq",
        "name": "John Doe",
        "created_at": "2023-03-01T03:39:54.548737311Z",
        "updated_at": "2023-03-01T03:39:54.548737441Z"
    },
    "Errors": null
}
```

Email already used 200
```json
{
    "Message": "email is already used",
    "Data": null,
    "Errors": null
}
```

Field error 422
```json
{
    "Message": "Invalid Data",
    "Data": null,
    "Errors": [
        {
            "Field": "Email",
            "Message": "required"
        },
        {
            "Field": "Password",
            "Message": "required"
        },
        {
            "Field": "Name",
            "Message": "required"
        },
        {
            "Field": "Password2",
            "Message": "required"
        }
    ]
}
```

Password confirmation error 400
```json
{
    "Message": "password is not the same as a confirmation password",
    "Data": null,
    "Errors": null
}
```

### Login
POST `/auth/login`

request field

```json
{
  "email": "example@mail.com",
  "password": "testPassword"
}
```

Success response 200
```json
{
  "Message": "User logged in successfully",
  "Data": {
    "token": "some_token"
  },
  "Errors": null
}
```

Unauthorized error 401
```json
{
    "Message": "Unauthorized",
    "Data": null,
    "Errors": null
}
```

Field error 422
```json
{
    "Message": "Invalid Data",
    "Data": null,
    "Errors": [
        {
            "Field": "Email",
            "Message": "required"
        },
        {
            "Field": "Password",
            "Message": "required"
        }
    ]
}
```


## Article

### Get all
GET `/article`

query params
- limit:int = limit the record want to be queried
- where:string = condition
- order_by:string = column name
- order:[asc, desc] = sort type

### Get by id
GET `/article/:id`

params :
- id:int = id of the article

### Add article
POST `/article`

request field
```json
{
  "title": "Title",
  "content": "Lorem ipsum dolor sit amet"
}
```

### Update article
PUT `/article/:id`

params :
- id:int = id of the article

request field
```json
{
  "title": "Title",
  "content": "Lorem ipsum dolor sit amet"
}
```

### Delete article
DELETE `/article/:id`

params :
- id:int = id of the article