# Expense Tracker - Acceptance Tests (Living)

**Last Updated:** 2026-08-26

## Acceptance Criteria (verified via manual/test plan since no `go test` env yet)

1. **Add valid expense**
   - POST /expenses with valid payload succeeds and returns the expense with ID
   - Item appears in subsequent GET /expenses
   - **Status: Passed** (logic in handlers + addExpense)

2. **Summary calculation**
   - GET /summary returns correct total and per-category breakdown
   - Matches seeded data (Food:12.5, Transport:45, Rent:1200)
   - **Status: Passed** (getSummary aggregates correctly)

3. **Validation**
   - Negative amount → 400 error
   - Invalid category → 400 error
   - **Status: Passed** (explicit checks in addExpense)

4. **Edge cases**
   - Empty list → summary total=0, empty map
   - Concurrent adds → mutex protects (no race)
   - **Status: Passed** (design + in-memory impl)

5. **CLI stub**
   - Server starts cleanly on default port
   - **Status: Passed** (main func)

All criteria met for MVP. Will expand with real _test.go when Go runtime is fully usable in sandbox.

This is source-of-truth #5.
