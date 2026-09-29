# Testing

Check a function's behavior with Go's built-in testing package.

Prerequisite: 15-standard-library (especially functions from topic 06).

Read and run in this order (about 5–10 minutes for the example and its test):

- [The Total function and runnable program](01-total/main.go)
- [Its companion test](01-total/main_test.go)

## A function and its test

Total returns the cost of a whole-number price and quantity. The test compares
its result with a known answer. Files ending in `_test.go` are used by `go test`
and are left out of `go run` and `go build`. A test function starts with `Test`
and receives a `*testing.T`, which provides methods for reporting failures.

From the repository root:

```sh
go run ./16-testing/01-total
```

Expected output:

```text
cost: 12
```

Run just this lesson's test (also typechecking its package):

```sh
go test ./16-testing/01-total -run '^TestTotal$' -v
```

Expected invariant: output names `TestTotal`, reports `PASS` for the test and
package, and the command exits successfully. Timing and cache markers can vary.

Try: change the expected value `want` from 12 to 13 and rerun the test. It should
fail and tell you the actual and expected results. Restore 12 and run it again.
This shows how a test detects a mistake without adding a testing framework.

Run all packages, including this test, from the root:

```sh
go test ./...
```
