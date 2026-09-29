# Defer

Schedule a call to run when the current function returns.

Prerequisite: 13-errors

Read and run in this order (about 5–10 minutes per example):

- [Deferred call order](01-order/main.go)

## Deferred call order

Defer is often used to pair acquiring a resource with its cleanup.

From the repository root:

```sh
go run ./14-defer/01-order
```

Expected output:

```text
start packing
current label: coffee
packing done
label: tea
put lid on
close box
back in main
```

Try: Swap the first two defer lines and predict the cleanup order.
