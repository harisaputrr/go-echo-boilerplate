package utils

import "net/http"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BadRequestError struct {
	ErrorResponse
}

type NotFoundError struct {
	ErrorResponse
}

type DuplicatedError struct {
	ErrorResponse
}

type UnauthorizedError struct {
	ErrorResponse
}

type ForbiddenError struct {
	ErrorResponse
}

type InternalServerError struct {
	ErrorResponse
}

func (err BadRequestError) Error() string {
	return err.Message
}

func (err *NotFoundError) Error() string {
	return err.Message
}

func (err *DuplicatedError) Error() string {
	return err.Message
}

func (err *UnauthorizedError) Error() string {
	return err.Message
}

func (err *ForbiddenError) Error() string {
	return err.Message
}

func (err *InternalServerError) Error() string {
	return err.Message
}

func NewBadRequestError(err error) BadRequestError {
	return BadRequestError{
		ErrorResponse: ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		},
	}
}

func NewNotFoundError(err error) NotFoundError {
	return NotFoundError{
		ErrorResponse: ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		},
	}
}

func NewDuplicatedError(err error) DuplicatedError {
	return DuplicatedError{
		ErrorResponse: ErrorResponse{
			Code:    http.StatusConflict,
			Message: err.Error(),
		},
	}
}

func NewUnauthorizedError(err error) UnauthorizedError {
	return UnauthorizedError{
		ErrorResponse: ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		},
	}
}

func NewForbiddenError(err error) ForbiddenError {
	return ForbiddenError{
		ErrorResponse: ErrorResponse{
			Code:    http.StatusForbidden,
			Message: err.Error(),
		},
	}
}

func NewInternalServerError(err error) InternalServerError {
	return InternalServerError{
		ErrorResponse: ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	}
}
