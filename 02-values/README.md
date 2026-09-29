# Variables, constants, and types

Store values, update them, and see how Go represents them.

Prerequisite: 01-programs

Read and run in this order (about 5–10 minutes per example):

- [Declarations and assignment](01-variables/main.go)
- [Constants](02-constants/main.go)
- [Basic types and zero values](03-types-zero/main.go)
- [Type conversion](04-conversion/main.go)

## Declarations and assignment

Compare an explicit type with an inferred type.

From the repository root:

```sh
go run ./02-values/01-variables
```

Expected output:

```text
Corner shop
apples: 5 baskets: 3
```

Try: Change the starting number of baskets and predict the final count.

## Constants

Use names for values that stay fixed.

From the repository root:

```sh
go run ./02-values/02-constants
```

Expected output:

```text
days: 14
```

Try: Change weeks to 3 and predict the number of days.

## Basic types and zero values

An uninitialized variable still has a usable default value.

From the repository root:

```sh
go run ./02-values/03-types-zero
```

Expected output:

```text
zero: 0 0.0 false ""
types: int float64 bool string
2 3.5 true tea
```

Try: Give ready a true value before printing the first line.

## Type conversion

Convert numbers explicitly when their types differ.

From the repository root:

```sh
go run ./02-values/04-conversion
```

Expected output:

```text
whole share: 3
decimal share: 3.5
back to int: 3
```

Try: Change total to 9 and predict both division results.
