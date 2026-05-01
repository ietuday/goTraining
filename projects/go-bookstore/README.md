# Go Bookstore API

A small Book Management System built with Go, MySQL, and Docker.

## Run with Docker

```bash
docker compose up --build
```

The API runs at `http://localhost:8000` and MySQL is exposed on `localhost:3307`.

## Endpoints

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/health` | Check API and database health |
| `GET` | `/api/books` | List all books |
| `GET` | `/api/books/{id}` | Get one book |
| `POST` | `/api/books` | Create a book |
| `PUT` | `/api/books/{id}` | Update a book |
| `DELETE` | `/api/books/{id}` | Delete a book |

## Example requests

```bash
curl http://localhost:8000/api/books
```

```bash
curl -X POST http://localhost:8000/api/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Go in Action",
    "author": "William Kennedy, Brian Ketelsen, and Erik St. Martin",
    "isbn": "9781617291784",
    "price": 34.99,
    "description": "A hands-on book for Go applications."
  }'
```

```bash
curl -X PUT http://localhost:8000/api/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan A. A. Donovan and Brian W. Kernighan",
    "isbn": "9780134190440",
    "price": 42.99,
    "description": "Updated price."
  }'
```

```bash
curl -X DELETE http://localhost:8000/api/books/1
```

## Run locally without Docker for Go

Start only MySQL:

```bash
docker compose up mysql
```

Run the API from this directory:

```bash
go run .
```

The local defaults match the Compose database credentials.
