# Expense Tracker

## Goal
Build a simple, functional expense tracking application that allows users to record, categorize, view, and summarize expenses. The app will support both a CLI for quick operations and a basic HTTP API for extensibility.

## Users
- Individual users managing personal or small-team expenses.

## Features
- Add expense (amount, description, category, date)
- List all expenses
- Filter expenses by category or date range
- Generate summary (total spent, by category)
- In-memory storage (no external DB for MVP)
- CLI commands and HTTP endpoints (/expenses, /summary)

## Workflows
1. User adds an expense via CLI or POST /expenses
2. User lists expenses or queries summary
3. Data persisted in-memory during runtime

## Data Model
- Expense: ID (uuid), Amount (float64), Description (string), Category (string), Date (time.Time)

## Business Rules
- Amount must be positive
- Category is required and one of: Food, Transport, Rent, Entertainment, Utilities, Other
- Date defaults to now if not provided

## Constraints
- Go only (per environment)
- In-memory only (no files/DB)
- No auth for MVP

## Non-Functional Requirements
- Runnable with `go run main.go`
- Clean, idiomatic Go code
- Living spec updated with any changes

## Edge Cases
- Zero/negative amount (reject)
- Invalid category (reject)
- Empty list (graceful summary = 0)
- Concurrent access (basic mutex protection)

## Acceptance Criteria
- Can add valid expense and see it in list
- Summary correctly aggregates by category
- Invalid inputs return clear errors
- CLI and API both functional

## Non-Goals
- Persistence, user accounts, advanced reporting, mobile UI, cloud deployment

## Architecture Rules

### Modules
- main: entrypoint (CLI + HTTP server)
- expense: core types and service
- store: in-memory repository
- handler: HTTP handlers

### Layers
- Handler → Service → Store

### Allowed Dependencies
- Standard library only (net/http, encoding/json, sync, time, fmt, flag)

### Forbidden Dependencies
- External packages (no gorilla, no uuid libs — use simple int ID for MVP)

### Data Flow
CLI/HTTP → Handler → Service (validate/business rules) → Store (CRUD)

### Public Interfaces
- CLI flags: add, list, summary
- HTTP: GET/POST /expenses, GET /summary

### Integration Points
- None (self-contained)

This spec is living — last updated 2026-08-26.
