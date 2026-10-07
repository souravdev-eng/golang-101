/**
 * The course outline shown in every page's left sidebar. Adding a page means
 * adding one entry here. An entry with `href: null` is planned but not
 * written yet, and shows as "coming". An entry may also carry a `children`
 * array of the same shape; those render as a collapsible nested list under it
 * (used for a topic's sub-notes, e.g. 18's fourteen context sub-topics).
 *
 * Paths are relative to a page one folder deep (lessons/, notes/, reference/).
 */
window.COURSE = {
  title: 'Go for JS/TS developers',
  subtitle: 'go-basic · 18 topics',
  home: '../notes/00-course-map.html',
  groups: [
    {
      label: 'Start here',
      items: [
        { num: 0, title: 'Course map', href: '../notes/00-course-map.html' },
      ],
    },
    {
      label: 'Topic notes',
      items: [
        { num: 1, title: 'Programs and output', href: '../notes/01-programs.html' },
        { num: 2, title: 'Variables, constants, types', href: '../notes/02-values.html' },
        { num: 3, title: 'Operators', href: '../notes/03-operators.html' },
        { num: 4, title: 'Conditions and switches', href: '../notes/04-conditions.html' },
        { num: 5, title: 'Loops', href: '../notes/05-loops.html' },
        { num: 6, title: 'Functions and scope', href: '../notes/06-functions.html' },
        { num: 7, title: 'Arrays and slices', href: '../notes/07-collections.html' },
        { num: 8, title: 'Maps', href: '../notes/08-maps.html' },
        { num: 9, title: 'Strings, bytes, runes', href: '../notes/09-text.html' },
        { num: 10, title: 'Structs and methods', href: '../notes/10-structs.html' },
        { num: 11, title: 'Pointers', href: '../notes/11-pointers.html' },
        { num: 12, title: 'Interfaces', href: '../notes/12-interfaces.html' },
        { num: 13, title: 'Errors', href: '../notes/13-errors.html' },
        { num: 14, title: 'Defer', href: '../notes/14-defer.html' },
        { num: 15, title: 'Standard library', href: '../notes/15-standard-library.html' },
        { num: 16, title: 'Testing', href: '../notes/16-testing.html' },
        { num: 17, title: 'Goroutines and channels', href: '../notes/17-concurrency.html' },
        {
          num: 18, title: 'The context package', href: '../notes/18-context.html',
          children: [
            { num: 1, title: 'Why context exists', href: '../notes/18-01-why-context.html' },
            { num: 2, title: 'The empty root', href: '../notes/18-02-background.html' },
            { num: 3, title: 'WithCancel', href: '../notes/18-03-with-cancel.html' },
            { num: 4, title: 'WithTimeout', href: '../notes/18-04-with-timeout.html' },
            { num: 5, title: 'WithDeadline', href: '../notes/18-05-with-deadline.html' },
            { num: 6, title: 'Done and Err', href: '../notes/18-06-done-and-err.html' },
            { num: 7, title: 'Propagation', href: null },
            { num: 8, title: 'The context tree', href: null },
            { num: 9, title: 'Fan-out', href: null },
            { num: 10, title: 'WithValue', href: null },
            { num: 11, title: 'HTTP server', href: null },
            { num: 12, title: 'HTTP client', href: null },
            { num: 13, title: 'Graceful shutdown', href: null },
            { num: 14, title: 'Rules and pitfalls', href: null },
          ],
        },
      ],
    },
    {
      label: 'Lessons',
      items: [
        { num: 1, title: "Find Go's entry point", href: '../lessons/0001-go-program-entry.html' },
      ],
    },
    {
      label: 'Reference',
      items: [
        { title: 'Program basics & glossary', href: '../reference/go-program-basics.html' },
      ],
    },
  ],
};
