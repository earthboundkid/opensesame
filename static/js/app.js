window.addEventListener("load", () => {
  let addBtn = document.getElementById("add-button");
  let container = document.getElementById("req-list");
  let cnt = document.querySelectorAll("[name=checkboxes]").length || 0;
  addBtn.addEventListener("click", () => {
    container.insertAdjacentHTML(
      "beforeend",
      `
      <li>
          <label for="alpha-${cnt}">
              Requirement
          </label>
          <input
              id="alpha-${cnt}"
              type="checkbox"
              name="checkboxes"
              value="alpha-${cnt}"
              checked>
          <input
              type="text"
              name="alpha-${cnt}"
              placeholder="!@#$%^&*">
      </li>
    `
    );
    cnt++;
  });

  let pwEl = document.getElementById("password");

  pwEl.addEventListener("click", () => {
    let range = document.createRange();
    range.selectNodeContents(pwEl);
    let selection = window.getSelection();
    selection.removeAllRanges();
    selection.addRange(range);
    if (document.execCommand("copy")) {
      alert("Copied");
    }
  });
});
