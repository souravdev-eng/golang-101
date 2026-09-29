# Conditions and switches

Choose which block of code should run.

Prerequisite: 03-operators

Read and run in this order (about 5–10 minutes per example):

- [If, else if, and else](01-if-else/main.go)
- [Combined, nested, and initialized conditions](02-condition-patterns/main.go)
- [Switch forms](03-switch/main.go)

## If, else if, and else

The first matching branch runs; later branches are skipped.

From the repository root:

```sh
go run ./04-conditions/01-if-else
```

Expected output:

```text
wear a light jacket
```

Try: Try temperatures 10 and 30 to reach the other branches.

## Combined, nested, and initialized conditions

Use a combined check, a check inside a check, and a short-lived variable.

From the repository root:

```sh
go run ./04-conditions/02-condition-patterns
```

Expected output:

```text
adult entry allowed
young visitor discount
seats left: 7
```

Try: Set hasTicket to false and see which messages disappear.

## Switch forms

Match a value, then use a switch that checks conditions.

From the repository root:

```sh
go run ./04-conditions/03-switch
```

Expected output:

```text
weekend
good progress
```

Try: Change day to Monday and score to 95.
