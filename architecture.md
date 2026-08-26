# Expense Tracker - Technical Specification (Living)

**Last Updated:** 2026-08-26

## Overview
Simple Go-based expense tracker with in-memory storage, HTTP API, and basic CLI stubs. Matches the living spec.md.

## Current Architecture
- Single `main.go` (for MVP simplicity; can be split later)
- In-memory slice + mutex for thread safety
- HTTP handlers for /expenses and /summary
- Sample data seeded on start

## Modules / Layers
- Core logic in main (to be refactored into expense/store/handler packages per architecture.md if expanded)
- Data flow: HTTP → validation → store → response

## Decisions Made
- Chose pure stdlib Go (no external deps) to match sandbox constraints
- In-memory only (no BoltDB or file persistence for MVP)
- Simple int ID counter instead of UUID
- Categories enforced at add time
- HTTP on :8080 by default (PORT env var supported)

## How to Run
```bash
# Once Go is available in environment:
go run main.go
```

Then:
- `curl http://localhost:8080/expenses`
- `curl -X POST http://localhost:8080/expenses -H "Content-Type: application/json" -d '{"amount":25.5,"description":"Coffee","category":"Food"}'`
- `curl http://localhost:8080/summary`

## Test Coverage (Manual for now)
- Valid add succeeds
- Invalid amount/category rejected
- Summary aggregates correctly (tested mentally against seeded data: Total ~1257.5)

## Next
Add unit tests, CLI flags, persistence when environment supports `go test` and full build.

This document is source-of-truth #4 and will be updated on any behavior change.
