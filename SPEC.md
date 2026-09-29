# Go basics through short, runnable lessons

## Problem Statement

The learner wants to learn Go from the very basics and build a broad understanding of how to write Go code. The repository currently contains a Go module and a single Hello World program, with no learning sequence or supporting notes. Long, text-heavy notes would make it harder to learn by reading and changing code.

## Solution

Create an ordered collection of small Go lessons grouped into topic folders. Each example is a self-contained, runnable file focused on one concept or a closely related set of concepts. Each folder has a short README explaining the topic, the suggested reading order, and how to run its examples.

Each example should take a beginner no more than about 10 minutes to read, run, and understand. Explain the code with short, simple comments placed near the relevant lines. Use familiar examples, predictable output, and a small optional change the learner can try. Split a topic into multiple examples when it would otherwise exceed the time budget.

## User Stories

1. As a beginner learning Go, I want an ordered topic index, so that I know where to start and what to study next.
2. As a beginner learning Go, I want brief setup and run instructions, so that I can execute examples without guessing commands.
3. As a beginner learning Go, I want independently runnable example files, so that I can study one concept at a time.
4. As a beginner learning Go, I want each example to fit within about 10 minutes, so that learning feels manageable.
5. As a beginner learning Go, I want a short README in each topic folder, so that I can understand its purpose without reading long notes.
6. As a beginner learning Go, I want comments in plain language, so that I can understand unfamiliar syntax beside the code.
7. As a beginner learning Go, I want expected output, so that I can compare what I run with what the example teaches.
8. As a beginner learning Go, I want a small suggested experiment, so that I can practice by changing working code.
9. As a beginner learning Go, I want to understand packages, imports, and the main function, so that I understand the structure of a basic program.
10. As a beginner learning Go, I want examples of printing and formatted output, so that I can display values and inspect results.
11. As a beginner learning Go, I want examples of variable declarations, short declarations, and assignment, so that I can store and update values.
12. As a beginner learning Go, I want examples of constants, basic types, and zero values, so that I understand how values are represented and initialized.
13. As a beginner learning Go, I want examples of type conversion, so that I can work with values of different types.
14. As a beginner learning Go, I want examples of arithmetic, comparison, and logical operators, so that I can calculate results and express conditions.
15. As a beginner learning Go, I want examples of if, else if, and else, so that I can choose between different actions.
16. As a beginner learning Go, I want examples of combined conditions, nested conditions, and if initializer statements, so that I can recognize common conditional patterns.
17. As a beginner learning Go, I want examples of switch statements, including expressionless switches, so that I can express multiple choices clearly.
18. As a beginner learning Go, I want examples of counted, condition-only, and infinite for loops with an exit, so that I understand Go's loop forms.
19. As a beginner learning Go, I want examples of break, continue, and nested loops, so that I can control repetition.
20. As a beginner learning Go, I want examples of range over collections, so that I can visit their elements.
21. As a beginner learning Go, I want examples of functions, parameters, and return values, so that I can organize reusable behavior.
22. As a beginner learning Go, I want examples of multiple return values, variadic parameters, and anonymous functions, so that I can recognize useful function patterns.
23. As a beginner learning Go, I want examples of local scope and shadowing, so that I can understand where a variable is available and avoid confusing names.
24. As a beginner learning Go, I want examples of arrays and slices, so that I can store ordered collections and understand their differences.
25. As a beginner learning Go, I want examples of slice indexing, slicing, append, length, and capacity, so that I can use growing collections.
26. As a beginner learning Go, I want examples of maps, including lookup, update, deletion, and missing keys, so that I can work with keyed data.
27. As a beginner learning Go, I want examples of strings, bytes, and runes, so that I can understand basic text handling.
28. As a beginner learning Go, I want examples of structs and methods, so that I can group related data and behavior.
29. As a beginner learning Go, I want examples of pointers and value versus pointer changes, so that I can understand how functions affect data.
30. As a beginner learning Go, I want an introductory interface example, so that I can see how different types share behavior.
31. As a beginner learning Go, I want examples of returning and checking errors, so that I can handle expected failures explicitly.
32. As a beginner learning Go, I want an example of defer, so that I can understand when delayed cleanup runs.
33. As a beginner learning Go, I want a small standard-library example, so that I can see how built-in packages help solve everyday tasks.
34. As a beginner learning Go, I want a simple test example, so that I can understand how Go checks function behavior.
35. As a beginner learning Go, I want introductory goroutine and channel examples after the foundations, so that I can begin understanding concurrent Go programs.

