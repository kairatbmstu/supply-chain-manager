class BaseValidator {
    errors;
    constructor() {
        this.errors = []
    }

    validate() {

    }

    hasErrors() {
        return false
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