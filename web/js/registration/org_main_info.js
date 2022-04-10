$(document).ready(function () {

    $("#form-submit-button").click(function () {
        console.log(" click submit button ")
        var validator = new OrgMainInfoValidator();
        validator.validate();
        if (!validator.hasErrors()) {
            $("#form").submit();
        }
    });
});


class OrgMainInfoValidator extends BaseValidator {
    constructor() {
        super();
        this.errors = [];
    }

    validate() {
        var orgNameVal = $("#org_name").val();
        var formOfLawVal = $("#form_of_law").val();
        var orgAddressVal = $("#org_address").val();

        console.log('orgNameVal : ', orgNameVal);
        console.log('formOfLawVal : ', formOfLawVal);
        console.log('orgAddressVal : ', orgAddressVal);

        if (orgNameVal == null || orgNameVal == "") {
            this.errors.push(new ErrorField("org_name", "Поле не может быть пустым"));
            $("#org_name-error").css("display","block");
        } else {
            $("#org_name-error").css("display","none");     
        }
        if (formOfLawVal == null || formOfLawVal == "") {
            this.errors.push(new ErrorField("form_of_law", "Поле не может быть пустым"));
            $("#form_of_law-error").css("display","block");
        } else {
            $("#form_of_law-error").css("display","none");
        }
        if (orgAddressVal == null || orgAddressVal == "") {
            this.errors.push(new ErrorField("org_address", "Поле не может быть пустым"));
            $("#org_address-error").css("display","block");
        } else {
            $("#org_address-error").css("display","none");
        }
    }

   
}