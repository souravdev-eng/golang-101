/**
 * Applies the saved reading theme before the page paints. Loaded in <head>,
 * so it must stay tiny and synchronous. The theme picker itself lives in
 * nav.js; it writes the same key.
 *
 *   auto   no data-theme: follow the OS light/dark setting
 *   light  data-theme="light"
 *   dark   data-theme="dark"
 *   paper  data-theme="paper": warm sepia page, serif body for long reads
 */
(function () {
  let theme = 'auto';
  try { theme = localStorage.getItem('go-basic:theme') || 'auto'; } catch (e) { /* storage blocked */ }
  if (theme !== 'auto') document.documentElement.setAttribute('data-theme', theme);
})();
