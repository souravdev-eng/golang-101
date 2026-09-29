# Strings, bytes, and runes

Understand the difference between encoded bytes and Unicode code points.

Prerequisite: 08-maps

Read and run in this order (about 5–10 minutes per example):

- [Bytes and runes](01-bytes-runes/main.go)
- [Editing copies of text](02-text-copies/main.go)

## Bytes and runes

Go strings contain bytes; ranging over a string decodes Unicode code points.

From the repository root:

```sh
go run ./09-text/01-bytes-runes
```

Expected output:

```text
bytes: 5
first byte: 99
runes: 4
byte 0: c
byte 1: a
byte 2: f
byte 3: é
```

Try: Try word := "éa" and predict the byte offset of a.

## Editing copies of text

Strings cannot be changed in place; convert to a slice to edit a copy.

From the repository root:

```sh
go run ./09-text/02-text-copies
```

Expected output:

```text
original: hello
edited: Hello
rune edit: cafs
joined: Hello Go
```

Try: Replace the final rune with é again and compare the result.
