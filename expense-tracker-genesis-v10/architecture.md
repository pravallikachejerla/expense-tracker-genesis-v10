# Expense Tracker Architecture

## Modules
- main.go contains all for MVP (handler, service, store co-located)

## Layers
- HTTP handlers → business logic (addExpense, getSummary) → in-memory store

## Allowed Dependencies
- stdlib only (http, json, sync, time, fmt, log, os)

## Forbidden Dependencies
- No external modules

## Data Flow
Request → validation → mutex-protected store → JSON response

## Public Interfaces
- GET /expenses
- POST /expenses
- GET /summary

## Integration Points
None.

Matches spec. Ready for expansion.
