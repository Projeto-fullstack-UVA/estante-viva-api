# TODO — Fixes & Improvements

Issues found during a review of the codebase, grouped by priority. Each item notes
the location and the suggested fix.

---

## 🔴 Critical (security / broken behavior)

### 1. Privilege escalation: clients can self-assign any role on registration
`internals/dtos/users/create_user_request.go:18` — `role` is bound directly from the
request body with `oneof=student teacher donator admin`. Anyone hitting
`POST /users` can register as `"role": "admin"`. The same applies to `score`
(`create_user_request.go:20`), which a client can set to any value.
**Fix:** force `role` to a safe default (`student`) on registration and ignore
client-supplied `role`/`score`; gate role changes behind an admin-only endpoint.

### 2. `Register` issues a token / response with `id = 0`
`internals/repositories/user_repository.go:82` (`CreateUser`) uses `Pool.Exec` and
never reads back the generated identity, so `user.ID` stays `0`. Then
`internals/services/user_service.go:50` signs a JWT with `user_id: 0` and
`NewRegisterUserResponse` returns `"id": 0`. New users get a useless token.
**Fix:** use `INSERT ... RETURNING id` and scan it into `user.ID` before generating
the token (mirror `CreateLoan`).

### 3. CORS does not allow the `Authorization` header
`cmd/app/main.go:28` — `Access-Control-Allow-Headers` is only `Content-Type`. Every
authenticated browser request sends `Authorization`, so the preflight fails and the
frontend cannot call any protected route.
**Fix:** add `Authorization` (and likely the `Authorization` response header to
`Access-Control-Expose-Headers`, since `Login` returns the token in that header).

### 4. No authorization layer despite the auth middleware
`cmd/app/main.go:60-69` — every authenticated user can list all users, create books,
and create loans for **any** `user_id`. `BorrowBook` trusts the body's `user_id`
(`loan_controller.go:50`) instead of the `user_id` the middleware put in context
(`middleware/authentication.go:42`).
**Fix:** add role-based checks (the role is already in the JWT claims) and derive the
borrower from the authenticated context, not the request body.

---

## 🟠 High

### 5. `JWT_SECRET_KEY` is not validated at startup
`cmd/app/main.go:45-48` validates `DATABASE_URL` but not `JWT_SECRET_KEY`. The server
boots fine and then every auth call fails with `500`
(`middleware/authentication.go:13`).
**Fix:** check `JWT_SECRET_KEY` at startup and `log.Fatal` if missing.

### 6. Internal error messages leaked to clients
Multiple controllers return `c.String(http.StatusInternalServerError, err.Error())`
(e.g. `user_controller.go:26,53,73`, `book_controller.go:17,36`,
`loan_controller.go:18,35`). This exposes raw DB/internal errors.
**Fix:** log the real error server-side, return a generic message to the client.

### 7. Duplicate-key registration returns 500 instead of 409
`email`, `document`, and `cellphone` are `UNIQUE` (`migrations/0001_create_users.sql`).
A duplicate makes `Register` return the generic "failed to register user" with status
`500` (`user_service.go:44`).
**Fix:** detect the Postgres unique-violation (`23505`) and return a `409 Conflict`
with a clear message.

### 8. Loan operations are not transactional
`BorrowBook` (`services/loan_service.go:47-54`) and `ReturnBook`
(`loan_service.go:73-79`) do `CreateLoan`/`ReturnLoan` and `UpdateBookStatus` as two
separate statements. If the second fails, the book status and loan state diverge.
**Fix:** wrap each operation in a single `pgx` transaction.

### 9. ISO timestamp written into a `DATE` column
`services/loan_service.go:11,45,71` formats dates as
`2006-01-02T15:04:05.000Z` and inserts them into `return_date` / `returned_at`, which
are `DATE` columns (`migrations/0003_create_loans.sql`). This relies on Postgres
coercing a full timestamp string to `date` and silently drops the time component.
**Fix:** either store `TIMESTAMPTZ` columns, or format as `2006-01-02`. Verify the
insert actually succeeds against Postgres.

---

## 🟡 Medium

### 10. `role` allow-list mismatch + typo between DTO and DB
DTO allows `student teacher donator admin`
(`create_user_request.go:18`); the DB `CHECK` allows
`student, teacher, donator, admin, voluteer` (`migrations/0001_create_users.sql:10`),
which also contains the typo **"voluteer"** (should be "volunteer").
**Fix:** reconcile the two lists and fix the spelling.

### 11. `Book.created_at` should be server-set
`CreateBookRequest` requires the client to send `created_at` (and `status`)
(`create_book_request.go:15-16`). `created_at` is bookkeeping the server should own.
**Fix:** set `created_at = time.Now()` server-side; default `status` to `available`.

### 12. Loan does not validate that `user_id` exists
`BorrowBook` checks the book but never the user (`services/loan_service.go:33-47`).
A bad `user_id` hits the FK constraint and surfaces as a generic `500`.
**Fix:** validate the user up front (or map the FK violation to `404`).

### 13. `Password` field exposed on the `User` entity
`internals/entities/user.go:9` tags the password with `json:"password"`. Current
handlers use DTOs, but any direct marshal of `User` would leak the hash.
**Fix:** use `json:"-"` on the password field.

### 14. Wrong-credentials returns `404 User not found` instead of `401`
`services/user_service.go:21-23` maps a bad password to `ErrUserNotFound`, and the
controller returns `404` (`user_controller.go:22-24`).
**Fix:** return `401 Unauthorized` for invalid credentials (a uniform
"invalid credentials" message is fine and avoids user enumeration).

---

## 🟢 Low (cleanup / polish)

### 15. Inconsistent `Book.ID` type
`entities.Book.ID` and `BookResponse.ID` are `string`
(`entities/book.go:8`, `dtos/books/book_response.go:10`), while loans reference books
as `int64` (`book_id`). Pick one type for consistency.

### 16. Dead / unused code
`UpdateUserPassword` (`repositories/user_repository.go:95`) and
`NewRegisterUserResponseList` (`dtos/users/user_response.go:88`) are never called.
**Fix:** remove or wire them up.

### 17. Hardcoded CORS origins
`cmd/app/main.go:16` hardcodes the allowed origins.
**Fix:** drive them from an environment variable.

### 18. No graceful shutdown / pool cleanup
`cmd/app/main.go` runs the server with no signal handling and never closes the pgx
pool.
**Fix:** add `http.Server` with `Shutdown` on `SIGINT`/`SIGTERM` and `Pool.Close()`.

### 19. Date fields require full RFC3339 timestamps
`birthDate`, `release_date` are `time.Time`, so clients must send
`1999-10-30T00:00:00Z` for what are logically plain dates.
**Fix:** accept a `YYYY-MM-DD` format with a custom type/binding.

### 20. No password strength / length validation
`create_user_request.go:14` only marks `password` as `required`.
**Fix:** add a minimum-length (and optionally complexity) binding.

### 21. Unsafe nil-deref pattern in list scanners
`GetUsers`/`GetBooks`/`GetLoans` do `append(list, *x)` where the scan helper can
return `(nil, nil)` on `ErrNoRows` (e.g. `user_repository.go:45`). Safe today because
it only runs after `rows.Next()`, but fragile.
**Fix:** use a dedicated row-scan that doesn't translate `ErrNoRows` to `nil` inside
loops.
