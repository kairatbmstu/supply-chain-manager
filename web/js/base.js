class BaseValidator {
    errors;
    constructor() {
        this.errors = []
    }

    validate() {

    }

    hasErrors() {
        console.log(this.errors);
        if (this.errors.length > 0) {
            return true;
        }

        return false;
    }

    isEmail(email) {
        var regex = /^([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9-])+\.)+([a-zA-Z0-9]{2,4})+$/;
        return regex.test(email);
    }
}

class ErrorField {
    fieldName;
    fieldMessage;
    constructor(fieldName, fieldMessage) {
        this.fieldName;
        this.fieldMessage;
    }
}