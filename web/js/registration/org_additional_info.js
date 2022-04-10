$(document).ready(function () {

    $("#form").submit(function(e) {
        e.preventDefault();
        e.stopPropagation();
        var validator = new OrgAdditionalInfoValidator();
        validator.validate();
        if (!validator.hasErrors()) {
            $("#form").submit();
        }
    });
});

class OrgAdditionalInfoValidator extends BaseValidator {
    constructor() {
        super()
    }

    validate() {
        var phoneNumber = $("#phoneNumber").val();
        var orgSite = $("#orgSite").val();
        var iinbin = $("#iinbin").val();
        var iban = $("#iban").val();
        var bic = $("#bic").val();
        var kbe = $("#kbe").val();

        console.log('phoneNumber : ', phoneNumber);
        console.log('orgSite : ', orgSite);
        console.log('iinbin : ', iinbin);
        console.log('iban : ', iban);
        console.log('bic : ', bic);
        console.log('kbe : ', kbe);

        if (phoneNumber == null || phoneNumber == "") {
            this.errors.push(new ErrorField("org_name", "Поле не может быть пустым"));
            $("#phoneNumber-error").css("display", "block");
        } else {
            $("#phoneNumber-error").css("display", "none");
        }
        if (orgSite == null || orgSite == "") {
            this.errors.push(new ErrorField("form_of_law", "Поле не может быть пустым"));
            $("#orgSite-error").css("display", "block");
        } else {
            $("#orgSite-error").css("display", "none");
        }
        if (iinbin == null || iinbin == "") {
            this.errors.push(new ErrorField("org_address", "Поле не может быть пустым"));
            $("#iinbin-error").css("display", "block");
        } else {
            $("#iinbin-error").css("display", "none");
        }
        if (iban == null || iban == "") {
            this.errors.push(new ErrorField("org_name", "Поле не может быть пустым"));
            $("#iban-error").css("display", "block");
        } else {
            $("#iban-error").css("display", "none");
        }
        if (bic == null || bic == "") {
            this.errors.push(new ErrorField("bic", "Поле не может быть пустым"));
            $("#bic-error").css("display", "block");
        } else {
            $("#bic-error").css("display", "none");
        }
        if (kbe == null || kbe == "") {
            this.errors.push(new ErrorField("kbe", "Поле не может быть пустым"));
            $("#kbe-error").css("display", "block");
        } else {
            $("#kbe-error").css("display", "none");
        }
    }

}