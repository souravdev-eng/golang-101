# 01: Scaffold the section and teach the core cancellation tools

**What to build:** A learner can open a new `18-context/` section, see it listed in the course index, and run the first six examples to go from *feeling the problem* to *using every core cancellation tool*. Running each example from the repo root prints exactly the output documented in the section README.

Concretely this delivers the section scaffolding (the prefactor every later ticket appends to) plus examples `01`–`06`:

- `01-why-context` — a goroutine with no way to stop it; shows the pain context removes (no `context` used yet).
- `02-background` — `context.Background()` as the empty root: `Done()` never fires, `Err()` is nil.
- `03-with-cancel` — a worker watching `ctx.Done()` in a `select`; caller calls `cancel()` to stop it.
- `04-with-timeout` — `context.WithTimeout` auto-cancels after a duration; introduces `defer cancel()`.
- `05-with-deadline` — `context.WithDeadline` cancels at a fixed instant; contrast duration vs moment.
- `06-done-and-err` — `ctx.Done()` as a channel and `ctx.Err()` returning `context.Canceled` vs `context.DeadlineExceeded`.

The README skeleton (title, one-line purpose, `Prerequisite: 17-concurrency`, ordered index, one subsection per example with run command + fenced Expected output + a "Try" line) must mirror `15-standard-library/README.md` and `17-concurrency/README.md`. Cross-link back to `15-standard-library/09-context` as the gentle intro.

**Blocked by:** None (can start immediately).

**Status:** done

- [x] `18-context/` exists at the repo root and is listed in the course index (root `README.md` / `NOTES.md`) consistent with how `17-concurrency` is listed.
- [x] Section `README.md` follows the established structure (purpose, prerequisite, ordered index, per-example run + Expected output + Try).
- [x] Each of `01`–`06` is its own folder with a `package main` `main.go`, runnable via `go run ./18-context/NN-name`, no external modules, builds on Go 1.22.1.
- [x] Every helper takes `ctx context.Context` as its first parameter named `ctx`; every derived context releases `cancel` via `defer cancel()`.
- [x] Comments are plain-language and sit beside the code, explicitly naming syntax a JS/TS learner may forget (`<-ch`, `:=`, `select`).
- [x] Each example's stdout is deterministic and matches its README Expected output block exactly (wide duration gaps / pre-filled buffered channels, no racy sleeps).
