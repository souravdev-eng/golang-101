# Structs and methods

Group related fields and attach behavior to a type.

Prerequisite: 09-text

Read and run in this order (about 5–10 minutes per example):

- [Struct fields](01-fields/main.go)
- [Value receiver methods](02-methods/main.go)

## Struct fields

A struct groups named values, which can have different types.

From the repository root:

```sh
go run ./10-structs/01-fields
```

Expected output:

```text
Go Basics 80
updated pages: 96
zero book: "" 0
```

Try: Add an Author string field and print a book with an author.

## Value receiver methods

A method is a function attached to a type through a receiver.

From the repository root:

```sh
go run ./10-structs/02-methods
```

Expected output:

```text
area: 12
```

Try: Change the width to 5 and predict the area.
