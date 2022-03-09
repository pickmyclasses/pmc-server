package shared

import "net/http"

type AppErr interface {
	Error() string
	Code() int
}

// ContentNotFoundErr represents resource not found error
// this should be used for whenever a content (eg, course) not found by the given criteria
type ContentNotFoundErr struct{}

func (err ContentNotFoundErr) Error() string {
	return NoInfoErr
}

func (err ContentNotFoundErr) Code() int {
	return http.StatusNotFound
}

// UserInfoNotFoundErr represents user info not found error
// this should only be used for user ID doesn't exist in the database
type UserInfoNotFoundErr struct{}

func (err UserInfoNotFoundErr) Error() string {
	return UserNotFoundErr
}

func (err UserInfoNotFoundErr) Code() int {
	return http.StatusUnauthorized
}

// ParamInsufficientErr represents post/get param body insufficiency
// this should be used for any request that has a param body
type ParamInsufficientErr struct {}

func (err ParamInsufficientErr) Error() string {
	return InsufficientParamErr
}

func (err ParamInsufficientErr) Code() int {
	return http.StatusBadRequest
}

// ParamIncompatibleErr represents incompatible parameters such as invalid page number, etc
type ParamIncompatibleErr struct {}

func (err ParamIncompatibleErr) Error() string {
	return BadParamErr
}

func (err ParamIncompatibleErr) Code() int {
	return http.StatusBadRequest
}

// MalformedIDErr represents the error when the provided ID is malformed or does not exist
type MalformedIDErr struct {}

func (err MalformedIDErr) Error() string {
	return BadIdErr
}

func (err MalformedIDErr) Code() int {
	return http.StatusUnprocessableEntity
}

// InternalErr represents an internal server errors
// this is mostly cost by internet delay or server issue
type InternalErr struct{}

func (err InternalErr) Error() string {
	return InternalServerErr
}

func (err InternalErr) Code() int {
	return http.StatusInternalServerError
}
