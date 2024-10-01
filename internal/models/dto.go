package models

type SuccessResponse struct {
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}
