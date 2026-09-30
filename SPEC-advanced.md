# Phase 2: Go beyond the basics, in the same short-lesson format

## Problem Statement

Topics 01–17 give the learner a grounding in Go syntax and the common patterns. Real Go projects also rely on composition, richer error handling, generics, multi-package modules, deeper testing, practical concurrency, and the standard-library pieces used for I/O, JSON, and HTTP. The initial course deliberately left these out. The learner does not yet know what kind of Go project comes next, so this phase has to prepare for general project work rather than one domain.

## Solution

Extend the course with topics 18–30 using the same format as phase 1: numbered topic folders, one runnable `main.go` per example, a topic README with expected output and a "Try:" change, and a context note in `notes/NN-topic.html` registered in `assets/course-map.js`. Each example still fits in about 10 minutes. The phase ends with a small capstone that combines the earlier topics.

Phase 2 starts only after topics 11–17 are complete, because it builds directly on pointers, interfaces, errors, testing, and goroutines.

## Curriculum

### A. Types in depth

- **18 Embedding and composition**: embedded structs, promoted fields and methods, composition in place of inheritance (JS/TS `extends` as the bridge).
- **19 Interfaces in depth**: implicit satisfaction, type assertions, type switches, `any`, `fmt.Stringer`, small interfaces such as `io.Reader` and `io.Writer`.
- **20 Errors in depth**: wrapping with `%w`, `errors.Is` and `errors.As`, sentinel errors, custom error types, and `panic`/`recover` and when to avoid them.
- **21 Generics**: type parameters, constraints, generic functions and types, and when a plain interface or concrete type is clearer.

### B. Project shape

- **22 Packages and modules**: several packages in one module, `internal/`, designing an exported API, package documentation.
- **23 Testing in depth**: table-driven tests, subtests, `t.Helper`, `Example` functions checked against their output, benchmarks, and a short fuzz test.

### C. Concurrency in practice

- **24 Synchronisation**: `sync.WaitGroup`, `sync.Mutex`, and finding data races with `-race`.
- **25 Select and context**: `select`, timeouts, and cancellation with `context`.
- **26 Concurrency patterns**: worker pool, pipeline, fan-out/fan-in, and stopping goroutines cleanly.

### D. Everyday standard library

- **27 Files and I/O**: `io`, `bufio`, `os`, and `io/fs`, using `testing/fstest` or embedded files to avoid depending on the environment.
- **28 JSON**: `encoding/json`, struct tags, and decoding into typed structs.
- **29 HTTP**: a small `net/http` handler, tested with `net/http/httptest` so no real network is needed.
- **30 Capstone**: a small, tested program that combines packages, errors, JSON, concurrency, and HTTP or I/O.

## Implementation Decisions

- Keep the existing format, file layout, and verification commands from SPEC.md, plus the navigation and brand requirements in NOTES.md.
- Stay on Go 1.22.1 and the standard library by default. Upgrading Go (for example to use range-over-func iterators) is an optional later decision.
- Use JS/TS comparisons in every context note, especially for composition, errors, `async`/`await` compared with goroutines, and `fetch`/Express compared with `net/http`.
- The main learning loop stays read–predict–run–change. Each topic may add one optional exercise: a skipped test the learner can enable and make pass. Skipping by default keeps `go test ./...` green.
- Examples must stay deterministic. Concurrency examples coordinate explicitly and do not rely on sleeps; HTTP uses `httptest`; file examples use in-memory or embedded file systems.
- Topics 22 and 30 may need several files or packages per example. The directory form of `go run` covers them.

## Testing Decisions

- Same verification boundary as phase 1: run each example through its entry point and compare its output with the README.
- `go fmt`, `go build`, `go vet`, and `go test ./...` must pass from the repository root, with optional exercises skipped.
- Topics 24–26 must also pass `go test -race ./...` and `go run -race`.
- Lesson size and clarity are reviewed by hand.

## Out of Scope

- Third-party frameworks, databases, deployment, and external services.
- Reflection, `unsafe`, runtime internals, and performance tuning beyond an introductory benchmark.
- Building the lessons as part of this planning step.

## Open Decisions

- **Target project type** (web API, command-line tool, infrastructure tooling): unknown. When it is known, deepen part D towards it; for example, add routing and middleware for an API, or flags and subcommands for a command-line tool.
- **Go version upgrade**: undecided; revisit before topic 21.
- **Optional exercises**: planned as optional; confirm when phase 2 begins.
