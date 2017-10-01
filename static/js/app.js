window.addEventListener("load", () => {
  let container = document.getElementById("alphabets");
  let addBtn = document.getElementById("add-button");
  let cnt = document.querySelectorAll("[name=checkboxes]").length || 0;
  addBtn.addEventListener("click", () => {
    let chbx = document.createElement("input");
    chbx.type = "checkbox";
    chbx.checked = true;
    chbx.name = "checkboxes";
    chbx.value = cnt;
    let input = document.createElement("input");
    input.type = "text";
    input.name = `alpha-${cnt}`;
    input.placeholder = "!@#$%^&*";
    container.insertBefore(chbx, addBtn);
    container.insertBefore(input, addBtn);
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
