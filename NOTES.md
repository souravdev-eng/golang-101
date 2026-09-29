# Teaching notes

- The learner wants Go familiarity for upcoming projects and knows JS/TS (2026-09-28).
- Use JS/TS comparisons as bridges, explaining where the Go model differs.
- Follow SPEC.md's short read–predict–run–change approach using existing topic examples.
- First lesson prepared: lessons/0001-go-program-entry.html. Completion and understanding have not yet been demonstrated.

## Format preferences (2026-09-28)

- Reading code cold doesn't help the learner. They want a short **context note per topic folder** (7–10 min) that says what the folder is for and what to look for in each example, *before* opening the code. These live in `notes/NN-topic.html`, with `notes/00-course-map.html` as home.
- Layout must be the hellointerview-style three-column layout: course menu on the left, "On this page" outline and reading progress on the right. It comes from `assets/nav.js` + `assets/course-map.js` + `assets/course.css`, ported from the learner's observability course. Every new page must include those three and be registered in `course-map.js`.
- Brand: Astria Digital tokens (navy #051367, Lexend/Inter/IBM Plex Mono) are already in course.css.

## Next

- Ask the learner which topic notes they read and have them answer the "Check yourself" questions from memory; then continue lessons from their weakest topic.
