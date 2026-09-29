# Maps

Associate keys with values and check whether a key exists.

Prerequisite: 07-collections

Read and run in this order (about 5–10 minutes per example):

- [Lookup, update, and deletion](01-lookup-update/main.go)
- [Map creation and ordered output](02-iteration/main.go)

## Lookup, update, and deletion

A map stores values by key rather than by position.

From the repository root:

```sh
go run ./08-maps/01-lookup-update
```

Expected output:

```text
apples: 3
pears: 0 exists: false
bananas: 0 exists: true
keys: 3
```

Try: Delete bananas and compare the lookup result.

## Map creation and ordered output

Map iteration has no guaranteed order; sort keys when displaying them.

From the repository root:

```sh
go run ./08-maps/02-iteration
```

Expected output:

```text
apples 3
pears 2
```

Try: Add bananas with a count of 4 and predict where it appears.
