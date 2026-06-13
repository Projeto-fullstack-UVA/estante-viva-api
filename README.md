# Estante Viva API

REST API for **Estante Viva**, a community library system that manages users, a book
catalog, and book loans. Built in Go with [Gin](https://github.com/gin-gonic/gin),
PostgreSQL (via [pgx](https://github.com/jackc/pgx)), and JWT-based authentication.

## Tech stack

- **Go** 1.26
- **Gin** — HTTP router / middleware
- **PostgreSQL** — accessed through `pgx/v5` connection pool
- **golang-jwt/v5** — HS256 JWT auth
- **golang.org/x/crypto** — password hashing
- **godotenv** — `.env` loading

## Project structure

```
cmd/app/main.go            Entry point: env loading, CORS, route registration
internals/
  auth/                    JWT generation & verification
  controllers/            HTTP handlers (users, books, loans)
  dtos/                    Request/response shapes per domain
  entities/                Domain models (User, Book, Loan)
  middleware/              Authentication middleware
  repositories/            Database access (pgx pool + queries)
  services/                Business logic + domain errors
  utils/                   Password helpers
migrations/                SQL schema (users, books, loans)
```

## Getting started

### Prerequisites

- Go 1.26+
- A PostgreSQL database

If you use Nix, a dev shell with Go tooling is provided:

```bash
nix-shell
```

### Configuration

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://user:password@localhost:5432/estante_viva
JWT_SECRET_KEY=your-secret-key
```

| Variable         | Description                                  |
| ---------------- | -------------------------------------------- |
| `DATABASE_URL`   | PostgreSQL connection string (required)      |
| `JWT_SECRET_KEY` | Secret used to sign/verify JWTs (required)   |

### Database setup

Apply the migrations in `migrations/` to your database, in order:

```bash
psql "$DATABASE_URL" -f migrations/0001_create_users.sql
psql "$DATABASE_URL" -f migrations/0002_create_books.sql
psql "$DATABASE_URL" -f migrations/0003_create_loans.sql
```

### Run

```bash
go run ./cmd/app
```

The server starts on **http://localhost:8080**.

## Authentication

Protected routes require a JWT in the `Authorization` header. Obtain a token via
`POST /login` (the token is returned in the response body and the `Authorization`
response header). Tokens are signed with HS256 and expire after **24 hours**.

```
Authorization: Bearer <token>
```

The `Bearer ` prefix is optional — the middleware strips it if present.

### CORS

Cross-origin requests are allowed from `http://localhost:5173` and
`http://localhost:4173` (typical Vite dev/preview origins).

## API reference

Base URL: `http://localhost:8080`. 🔒 = requires authentication.

### Auth & users

| Method | Path          | Auth | Description                  |
| ------ | ------------- | ---- | ---------------------------- |
| POST   | `/login`      |      | Authenticate, returns a JWT  |
| POST   | `/users`      |      | Register a new user          |
| GET    | `/users`      | 🔒   | List all users               |
| GET    | `/users/:id`  | 🔒   | Get a user by ID             |

### Books

| Method | Path          | Auth | Description          |
| ------ | ------------- | ---- | -------------------- |
| GET    | `/books`      | 🔒   | List all books       |
| POST   | `/books`      | 🔒   | Create a book        |
| GET    | `/books/:id`  | 🔒   | Get a book by ID     |

### Loans

| Method | Path          | Auth | Description                        |
| ------ | ------------- | ---- | --------------------------------- |
| GET    | `/loans`      | 🔒   | List all loans                    |
| POST   | `/loans`      | 🔒   | Borrow a book (creates a loan)    |
| GET    | `/loans/:id`  | 🔒   | Get a loan by ID                  |
| PATCH  | `/loans/:id`  | 🔒   | Return a borrowed book            |

## Examples

### Register

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ada Lovelace",
    "email": "ada@example.com",
    "birthDate": "1815-12-10T00:00:00Z",
    "password": "s3cret",
    "address": "Rua das Flores, 100",
    "document": "12345678901",
    "cellphone": "11999998888",
    "role": "student",
    "campus": "Centro"
  }'
```

`role` must be one of `student`, `teacher`, `donator`, `admin`.

### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{ "email": "ada@example.com", "password": "s3cret" }'
```

Response:

```json
{ "id": 1, "token": "<jwt>" }
```

### Create a book

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <jwt>" \
  -d '{
    "title": "The Pragmatic Programmer",
    "author": "Hunt & Thomas",
    "release_date": "1999-10-30T00:00:00Z",
    "edition": "1st",
    "status": "available",
    "created_at": "2026-06-13T00:00:00Z"
  }'
```

`status` must be one of `available`, `lent`.

### Borrow a book

```bash
curl -X POST http://localhost:8080/loans \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <jwt>" \
  -d '{ "user_id": 1, "book_id": 1 }'
```

Borrowing sets the book's status to `lent` and assigns a **14-day** return date.
A book that is not `available` returns `409 Conflict`.

### Return a book

```bash
curl -X PATCH http://localhost:8080/loans/1 \
  -H "Authorization: Bearer <jwt>"
```

Returning marks the loan's `returned_at` and sets the book back to `available`.
A loan that was already returned responds with `409 Conflict`.

## Data model

- **users** — id, name, birth_date, email (unique), password, address, document
  (unique), cellphone (unique), role, campus, score, created_at.
- **books** — id, title, author, release_date, edition, status, created_at.
- **loans** — id, user_id → users, book_id → books, return_date, returned_at.
  Deleting a user or book cascades to their loans.

## Error responses

Errors are returned as plain text with an appropriate HTTP status. Common cases:

- `400 Bad Request` — invalid payload or malformed ID
- `401 Unauthorized` — missing/invalid token on a protected route
- `404 Not Found` — user, book, or loan does not exist
- `409 Conflict` — book unavailable, or loan already returned
- `500 Internal Server Error` — unexpected failure
