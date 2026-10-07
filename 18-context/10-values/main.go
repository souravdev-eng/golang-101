package main

import (
	"context"
	"fmt"
)

// context.WithValue carries request-scoped data (a request ID, a user, a trace)
// down a call chain without adding a parameter to every function. The danger is
// key collisions: if two packages both use the string "requestID" as a key, one
// silently overwrites the other. The fix is a PRIVATE key type.

// ctxKey is unexported, so no other package can create a value of this type.
// That makes our keys unique by construction — nobody else can collide with us.
type ctxKey int

// requestIDKey is the one key value this package uses. Giving it a named
// constant (instead of a bare literal) keeps the set and get sides in sync.
const requestIDKey ctxKey = iota

// withRequestID returns a child context carrying the request ID. Wrapping
// WithValue in a helper keeps the key private to this file.
func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

// requestID is the typed getter. ctx.Value returns an any (interface{}), so we
// type-assert it back to string. The comma-ok form reports whether the key was
// present and the type matched, so callers never get a surprise panic.
func requestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey).(string)
	return id, ok
}

// handle is deep in the call chain. It never received the request ID as an
// argument; it reads it from the context instead.
func handle(ctx context.Context) {
	if id, ok := requestID(ctx); ok {
		fmt.Println("handle: request id is", id)
	} else {
		fmt.Println("handle: no request id in context")
	}
}

func main() {
	// Attach the request ID once at the top, then pass ctx down.
	ctx := withRequestID(context.Background(), "req-42")
	handle(ctx)

	// A plain Background has no request ID, so the typed getter reports absent
	// rather than returning a misleading empty string.
	handle(context.Background())
}
