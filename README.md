# Documentation

Tech stack: Golang (Gin, Gorm), PostgresSQL, RabbitMQ

## Architecture

![Microservices Architecture](/microservices-architecture.png)

## API Specification

**1. USER SERVICE**

1.1 Register User

Type: POST

URL: http://localhost:8000/user/register

Payload:

```
{
    "email": "youremail@gmail.com",
    "whatsapp": "085123456789",
    "password": "yourpassword"
}
```

Return:

```
{
    "data": {
        "id": "user-xxxxxxxxxxxxxxxx"
    },
    "message": "success"
}
```

2.2 Login User

Type: POST

URL: http://localhost:8000/user/login

Payload:

```
{
    "email": "youremail@gmail.com",
    "password": "yourpassword"
}
```

Return:

```
{
    "data": {
        "token": "xxxxxxxxxxxxxxxx"
    },
    "message": "success"
}
```

**2. TICKET SERVICE**

2.1 Insert Ticket

Type: POST

Url: http://localhost:8000/ticket

Authorization: "Bearer "+token

Payload:

```
{
    "event": "event name",
    "qty": 1
}
```

Return:

```
{
    "data": {
        "id": "xxxxxxxxxxxxxxxx"
    },
    "message": "success"
}
```

2.2 Get Ticket Data

Type: GET
Url: http://localhost:8000/ticket/{id}

Return:

```
{
    "data": {
        "id": "ticket-xxxxxxxxxxxxxxxx",
        "owner": {
            "id": "user-xxxxxxxxxxxxxxxx",
            "email": "youremail@gmail.com",
            "whatsapp": "085123456789",
            "registered_on": "2023-01-01T00:00:00.000000+07:00",
            "updated_on": "2023-01-01T00:00:00.000000+07:00"
        },
        "event": "event name",
        "qty": 1,
        "status": "sold",
        "registered_on": "2023-01-01T00:00:00.000000Z",
        "updated_on": "2023-01-01T00:00:00.000000Z"
    },
    "message": "success"
}
```

2.3 Update Ticket Data

Type: PUT

Url: http://localhost:8000/ticket/{id}

Authorization: "Bearer "+token

Payload:

```
{
    "event": "event name",
    "qty": 1
}
```

Return:

```
{
    "message": "success"
}
```

2.4 Update Ticket Status

Type: PUT

Url: http://localhost:8000/ticket/{id}/{status}

Authorization: "Bearer "+token

Return:

```
{
    "message": "success"
}
```

**3. ORDER SERVICE**

3.1 Insert Order

Type: POST

Url: http://localhost:8000/order

Authorization: "Bearer "+token

Payload:

```
{
    "ticket_id": "ticket-xxxxxxxxxxxxxxxx",
    "qty": 1
}
```

Return:

```
{
    "data": {
        "id": "xxxxxxxxxxxxxxxx"
    },
    "message": "success"
}
```

2.2 Get Ticket Data

Type: GET

Url: http://localhost:8000/order/{id}

Return:

```
{
    "data": {
        "id": "order-xxxxxxxxxxxxxxxx",
        "buyer": {
            "id": "user-xxxxxxxxxxxxxxxx",
            "email": "youremail@gmail.com",
            "whatsapp": "085123456789",
            "registered_on": "2023-01-01T00:00:00.000000+07:00",
            "updated_on": "2023-01-01T00:00:00.000000+07:00"
        },
        "ticket": {
            "id": "ticket-xxxxxxxxxxxxxxxx",
            "owner": {
                "id": "user-xxxxxxxxxxxxxxxx",
                "email": "youremail@gmail.com",
                "whatsapp": "085123456789",
                "registered_on": "2023-01-01T00:00:00.000000+07:00",
                "updated_on": "2023-01-01T00:00:00.000000+07:00"
            },
            "event": "event name",
            "qty": 1,
            "status": "sold",
            "registered_on": "2023-01-01T00:00:00.000000Z",
            "updated_on": "2023-01-01T00:00:00.000000Z"
        },
        "qty": 1,
        "status": "canceled",
        "registered_on": "2023-01-01T00:00:00.000000Z",
        "updated_on": "2023-01-01T00:00:00.000000Z"
    },
    "message": "success"
}
```

2.3 Update Order Status

Type: PUT

Url: http://localhost:8000/ticket/{id}/{status}

Authorization: "Bearer "+token

Return:

```
{
    "message": "success"
}
```
