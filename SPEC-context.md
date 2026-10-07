# The context package, learned through small runnable examples

## Problem Statement

The learner wants to go deep on Go's `context` package because it is central to
Go's concurrency patterns, but the repository currently teaches it with a single
example (`15-standard-library/09-context/main.go`). One example is enough for a
quick tour of the standard library, but not enough to build real understanding of
*what* context is, *how* its pieces fit together, the *design patterns* built on
it, and *when* to reach for it in real programs.

The learner comes from JS/TS, learns best by reading and changing small working
code rather than long prose, and sometimes forgets Go syntax — so explanations
must sit beside the code as plain-language comments, not in separate notes.

A good learning experience here means: a motivating problem shown *before* the
fix, one idea per runnable file, deterministic output to check against, and a
small "try this" change per example.

## Solution

Create a new top-level section `18-context/` — a dedicated deep-dive that sits
after `17-concurrency` in the course ordering. The existing
`15-standard-library/09-context` stays as the gentle one-example introduction;
this section is where the learner goes deep.

The section is an ordered sequence of small, independently `go run`-able `main.go`
files, each in its own numbered folder, each teaching one concept and each fitting
in about 5–10 minutes. A section `README.md` gives the reading order, how to run
each example, its expected output, and a small "Try" experiment — matching the
conventions already used by `15-standard-library` and `17-concurrency`.

Every example uses familiar, concrete scenarios (an inventory lookup, a slow API
call, a batch of workers, an HTTP handler) and heavy plain-language comments
placed next to the lines they explain, so the learner can recall forgotten syntax
without leaving the file.

The arc moves deliberately: first *feel the problem* (a goroutine that never
stops), then *the core tools* (cancel, timeout, deadline, Done/Err), then
*propagation and patterns* (passing context down a call chain, fan-out
cancellation, request-scoped values), then *real scenarios* (HTTP server and
client, graceful shutdown), and finally *the rules and anti-patterns* collected in
one place.

## User Stories

1. As a Go learner, I want a dedicated context section separate from the quick std-lib tour, so that I can study the topic in depth without the material feeling rushed.
2. As a Go learner, I want an ordered index with a suggested reading order, so that I know where to start and what comes next.
3. As a Go learner, I want each example in its own runnable folder, so that I can study and change one idea at a time.
4. As a Go learner, I want the expected output printed in the README, so that I can confirm my run matches the lesson.
5. As a Go learner, I want a small "Try" experiment per example, so that I can learn by changing working code.
6. As a Go learner, I want plain-language comments beside the code, so that I can recall syntax I have forgotten without reading separate notes.
7. As a Go learner, I want to first see a goroutine that never stops, so that I understand the problem context solves before I see the solution.
8. As a Go learner, I want to see `context.Background()` as the empty root context, so that I understand where every context tree begins.
9. As a Go learner, I want a `context.WithCancel` example, so that I can let a caller explicitly signal "stop" to running work.
10. As a Go learner, I want a `context.WithTimeout` example, so that I can bound how long I wait for slow work.
11. As a Go learner, I want a `context.WithDeadline` example, so that I understand how it differs from a timeout (a fixed moment vs a duration).
12. As a Go learner, I want to understand `ctx.Done()` as a channel and the `select` pattern around it, so that I can write code that reacts to cancellation.
13. As a Go learner, I want to see `ctx.Err()` return `context.Canceled` vs `context.DeadlineExceeded`, so that I can tell *why* work stopped.
14. As a Go learner, I want to see context passed down a chain of function calls, so that I understand how cancellation propagates through a program.
15. As a Go learner, I want to see that cancelling a parent context cancels its children, so that I understand the context tree.
16. As a Go learner, I want a fan-out example where one cancel stops many workers, so that I can coordinate shutdown of concurrent work.
17. As a Go learner, I want a `context.WithValue` example with a request ID, so that I can carry request-scoped data through a call chain.
18. As a Go learner, I want to see the correct way to define context value keys (a private key type), so that I avoid collisions and follow the idiom.
19. As a Go learner, I want an HTTP server handler using `r.Context()`, so that I see how a client disconnect cancels server-side work.
20. As a Go learner, I want an outbound HTTP call using `http.NewRequestWithContext`, so that I can apply a timeout to a request to another service.
21. As a Go learner, I want a graceful-shutdown example using `signal.NotifyContext`, so that I can stop a program cleanly on Ctrl+C.
22. As a Go learner, I want a final "rules and anti-patterns" example, so that I can see the do's and don'ts collected in one place.
23. As a Go learner, I want to see why `ctx` is always the first parameter and named `ctx`, so that I write idiomatic signatures.
24. As a Go learner, I want to see why `cancel` must always be called (via `defer`), so that I avoid leaking resources even when a timeout fires first.
25. As a Go learner, I want to see why context should not be stored in a struct, so that I avoid a common mistake.
26. As a Go learner, I want to see why `WithValue` is not for passing optional function parameters, so that I use it only for request-scoped data.
27. As a Go learner, I want deterministic output in every example despite concurrency and timing, so that my runs reliably match the expected output.
28. As a Go learner, I want the section listed from the repository READMEs/course index, so that I can find it as part of the course.
29. As a Go learner, I want prerequisites named (goroutines and channels), so that I know what to learn first.
30. As a Go learner, I want each example to compile and run on Go 1.22.1 with no external dependencies, so that I can run everything with the standard toolchain.

