window.addEventListener("load", () => {
  let addBtn = document.getElementById("add-button");
  let cnt = document.querySelectorAll("[name=checkboxes]").length || 0;
  addBtn.addEventListener("click", () => {
    addBtn.insertAdjacentHTML(
      "beforebegin",
      `
      <label>
        <input
          type="checkbox"
          name="checkboxes"
          value="alpha-${cnt}"
          checked>
        Custom
      </label>
      <input
        type="text"
        name="alpha-${cnt}"
        placeholder="!@#$%^&*">
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
      alert("copied");
    }
  });
});
