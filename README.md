# Tasks API in Go

This is a simple CRUD api for a tasks app.

## Endpoints

- `GET /` - Get all tasks
- `POST /` - Create new task
  ```json
  {
    "title": "First task",
    "description": "This is my first task",
    "completed": true,
    "user_id": "janedoe123"
  }
  ```
- `GET /:id` - Get one task by id
- `PUT /:id` - Update one task by id
- `DELETE /:id` - Delete one task by id
