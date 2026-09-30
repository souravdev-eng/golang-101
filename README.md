# Go basics in small steps

Learn Go by reading, running, and changing short programs. Each lesson focuses on
one concept or a small related group and should take about 5–10 minutes. No input,
external services, or extra dependencies are needed.

## Setup and first run

Install Go 1.22.1 or later from [go.dev](https://go.dev/dl/). This module declares
Go 1.22.1; all lessons use features available in that version.

Open a terminal in this repository's root directory (the directory containing
`go.mod`), then run:

```sh
go version
go run .
```

The original starter program prints:

```text
Hello, World!
```

`go.mod` names the module and declares its Go version. A package groups Go source
files in a directory. Each lesson below has its own directory and `main` entry
point, so it runs independently. You do not need to finish earlier programs to
run a later one; prerequisites describe what to understand first.

## Learning order

Before opening a topic's code, read its short context note: open
[notes/00-course-map.html](notes/00-course-map.html) in a browser.

Follow each topic README's example list from top to bottom. Read the comments,
predict the output, run the exact command, and try the small suggested change.
All commands in this repository are relative to the repository root.

| Step | Topic | What you will practice |
| --- | --- | --- |
| 01 | [Programs and output](01-programs/README.md) | Packages, imports, main, printing |
| 02 | [Variables, constants, and types](02-values/README.md) | Declarations, assignment, zero values, conversion |
| 03 | [Operators](03-operators/README.md) | Arithmetic, comparisons, boolean logic |
| 04 | [Conditions and switches](04-conditions/README.md) | If branches, nested checks, switch forms |
| 05 | [Loops](05-loops/README.md) | For forms, break, continue, nested loops, range |
| 06 | [Functions and scope](06-functions/README.md) | Parameters, returns, variadic and anonymous functions, shadowing |
| 07 | [Arrays and slices](07-collections/README.md) | Fixed arrays, indexing, shared slices, append, capacity |
| 08 | [Maps](08-maps/README.md) | Lookup, update, delete, missing keys, sorted output |
| 09 | [Strings, bytes, and runes](09-text/README.md) | UTF-8, byte indexes, Unicode code points, text copies |
| 10 | [Structs and methods](10-structs/README.md) | Fields and value receivers |
| 11 | [Pointers and changes](11-pointers/README.md) | Addresses, copies, pointer receivers |
| 12 | [Interfaces](12-interfaces/README.md) | Different types providing the same behavior |
| 13 | [Errors](13-errors/README.md) | Returning failures, checking errors, parsing text |
| 14 | [Defer](14-defer/README.md) | Cleanup order and argument evaluation |
| 15 | [Everyday standard library](15-standard-library/README.md) | Text, dates, slices, paths, flags, I/O, JSON, HTTP, and context |
| 16 | [Testing](16-testing/README.md) | A small function and its companion test |
| 17 | [Goroutines and channels](17-concurrency/README.md) | Concurrent work, waiting, sending and receiving |

For example, run just one lesson:

```sh
go run ./02-values/01-variables
```

Compare it with the expected output in that topic's README. You can also run a
single example file, such as `go run ./02-values/01-variables/main.go`. Use the
directory form for lessons with multiple source files. Save your edits before
running again. If Go reports a filename and line number, start checking there.

## Check your changes

From the repository root:

```sh
go fmt ./...
go build ./...
go vet ./...
go test ./...
```

`go fmt` formats source; `go build` checks types and compiles packages; `go vet`
checks for common mistakes. `go test` runs the teaching test and also compiles
all lesson packages. Most packages report `[no test files]`; that is expected.
The testing topic shows how to run its test on its own.

Each example has a separate main package, so module-wide commands do not create
conflicting main functions. The concurrency lessons coordinate completion and
use no sleeps; their documented output order is predictable.
