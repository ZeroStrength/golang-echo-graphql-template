# golang-echo-graphql-template

## Install

```
go get ./...
```

## Run

```
go run cmd/backend/main.go
```

## GraphQL test

POST http://localhost:3000/graphql

```graphql
query {
    Users {
        id
        name
        age
    }
}
```
-> Result
```json
{
  "data": {
    "Users": [
      {
        "age": "20",
        "id": "1",
        "name": "name1"
      },
      {
        "age": "30",
        "id": "2",
        "name": "name2"
      },
      {
        "age": "40",
        "id": "3",
        "name": "name3"
      }
    ]
  }
}
```

#### with id filter
```graphql
query {
    Users(id: 3) {
        id
        name
        age
    }
}
```
-> result
```json
{
  "data": {
    "Users": [
      {
        "age": "40",
        "id": "3",
        "name": "name3"
      }
    ]
  }
}
```


