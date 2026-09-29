# Goroutines and channels

Start concurrent work and coordinate it explicitly.

Prerequisite: 16-testing

Read and run in this order (about 5–10 minutes per example):

- [Start and wait for a goroutine](01-goroutine/main.go)
- [Send, receive, and close](02-channels/main.go)

## Start and wait for a goroutine

A goroutine runs a function concurrently; main must wait for work it needs.

From the repository root:

```sh
go run ./17-concurrency/01-goroutine
```

Expected output:

```text
worker: packed lunch
main: ready to leave
```

Try: Change the worker message; check that it still prints before the main message.

## Send, receive, and close

A channel carries values between goroutines and coordinates their progress.

From the repository root:

```sh
go run ./17-concurrency/02-channels
```

Expected output:

```text
received: sandwich
received: fruit
all lunches received
```

Try: Send "water" before close and predict the extra output line.
