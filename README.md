# Estante Viva API

REST API for **Estante Viva**, a community library system that manages users, a book
catalog, book loans, partner institutions, and events. Built in Go with
[Gin](https://github.com/gin-gonic/gin), PostgreSQL (via
[pgx](https://github.com/jackc/pgx)), and JWT-based authentication.

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
  controllers/             HTTP handlers (users, books, loans, institutions, events)
  dtos/                    Request/response shapes per domain
  entities/                Domain models (User, Book, Loan, Institution, Event)
  middleware/              Authentication & role-based authorization
  repositories/            Database access (pgx pool + queries)
  services/                Business logic + domain errors
  utils/                   Password helpers
migrations/                SQL schema (users, books, loans, institutions, events)
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
PORT=:8080
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:4173
```

| Variable          | Description                                                        |
| ----------------- | ----------------------------------------------------------------- |
| `DATABASE_URL`    | PostgreSQL connection string (required)                           |
| `JWT_SECRET_KEY`  | Secret used to sign/verify JWTs (required)                        |
| `PORT`            | Address the server listens on, e.g. `:8080` (required)           |
| `ALLOWED_ORIGINS` | Comma-separated list of allowed CORS origins (required)          |

All four variables are required — the app exits on startup if any is missing.

### Database setup

Apply the migrations in `migrations/` to your database, in order:

```bash
psql "$DATABASE_URL" -f migrations/0001_create_users.sql
psql "$DATABASE_URL" -f migrations/0002_create_books.sql
psql "$DATABASE_URL" -f migrations/0003_create_loans.sql
psql "$DATABASE_URL" -f migrations/0004_create_institutions.sql
psql "$DATABASE_URL" -f migrations/0005_create_events.sql
psql "$DATABASE_URL" -f migrations/0006_alter_users.sql
```

### Run

```bash
go run ./cmd/app
```

The server starts on the address given by `PORT` (e.g. **http://localhost:8080**).

## Authentication & authorization

Protected routes require a JWT in the `Authorization` header. Obtain a token via
`POST /login` (the token is returned in the response body and the `Authorization`
response header). Tokens are signed with HS256 and expire after **24 hours**.

```
Authorization: Bearer <token>
```

The `Bearer ` prefix is optional — the middleware strips it if present.

Some routes additionally require a specific **role**. The token carries the user's
role and the authorization middleware enforces it:

- `admin` — full access (user management, institutions, deletes).
- `teacher` — may create/manage books and events alongside `admin`.
- `student` — standard authenticated access (browse catalog, borrow books).

### CORS

Cross-origin requests are allowed from the origins listed in `ALLOWED_ORIGINS`
(typically the Vite dev/preview origins `http://localhost:5173` and
`http://localhost:4173`).

## API reference

Base URL: `http://localhost:8080`. 🔒 = requires authentication; the **Roles**
column lists the roles allowed when access is role-restricted.

### Auth & users

| Method | Path         | Auth | Roles   | Description                 |
| ------ | ------------ | ---- | ------- | --------------------------- |
| POST   | `/login`     |      |         | Authenticate, returns a JWT |
| POST   | `/users`     |      |         | Register a new user         |
| GET    | `/me`        | 🔒   |         | Get the current user        |
| GET    | `/users`     | 🔒   | `admin` | List all users              |
| GET    | `/users/:id` | 🔒   | `admin` | Get a user by ID            |
| PATCH  | `/users/:id` | 🔒   | `admin` | Update a user               |
| DELETE | `/users/:id` | 🔒   | `admin` | Delete a user               |

### Books

| Method | Path         | Auth | Roles              | Description      |
| ------ | ------------ | ---- | ------------------ | ---------------- |
| GET    | `/books`     | 🔒   |                    | List all books   |
| GET    | `/books/:id` | 🔒   |                    | Get a book by ID |
| POST   | `/books`     | 🔒   | `admin`, `teacher` | Create a book    |
| PATCH  | `/books/:id` | 🔒   | `admin`, `teacher` | Update a book    |
| DELETE | `/books/:id` | 🔒   | `admin`            | Delete a book    |

### Loans

| Method | Path         | Auth | Roles   | Description                    |
| ------ | ------------ | ---- | ------- | ------------------------------ |
| GET    | `/loans`     | 🔒   |         | List all loans                 |
| GET    | `/loans/:id` | 🔒   |         | Get a loan by ID               |
| POST   | `/loans`     | 🔒   |         | Borrow a book (creates a loan) |
| PATCH  | `/loans/:id` | 🔒   | `admin` | Return a borrowed book         |
| DELETE | `/loans/:id` | 🔒   | `admin` | Delete a loan                  |

### Institutions

| Method | Path                | Auth | Roles   | Description             |
| ------ | ------------------- | ---- | ------- | ----------------------- |
| GET    | `/institutions`     |      |         | List all institutions   |
| GET    | `/institutions/:id` | 🔒   |         | Get an institution by ID |
| POST   | `/institutions`     | 🔒   | `admin` | Create an institution   |
| DELETE | `/institutions/:id` | 🔒   | `admin` | Delete an institution   |

### Events

| Method | Path          | Auth | Roles              | Description       |
| ------ | ------------- | ---- | ------------------ | ----------------- |
| GET    | `/events`     | 🔒   |                    | List all events   |
| GET    | `/events/:id` | 🔒   |                    | Get an event by ID |
| POST   | `/events`     | 🔒   | `admin`, `teacher` | Create an event   |
| DELETE | `/events/:id` | 🔒   | `admin`, `teacher` | Delete an event   |

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
    "institution_id": 1
  }'
