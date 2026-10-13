/**
 * Small, offline syntax highlighter for the code blocks in notes and lessons.
 *
 * Every <pre> (or <pre><code>) holding plain text gets coloured spans. The
 * language comes from `data-lang` on the <pre> when present ("go", "js", "sh",
 * "text"); otherwise from the label or sentence just before the block:
 *   "Expect…" / "…output…"   → text (left as is: it's program output)
 *   "Run…" or a `go …` line  → sh
 *   "JS/TS…" / "JavaScript"  → js
 *   anything else            → go
 * Blocks that already contain markup (e.g. pre.diff) are skipped.
 *
 * Colours are the --hl-* tokens in course.css, so every theme restyles them.
 */
(function () {
  const words = (list) => new Set(list.split(' '));

  const LANGS = {
    go: {
      keywords: words('break case chan const continue default defer else fallthrough for func go goto if import interface map package range return select struct switch type var'),
      constants: words('true false nil iota'),
      types: words('bool byte complex64 complex128 error float32 float64 int int8 int16 int32 int64 rune string uint uint8 uint16 uint32 uint64 uintptr any comparable'),
      builtins: words('append cap clear close complex copy delete imag len make max min new panic print println real recover'),
      strings: /"(?:\\.|[^"\\\n])*"|`[^`]*`|'(?:\\.|[^'\\\n])*'/y,
    },
    js: {
      keywords: words('async await break case catch class const continue default delete do else export extends finally for from function if import in instanceof let new of return static super switch throw try typeof var void while yield interface type implements'),
      constants: words('true false null undefined this NaN Infinity'),
      types: words('string number boolean void any unknown never Promise Error Map Set Array Object'),
      builtins: words('console setTimeout clearTimeout fetch'),
      strings: /"(?:\\.|[^"\\\n])*"|'(?:\\.|[^'\\\n])*'|`(?:\\.|[^`\\])*`/y,
    },
  };

  const COMMENT = /\/\/[^\n]*|\/\*[\s\S]*?\*\//y;
  const NUMBER = /(?:0[xX][\da-fA-F_]+|\d[\d_]*(?:\.\d+)?(?:[eE][+-]?\d+)?)\b/y;
  const IDENT = /[A-Za-z_$][\w$]*/y;

  const escape = (s) => s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
  const span = (cls, text) => `<span class="hl-${cls}">${escape(text)}</span>`;

  /** Try `re` at `pos`; return the matched text or null. */
  const at = (re, src, pos) => { re.lastIndex = pos; const m = re.exec(src); return m ? m[0] : null; };

  const highlightCode = (src, lang) => {
    const L = LANGS[lang];
    let out = '';
    let pos = 0;
    let prevWord = '';
    while (pos < src.length) {
      let m;
      if ((m = at(COMMENT, src, pos))) { out += span('com', m); pos += m.length; continue; }
      if ((m = at(L.strings, src, pos))) { out += span('str', m); pos += m.length; continue; }
      if (/\d/.test(src[pos]) && !/[\w$]/.test(src[pos - 1] || '') && (m = at(NUMBER, src, pos))) {
        out += span('num', m); pos += m.length; continue;
      }
      if ((m = at(IDENT, src, pos))) {
        const next = src.slice(pos + m.length).match(/^\s*(.)/);
        const calls = next && next[1] === '(';
        let cls = null;
        if (L.keywords.has(m)) cls = 'kw';
        else if (L.constants.has(m)) cls = 'const';
        else if (L.types.has(m) || prevWord === 'type' || prevWord === 'class' || prevWord === 'interface') cls = 'type';
        else if (L.builtins.has(m)) cls = 'fn';
        else if (calls) cls = 'fn';
        out += cls ? span(cls, m) : escape(m);
        prevWord = m;
        pos += m.length;
        continue;
      }
      if (!/\s/.test(src[pos])) prevWord = '';
      out += escape(src[pos]);
      pos += 1;
    }
    return out;
  };

  /** Shell: first word of a line is the command, -flags and quoted args stand out, # comments. */
  const highlightShell = (src) => src.split('\n').map((line) => {
    const hash = line.search(/(^|\s)#/);
    const code = hash >= 0 ? line.slice(0, hash) : line;
    const comment = hash >= 0 ? line.slice(hash) : '';
    let first = true;
    const body = code.split(/(\s+)/).map((part) => {
      if (!part || /^\s+$/.test(part)) return part;
      if (part === '$') return span('com', part);
      if (first) { first = false; return span('fn', part); }
      if (/^--?[\w-]/.test(part)) return span('const', part);
      if (/^["'].*["']$/.test(part)) return span('str', part);
      return escape(part);
    }).join('');
    return body + (comment ? span('com', comment) : '');
  }).join('\n');

  const detect = (pre, text) => {
    if (pre.dataset.lang) return pre.dataset.lang;
    const before = pre.previousElementSibling
      || (pre.parentElement && pre.parentElement.previousElementSibling);
    const hint = before ? before.textContent.trim() : '';
    if (/expect|output/i.test(hint)) return 'text';
    if (/^run\b/i.test(hint)) return 'sh';
    if (/^\s*(\$ )?(go|cd|git|curl) [\w./-]/.test(text) && text.trim().split('\n').length <= 3) return 'sh';
    if (/\bJS\b|\bTS\b|JavaScript|TypeScript|Jest/i.test(hint)) return 'js';
    return 'go';
  };

  document.querySelectorAll('main pre').forEach((pre) => {
    const host = pre.children.length === 1 && pre.firstElementChild.tagName === 'CODE' ? pre.firstElementChild : pre;
    if (host.children.length) return;
    const text = host.textContent;
    const lang = detect(pre, text);
    pre.dataset.lang = lang;
    if (lang === 'text') { pre.classList.add('is-output'); return; }
    host.innerHTML = lang === 'sh' ? highlightShell(text) : highlightCode(text, LANGS[lang] ? lang : 'go');
  });
})();
