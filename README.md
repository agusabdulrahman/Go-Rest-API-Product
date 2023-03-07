# Go-Rest-API-Product Using Gin

## REST API for Product entity with feature:

- Create, Read, Update, Delete
- Filtering
- Pagination

### Installation
#### Clone this repository:

```bash
https://github.com/agusabdulrahman/Go-Rest-API-Product
```

Change into the directory:
```bash
cd Go-Rest-API-Product
```

Start the server:
```bash
go run main.go
```
By default, the server listens on port 8080.

Endpoints

`GET /products`
Retrieves a list of products.

Response:
- `page` - the page number of the results to return (default: 1)

Start the server:
```bash
{
    "page": 1,
    "pageSize": 10,
    "total": 25,
    "products": [
        {
            "id": 1,
            "name": "Dell XPS 13",
            "description": "13-inch laptop",
            "price": 1500,
            "completed": false
        },
        ...
    ]
}

```

`POST /products`

Adds a new product.

Request Body:

``` 
{
    "name": "New Product",
    "description": "A new product",
    "price": 1000,
    "completed": false
}

```
Response:
```
{
    "id": 26,
    "name": "New Product",
    "description": "A new product",
    "price": 1000,
    "completed": false
}
```

`GET /products/:id`
Retrieves a single product by ID.

Response:

```
{
    "id": 1,
    "name": "Dell XPS 13",
    "description": "13-inch laptop",
    "price": 1500,
    "completed": false
}
```

`PATCH /products/:id`
Toggles the completion status of a product by ID.

Response:
```
{
    "id": 1,
    "name": "Dell XPS 13",
    "description": "13-inch laptop",
    "price": 1500,
    "completed": true
}
```

`DELETE /products/20`
Delete.

Response:
```
{
    "message": "product deleted"
}
```




