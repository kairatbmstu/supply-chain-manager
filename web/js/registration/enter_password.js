$(document).ready(function () {

    $("#form-submit-button").click(function () {
        console.log(" click submit button ");
        var validator = new EnterPasswordValidator();
        validator.validate()
        if (!validator.hasErrors()){
            $("form").submit();
        }
    });
});

class EnterPasswordValidator extends BaseValidator {
    constructor() {
        super()
    }

    validate() {
        var password = $("#password").val();
        var confirm_password = $("#confirm_password").val();

        console.log('password : ', password);
        console.log('confirm_password : ', confirm_password);
        

        if (password == null || password == "") {
            this.errors.push("Поле password не может быть пустым");
            $("#error-password").css("display", "block");
            $("#error-password").text("Поле password не может быть пустым");
            return;
        } else {
            $("#error-password").css("display", "none");
            $("#error-password").text("");
        }

        if (confirm_password == null || confirm_password == "") {
            this.errors.push("Поле confirm_password не может быть пустым");
            $("#error-confirm_password").css("display", "block");
            $("#error-confirm_password").text("Поле confirm_password не может быть пустым");
            return;
        } else {
            $("#error-confirm_password").css("display", "none");
            $("#error-confirm_password").text("");
        }


        if (password != null){
            if (password.length < 8){
                this.errors.push("Пароль должен состоять из минимум 8 символов");
                $("#error-password").css("display", "block");
                $("#error-password").text("Пароль должен состоять из минимум 8 символов");
                return;
            } else {
                $("#error-password").css("display", "none");
                $("#error-password").text("");
            }
        }

        if (password != null && confirm_password != null) {
            if (password !== confirm_password){
                this.errors.push("Введенные пароли должны быть одинаковы");
                $("#error-password").css("display", "block");
                $("#error-password").text("Введенные пароли должны быть одинаковы");
                return;
            } else {
                $("#error-password").css("display", "none");
                $("#error-password").text("");
            }
        }
      
    }
}