# 🛒 Go E-Shop Service

A modular, clean-architecture-based backend for an e-commerce service written in **Go**.  
It follows **Clean Architecture principles** for better maintainability, testability, and separation of concerns.

---

## 🚀 Tech Stack

- **Language:** Go (Golang)
- **Architecture:** Clean Architecture
- **Database:** PostgreSQL
- **HTTP Router:** net/http (can be replaced with Chi or Gin)
- **Environment:** `.env` using `godotenv`
- **ORM/DB Layer:** Custom repository pattern
- **Dependency Management:** Go Modules

---

## 🏗️ Project Architecture Overview

Clean Architecture layers used:

```
+-------------------------------+
| Frameworks & Drivers | -> HTTP, Database, etc.
+-------------------------------+
| Interface Adapters | -> Handlers, Repository Implementations
+-------------------------------+
| Application / Use Cases | -> Business Logic
+-------------------------------+
| Entities / Domain | -> Core Models
+-------------------------------+
```

Dependencies always point **inward** →
`Infrastructure → Repository → UseCase → Domain`

---

## ▶️ Project run command

First build-

```bash
make build
```

then (This will run build project)-

```bash
make run
```

if need to run migrations
for up-

```bash
make migrate-up
```

for down-

```bash
make migrate-down
```

---

## 📂 Folder Structure

```
go-e-shop-service/
│
├── cmd/
│ |── server/
│ | └── main.go # Application entry point
| |── migration/
│ | └── main.go # Database migration logic
├── internal/
│ ├── domain/ # Core business models
│ │ └── user.go # Defines User entity
│ │
│ ├── usecase/ # Application business rules
│ │ └── auth_usecase.go # Login/Register business logic
│ │
│ ├── repository/ # Repository interfaces & implementations
│ │ ├── user_repo.go # Interface (abstract)
│ │ └── user_repo_postgres.go # Postgres implementation
│ │
│ ├── infrastructure/ # Frameworks, DB, HTTP
│ │ ├── db/
│ │ │ └── postgres.go # Postgres connection setup
│ │ │
│ │ ├── http/
│ │ │ ├── router.go # Sets up HTTP routes
│ │ │ └── handler/
│ │ │ ├── auth_handler.go # HTTP layer for auth endpoints
│ │ │ ├── user_handler.go # HTTP layer for user endpoints
│ │ │ └── ... # Other feature handlers
│ │
│ └── utils/ # Optional helpers/utilities
│ └── response.go # Standardized JSON responses
│
├── migrations/ # SQL migration files
│ ├── 001_create_users_table.up.sql
│ └── 001_create_users_table.down.sql
│
├── .env # Environment variables
├── go.mod
├── go.sum
└── README.md
```

---

## 🧩 Clean Architecture Layers

| Layer              | Purpose                                               | Example                         |
| ------------------ | ----------------------------------------------------- | ------------------------------- |
| **Domain**         | Defines core business models/entities.                | `domain/user.go`                |
| **UseCase**        | Contains application logic independent of frameworks. | `usecase/auth_usecase.go`       |
| **Repository**     | Handles data persistence (DB, cache, etc).            | `repository/user_repository.go` |
| **Infrastructure** | Handles technical implementations (HTTP, DB, etc).    | `infrastructure/db/postgres.go` |
| **Handler**        | Manages HTTP requests and responses.                  | `handler/auth_handler.go`       |

---

## ⚙️ Environment Variables (`.env`)

Create a `.env` file in the project root:

```env
# PostgreSQL connection
POSTGRES_DSN=postgres://{user_name}:{password}@localhost:{port}/{db}?sslmode=disable

# App configuration
APP_PORT=8080
APP_ENV=development
```
