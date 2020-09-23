require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

const role = document.getElementById("regularUser"),
    divEmail = document.getElementById("emailManager"),
    manager = document.getElementById("manager"),
    divRadio = document.getElementById("radioUser"),
    managerEmail = document.getElementById("managerEmail"),
    users = document.getElementById("users"),
    devices = document.getElementById("devices"),
    divManager = document.getElementById("radioManager");

divRadio.classList.add("borde")

role.addEventListener("click", () => {
    if (role.checked) {
        divEmail.style.display = "block";
        divRadio.classList.add("borde");
        divManager.classList.remove("borde");
    }
});

manager.addEventListener("click", () => {
    if (manager.checked) {
        divEmail.style.display = "none"
        divRadio.classList.remove("borde");
        divManager.classList.add("borde");
        console.log(managerEmail.value)
    }
});