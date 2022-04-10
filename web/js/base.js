class Validator {

    constructor(form) {
        this.errors = []
        this.form = form
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