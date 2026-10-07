# 02: Propagation, the context tree, fan-out, and request-scoped values

**What to build:** A learner who has met the core tools can now run the four "patterns" examples and understand how context flows through a real program: down a call chain, across a parent/child tree, out to many workers at once, and as a carrier for request-scoped data. Each example is runnable from the repo root and prints exactly the output documented in its README section.

Delivers examples `07`–`10`, each appended to the existing section README (index + per-example subsection):

- `07-propagation` — one `ctx` passed down a chain (A → B → C); cancelling at the top stops work at the bottom. Reinforces the "first parameter, named `ctx`" convention.
- `08-context-tree` — derive a child from a parent: cancelling the parent cancels the child; cancelling the child leaves the parent alive.
- `09-fan-out` — several workers share one context; a single `cancel()` (or one worker's error triggering cancel) stops all of them. Standard library only (`sync.WaitGroup` + shared `context`); `errgroup` may be named in a comment but not used.
- `10-values` — `context.WithValue` carries a request ID through a call chain, using a **private key type** (unexported) to avoid collisions, plus a typed getter helper.

**Blocked by:** 01 (needs the section scaffold, README structure, and the core-tool concepts).

**Status:** done

- [x] `07`–`10` each exist as their own runnable folder, listed in the section README index with a run command, Expected output block, and a Try line.
- [x] `09-fan-out` uses only the standard library; its printed output is serialized to a fixed order so it is deterministic despite concurrency.
- [x] `10-values` defines value keys with a private key type (not a bare string/`int`) and reads them through a typed helper.
- [x] All examples keep `ctx` as the first parameter named `ctx`, use `defer cancel()` for derived contexts, and carry plain-language comments beside the code.
- [x] Each example builds on Go 1.22.1 with no external modules and matches its README Expected output exactly.
