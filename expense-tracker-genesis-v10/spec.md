# Expense Tracker - Updated Specification (Living Document)

## Goal
Build a complete, production-ready expense tracking application with full CRUD operations, persistence, authentication, a responsive frontend UI with visualizations, and proper API integration. Supports multi-user scenarios with secure data isolation.

## Users
- Individual users managing personal or family expenses.
- Multiple users with separate accounts (authenticated).

## Features
- **CRUD for Expenses**: Create, read, update, delete expenses with amount, description, category, date.
- **Filtering and Search**: By date range, category, amount.
- **Summaries and Reports**: Total spending, breakdown by category, monthly trends (with charts).
- **Categories**: Predefined + ability to add custom.
- **Authentication**: User registration, login, JWT-based protected routes (multi-user support).
- **Frontend UI**: Responsive single-page app with forms, lists, dashboard, charts (using Chart.js).
- **Database**: SQLite with proper schema, indexes, and migrations.
- **API**: RESTful endpoints with validation, error handling, CORS.

## Workflows
- User registers/logs in → JWT token.
- Authenticated user adds/views/edits/deletes expenses via UI or API.
- Dashboard shows real-time summaries and charts.
- Data persisted across restarts.

## Data Model
- **User**: ID, Username, PasswordHash, CreatedAt.
- **Expense**: ID, UserID, Amount, Description, Category, Date, CreatedAt.
- **Category**: ID, Name, UserID (optional for custom).

## Business Rules
- Amount must be > 0.
- Date cannot be in future (or allow with warning).
- Categories validated against allowed list or user-specific.
- Users can only access their own expenses.
- All operations logged with timestamps.

## Constraints
- Go backend with SQLite (modernc.org/sqlite).
- React/Vite + Tailwind frontend (or vanilla JS if simpler for integration).
- JWT for auth (golang-jwt/jwt).
- No external cloud services; self-contained.

## Non-Functional Requirements
- Runs on port 8080 (backend), 5173 (frontend dev).
- Secure (hashed passwords, JWT).
- Responsive UI, accessible.
- Thread-safe, error-resilient.
- Performance: Efficient queries with indexes.

## Edge Cases
- Invalid auth, duplicate categories, zero/negative amounts, malformed dates, empty DB, concurrent updates, large datasets.
- Graceful error messages in UI and API (4xx/5xx).

## Acceptance Criteria
- All CRUD operations work via UI and API.
- Auth protects endpoints; data isolated per user.
- Summaries and charts update correctly.
- Persistence survives restarts.
- Frontend integrates seamlessly with backend APIs.
- All living docs updated; tests/build pass.
- Verified via manual curl + browser.

## Non-Goals
- Advanced analytics (ML predictions), mobile app, multi-currency (unless extended), cloud deployment in this iteration.

**Updated**: 2026-08-26 with best-practice features per user clarification. Integrated across frontend, backend, APIs, DB. Matches architecture rules.

Living document — update on any behavior change.