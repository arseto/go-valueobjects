package valueobjects

import (
	"fmt"
	"net/http"
)

type HttpResponse interface {
	GetStatus() int
}

type ErrorResponse struct {
	Status           int    `json:"status"`
	DeveloperMessage string `json:"developer_message"`
	UserMessage      string `json:"user_message"`
	ErrorCode        string `json:"error_code"`
	MoreInfo         string `json:"more_info"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Error %s: %s", e.ErrorCode, e.DeveloperMessage)
}

func (e *ErrorResponse) GetStatus() int {
	return e.Status
}

func NewDatabaseError(err error, errCode string) *ErrorResponse {
	errorResponse := &ErrorResponse{
		Status:           http.StatusInternalServerError,
		DeveloperMessage: err.Error(),
		UserMessage:      "Internal Database Error",
		ErrorCode:        errCode,
		MoreInfo:         "@arseto",
	}
	return errorResponse
}

func NewValidationError(fieldName string, vt ValidationType,
	errCode string, fieldValue string) *ErrorResponse {
	err := &ErrorResponse{
		Status:           http.StatusUnprocessableEntity,
		DeveloperMessage: fieldName + " " + vt.GetMessage(fieldValue),
		UserMessage:      fieldName + " " + vt.GetMessage(fieldValue),
		ErrorCode:        errCode,
		MoreInfo:         "@arseto",
	}
	return err
}

type SuccessResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Status  int         `json:"status"`
}

func (r *SuccessResponse) GetStatus() int {
	return r.Status
}
