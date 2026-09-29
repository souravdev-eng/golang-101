# Functions and scope

Give a piece of behavior a name and pass data to it.

Prerequisite: 05-loops

Read and run in this order (about 5–10 minutes per example):

- [Parameters and return values](01-parameters-returns/main.go)
- [Multiple return values](02-multiple-returns/main.go)
- [Variadic and anonymous functions](03-variadic-anonymous/main.go)
- [Scope and shadowing](04-scope/main.go)

## Parameters and return values

Call functions with inputs and use their results.

From the repository root:

```sh
go run ./06-functions/01-parameters-returns
```

Expected output:

```text
Hello, Mira
cost: 12
```

Try: Call greet with your name and total with a different quantity.

## Multiple return values

Receive more than one result from a function.

From the repository root:

```sh
go run ./06-functions/02-multiple-returns
```

Expected output:

```text
each: 2 left: 1
```

Try: Split 10 apples among 4 people.

## Variadic and anonymous functions

Accept any number of inputs and store a function in a variable.

From the repository root:

```sh
go run ./06-functions/03-variadic-anonymous
```

Expected output:

```text
sum: 9
empty sum: 0
double: 10
```

Try: Pass an extra value to sum and predict the result.

## Scope and shadowing

A variable belongs to the block where it is declared.

From the repository root:

```sh
go run ./06-functions/04-scope
```

Expected output:

```text
outside
inside
outside
updated
```

Try: Change the first inner := to = and compare the output.
