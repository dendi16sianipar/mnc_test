# MNC Test

## Run application

```
go run .
```

## API

### Register

#### POST /register

Request

```
{
	"name" : "John",
	"email" : "john@gmail.com",
	"password" : "john123"
}
```

Response

```
{
    "data": {
        "id": 3,
        "name": "john",
        "email": "john@gmail.com",
        "password": "john123",
        "balance": 0
    }
}
```

### Login

#### POST /login

Request

```
{
	"email" : "john@gmail.com",
	"password" : "john123"
}
```

Response

```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Payment

#### POST /payment
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

Request

```
{
	"customer_id" : 1,
	"bill" : 10000,
	"description" : "pay lunch"
}
```

Response

```
{
    "data": {
        "id": 1,
        "customer_id": 1,
        "bill": 1,
        "description": "pay lunch",
        "date": "2023-04-19T04:01:12.309451Z"
    }
}
```
