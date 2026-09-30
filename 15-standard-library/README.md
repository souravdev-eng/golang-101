# Everyday standard library

Use small examples to meet common standard-library packages for text, dates,
collections, paths, command-line options, I/O, JSON, HTTP, and cancellation.

Prerequisite: 14-defer

Read and run in this order (about 5–10 minutes per example):

- [Cleaning a shopping list](01-strings/main.go)
- [Parsing and formatting a date](02-time/main.go)
- [Sorting a copy of a slice](03-slices/main.go)
- [Working with a file path](04-filepath/main.go)
- [Reading command-line flags](05-flag/main.go)
- [Copying to standard output](06-os-io/main.go)
- [Reading and writing JSON](07-json/main.go)
- [Handling an HTTP request](08-http/main.go)
- [Cancelling a context](09-context/main.go)

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

## Parsing and formatting a date

Use a fixed date so the result is the same whenever you run the example. Go's
date layouts use the reference date `Mon Jan 2 15:04:05 MST 2006` instead of
tokens like `YYYY` and `MM`.

From the repository root:

```sh
go run ./15-standard-library/02-time
```

Expected output:

```text
date: Mon, 01 Jan 2024
one week later: 2024-01-08
```

Try: Change the input to `2024-01-02` and predict both output lines.

## Sorting a copy of a slice

`slices.Sort` changes its input. Clone the slice first to keep the original order.

From the repository root:

```sh
go run ./15-standard-library/03-slices
```

Expected output:

```text
original: [12 5 9]
sorted: [5 9 12]
has 9: true
```

Try: Add `3` to `prices` and predict the new sorted slice.

## Working with a file path

`path/filepath` joins and inspects paths using the current operating system's
separator. `ToSlash` makes the displayed result the same on every system.
These functions work with path text; this example does not create a file.

From the repository root:

```sh
go run ./15-standard-library/04-filepath
```

Expected output:

```text
path: reports/2024/sales.csv
folder: reports/2024
file: sales.csv
extension: .csv
```

Try: Change `sales.csv` to `summary.txt` and predict the file and extension.

## Reading command-line flags

The `flag` package gives a command defaults and lets the person running it
override them. `String` and `Int` return pointers, so read their values with `*`.

From the repository root:

```sh
go run ./15-standard-library/05-flag
```

Expected output:

```text
Hello, Gopher
Hello, Gopher
```

Try: Run `go run ./15-standard-library/05-flag -name Go -times 1` and predict
how many lines it prints.

## Copying to standard output

`io.Copy` moves data from a reader to a writer. Here `strings.NewReader` supplies
the text and `os.Stdout` is the writer. The later I/O topic covers files.

From the repository root:

```sh
go run ./15-standard-library/06-os-io
```

Expected output:

```text
receipt: tea, milk
```

Try: Add `, bread` to the message and predict the output.

## Reading and writing JSON

`encoding/json` converts between JSON text and a Go struct. The field tags set
the JSON key names. Parsing and encoding can fail, so check both errors.

From the repository root:

```sh
go run ./15-standard-library/07-json
```

Expected output:

```text
product: tea 12
json: {"name":"tea","price":12}
```

Try: Change the JSON price to `15` and predict both output lines.

## Handling an HTTP request

`net/http` defines a handler. `net/http/httptest` sends it a local, in-memory
request, so this example needs no server or internet connection.

From the repository root:

```sh
go run ./15-standard-library/08-http
```

Expected output:

```text
status: 200
body: Hello, Go
```

If Go 1.22.1 on macOS reports `missing LC_UUID` when running this example,
use `go run -ldflags=-linkmode=external ./15-standard-library/08-http`.

Try: Change `name=Go` to `name=Gopher` and predict the body.

## Cancelling a context

`context` carries a cancellation signal. The full concurrency use comes later;
this example only shows that cancellation closes `Done` and sets `Err`.

From the repository root:

```sh
go run ./15-standard-library/09-context
```

Expected output:

```text
before cancel: true
after cancel: context canceled
```

Try: Predict what would happen if `cancel()` moved below `<-ctx.Done()`.
