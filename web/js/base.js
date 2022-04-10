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
}

class ErrorField {
    fieldName;
    fieldMessage;
    constructor(fieldName,fieldMessage){
        this.fieldName;    
        this.fieldMessage;
    }
}