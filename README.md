# product-api-assesment

**Product-API-Assesment is an assesment project for Erajaya technical test**
#Golang

**Tech Stack**

- Docker
- Postgresql
- Redis
- SQLC

**Feature**
_Store cached data on Redis Client_, _Pagination_, _Sorting Result_

**List of Instant Commands**

- make postgres <br />
  to pull and run postgres service on docker <br />
- make createdb <br />
  to create database on postgres <br />
- make dropdb <br />
  to delete database on postgres <br />
- make redis <br />
  to pull and run redis service on docker <br />
- make migrate-up <br />
  to run SQL up migrations, create table "product" <br />
- make migrate-down <br />
  to run SQL down migrations, drop table "product" <br />
- make test <br />
  to run testing on application <br />
- make run <br />
  to run the application <br />

This app has feature to <br />

- Add Product
- Get list of Product (implement Pagination, Sort By Name, Price, Quantity, and CreatedAt)

**API**

**_Create New Product_**

POST http://localhost:8080/v1/products
Endpoint to create new Product and save it to database

Parameter :

- required: optional
- type: json
- field: { <br />
  &emsp; name string <br />
  &emsp; price int <br />
  &emsp; description string <br />
  &emsp; quantity int <br />
  }
- example: { <br />
  &emsp; "name": "Payung", <br />
  &emsp; "price": 12000, <br />
  &emsp; "description": "import dari Jepang", <br />
  &emsp; "quantity": 20 <br />
  }

Response :

**200 OK** <br />

- Body: { <br />
  &emsp; "name": "Payung", <br />
  &emsp; "price": 12000, <br />
  &emsp; "description": "import dari Jepang", <br />
  &emsp; "quantity": 20 <br />
  }

**500 Internal Server Error**

- Body: <br />
  &emsp; "Internal Server Error" <br />

**_Get List of Product_**

GET http://localhost:8080/v1/products
Endpoint to create new Product and save it to database

Parameter :

- required: yes
- type: query
- field: { <br />
  &emsp; limit int <DEFAULT: 5> <br />
  &emsp; offset int <DEFAULT: 0> <br />
  &emsp; name ASC/DESC <br />
  &emsp; price ASC/DESC <br />
  &emsp; quantity ASC/DESC <br />
  }
- example: <br />
  &emsp; "URL": "/v1/products?limit=2&offset=0&name=desc", <br />

Response :

**200 OK** <br />

- Body: [ <br />
  &emsp; { <br />
  &emsp; &emsp; "name": "Payung", <br />
  &emsp; &emsp; "price": 12000, <br />
  &emsp; &emsp; "description": "import dari Jepang", <br />
  &emsp; &emsp; "quantity": 20 <br />
  &emsp; }, <br />
  &emsp; { <br />
  &emsp; &emsp; "name": "Kursi", <br />
  &emsp; &emsp; "price": 12000, <br />
  &emsp; &emsp; "description": "terbuat dari kayu Jati asli", <br />
  &emsp; &emsp; "quantity": 10 <br />
  &emsp; } <br />
  ] <br />

**500 Internal Server Error**

- Body: <br />
  &emsp; "Internal Server Error" <br />
