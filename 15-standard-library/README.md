# Everyday standard library

Use the strings package to clean and split simple text.

Prerequisite: 14-defer

Read and run in this order (about 5–10 minutes per example):

- [Cleaning a shopping list](01-strings/main.go)

## Cleaning a shopping list

Combine a few small string helpers without external dependencies.

From the repository root:

```sh
go run ./15-standard-library/01-strings
```

Expected output:

```text
items: [tea milk bread]
has milk: true
display: tea | milk | bread
```

Try: Add Eggs to input and predict the displayed list.
