# Arrays and slices

Store ordered values and learn how slices refer to arrays.

Prerequisite: 06-functions

Read and run in this order (about 5–10 minutes per example):

- [Fixed-size arrays](01-arrays/main.go)
- [Indexing, slicing, and sharing](02-slices/main.go)
- [Append, length, and capacity](03-append-capacity/main.go)

## Fixed-size arrays

An array has a fixed length that is part of its type.

From the repository root:

```sh
go run ./07-collections/01-arrays
```

Expected output:

```text
scores: [10 20 30]
first: 10 length: 3
original: [10 20 30]
copy: [99 20 30]
zero array: [0 0]
```

Try: Change the second element of copyOfScores and inspect the original.

## Indexing, slicing, and sharing

A slice describes part of an underlying array.

From the repository root:

```sh
go run ./07-collections/02-slices
```

Expected output:

```text
first: apple length: 4
middle: [banana pear]
fruits: [apple orange pear plum]
first two: [apple orange]
last two: [pear plum]
```

Try: Change the slice bounds to fruits[2:4] and predict which fruit changes.

## Append, length, and capacity

Length counts visible elements; capacity counts room in the underlying array.

From the repository root:

```sh
go run ./07-collections/03-append-capacity
```

Expected output:

```text
start: 0 3
scores: [10 20]
length/capacity: 2 3
full: [10 20 30]
grown: [10 20 30 40] length: 4
```

Try: Append 50 too and predict the final length.
