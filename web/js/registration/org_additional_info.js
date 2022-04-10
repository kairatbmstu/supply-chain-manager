$(document).ready(function () {

    $("#form-submit-button").click(function () {
        console.log(" click submit button ")
    });
});

class OrgAdditionalInfoValidator extends BaseValidator {
    constructor(){
        super()
    }

    validate(){
        var phoneNumber = $("#phoneNumber").val();
        var orgSite = $("#orgSite").val();
        var iinbin = $("#iinbin").val();
        var accoutNumber = $("#accoutNumber").val();
        var bic = $("#bic").val();
        var kbe = $("#kbe").val();

        console.log('phoneNumber : ', phoneNumber);
        console.log('orgSite : ', orgSite);
        console.log('iinbin : ', iinbin);
        console.log('accoutNumber : ', accoutNumber);
        console.log('bic : ', bic);
        console.log('kbe : ', kbe);

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

    hasErrors(){

    }
}