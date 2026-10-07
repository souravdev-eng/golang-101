package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// The whole section in one file: the five rules that keep context code correct,
// each demonstrated beside a comment explaining it. Run it, then read it.

// RULE 3: Don't store a Context in a struct. A context describes one call's
// lifetime; a struct outlives the call. Keep configuration in the struct and
// pass ctx to each method instead. This type holds a name — and no ctx field.
type fetcher struct {
	name string
}

// RULE 1: ctx is the FIRST parameter, and it is named ctx. Every function in the
// standard library follows this, so callers always know where the context goes.
// RULE 5: return ctx.Err() so the caller can learn WHY the work stopped.
func (f fetcher) fetch(ctx context.Context) error {
	select {
	case <-time.After(50 * time.Millisecond): // The real work finishing.
		return nil
	case <-ctx.Done(): // Cancelled or timed out first.
		return ctx.Err()
	}
}

// RULE 4: WithValue is for request-scoped data that rides along with a request —
// a user, a request ID, a trace. It is NOT a back door for optional parameters
// like page size or retry count; those belong in the function signature.
type ctxKey int

const userKey ctxKey = iota

func main() {
	// RULE 2: always defer cancel(), even here where the timeout will fire on its
	// own. cancel() releases the timer immediately; skipping it leaks resources
	// until the deadline passes. "A timeout fires first" is not an excuse to omit
	// it — the vet tool flags a missing cancel.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	f := fetcher{name: "inventory"} // ctx is NOT stored on the struct (rule 3).
	err := f.fetch(ctx)             // ctx passed as the first argument (rule 1).

	// RULE 5: check ctx-derived errors to learn why work stopped. errors.Is sees
	// through wrapping, so it works even when a deeper layer wrapped the cause.
	switch {
	case err == nil:
		fmt.Println("fetch: finished in time")
	case errors.Is(err, context.DeadlineExceeded):
		fmt.Println("fetch: stopped because it ran out of time")
	case errors.Is(err, context.Canceled):
		fmt.Println("fetch: stopped because a caller cancelled it")
	}

	// RULE 4: attach request-scoped data, then read it back with the comma-ok
	// form so a missing or wrong-typed value is reported, not guessed.
	ctx = context.WithValue(ctx, userKey, "alice")
	if user, ok := ctx.Value(userKey).(string); ok {
		fmt.Println("request user:", user)
	}
}
