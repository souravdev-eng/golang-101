# Programs and output

Start with the shape of a Go program, then display values.

Prerequisite: None; follow the setup in the root README.

Read and run in this order (about 5–10 minutes per example):

- [Hello World](01-hello/main.go)
- [Formatted output](02-printing/main.go)

## Hello World

Read `package`, `import`, and `main` in that order.

From the repository root:

```sh
go run ./01-programs/01-hello
```

Expected output:

```text
Hello, World!
```

Try: Change the greeting inside the quotes and run again.

## Formatted output

Compare Print, Println, and Printf.

From the repository root:

```sh
go run ./01-programs/02-printing
```

Expected output:

```text
Shopping: apples 3
apples: 3 items, 1.50 each
```

Try: Change %.2f to %.1f to show one decimal place.
