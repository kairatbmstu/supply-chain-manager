$(document).ready(function () {

    $("#goto_registration").click(function () {
        window.location = "/register/resident_type";
    });

    $("#login-form > button[type='button']").click(function () {

        if (!hasErrors()) {
            console.log("form submit");
            $("#login-form").submit();
        }


    });


});

function hasErrors() {
    var usernameHasErrors = false
    var passwordHasErrors = false
    var usernameEmailHasErrors = false

    var username = $("#username").val();
    var password = $("#password").val();

    console.log("username: " + username);
    console.log("password: " + password);

    if (username == null || username == "") {
        $("#username-error").css("display", "block")
        $("#username-valid").css("display", "none")
        usernameHasErrors = true
    } else {
        $("#username-error").css("display", "none")
        $("#username-valid").css("display", "block")
        usernameHasErrors = false
    }

    if (!isEmail(username)) {
        $("#username-email-error").css("display", "block")
        usernameEmailHasErrors = true
    } else {
        $("#username-email-error").css("display", "none")
        usernameEmailHasErrors = false
    }

    if (password == null || password == "") {
        $("#password-error").css("display", "block")
        $("#password-valid").css("display", "none")
        passwordHasErrors = true
    } else {
        $("#password-error").css("display", "none")
        $("#password-valid").css("display", "block")
        passwordHasErrors = false
    }

    return usernameHasErrors || passwordHasErrors || usernameEmailHasErrors
}


function isEmail(email) {
    var regex = /^([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    return regex.test(email);
}
