# TODO — Estante Viva API

Improvements and implementations to be made, grouped by priority. Items reference
real spots in the codebase so they can be picked up directly.

## Correctness & data integrity

- [ ] **Make borrow/return atomic with a DB transaction.** In
  `services/loan_service.go`, `BorrowBook` runs `CreateLoan` then
  `UpdateBookStatus("lent")` as two separate statements — if the second fails,
  a loan exists but the book stays `available`. `ReturnBook` and `DeleteLoan`
  have the same gap. Wrap each flow in a `pgx.Tx` (begin → create loan + update
  status → commit, rollback on error).
- [ ] **Guard against double-borrow races.** Two concurrent `POST /loans` for the
  same book can both read `status = available` before either updates it. Use
  `SELECT ... FOR UPDATE` inside the transaction, or a conditional
  `UPDATE books SET status='lent' WHERE id=$1 AND status='available'` and treat
  0 rows affected as `ErrBookNotAvailable`.
- [ ] **Distinguish duplicate-key errors on register.** `services.Register`
  returns a generic 500 when email/document/cellphone are already taken (unique
  constraints). Detect pgx `23505` and return `409 Conflict` with a clear field.

## Security & auth

- [ ] **Centralize the JWT secret.** `middleware/authentication.go` reads
  `os.Getenv("JWT_SECRET_KEY")` directly on every request instead of using the
  `environment` package that the rest of the app loads from. Use
  `environment.JwtSecretKey` for one source of truth.
- [ ] **Enforce loan ownership.** `GET /loans` and `GET /loans/:id` let any
  authenticated user read every loan. Scope non-admins to their own loans
  (filter by `user_id` from the token), or require admin/teacher.
- [ ] **Validate JWT claims explicitly.** `auth.VerifyToken` doesn't assert
  `token.Valid` or check `exp`/expected claim types. Verify validity and the
  presence/type of `user_id` and `role` before trusting them.
- [ ] **Don't leak internal error detail in token errors.** `VerifyToken`
  wraps the raw parser error into the message; return a generic
  "invalid token" to clients and log the detail server-side.
- [ ] **Add request body size limits and basic rate limiting** (especially on
  `/login` and `/users`) to reduce brute-force / abuse surface.
- [ ] **Set Gin to release mode in production** (`gin.SetMode(gin.ReleaseMode)`
  driven by an env var) instead of always `gin.Default()` in debug.

## API consistency & validation

- [ ] **Standardize error responses as JSON.** Controllers mix `c.String(...)`
  for errors with `c.JSON(...)` for success. Adopt a single error envelope,
  e.g. `{ "error": "...", "code": "..." }`, and a small helper.
- [ ] **Remove `created_at` from request bodies / docs.** The book repo already
  overrides it with `time.Now()`; the README "Create a book" example still
  sends `created_at`. Drop it from the public contract.
- [ ] **Add validation to remaining DTOs.** `oneof` is used for book `status`;
  apply the same to user `role` (`student|teacher|donator|admin`) and add
  `required`/format constraints to user create/update and institution DTOs.
- [ ] **Add pagination to list endpoints** (`/users`, `/books`, `/loans`,
  `/events`, `/institutions`) with `limit`/`offset` (or cursor) query params —
  the queries currently `SELECT ... ORDER BY id` with no bound.
- [ ] **Add filtering/search** for books (by title/author/status) and loans
  (by user, active vs returned).

## Operations & lifecycle

- [ ] **Graceful shutdown.** `main.go` calls `router.Run` and never closes the
  pgx pool. Use `http.Server` with `signal.NotifyContext`, drain in-flight
  requests, then `repositories.Pool.Close()`.
- [ ] **Make the port configurable** via env (`PORT`, default `8080`) instead of
  the hardcoded `:8080`.
- [ ] **Add a health/readiness endpoint** (`GET /health`) that pings the DB pool
  for liveness checks and container orchestration.
- [ ] **Adopt a migration tool** (golang-migrate / goose) and a `make migrate`
  target instead of applying `migrations/*.sql` by hand. Update the README,
  which only lists migrations 0001–0003 (0004–0006 now exist).
- [ ] **Structured logging.** Replace the ad-hoc `log.Println` calls with a
  leveled/structured logger (slog) and add a request-ID/log middleware.

## Testing & CI

- [ ] **Add tests — there are currently zero `*_test.go` files.** Start with:
  - unit tests for `utils/password.go` and `auth/token.go` (generate/verify,
    expiry, tampered token, bad signing method);
  - service-layer tests for loan flows (borrow when unavailable, double return);
  - handler tests with `httptest` covering auth/authorization and status codes.
- [ ] **Add an integration test suite** against an ephemeral Postgres
  (testcontainers or a CI service) exercising the migrations + repositories.
- [ ] **Set up CI** (GitHub Actions): `go vet`, `gofmt -l`, `golangci-lint`,
  `go test ./...` on every PR.

## Documentation

- [ ] **Bring the README in sync with the actual API.** It's missing the
  `/me`, `/institutions`, and `/events` routes, the role-based authorization
  (`admin`/`teacher`), the `ALLOWED_ORIGINS` env var (required by
  `environment.go` but undocumented), and the PATCH/DELETE user & book routes.
  Also fix the borrow example: `POST /loans` now takes the user from the JWT,
  not a `user_id` in the body.
- [ ] **Publish an OpenAPI/Swagger spec** and serve it (e.g. swaggo) so the
  contract is machine-readable and the frontend can generate a client.

## Nice-to-haves

- [ ] Overdue-loan reporting (loans past `return_date` and not yet returned).
- [ ] User loan history endpoint and a `score` adjustment on late returns
  (the `users.score` column already exists).
- [ ] Dockerize end-to-end: verify the existing `Dockerfile` builds, add a
  `docker-compose.yml` with Postgres for local dev.
- [ ] Configurable loan period (currently a hardcoded 14 days in
  `loan_service.go`).
