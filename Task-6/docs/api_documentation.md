# Task Management API Documentation

## Base URL
`http://localhost:8080/api`

## Authentication

### Register a new user
**URL**: `/register`  
**Method**: `POST`  
**Request Body**:  
```json
{
    "username": "string (required, unique)",
    "password": "string (required)",
    "role": "string (optional, defaults to 'user')"
}
```
**Response**:  
```json
{
    "id": "string",
    "username": "string",
    "role": "string"
}
```

### Login
**URL**: `/login`  
**Method**: `POST`  
**Request Body**:  
```json
{
    "username": "string (required)",
    "password": "string (required)"
}
```
**Response**:  
```json
{
    "token": "jwt-token-string"
}
```

## Tasks (Requires Authentication)

### Get all tasks
**URL**: `/tasks`  
**Method**: `GET`  
**Headers**: `Authorization: Bearer <token>`  
**Response**:  
```json
[
    {
        "id": "string",
        "title": "string",
        "description": "string",
        "due_date": "ISO8601",
        "status": "pending|in_progress|completed",
        "created_by": "string",
        "created_at": "ISO8601",
        "updated_at": "ISO8601"
    }
]
```

### Create a task
**URL**: `/tasks`  
**Method**: `POST`  
**Headers**: `Authorization: Bearer <token>`  
**Request Body**:  
```json
{
    "title": "string (required)",
    "description": "string",
    "due_date": "ISO8601 (required)",
    "status": "pending|in_progress|completed (required)"
}
```

### Update a task
**URL**: `/tasks/:id`  
**Method**: `PUT`  
**Headers**: `Authorization: Bearer <token>`  
**Request Body**: Same as create, but all fields optional

### Delete a task
**URL**: `/tasks/:id`  
**Method**: `DELETE`  
**Headers**: `Authorization: Bearer <token>`