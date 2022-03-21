# books-api

To run its

`go run main.go`

Persistence is backed by a filestore

### List Books

`GET /books`

```
curl http://localhost:8000/books?page=1
```

Takes an optional `page` query parameter. Limit per page is currently set at 10


`POST /books`

```
curl -X POST http://localhost:8000/books \
   -H 'Content-Type: application/json' \
   -d '{"title": "Harry Potter", "authors": [{"name": "J.K Rowling"}],"publisher": {"name": "Hodder"},"published_at": "1st January 1999","pages": 2000}'
```