package errors

import (
	"fmt"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	error   string `json:"error"`
	status  int    `json:"status"`
}

func NewBadRequest(message string) *RestErr {
	return New(message, "bad_request", http.StatusBadRequest)
}

func NewInternalError(message string) *RestErr {
	return New(message, "internal_error", http.StatusInternalServerError)
}

func NewNotFoundRequest(message string) *RestErr {
	return New(message, "not_found_error", http.StatusNotFound)
}

func New(msg string, err string, stat int) *RestErr {
	return &RestErr{
		Message: msg,
		error:   err,
		status:  stat,
	}
}

func (e *RestErr) Error() string {
	return fmt.Sprintf("Rest Error: %s , %s, with status code %d", e.Message, e.error, e.status)
}

func (re *RestErr) Status() int {
	return re.status
}
