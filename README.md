# Go REST API

# Stack

- Go 1.21.7
- Docker
- Docker compose
- Postgres 14.1-alpine
- Gorm 1.25.7
- Mux 1.8.0
- Sqlmock 1.5.2
- Testify 1.8.4

# Docker

## Build & Run

```shell
docker-compose -p rest-api up --remove-orphans --force-recreate -d --build
```

## Down

```shell
docker-compose -p rest-api down --remove-orphans --rmi local --volumes
```

# Swagger

- go get -u github.com/swaggo/swag/cmd/swag
- $GOPATH/bin/swag init
- go get -u github.com/swaggo/http-swagger

# Swagger UI

- [HTML](http://localhost:3000/swagger/index.html)