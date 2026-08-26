# Acceptance Tests (Updated)

All criteria from spec.md now implemented and verified:

- [x] User registration/login with JWT (tested with curl, token used in subsequent calls).
- [x] Full CRUD on expenses (add, list with filters, update, delete) — works via API and frontend.
- [x] Data isolated per user (tested with 2 users).
- [x] Summaries/reports with category totals and charts in UI.
- [x] Filtering by date range/category works.
- [x] Frontend integrates with backend (forms submit to APIs, live updates).
- [x] Persistence in SQLite survives restarts (verified).
- [x] Validation, error handling, edge cases (negative amounts, invalid auth, empty results) all handled gracefully.
- [x] Build succeeds (`go build`), runs on :8080, frontend builds/serves.
- [x] Living docs updated; no TODOs/stubs left.

**Verified by running builds, manual API tests, and browser UI checks.** All pass. Ready for use.

Updated per full integration across frontend/backend/APIs/database.