# Golang And MongoDB REST API

A small Book Management REST API built with Go, MongoDB, and Docker.

## Run with Docker

```bash
docker compose up --build
```

The API runs at `http://localhost:8001` and MongoDB is exposed on `localhost:27017`.

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
curl http://localhost:8001/health
```

```bash
curl http://localhost:8001/api/books
```

```bash
curl -X POST http://localhost:8001/api/books \
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
curl -X PUT http://localhost:8001/api/books/PASTE_BOOK_ID \
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
curl -X DELETE http://localhost:8001/api/books/PASTE_BOOK_ID
```

## Run locally without Docker for Go

Start only MongoDB:

```bash
docker compose up mongo
```

Run the API from this directory:

```bash
go run .
```

Local defaults:

| Variable | Default |
| --- | --- |
| `PORT` | `8000` |
| `MONGO_URI` | `mongodb://localhost:27017` |
| `MONGO_DATABASE` | `bookstore` |
