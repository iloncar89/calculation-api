package restErrors

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
}

type restErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

const MissingBothCalculationParamsErrorMessage = "Missing parameters x and y"
const FoundMoreThanOneBothParamsErrorMessage = "Found more than one x and y parameters"
const FoundMoreThanOneXParamsErrorMessage = "Found more than one x parameters"
const FoundMoreThanOneYParamsErrorMessage = "Found more than one y parameters"
const MissingParameterXErrorMessage = "Missing parameter x."
const MissingParameterYErrorMessage = "Missing parameter y."
const ErrorParsingStringToNumber = "One or more parameters in request aren't integer or float numbers"
const MissingParametersInPostRequest = "Missing both parameters in request"
const CannotDivideByZero = "You can't divide by zero"
const ErrorLeadingZeros = "One of the variables contains leading zero"
const BadRequestError = "bad_request"
const InternalServerError = "internal_server_error"

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s",
		e.ErrMessage, e.ErrStatus, e.ErrError)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

//NewBadRequestError function for given string returns restErr struct for error bad request.
//Sets received string as error message.
func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   BadRequestError,
	}
}

//NewInternalServerError function for given string returns restErr struct for internal server error.
//Sets received string as error message.
func NewInternalServerError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   InternalServerError,
	}
}
