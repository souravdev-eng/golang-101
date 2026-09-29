# Pointers and changes

Use a pointer when code needs to change the caller's value.

Prerequisite: 10-structs

Read and run in this order (about 5–10 minutes per example):

- [Values versus pointers](01-value-pointer/main.go)
- [Pointer receiver methods](02-pointer-methods/main.go)

## Values versus pointers

Passing a value copies it; passing a pointer gives access to the original.

From the repository root:

```sh
go run ./11-pointers/01-value-pointer
```

Expected output:

```text
after value: 3
after pointer: 4
missing: true
```

Try: Call addThroughPointer twice and predict the count.

## Pointer receiver methods

Choose a pointer receiver when a method must update the struct.

From the repository root:

```sh
go run ./11-pointers/02-pointer-methods
```

Expected output:

```text
after value method: 3
after pointer method: 4
```

Try: Call Increment one more time.
