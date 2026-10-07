# Teaching notes

- The learner wants Go familiarity for upcoming projects and knows JS/TS (2026-09-28).
- Use JS/TS comparisons as bridges, explaining where the Go model differs.
- Follow SPEC.md's short read–predict–run–change approach using existing topic examples.
- First lesson prepared: lessons/0001-go-program-entry.html. Completion and understanding have not yet been demonstrated.

## Format preferences (2026-09-28)

- Reading code cold doesn't help the learner. They want a short **context note per topic folder** (7–10 min) that says what the folder is for and what to look for in each example, *before* opening the code. These live in `notes/NN-topic.html`, with `notes/00-course-map.html` as home.
- Layout must be the hellointerview-style three-column layout: course menu on the left, "On this page" outline and reading progress on the right. It comes from `assets/nav.js` + `assets/course-map.js` + `assets/course.css`, ported from the learner's observability course. Every new page must include those three and be registered in `course-map.js`.
- Brand: Astria Digital tokens (navy #051367, Lexend/Inter/IBM Plex Mono) are already in course.css.

## Topic 18 context — restructured to sub-topic notes (2026-10-07)

- Learner pushed back on a single big note for 14 runnable sub-folders: "it should be designed sub-topic wise." Agreed. **One note per sub-folder** is the right grain (also matches the one-win-per-lesson rule).
- New structure:
  - `notes/18-context.html` = **section hub/map** only (the one idea + JS `AbortController` bridge + a linked reading path of 5 acts). No per-example detail.
  - `notes/18-NN-name.html` = one short note per sub-topic, kept flat (one folder deep) so `../assets/` paths keep working. Each: idea → "the code that matters" excerpt → run/expect/try → one check-yourself → prev/next chain.
  - `course-map.js`: new sidebar group "Topic 18 · context" with all 14; built ones linked, 07–14 as `href: null` ("coming").
- **Built so far:** hub + 01–06 (Act 1 "feel the problem" + Act 2 "the four constructors"). Outputs for 01,02,03,04,05,06 all run and verified on Go 1.22.1.
- **Deliberate spacing:** 07–14 are left as "coming" and will be added in small batches as the learner progresses, not dumped at once.

## Sidebar structure (2026-10-07)

- Sidebar groups are collapsible accordions (click the group label). **Main topic groups open by default.**
- A topic with sub-notes is modelled as **nested children**, not a separate group: give its `course-map.js` item a `children: [...]` array (same item shape). It renders as a caret-toggled nested list under the parent item. The nest auto-opens when you're on the parent or any child.
- So Topic 18's 14 context sub-notes live under the "18 · The context package" item in the Topic notes group. Add future per-topic sub-notes the same way.

## Next

- Learner reads hub → 01 → … → 06, running each folder and doing each "Try" (predict before running).
- When they report Act 2 feels solid, build Act 3 notes: 07-propagation, 08-context-tree, 09-fan-out (same sub-note template; outputs already verified for 08).
- Then Act 4 (10–13) and Act 5 (14). Capstone idea unchanged: inventory lookup with request ID under a deadline (timeout + propagation + value).