## Implementation Decisions

- Retain the existing module and keep examples compatible with its declared Go 1.22.1 version. Use the standard library without third-party dependencies.
- Add an ordered learning index and numbered topic folders. Teach program structure, values, operators, conditions, loops, and functions before collections and the later topics listed in the user stories.
- Treat the broader topics as introductory coverage, not exhaustive reference material. “All” conditions and loops means the common Go forms and control-flow patterns, not every possible combination.
- Keep each runnable example in its own child directory within its topic folder. Each directory contains one main package with one entry point, allowing both individual execution and module-wide tooling without duplicate main declarations.
- Prefer one example file per lesson. A lesson that introduces tests may also include a small companion test file.
- Give each topic README a brief explanation, ordered example list, exact run commands, and any prerequisite topic. Keep explanations short enough to review quickly.
- Document commands relative to the repository root. Support running individual examples with go run and running tests with go test.
- Explain syntax and purpose with short comments near the relevant code. Avoid commenting every obvious line, dense terminology, or copying long explanations into every file.
- Choose descriptive identifiers and familiar scenarios. Introduce only concepts already covered or explicitly explained in the current lesson.
- Include expected output beside each example or in its topic README, plus one small optional experiment. For nondeterministic output, explain the invariant instead of promising an exact order.
- Keep the roughly 10-minute limit as an editorial acceptance criterion for every example file. Split larger examples and avoid turning the folder READMEs into textbooks.
- Keep the existing Hello World program as the starting example or incorporate it into the opening lesson without losing its behavior.

## Testing Decisions

- Proposed primary verification boundary: execute each lesson as a learner would, through its command-line entry point. Verify successful execution and the observable output or documented invariant.
- Check example output against the expected result. Use deterministic inputs; sort map keys when output order matters. Concurrency examples should finish predictably and demonstrate coordination without relying on sleep timing.
- Run module-wide compilation and tests to catch package conflicts and broken examples. Format Go source using gofmt.
- Check the README commands from the repository root and confirm each example works independently of running any previous lesson.
- Review comments, topic order, and lesson size manually. Automated checks cannot establish that a beginner will understand a file in 10 minutes.
- Add a small unit-test demonstration as learning material. Avoid a test framework or a separate test suite that mirrors every teaching line.
- There is no existing test suite or testing seam in this repository. No additional internal abstraction is needed for verification; runnable lesson entry points provide the highest practical boundary.
- Confirmation of the proposed verification boundary is pending, as required by the to-spec workflow.

## Out of Scope

- A text-heavy Go textbook, exhaustive language reference, or certification curriculum.
- A large application combining all topics, web frameworks, production services, databases, or deployment workflows.
- Advanced concurrency patterns, generics, reflection, unsafe operations, performance tuning, and deep runtime internals in the initial curriculum.
- External services, interactive input requirements, third-party dependencies, and environment-dependent examples.
- Implementing the lessons as part of this specification-writing task.

## Further Notes

- The repository currently has no Git metadata, configured remote, issue-tracker configuration, domain glossary, ADRs, or existing tests.
- The topic breadth beyond the learner's explicit examples is a proposed curriculum derived from the request for a broad grounding in Go basics.
- This is a local draft. Publication to the project issue tracker with the ready-for-agent label is pending tracker setup and confirmation of the proposed testing boundary.
- Run /setup-matt-pocock-skills to supply the issue tracker and triage label vocabulary required by the to-spec skill.
