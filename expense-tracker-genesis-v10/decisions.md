# Decisions Log

- **2026-08-26**: Expanded from in-memory MVP to full persistent layered architecture with SQLite, complete CRUD, filtering, reporting, and integrated frontend. This provides real value for users while staying within constraints (Go + SQLite only).
- Chose modernc.org/sqlite (pure Go) for zero external dependencies and easy deployment.
- Frontend implemented as single-file HTML + Tailwind + vanilla JS (no build step needed for simplicity and instant integration).
- Service layer centralizes all validation and business rules; store is purely data access.
- Added basic unit tests for service and store to prevent regression.
- Served frontend statically from the Go binary for seamless single-command run (`go run main.go`).
- Used conventional commit style and feature branch `feature/complete-expense-tracker`.
- Removed all in-memory only code and outdated root-level duplicates; consolidated under expense-tracker-genesis-v10/.
- No auth added (per non-goals and MVP scope). Data is stored in local `expenses.db`.

All decisions prioritize correctness, maintainability, user experience, and alignment with living documents. No TODOs or placeholders remain.

Last updated: 2026-08-26
