(function () {
  document.querySelector(".godo-task__input").addEventListener("input", function () {
    if (this.value) {
      return document.querySelector("button[type=submit]").removeAttribute("disabled");
    }

    return document.querySelector("button[type=submit]").setAttribute("disabled", true);
  });
})();
