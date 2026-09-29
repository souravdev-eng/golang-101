# Loops

Repeat work using Go's for loop and its common forms.

Prerequisite: 04-conditions

Read and run in this order (about 5–10 minutes per example):

- [Three forms of for](01-for-forms/main.go)
- [Break and continue](02-break-continue/main.go)
- [Nested loops](03-nested/main.go)
- [Range over a collection](04-range/main.go)

## Three forms of for

Read a counted loop, a condition-only loop, and a loop with an explicit exit.

From the repository root:

```sh
go run ./05-loops/01-for-forms
```

Expected output:

```text
step: 1
step: 2
step: 3
remaining: 2
remaining: 1
attempt: 1
attempt: 2
```

Try: Make the counted loop stop at 4.

## Break and continue

Skip one turn or stop the loop completely.

From the repository root:

```sh
go run ./05-loops/02-break-continue
```

Expected output:

```text
number: 1
number: 3
number: 4
```

Try: Change the break condition to number == 6.

## Nested loops

The inner loop runs fully for each turn of the outer loop.

From the repository root:

```sh
go run ./05-loops/03-nested
```

Expected output:

```text
row 1 seat 1
row 1 seat 2
row 1 seat 3
row 2 seat 1
row 2 seat 2
row 2 seat 3
```

Try: Use two seats per row and predict the four lines.

## Range over a collection

A slice is an ordered collection; topic 07 will explore it further.

From the repository root:

```sh
go run ./05-loops/04-range
```

Expected output:

```text
0 apple
1 banana
2 pear
packed: apple
packed: banana
packed: pear
```

Try: Add orange to the slice and run again.
