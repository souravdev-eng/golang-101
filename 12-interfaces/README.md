# Interfaces

Describe behavior that different types can provide.

Prerequisite: 11-pointers

Read and run in this order (about 5–10 minutes per example):

- [Implicit interface satisfaction](01-shared-behavior/main.go)

## Implicit interface satisfaction

A type satisfies an interface by having its required methods.

From the repository root:

```sh
go run ./12-interfaces/01-shared-behavior
```

Expected output:

```text
Hello, Mira
Beep, hello
```

Try: Change Robot.Greeting to return a different message.