## Implementation Decisions

**New section and ordering**

- Create `18-context/` at the repository root as the next course section after `17-concurrency`. Keep `15-standard-library/09-context` unchanged as the quick intro; the new section's README cross-links back to it as the "you've already met this" starting point.
- State the prerequisite as `17-concurrency` (goroutines and channels), since every example relies on `select`, goroutines, and channels.

**Example sequence** (one concept per folder; folder names are the concept):

1. `01-why-context` — the motivating problem: start a goroutine with no way to stop it and show it would run past the point the caller cares. Introduces nothing from `context` yet except the pain it removes.
2. `02-background` — `context.Background()` as the empty root; show that `ctx.Done()` on it never fires and `ctx.Err()` is nil. Establishes the base of every context tree.
3. `03-with-cancel` — `context.WithCancel`: a worker watches `ctx.Done()`; the caller calls `cancel()` to stop it. Introduces the `select { case <-ctx.Done(): ... }` shape.
4. `04-with-timeout` — `context.WithTimeout`: auto-cancel after a duration. `defer cancel()` introduced here.
5. `05-with-deadline` — `context.WithDeadline`: cancel at a fixed instant; contrast with timeout (duration vs moment).
6. `06-done-and-err` — focus on `ctx.Done()` (a channel) and `ctx.Err()` returning `context.Canceled` vs `context.DeadlineExceeded`; show how to branch on the reason.
7. `07-propagation` — pass one `ctx` down a chain of functions (A calls B calls C); cancelling at the top stops work at the bottom. Shows the "first parameter, named ctx" convention.
8. `08-context-tree` — derive a child context from a parent; cancelling the parent cancels the child, but cancelling the child leaves the parent alive.
9. `09-fan-out` — start several workers sharing one context; one `cancel()` (or one worker's error triggering cancel) stops all of them. The "cancel the rest" pattern, using only the standard library (`sync.WaitGroup` + a shared `context`).
10. `10-values` — `context.WithValue` carrying a request ID through a call chain, with a **private key type** (`type ctxKey string` or an unexported struct) to avoid collisions, and a typed getter helper.
11. `11-http-server` — an `http.HandlerFunc` that reads `r.Context()` and abandons slow work when the request's context is done (client disconnect / server timeout). Runnable without a browser by driving it with `httptest` in `main` so the output is deterministic.
12. `12-http-client` — an outbound request built with `http.NewRequestWithContext` and a timeout; show the request aborting when the (locally served, deliberately slow) endpoint is slower than the timeout. Use `httptest.NewServer` so no network is required.
13. `13-graceful-shutdown` — `signal.NotifyContext` to turn Ctrl+C (SIGINT) into a cancelled context; the program drains work and exits cleanly. Drive the signal programmatically (send SIGINT to self or cancel) so the example terminates deterministically for the reader.
14. `14-rules-and-pitfalls` — a short runnable file that demonstrates and comments the rules in one place: ctx is the first parameter named `ctx`; always `defer cancel()`; don't store ctx in a struct; `WithValue` is for request-scoped data only, not optional parameters; always check `ctx.Err()` after `Done()`.

- This list is the intended scope; the exact number may shrink by merging adjacent ideas (e.g. deadline folded into timeout) if any example would otherwise fall below the "worth its own folder" bar. Splitting is preferred over any single file exceeding the ~10-minute budget.

**Example shape (applies to every file)**

- `package main` with a `main()` the learner runs via `go run ./18-context/NN-name`.
- `ctx context.Context` is always the first parameter of any helper that takes it, named `ctx`.
- Every derived context's `cancel` is released with `defer cancel()`.
- Comments are plain language, placed on the line or the line above what they explain, and explicitly call out syntax a JS/TS learner may forget (e.g. `<-ch` is a channel receive; `:=` declares-and-assigns; `select` waits on whichever channel is ready first).
- No external modules; standard library only; compiles on Go 1.22.1.

**Section README**

- Mirror the structure of `15-standard-library/README.md` and `17-concurrency/README.md`: title, one-line purpose, prerequisite, an ordered bulleted index linking each `main.go`, then one subsection per example with its run command, a fenced **Expected output** block, and a **Try** line.

**Course index wiring**

- Add `18-context` to the repository's top-level course listing wherever sections are enumerated (root `README.md` and/or `NOTES.md`), consistent with how `17-concurrency` is listed.
- Note in `MISSION.md`'s Phase 2 ("practical concurrency") that the context deep-dive realizes part of that phase, if an update there fits the existing wording. (Optional wiring, not required for the section to stand alone.)

## Testing Decisions

- **What a good "test" is here:** this repo verifies learning examples by their *documented external behavior* — the exact text an example prints when run — not by unit tests of internal functions. The README's **Expected output** block is the golden record. A good example therefore produces the same output on every run.
- **Determinism despite concurrency/timing:** design each timing example so the outcome is never a race. Separate the competing durations widely (e.g. work takes clearly longer than a timeout, or a reply is pre-loaded into a buffered channel so it always wins). Prefer pre-filled buffered channels and `time.AfterFunc`/fixed durations over bare `time.Sleep` guesses. Where ordering between goroutines would otherwise vary, serialize the printed output (collect results then print in a fixed order).
- **Modules covered:** every `18-context/NN-*/main.go` must `go build` and `go run` cleanly on Go 1.22.1 and produce output matching its README block. `11-http-server` and `12-http-client` use `net/http/httptest` so they are self-contained and deterministic with no real network. `13-graceful-shutdown` triggers its own cancellation/signal so it terminates without human input.
- **Prior art:** follow the existing golden-output style of `15-standard-library` (especially `02-time`, which deliberately uses a fixed date for stable output) and `17-concurrency` (which documents exact expected lines). Reuse the inventory-lookup framing already present in `15-standard-library/09-context`.
- **Suggested verification step for the author/agent:** run each example and diff its stdout against the README's Expected output block before considering the section done.

## Out of Scope

- Third-party libraries, including `golang.org/x/sync/errgroup` — the fan-out example uses only the standard library (`sync` + `context`). errgroup may be mentioned in a comment as "what you'd reach for in production", but not used.
- Context internals / runtime implementation (how `Done()` channels are propagated inside the runtime, `context.AfterFunc` edge cases, custom `Context` implementations beyond a brief mention).
- Generics, reflection, and performance tuning (consistent with `MISSION.md` scope).
- Production concerns: real servers, databases, tracing/observability integrations, distributed cancellation across services.
- Changing the existing `15-standard-library/09-context` example beyond leaving it as the intro (the learner currently has uncommitted edits there; this spec does not depend on or alter that work).
- Setting up an external issue tracker / triage labels (the `/to-spec` default). This spec is delivered as a local Markdown file per the repository's existing `SPEC.md` / `SPEC-advanced.md` convention.

## Further Notes

- Design patterns the section should make explicit (woven into examples and collected in `14-rules-and-pitfalls`): accept `ctx` as the first parameter named `ctx`; never store a `Context` in a struct — pass it explicitly; always `defer cancel()` even when a timeout will fire first; derive child contexts rather than creating a new `Background()` mid-chain; watch `ctx.Done()` in `select`; use `WithValue` only for request-scoped data with a private key type; check `ctx.Err()` to learn *why* work stopped.
- "When to use it" real scenarios to anchor the lessons: bounding a slow outbound API/DB call (timeout), abandoning server work when the client disconnects (HTTP handler), stopping a batch of workers together (fan-out), and shutting down cleanly on a signal (graceful shutdown).
- Suggested follow-on once the section lands: a tiny capstone that combines a timeout, propagation, and a value (e.g. an inventory lookup with a request ID that calls a slow service under a deadline), aligning with the capstone idea already noted in `SPEC-advanced.md`.
- The `/setup-matt-pocock-skills` command can be run later to configure an issue tracker; this spec could then be published there with a `ready-for-agent` label without changes.
