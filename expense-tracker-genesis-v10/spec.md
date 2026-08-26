# Expense Tracker

## Goal
Build a simple, functional expense tracking application that allows users to record, categorize, view, and summarize expenses. The app supports a basic HTTP API.

## Users
- Individual users managing personal expenses.

## Features
- Add expense (amount, description, category, date)
- List expenses
- Get spending summary by category
- In-memory storage with sample data

## Workflows
- POST to add expense (validated)
- GET to list or summarize

## Data Model
- Expense struct with ID, Amount, Description, Category, Date

## Business Rules
- Amount > 0
- Category from fixed list
- Date defaults to now

## Constraints
- Go + stdlib only
- In-memory only

## Non-Functional Requirements
- Runs on port 8080
- Thread-safe
- Clear error messages

## Edge Cases
- Invalid inputs rejected with 400
- Empty state handled

## Acceptance Criteria
- All listed in acceptance_tests.md — all pass.

## Non-Goals
- Persistence, auth, UI, advanced filtering

Living document — updated with implementation.
