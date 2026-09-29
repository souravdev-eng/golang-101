// Reusable local feedback. No Go code is executed by these widgets.
document.querySelectorAll('form[data-practice]').forEach(form => {
  form.addEventListener('submit', event => {
    event.preventDefault();
    const feedback = form.querySelector('[role="status"]');
    if (form.dataset.practice === 'choice') {
      const selected = form.querySelector('input:checked');
      feedback.textContent = selected
        ? selected.dataset.feedback
        : 'Choose an answer first, then check it.';
    } else if (form.dataset.practice === 'output') {
      const normalize = value => value.replace(/\r\n/g, '\n').trim();
      const answer = form.querySelector('textarea').value;
      feedback.textContent = normalize(answer) === normalize(form.dataset.expected)
        ? form.dataset.success
        : 'Compare the statements in order. Each Println finishes its own line. Try again before revealing the output.';
    }
  });
});
