# You - commerce Api

This project is the starting point for an e-commerce SaaS i want to build.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

- golang
- Postgres

### Installing

```golang
go run main.go
```

## Available Endpoints

- Add product */v1/addproduct*
- Get all products */v1/getallproducts*
- Get products by category */v1/getproductsbycategory*
- Update product */v1/updateproduct*
- Delete Product */v1/deleteproduct*
- Add category */v1/addcategory*
- Get categories */v1/getcategories*

## Response types
The api request and response formats is strictly json

response format 
```json
{
    "status":"success",
    "payload":[]
}
```

where payload differs


### Add Product

request body 

```json
{
    "name":"shirt",
    "category":"2",
    "quantity":"4",
    "price":"200",
    "image":"kksus//skhs"
}
```
***The image field should contain a base64 encoded image***

typical response 

```json
{
    "status":"success"
}
{
    "status":"error",
    "msg":"cause of error"
}
```
### Delete Product

request body 

```json
{
    "id": 1
}
```

typical response 

```json
{
    "status":"success"
}
{
    "status":"error",
    "msg":"cause of error"
}
```
### Get all products
make a get request to this endpoint ***/v1/getallproducts***

typical response 

```json
{
    "status":"success",
    "products":[]
}
{
    "status":"error",
    "msg":"cause of error"
}
```

### Get products by category
 - Post request
request body 

```json
{
    "category":"2",
}
```

typical response 

```json
{
    "status":"success",
    "products":[]
}
{
    "status":"error",
    "msg":"cause of error"
}
```

### Add Category

- Post request

request body 

```json
{
    "name":"shoes",
}
```

typical response 

```json
{
    "status":"success"
}
{
    "status":"error",
    "msg":"cause of error"
}
```

### Add Product

request body 

```json
{
    "id":"1",
    "name":"shirt",
    "category":"2",
    "quantity":"4",
    "price":"200",
}
```

typical response 

```json
{
    "status":"success"
}
{
    "status":"error",
    "msg":"cause of error"
}
```

### Get all categories
make a get request to this endpoint ***/v1/getcategories***

typical response 

```json
{
    "status":"success",
    "categories":[]
}
{
    "status":"error",
    "msg":"cause of error"
}
```