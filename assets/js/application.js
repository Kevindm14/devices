require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");


const role = document.getElementById("regularUser"),
    divEmail = document.getElementById("emailManager"),
    manager = document.getElementById("manager"),
    divRadio = document.getElementById("radioUser"),
    divManager = document.getElementById("radioManager");

$("#radioUser").addClass("borde")

$("#regularUser").on("click", () => {
    if (role.checked) {
        divEmail.style.display = "block";
        divRadio.classList.add("borde");
        divManager.classList.remove("borde");
    }
});

$("#manager").on("click", () => {
    if (manager.checked) {
        divEmail.style.display = "none"
        divRadio.classList.remove("borde");
        divManager.classList.add("borde");
    }
});

$("#addDevice").on("click", () => {
    const deviceVersion = `
        <div class="row text-muted">
            <div class="col-lg-3">
                <label>Storage</label>
                <input type="text" name="Storage" class="form-control" %>
            </div>
            <div class="col-lg-3">
                <label>Cost</label>
                <input type="number" step="0.01" name="Cost" class="form-control" %>
            </div>
            <div class="col-lg-3">
                <label>Color</label>
                <input type="text" name="Color" class="form-control" %>
            </div>
            <div class="col-lg-2">
                <label>Image</label>
                <input type="text" name="Image" class="form-control" %>
            </div>
            <div class="col-lg-1 text-center">
                <button class="btn btn-version-delete"><span><i class="far fa-trash-alt"></i></span></button>
            </div>
        </div>
    `
    $( "#versions" ).before(deviceVersion);
})

