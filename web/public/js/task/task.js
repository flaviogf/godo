(function () {
  document.querySelector(".godo-task__input").addEventListener("input", function () {
    if (this.value) {
      return document.querySelector("button[type=submit]").removeAttribute("disabled");
    }

    return document.querySelector("button[type=submit]").setAttribute("disabled", true);
  });

  document.querySelectorAll(".godo-task__checkbox").forEach(function (it) {
    it.addEventListener("change", function () {
      this.parentNode.submit();
    });
  });

  document.querySelectorAll(".godo-task__trash").forEach(function (it) {
    it.addEventListener("click", function () {
      this.parentNode.submit();
    });
  });
})();
