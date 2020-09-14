require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

const role = document.getElementById("regularUser"),
    divEmail = document.getElementById("emailManager"),
    manager = document.getElementById("manager");

divEmail.style.display = "none"

role.addEventListener("click", () => {
    if (role.checked) {
        console.log("hola")
        divEmail.style.display = "block"
    }
})

manager.addEventListener("click", () => {
    if (manager.checked) {
        console.log("hola")
        divEmail.style.display = "none"
    }
})
