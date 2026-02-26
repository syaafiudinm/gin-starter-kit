# Go Starter Kit

A clean and well-structured Go REST API starter kit built with **Gin** framework, **GORM** ORM, and **MySQL** database.

## Project Structure

```
go-starter-kit/
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   ├── config.go               # Configuration loader (env vars)
│   └── database.go             # MySQL database connection & migration
├── internal/
│   ├── dto/
│   │   └── user_dto.go         # Data Transfer Objects (request/response)
│   ├── handler/
│   │   └── user_handler.go     # HTTP handlers (controllers)
│   ├── middleware/
│   │   ├── cors.go             # CORS middleware
│   │   └── recovery.go         # Panic recovery middleware
│   ├── model/
│   │   └── user.go             # Database models (GORM)
│   ├── repository/
│   │   └── user_repository.go  # Data access layer
│   └── service/
│       └── user_service.go     # Business logic layer
├── pkg/
│   └── response/
│       └── response.go         # Standardized JSON response helpers
├── routes/
│   └── routes.go               # Route definitions & dependency wiring
├── .env                        # Environment variables (do not commit)
├── .env.example                # Environment variables template
├── .gitignore                  # Git ignore rules
├── go.mod                      # Go module definition
├── go.sum                      # Go module checksums
└── README.md                   # This file
```

## Architecture

This project follows a **layered architecture** pattern:

```
Request → Handler → Service → Repository → Database
                       ↓
                     Model / DTO
```

| Layer          | Responsibility                                      |
| -------------- | --------------------------------------------------- |
| **Handler**    | Parse HTTP requests, validate input, return response |
| **Service**    | Business logic, data transformation                 |
| **Repository** | Database operations (CRUD via GORM)                 |
| **Model**      | Database table definitions (GORM models)            |
| **DTO**        | Request/response data structures                    |

## Prerequisites

- **Go** 1.21 or higher
- **MySQL** 5.7 or higher
- **Git**

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/syaafiudinm/go-starter-kit.git
cd go-starter-kit
```

### 2. Set up environment variables

```bash
cp .env.example .env
```

Edit `.env` and configure your database credentials:

```
APP_NAME=go-starter-kit
APP_PORT=8080
APP_ENV=development

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=go_starter_kit
```

### 3. Create the MySQL database

```sql
CREATE DATABASE go_starter_kit CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 4. Install dependencies

```bash
go mod tidy
```

### 5. Run the application

```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080` (or whichever port you configured).

## API Endpoints

### Health Check

| Method | Endpoint           | Description          |
| ------ | ------------------ | -------------------- |
| GET    | `/api/v1/health`   | Service health check |

### Users

| Method | Endpoint             | Description        |
| ------ | -------------------- | ------------------ |
| POST   | `/api/v1/users`      | Create a new user  |
| GET    | `/api/v1/users`      | Get all users      |
| GET    | `/api/v1/users/:id`  | Get user by ID     |
| PUT    | `/api/v1/users/:id`  | Update user by ID  |
| DELETE | `/api/v1/users/:id`  | Delete user by ID  |

### Request & Response Examples

#### Create User

**Request:**

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "08123456789"
  }'
```

**Response (201 Created):**

```json
{
  "code": 201,
  "status": "success",
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "08123456789",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Get All Users (Paginated)

**Request:**

```bash
curl http://localhost:8080/api/v1/users?page=1&limit=10
```

**Response (200 OK):**

```json
{
  "code": 200,
  "status": "success",
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "08123456789",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "current_page": 1,
    "per_page": 10,
    "total_items": 1,
    "total_pages": 1
  }
}
```

#### Update User

**Request:**

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com"
  }'
```

#### Delete User

**Request:**

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

### Error Response Format

```json
{
  "code": 400,
  "status": "error",
  "message": "Invalid request body",
  "errors": "Key: 'CreateUserRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"
}
```

## Adding a New Resource

To add a new resource (e.g., `Product`), follow these steps:

1. **Model** — Create `internal/model/product.go` with GORM struct tags
2. **DTO** — Create `internal/dto/product_dto.go` with request/response structs
3. **Repository** — Create `internal/repository/product_repository.go` with interface and implementation
4. **Service** — Create `internal/service/product_service.go` with business logic
5. **Handler** — Create `internal/handler/product_handler.go` with Gin handlers
6. **Routes** — Register routes in `routes/routes.go`
7. **Migration** — Add the model to `config.AutoMigrate()` in `cmd/main.go`

## Build for Production

```bash
# Build the binary
go build -o bin/server cmd/main.go

# Run in production
APP_ENV=production ./bin/server
```

## License

This project is open-sourced under the [MIT License](LICENSE).