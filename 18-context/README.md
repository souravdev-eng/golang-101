# The context package

Go deep on `context`: cancellation, timeouts, deadlines, and the signals that
tell work when and why to stop.

Prerequisite: 17-concurrency

You already met `context` briefly in
[15-standard-library/09-context](../15-standard-library/09-context/main.go).
That was the quick tour. This section is the deep dive: one idea per runnable
file, each built on `select`, goroutines, and channels from 17-concurrency.

Read and run in this order (about 5–10 minutes per example):

- [The problem context solves](01-why-context/main.go)
- [The empty root context](02-background/main.go)
- [Cancelling explicitly with WithCancel](03-with-cancel/main.go)
- [Bounding work with WithTimeout](04-with-timeout/main.go)
- [Cancelling at a fixed moment with WithDeadline](05-with-deadline/main.go)
- [Reading Done and Err](06-done-and-err/main.go)
- [Propagating one context down a chain](07-propagation/main.go)
- [The parent/child context tree](08-context-tree/main.go)
- [Fanning out to many workers](09-fan-out/main.go)
- [Carrying request-scoped values](10-values/main.go)

## The problem context solves

A goroutine with no stop signal runs every step, even after the caller stops
caring. This example has no `context` yet — it shows the pain the rest of the
section removes.

From the repository root:

```sh
go run ./18-context/01-why-context
```

Expected output:

```text
worker: step 1
worker: step 2
worker: step 3
worker: step 4
worker: step 5
main: the worker ran all 5 steps; I had no way to stop it
```

Try: Change `countUp(5, ...)` to `countUp(8, ...)` and predict how many
`worker: step` lines print before the final line.

## The empty root context

`context.Background()` is the empty root every context tree starts from. Its
`Done()` channel is nil and never fires, and its `Err()` is always nil.

From the repository root:

```sh
go run ./18-context/02-background
```

Expected output:

```text
Done channel is nil: true
Err is nil: true
```

Try: Add `fmt.Println(context.TODO().Done() == nil)` and predict the result
(`TODO` is the placeholder root when you are not sure which context to use yet).

## Cancelling explicitly with WithCancel

`context.WithCancel` returns a context plus a `cancel` function. A worker watches
`ctx.Done()` inside a `select`; the caller calls `cancel()` to stop it.

From the repository root:

```sh
go run ./18-context/03-with-cancel
```

Expected output:

```text
worker: handled restock tea
worker: cancelled, stopping
```

Try: Send a second job (`jobs <- "restock milk"` with its own `<-report`) before
`cancel()` and predict the new middle line.

## Bounding work with WithTimeout

`context.WithTimeout` cancels the context automatically after a duration. The
work here takes a full second but the timeout is 20 milliseconds, so the timeout
always wins. This is where `defer cancel()` is introduced.

From the repository root:

```sh
go run ./18-context/04-with-timeout
```

Expected output:

```text
gave up: context deadline exceeded
```

Try: Change the timeout to `2 * time.Second` and predict which `select` case
wins and what prints.

## Cancelling at a fixed moment with WithDeadline

`context.WithDeadline` cancels at a specific instant rather than after a
duration. A timeout is "20ms from now"; a deadline is "at exactly this moment".

From the repository root:

```sh
go run ./18-context/05-with-deadline
```

Expected output:

```text
missed the deadline: context deadline exceeded
```

Try: Change the deadline to `time.Now().Add(2 * time.Second)` and predict which
line prints.

## Reading Done and Err

`ctx.Done()` is a channel that closes when a context finishes; `ctx.Err()` then
says why. An explicit `cancel()` gives `context.Canceled`; a timeout or deadline
gives `context.DeadlineExceeded`.

From the repository root:

```sh
go run ./18-context/06-done-and-err
```

Expected output:

```text
manual: was cancelled by a caller
timeout: ran out of time
```

Try: Remove `cancel1()` and predict what happens when `describe(ctx1, ...)`
waits on a context that is never cancelled.

## Propagating one context down a chain

One context is created at the top and passed as the first argument down a call
chain (`stepA` → `stepB` → `stepC`). The middle functions just forward it;
cancelling once at the top is felt at the bottom, where `stepC` checks
`ctx.Done()` before starting work.

From the repository root:

```sh
go run ./18-context/07-propagation
```

Expected output:

```text
stepA: passing ctx down to stepB
stepB: passing ctx down to stepC
stepC: ctx cancelled, not starting work
```

Try: Move `cancel()` to after `stepA(ctx)` returns and predict which `stepC`
line prints instead.

## The parent/child context tree

A child context is derived from a parent, forming a tree. Cancellation flows
down only: cancelling the parent cancels the child, but cancelling the child
leaves the parent alive. This example shows both directions.

From the repository root:

```sh
go run ./18-context/08-context-tree
```

Expected output:

```text
cancel parent -> parent: done child: done
cancel child  -> parent: alive child: done
```

Try: Add a grandchild (`context.WithCancel(child1)`) and predict its state after
`cancelParent1()` runs.

## Fanning out to many workers

Several workers share one context. A single `cancel()` stops all of them; here
one worker fails and calls `cancel()`, which stops the rest. Standard library
only (`sync.WaitGroup` + a shared `context`). The results are collected and
sorted before printing, so the output is the same on every run despite the
concurrency.

From the repository root:

```sh
go run ./18-context/09-fan-out
```

Expected output:

```text
worker 1: stopped (context canceled)
worker 2: failed, cancelling the group
worker 3: stopped (context canceled)
```

Try: Change `failAt` to `1` and predict which worker now reports the failure
(the sorted output keeps the lines in worker-number order either way).

## Carrying request-scoped values

`context.WithValue` carries a request ID down a call chain without adding a
parameter to every function. The key uses a private (unexported) type so it
cannot collide with keys from other packages, and a typed getter reads it back
with the comma-ok form so a missing value is reported, not guessed.

From the repository root:

```sh
go run ./18-context/10-values
```

Expected output:

```text
handle: request id is req-42
handle: no request id in context
```

Try: Add a second key (another `ctxKey` constant) for a user name, give it its
own getter, and predict whether it collides with the request ID key.
