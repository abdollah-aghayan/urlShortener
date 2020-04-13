package errors

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	
	restErr := RestErr{
		Message: "error message",
		error: "error string",
		status: http.StatusOK,
	}

	err := New("error message", "error string", 200)

	if err.Message != restErr.Message {
		t.Errorf("Expected %s error message got %s", restErr.Message, err.Message)
	}

	error := fmt.Sprintf("Rest Error: %s , %s, with status code %d", restErr.Message, restErr.error, restErr.status)
	if err.Error() != restErr.Error() {
		t.Errorf("Expected %s error string got %s", err.Error(), error)
	}

	if err.Status() != restErr.status {
		t.Errorf("Expected %d error string got %d", restErr.status, err.Status())
	}
}

func TestBadRequest(t *testing.T) {
	message := "error message"
	errStr := "bad_request"
	status := 400
	err := NewBadRequest(message)

	if err.Message != message {
		t.Errorf("Expected %s error message got %s", message, err.Message)
	}

	error := fmt.Sprintf("Rest Error: %s , %s, with status code %d", message, errStr, status)
	if err.Error() != error {
		t.Errorf("Expected %s error string got %s", err.Error(), error)
	}

	if err.Status() != status {
		t.Errorf("Expected %d error string got %d", status, err.Status())
	}
}

func TestInternalError(t *testing.T) {
	message := "error message"
	errStr := "internal_error"
	status := 500
	err := NewInternalError(message)

	if err.Message != message {
		t.Errorf("Expected %s error message got %s", message, err.Message)
	}

	error := fmt.Sprintf("Rest Error: %s , %s, with status code %d", message, errStr, status)
	if err.Error() != error {
		t.Errorf("Expected %s error string got %s", err.Error(), error)
	}

	if err.Status() != status {
		t.Errorf("Expected %d error string got %d", status, err.Status())
	}
}

func TestNotFoundError(t *testing.T) {
	message := "error message"
	errStr := "not_found_error"
	status := 404
	err := NewNotFoundRequest(message)

	if err.Message != message {
		t.Errorf("Expected %s error message got %s", message, err.Message)
	}

	error := fmt.Sprintf("Rest Error: %s , %s, with status code %d", message, errStr, status)
	if err.Error() != error {
		t.Errorf("Expected %s error string got %s", err.Error(), error)
	}

	if err.Status() != status {
		t.Errorf("Expected %d error string got %d", status, err.Status())
	}
}