```

`role` must be one of `student`, `teacher`, `admin`. `institution_id` is optional.

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
    "status": "available"
  }'
```

`status` must be one of `available`, `lent`. Requires an `admin` or `teacher` token.

### Borrow a book

```bash
curl -X POST http://localhost:8080/loans \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <jwt>" \
  -d '{ "book_id": 1, "return_date": "2026-07-10T00:00:00Z" }'
```

The borrower is taken from the JWT, not the request body. Borrowing assigns the
provided `return_date`. A book that is not `available` returns `409 Conflict`, and an
unknown `book_id` returns `404 Not Found`.

### Return a book

```bash
curl -X PATCH http://localhost:8080/loans/1 \
  -H "Authorization: Bearer <jwt>"
```

Returning marks the loan's `returned_at` and sets the book back to `available`.
A loan that was already returned responds with `409 Conflict`. Requires an `admin` token.

### Create an institution

```bash
curl -X POST http://localhost:8080/institutions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <jwt>" \
  -d '{
    "name": "Universidade Veiga de Almeida",
    "abbreviation": "UVA",
    "city": "Rio de Janeiro",
    "address": "Rua Ibituruna, 108"
  }'
```

All fields are required. Requires an `admin` token.

### Create an event

```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <jwt>" \
  -d '{
    "name": "Book Fair 2026",
    "description": "Annual community book fair",
    "date": "2026-09-12T14:00:00Z",
    "location": "Main Campus Auditorium",
    "institution_id": 1
  }'
```

All fields are required. `institution_id` must reference an existing institution.
Requires an `admin` or `teacher` token.

## Data model

- **users** — id, name, birth_date, email (unique), password, address, document
  (unique), cellphone (unique), role, institution_id → institutions, score,
  created_at.
- **books** — id, title, author, release_date, edition, status, created_at.
- **loans** — id, user_id → users, book_id → books, return_date, returned_at.
  Deleting a user or book cascades to their loans.
- **institutions** — id, name, abbreviation, city, address, created_at.
- **events** — id, name, description, date, location, institution_id → institutions,
  created_at.

## Error responses

Errors are returned as plain text with an appropriate HTTP status. Common cases:

- `400 Bad Request` — invalid payload or malformed ID
- `401 Unauthorized` — missing/invalid token on a protected route
- `403 Forbidden` — authenticated but lacking the required role
- `404 Not Found` — resource does not exist
- `409 Conflict` — book unavailable, or loan already returned
- `500 Internal Server Error` — unexpected failure
