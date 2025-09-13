# Checkpoint Backend

Checkpoint Backend is a Go-based project that provides APIs and services for managing users, games, reviews.
It follows a layered architecture with clear separation between controllers, handlers, repositories, and services.

Games database populated using IGDB API.

https://checkpoint-frontend-six.vercel.app/

## Features
- User authentication with JWT
- Game and review management
- Swagger documentation
- Docker support for database

## Setup

Clone the repository and install dependencies:

```bash
go mod tidy
```

Generate Swagger documentation and run the API:

```bash
make run-api
```

## Documentation
Swagger documentation is available in the `docs/` folder and can be generated with:

```bash
swag init -g cmd/api/main.go
```

## License
This project is licensed under the MIT License.
