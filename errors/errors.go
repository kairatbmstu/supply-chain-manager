package errors

type ErrorDTO struct {
	HttpStatus   int
	ErrorCode    string
	ErrorMessage string
	FieldValues  []FieldValueDTO
}

func (e ErrorDTO) Error() string {
	return e.ErrorMessage
}

type FieldValueDTO struct {
	FieldName    string
	ErrorMessage string
}

type ErrorType int

const (
	UserNotFound              ErrorType = 1
	UserAlreadyExists         ErrorType = 2
	OrganizationAlreadyExists ErrorType = 3
	InternalServerError       ErrorType = 4
)

func NewError(errorType ErrorType) ErrorDTO {
	var errorDto = ErrorDTO{}

	switch errorType {
	case UserNotFound:
		errorDto.HttpStatus = 404
		errorDto.ErrorCode = "user_not_Found"
		errorDto.ErrorMessage = "User was not found"
	case UserAlreadyExists:
		errorDto.HttpStatus = 409
		errorDto.ErrorCode = "user_already_exists"
		errorDto.ErrorMessage = "User already exists"
	}

	return errorDto
}
