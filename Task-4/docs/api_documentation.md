# Task Management API Documentation

This API allows you to manage tasks with basic CRUD operations.

## Base URL
`http://localhost:8080/api`

## Endpoints

### 1. Get All Tasks
**URL**: `/tasks`  
**Method**: `GET`  
**Response**:  
```json
[
    {
        "id": "string",
        "title": "string",
        "description": "string",
        "due_date": "string (ISO 8601 format)",
        "status": "string (pending|in_progress|completed)"
    }
]
```

### 2. Create a Task
**URL**: `/tasks`  
**Method**: `POST`  
**Request Body**:  
```json
{
    "title": "string (required)",
    "description": "string",
    "due_date": "string (ISO 8601 format, required)",
    "status": "string (pending|in_progress|completed, required)"
}
```
**Response**:  
Status: `201 Created`  
Body: Created task object

### 3. Get a Specific Task
**URL**: `/tasks/:id`  
**Method**: `GET`  
**Response**:  
```json
{
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "string (ISO 8601 format)",
    "status": "string (pending|in_progress|completed)"
}
```

### 4. Update a Task
**URL**: `/tasks/:id`  
**Method**: `PUT`  
**Request Body**:  
```json
{
    "title": "string",
    "description": "string",
    "due_date": "string (ISO 8601 format)",
    "status": "string (pending|in_progress|completed)"
}
```
**Response**:  
Updated task object

### 5. Delete a Task
**URL**: `/tasks/:id`  
**Method**: `DELETE`  
**Response**:  
Status: `204 No Content`