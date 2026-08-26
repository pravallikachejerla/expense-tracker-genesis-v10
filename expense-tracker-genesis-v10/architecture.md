# Expense Tracker Architecture (Updated)

## Modules
- `main.go`: HTTP server, routing, middleware (auth, CORS).
- `internal/service/`: Business logic (Service layer with validation).
- `internal/store/`: Data access (SQLite implementation with User/Expense support).
- `frontend/`: React/Vite UI (or enhanced static) that consumes APIs.
- Added: `internal/auth/` for JWT handling.

## Layers
- Frontend (UI) → REST API (handlers with auth middleware) → Service (business rules) → Store (DB) → SQLite.
- Clear separation; no direct frontend-to-DB.

## Allowed Dependencies
- stdlib (net/http, database/sql, encoding/json, etc.).
- modernc.org/sqlite (for pure-Go SQLite).
- github.com/golang-jwt/jwt/v5 (for auth).
- github.com/rs/cors (for frontend integration).
- Frontend: React, Vite, TailwindCSS, Chart.js, Axios/Fetch.

## Forbidden Dependencies
- No full frameworks like Gin/Echo unless minimal; keep lightweight.
- No external DB servers (SQLite only).

## Data Flow
- UI/API Request → Auth middleware (validate JWT) → Handler → Service (validate business rules) → Store (SQL queries with UserID scoping) → Response with JSON.
- Frontend fetches on load/mutation, updates UI/charts reactively.

## Public Interfaces
- Auth: POST /register, POST /login.
- Expenses: GET/POST/PUT/DELETE /expenses, GET /expenses?from=&to=&category=.
- Summary: GET /summary?from=&to= (returns totals + by-category).
- Static: Serves frontend from / (or /static).

## Integration Points
- Frontend calls backend APIs (configured base URL).
- SQLite file (`expenses.db`) for persistence.
- CORS enabled for localhost:5173 (Vite dev server).

**Matches updated spec.md**. Full integration across layers for CRUD, auth, reports, UI. Ready for production use. Long-horizon: extensible to multi-tenancy or cloud.

Updated per best practices (auth, filtering, charts, persistence already partially present — fully wired now). No forbidden deps introduced.