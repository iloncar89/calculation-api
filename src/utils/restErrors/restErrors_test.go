package restErrors

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	message := "something bad happen"
	err := NewBadRequestError(message)
	errToString := fmt.Sprintf("message: %s - status: %d - error: %s",
		err.Message(), err.Status(), BadRequestError)

	if err.Message() != message || err.Status() != http.StatusBadRequest || err.Error() != errToString {
		t.Errorf("Expected error message %s, status %d, error to string message %s but got %+v", message, http.StatusBadRequest, errToString, err)
	}
}

func TestNewInternalServerError(t *testing.T) {
	message := "something bad happen"
	err := NewInternalServerError(message)
	errToString := fmt.Sprintf("message: %s - status: %d - error: %s",
		err.Message(), err.Status(), InternalServerError)

	if err.Message() != message || err.Status() != http.StatusInternalServerError || err.Error() != errToString {
		t.Errorf("Expected error message %s, status %d, error to string message %s but got %+v", message, http.StatusInternalServerError, errToString, err)
	}
}

func TestRestErr_Error(t *testing.T) {
	message := "something bad happen"
	err := NewBadRequestError(message)
	errToString := fmt.Sprintf("message: %s - status: %d - error: %s",
		err.Message(), err.Status(), BadRequestError)

	if err.Error() != errToString {
		t.Errorf("Expected error message %s, status %d, and error %s but got %+v", message, http.StatusBadRequest, BadRequestError, err)
	}
}

func TestRestErr_Message(t *testing.T) {
	message := "something bad happen"
	err := NewBadRequestError(message)

	if err.Message() != message {
		t.Errorf("Expected error message %s, status %d, and error %s but got %+v", message, http.StatusBadRequest, BadRequestError, err)
	}
}

func TestRestErr_Status(t *testing.T) {
	message := "something bad happen"
	err := NewInternalServerError(message)

	if err.Status() != http.StatusInternalServerError {
		t.Errorf("Expected error message %s, status %d, and error %s but got %+v", message, http.StatusInternalServerError, InternalServerError, err)
	}
}
