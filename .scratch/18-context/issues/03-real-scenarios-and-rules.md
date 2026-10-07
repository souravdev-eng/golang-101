# 03: Real scenarios (HTTP, graceful shutdown) and the rules-and-pitfalls capstone

**What to build:** A learner who understands the tools and patterns can now see context in the places it actually earns its keep — serving an HTTP request, calling another service, and shutting a program down cleanly — and then read one commented file that collects the rules and anti-patterns. Every example is self-contained (no real network, no human input) and prints exactly the output documented in its README section.

Delivers examples `11`–`14`, each appended to the section README, and completes the section:

- `11-http-server` — an `http.HandlerFunc` that reads `r.Context()` and abandons slow work when the request's context is done (client disconnect / server timeout). Driven with `net/http/httptest` from `main` so output is deterministic.
- `12-http-client` — an outbound request built with `http.NewRequestWithContext` and a timeout, aborting when a deliberately slow endpoint (a local `httptest.NewServer`) is slower than the timeout.
- `13-graceful-shutdown` — `signal.NotifyContext` turns SIGINT into a cancelled context; the program drains work and exits cleanly. The signal/cancel is triggered programmatically so the example terminates deterministically.
- `14-rules-and-pitfalls` — a short runnable file demonstrating and commenting the rules in one place: `ctx` is the first parameter named `ctx`; always `defer cancel()` even when a timeout fires first; don't store a `Context` in a struct — pass it; `WithValue` is for request-scoped data only, not optional parameters; always check `ctx.Err()` to learn *why* work stopped.

**Blocked by:** 02 (the capstone summarizes propagation and values; the HTTP/shutdown examples build on timeout and `Done`/`Err` from 01 via 02).

**Status:** done

- [x] `11`–`14` each exist as their own runnable folder, listed in the section README index with a run command, Expected output block, and a Try line.
- [x] `11-http-server` and `12-http-client` use `httptest` and require no real network; both terminate on their own.
- [x] `13-graceful-shutdown` triggers its own signal/cancellation and exits without human input.
- [x] `14-rules-and-pitfalls` demonstrates each listed rule with a beside-the-code comment explaining it.
- [x] All examples build on Go 1.22.1 with no external modules; each matches its README Expected output exactly.
- [x] The section README index is complete (`01`–`14`) and the course index reflects the finished `18-context` section.
