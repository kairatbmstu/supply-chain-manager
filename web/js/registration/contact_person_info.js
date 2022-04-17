$(document).ready(function () {

    $("#form-submit-button").click(function () {
        console.log(" click submit button ")
        var validator = new ContactPersonValidator();
        console.log("errors : ", validator.hasErrors())
        validator.validate();
        console.log("errors : ", validator.hasErrors())
        if (!validator.hasErrors()) {
            console.log("go submit");
            $("form").submit();
        }

    });
});

class ContactPersonValidator extends BaseValidator {
    constructor() {
        super()
    }


    validate() {
        var email = $("#email").val();
        var contactFullname = $("#contactFullname").val();
        var position = $("#position").val();

        console.log('email : ', email);
        console.log('contactFullname : ', contactFullname);
        console.log('position : ', position);


        if (email == null || email == "") {
            this.errors.push("Поле email не может быть пустым");
            $("#error-email-required").css("display", "block");

        } else {

            if (!this.isEmail(email)) {
                console.log("email true");
                this.errors.push("Заполните email пожалуйста");
                $("#error-not-email").css("display", "block");
            } else {
                console.log("email false");
                $("#error-not-email").css("display", "none");
            }
            $("#error-email-required").css("display", "none");
        }


        if (contactFullname == null || contactFullname == "") {
            this.errors.push("Поле contactFullname не может быть пустым");
            $("#error-contactFullname-required").css("display", "block");
        } else {
            $("#error-contactFullname-required").css("display", "none");
        }
        if (position == null || position == "") {
            this.errors.push("Поле position не может быть пустым");
            $("#error-position-required").css("display", "block");
        } else {
            $("#error-position-required").css("display", "none");
        }

    }
}