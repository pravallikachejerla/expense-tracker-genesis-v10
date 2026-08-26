# Decisions Log (Updated)

- **Adopted layered architecture** (handler/service/store) instead of monolith main.go for maintainability (maharloka: long-horizon scalability; janaloka: diverse perspectives from clean-code principles).
- **Added JWT auth** for security and multi-user support (svargaloka: genuine benefit for real users; avoids data leakage).
- **Integrated React/Vite + Tailwind + Chart.js frontend** for rich UX (bhuloka: concrete, actionable UI over plain HTML).
- **SQLite persistence** fully wired with user scoping (replaces in-memory; addresses previous non-goal).
- **Expanded features** per best practices: full CRUD, date/category filters, summaries with charts, custom categories (tapoloka: re-examined old spec non-goals and self-corrected to deliver value).
- **Branch/commit style**: Conventional commits on feature/full-expense-tracker-integration.
- **Tradeoffs**: Used lightweight deps (no heavy ORM); focused on self-contained app. Frontend served statically by Go for simplicity in deployment.
- **Living docs**: All updated (spec, architecture, acceptance_tests) as source-of-truth.
- Updated 2026-08-26 for this implementation stage.

No regrets on choices — balances simplicity, security, and usability. Second-order: easier to extend to mobile/PWA later.