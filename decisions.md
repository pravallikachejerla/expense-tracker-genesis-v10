# Expense Tracker - Decisions Log (Living)

**Last Updated:** 2026-08-26

## Key Decisions
1. **Language & Stack**: Pure Go using only stdlib (net/http, sync, encoding/json). Matches sandbox provisioning and "go only" note. No Node/Python/JS parallel impl.

2. **Storage**: In-memory slice with mutex. Simple for MVP; avoids file/DB setup issues in sandbox. Can be replaced with embedded store later.

3. **Architecture**: Single main.go for initial delivery (easy to run). Matches current architecture.md layers (can split into packages without breaking API).

4. **ID Generation**: Simple incrementing int (not UUID). Avoids external deps.

5. **CLI vs API**: Prioritized HTTP API with CLI stub in main. Full CLI flags can be added via flag package if needed.

6. **Branch & Commit**: Using the pre-created feature branch. Conventional commit style (feat:, fix:). One commit for initial scaffold + living docs.

7. **Living Docs**: spec.md, architecture.md, acceptance_tests.md, decisions.md all created/updated as source-of-truth. Updated on every behavior change.

Assumptions: Sandbox will eventually have Go on PATH for `go run`/`go build`. If not, this is a runnable artifact once fixed.

Non-goals (from spec): persistence, auth, UI.

This is source-of-truth #6.
