# books-api

To run its

`go run main.go`

Persistence is backed by a filestore

## List Books

`GET /books`

```
curl http://localhost:8000/books?page=1
```

Takes an optional `page` query parameter. Limit per page is currently set at 10.

Example response


```
{
    "books": [
        {
            "id": "98e4e8be-a790-41cb-8555-899e58840d20",
            "title": "Harry Potter",
            "authors": [
                {
                    "name": "J.K Rowling"
                }
            ],
            "publisher": {
                "name": "Hodder"
            },
            "published_at": "1999-01-01",
            "pages": 1000
        },
        {
            "id": "07f67d9c-c55f-4459-bdd7-551e985ec3dd",
            "title": "Harry Potter 2",
            "authors": [
                {
                    "name": "J.K Rowling"
                }
            ],
            "publisher": {
                "name": "Hodder"
            },
            "published_at": "2000-01-01",
            "pages": 2000
        }
    ],
    "total_count": 2
}
```
```
status: 200
```

### Create Book

`POST /books`

```
curl -X POST http://localhost:8000/books \
   -H 'Content-Type: application/json' \
   -d '{"title": "Harry Potter", "authors": [{"name": "J.K Rowling"}],"publisher": {"name": "Hodder"},"published_at": "1999-01-01","pages": 2000}'
```

Example response

```
{
    "id": "ba14a180-6ab9-4244-98e8-1c387711c78b",
    "title": "Harry Potter",
    "authors": [
        {
            "name": "J.K Rowling"
        }
    ],
    "publisher": {
        "name": "Hodder"
    },
    "published_at": "1999-01-01",
    "pages": 2000
}
```

```
status: 200
```

### Update Book

`PATCH /books/:id`

```
curl -X PATCH http://localhost:8000/books/some-id \
   -H 'Content-Type: application/json' \
   -d '{"title": "Goblet of Fire"}'
```

Example response

```
{
    "id": "98e4e8be-a790-41cb-8555-899e58840d20",
    "title": "Goblet of Fire",
    "authors": [
        {
            "name": "J.K Rowling"
        }
    ],
    "publisher": {
        "name": "Hodder"
    },
    "published_at": "1900-01-01",
    "pages": 100
}
```

```
status: 200
```

### Delete Book

`DELETE /books/:id`

```
curl -X DELETE http://localhost:8000/books/some-id
```

Example response

```
status: 204
```

