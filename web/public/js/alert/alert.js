(function () {
  document
    .querySelector(".godo-alert span")
    .addEventListener("click", function () {
      this.parentNode.classList.add("godo-alert--hidden");
    });
})();
