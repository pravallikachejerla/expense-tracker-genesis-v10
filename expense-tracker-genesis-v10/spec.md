# Expense Tracker - Updated Specification (Living Document)

## Goal
Build a complete, functional expense tracking application with full CRUD operations, categorization, filtering, reporting, SQLite persistence, a REST API, and an integrated responsive frontend.

## Users
- Individual users managing personal or small-team expenses.

## Features
- Full CRUD for expenses (create, read, update, delete)
- Predefined and dynamic categories
- Filter expenses by category, date range, or search
- Real-time summary (total spent, breakdown by category, basic bar chart)
- SQLite database for persistence
- RESTful JSON API
- Modern single-page frontend (Tailwind + vanilla JS) fully integrated with backend
- Seeded sample data

## Workflows
1. User adds/edits expenses via UI or POST/PUT /expenses
2. Expenses are validated, persisted to SQLite, and reflected in UI
3. User can filter, view list, delete items
4. Summary and visualizations update live via API calls
5. Categories can be managed

## Data Model
- Expense: ID (int), Amount (float64 > 0), Description (string), Category (string), Date (time.Time)
- Categories stored in DB (with defaults)

## Business Rules
- Amount must be positive
- Category must be valid (from list or added)
- Date defaults to now
- All operations are validated in service layer

## Constraints
- Go 1.21 + standard library + modernc.org/sqlite
- Single binary + embedded frontend
- No authentication for MVP

## Non-Functional Requirements
- Runs on port 8080 (`go run main.go`)
- Responsive UI, clean idiomatic Go, thread-safe
- Living documents updated on every change
- Tests for core logic

## Edge Cases
- Negative/zero amount (400 error)
- Invalid category (rejected)
- Empty DB (graceful 0 totals, empty list)
- Concurrent access (handled by DB)
- Malformed JSON, missing fields

## Acceptance Criteria
- All CRUD operations work via UI and API (verified)
- Filters and search function correctly
- Summary and chart update accurately
- Data persists across restarts
- All tests pass
- Frontend and backend fully integrated

## Non-Goals
- User authentication, multi-tenancy, mobile app, cloud deployment, advanced analytics

## Architecture Rules

### Modules
- main: HTTP server, routing, static file serving
- internal/service: business logic and validation
- internal/store: SQLite persistence layer

### Layers
- Frontend (JS) → HTTP API → Handler → Service (validation/rules) → Store (DB) → SQLite

### Allowed Dependencies
- net/http, encoding/json, database/sql, modernc.org/sqlite, time, sync, fmt, log, os

### Forbidden Dependencies
- No heavy web frameworks, no external auth libs

### Data Flow
Browser/UI → API endpoint → Service → Store → SQLite file (expenses.db). Responses flow back for live updates.

### Public Interfaces
- GET/POST/PUT/DELETE /expenses
- GET /summary
- GET/POST /categories
- Static / (serves index.html)

### Integration Points
- Frontend calls backend API directly (same origin when served together)
- SQLite file for persistence

Last updated: 2026-08-26. This document is the source of truth and has been updated to match the implemented features.
