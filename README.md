# API Gateway

A simple API Gateway written in Go that provides:
- User authentication with JWT tokens
- PostgreSQL database integration
- Protected routes with middleware

## Features

- User registration and login
- JWT token generation and validation
- Protected API endpoints
- PostgreSQL integration

## Prerequisites

- Go 1.21 or later
- PostgreSQL database

## Environment Variables

Copy the `.env.example` file to `.env` and adjust the values as needed:

```
# API Gateway Configuration
PORT=8080

# JWT Configuration
JWT_SECRET_KEY=your_jwt_secret_key_change_this_in_production

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=apigateway
```

## Getting Started

1. Clone the repository
2. Install dependencies: `go mod download`
3. Configure your environment variables in `.env`
4. Create the PostgreSQL database: `createdb apigateway`
5. Run the application: `go run main.go`

## API Endpoints

### Public endpoints

- `POST /register`: Register a new user
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword",
    "first_name": "John",
    "last_name": "Doe"
  }
  ```

- `POST /login`: Login and get JWT token
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```

### Protected endpoints (require Authorization header with Bearer token)

- `GET /api/profile`: Get user profile information 