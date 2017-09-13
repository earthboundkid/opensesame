window.addEventListener("load", () => {
  let container = document.getElementById("alphabets");
  let addBtn = document.getElementById("add-button");
  let createAlpha = alpha => {
    let ta = document.createElement("textarea");
    ta.value = alpha;
    ta.name = "alpha";

    let button = document.createElement("button");
    button.attributes.type = "button";
    button.textContent = "-";
    button.addEventListener("click", () => {
      ta.remove();
      button.remove();
    });

    container.insertBefore(ta, addBtn);
    container.insertBefore(button, addBtn);
  };

  window.defaultAlphabets.forEach(createAlpha);

  addBtn.addEventListener("click", () => createAlpha(""));

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
