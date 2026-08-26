# Expense Tracker Architecture (Updated)

## Modules
- `main.go`: HTTP server, API handlers, static frontend serving, wiring of Service + Store.
- `internal/service/service.go`: Business logic, input validation, orchestration.
- `internal/store/sqlite.go`: All database operations (CRUD, queries, schema).

## Layers
- Frontend (HTML/JS/Tailwind) → HTTP handlers (in main) → Service layer → Store layer → SQLite.

## Allowed Dependencies
- Standard library (net/http, database/sql, encoding/json, time, sync, etc.)
- modernc.org/sqlite (pure-Go SQLite driver — already in go.mod)

## Forbidden Dependencies
- No web frameworks (no Gin, Echo), no ORM, no external auth or charting libraries.

## Data Flow
1. User interacts with frontend (add/edit/delete/filter).
2. JS fetches/posts to API endpoints.
3. Handlers delegate to Service (validation + rules).
4. Service calls Store methods.
5. Store executes parameterized SQL against expenses.db.
6. Results flow back as JSON; frontend updates DOM and chart live.

## Public Interfaces
- `GET    /expenses` — list (supports ?category= & ?start= & ?end=)
- `POST   /expenses` — create
- `GET    /expenses/{id}` — get one
- `PUT    /expenses/{id}` — update
- `DELETE /expenses/{id}` — delete
- `GET    /summary` — totals + by-category
- `GET    /categories` — list categories
- `POST   /categories` — add new category
- Static file server for frontend at `/`

## Integration Points
- Frontend and backend served from same origin (no CORS issues).
- SQLite file (`expenses.db`) created in working directory for persistence across restarts.
- Sample data seeded on first run.

This matches the current implementation. The previous in-memory only version has been replaced with a full persistent, integrated stack. All layers are respected and tests cover the service and store.

Last updated: 2026-08-26.
