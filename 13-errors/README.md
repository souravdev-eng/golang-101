# Errors

Return an error for an expected failure and check it at the call site.

Prerequisite: 12-interfaces

Read and run in this order (about 5–10 minutes per example):

- [Returning and checking errors](01-return-check/main.go)
- [Handling a library error](02-parse/main.go)

## Returning and checking errors

The usual Go pattern is a useful value followed by an error.

From the repository root:

```sh
go run ./13-errors/01-return-check
```

Expected output:

```text
each: 4
cannot share: people must be positive
```

Try: Try a negative number of people in the second call.

## Handling a library error

Standard-library functions use the same value-and-error pattern.

From the repository root:

```sh
go run ./13-errors/02-parse
```

Expected output:

```text
quantity: 12
invalid quantity: "many"
```

Try: Try "3.5" and observe that Atoi expects a whole number.
