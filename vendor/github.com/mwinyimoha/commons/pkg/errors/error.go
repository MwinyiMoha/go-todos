package errors

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/runtime/protoiface"
)

type ErrorCode int

const (
	Unknown ErrorCode = iota + 1
	NotFound
	InvalidArgument
	Internal
	Unauthenticated
	Unauthorized
	Conflict
	QuotaExceeded
	BadRequest
	NotImplemented
	ServiceUnavailable
	DeadlineExceeded
	PreconditionFailed
	TooEarly
	FailedDependency
	Gone
	UnprocessableEntity
	PayloadTooLarge
	UnsupportedMediaType
)

// GRPCDetailsFunc allows customizing how error details are attached to gRPC response.
// This is particularly useful with gRPC rich error reporting.
type GRPCDetailsFunc func(st *status.Status, details ...protoiface.MessageV1) (*status.Status, error)

// FieldViolation represents a validation field violation.
type FieldViolation struct {
	Field       string
	Description string
}

// Error is the unified application error type.
// It supports gRPC status details, HTTP status mapping, and validation field violations.
type Error struct {
	Original         error
	Message          string
	ErrCode          ErrorCode
	FieldViolations  []*FieldViolation
	ErrorDetailsFunc GRPCDetailsFunc
}

func (e *Error) Error() string {
	if len(e.FieldViolations) > 0 {
		var msgs []string
		for _, v := range e.FieldViolations {
			msgs = append(msgs, fmt.Sprintf("%s: %s", v.Field, v.Description))
		}
		return fmt.Sprintf("validation error(s): %v", msgs)
	}

	if e.Original != nil {
		return fmt.Sprintf("code: %d, message: %s, original error: %s", e.ErrCode, e.Message, e.Original)
	}

	return fmt.Sprintf("code: %d, message: %s", e.ErrCode, e.Message)
}

func (e *Error) Code() ErrorCode {
	return e.ErrCode
}

func (e *Error) Unwrap() error {
	return e.Original
}

// GRPCCode maps the internal error code to a gRPC status code.
func (e *Error) GRPCCode() codes.Code {
	switch e.ErrCode {
	case NotFound:
		return codes.NotFound
	case InvalidArgument, BadRequest, UnprocessableEntity:
		return codes.InvalidArgument
	case Internal, FailedDependency:
		return codes.Internal
	case Unauthenticated:
		return codes.Unauthenticated
	case Unauthorized:
		return codes.PermissionDenied
	case Conflict, PreconditionFailed:
		return codes.Aborted
	case QuotaExceeded:
		return codes.ResourceExhausted
	case NotImplemented, UnsupportedMediaType:
		return codes.Unimplemented
	case ServiceUnavailable:
		return codes.Unavailable
	case DeadlineExceeded, TooEarly:
		return codes.DeadlineExceeded
	case Gone:
		return codes.NotFound
	case PayloadTooLarge:
		return codes.OutOfRange
	default:
		return codes.Unknown
	}
}

// GRPCStatus converts the error into a gRPC status, with field or metadata details.
func (e *Error) GRPCStatus() *status.Status {
	st := status.New(e.GRPCCode(), e.Message)

	if len(e.FieldViolations) > 0 {
		violations := []*errdetails.BadRequest_FieldViolation{}
		for _, v := range e.FieldViolations {
			violations = append(
				violations,
				&errdetails.BadRequest_FieldViolation{
					Field:       v.Field,
					Description: v.Description,
				},
			)
		}

		detailed, err := e.ErrorDetailsFunc(st, &errdetails.BadRequest{FieldViolations: violations})
		if err != nil {
			return st
		}
		return detailed
	}

	if e.Original != nil {
		details := &errdetails.ErrorInfo{
			Reason:   e.Original.Error(),
			Metadata: map[string]string{"custom_code": fmt.Sprintf("%d", e.ErrCode)},
		}

		detailed, err := e.ErrorDetailsFunc(st, details)
		if err != nil {
			return st
		}
		return detailed
	}

	return st
}

// HTTPCode maps the internal error code to an HTTP status code.
func (e *Error) HTTPCode() int {
	switch e.ErrCode {
	case NotFound:
		return http.StatusNotFound
	case InvalidArgument, BadRequest:
		return http.StatusBadRequest
	case Internal, FailedDependency:
		return http.StatusInternalServerError
	case Unauthenticated:
		return http.StatusUnauthorized
	case Unauthorized:
		return http.StatusForbidden
	case Conflict, PreconditionFailed:
		return http.StatusConflict
	case QuotaExceeded:
		return http.StatusTooManyRequests
	case NotImplemented:
		return http.StatusNotImplemented
	case ServiceUnavailable:
		return http.StatusServiceUnavailable
	case DeadlineExceeded:
		return http.StatusGatewayTimeout
	case TooEarly:
		return http.StatusTooEarly
	case Gone:
		return http.StatusGone
	case UnprocessableEntity:
		return http.StatusUnprocessableEntity
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	case UnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	default:
		return http.StatusInternalServerError
	}
}

// HTTPStatus returns a standardized HTTP status code and JSON-safe error body.
// This is meant to be transport-agnostic, usable across multiple HTTP/REST frameworks.
func (e *Error) HTTPStatus() (int, map[string]any) {
	status := e.HTTPCode()

	body := map[string]any{
		"message": e.Message,
		"code":    e.ErrCode,
	}

	if len(e.FieldViolations) > 0 {
		var violations []map[string]string
		for _, v := range e.FieldViolations {
			violations = append(violations, map[string]string{
				"field":       v.Field,
				"description": v.Description,
			})
		}
		body["violations"] = violations
	}

	if e.Original != nil {
		body["reason"] = e.Original.Error()
	}

	return status, body
}

// addErrorDetails is the default detail attachment function for gRPC.
func addErrorDetails(st *status.Status, details ...protoiface.MessageV1) (*status.Status, error) {
	return st.WithDetails(details...)
}

// WrapError wraps an existing error into the unified Error type.
func WrapError(original error, code ErrorCode, format string, a ...interface{}) error {
	message := format
	if len(a) > 0 {
		message = fmt.Sprintf(format, a...)
	}

	return &Error{
		ErrCode:          code,
		Original:         original,
		Message:          message,
		ErrorDetailsFunc: addErrorDetails,
	}
}

// NewErrorf creates a new Error without wrapping an existing one.
func NewErrorf(code ErrorCode, format string, a ...interface{}) error {
	return WrapError(nil, code, format, a...)
}

// BuildViolations converts validator.ValidationErrors into []*FieldViolation.
func BuildViolations(verrs validator.ValidationErrors) []*FieldViolation {
	var violations []*FieldViolation
	for _, err := range verrs {
		violations = append(
			violations,
			&FieldViolation{
				Field:       err.StructField(),
				Description: err.Error(),
			},
		)
	}
	return violations
}

// NewValidationError builds an InvalidArgument error with validation details.
func NewValidationError(violations []*FieldViolation, msg ...string) *Error {
	message := "bad request"
	if len(msg) > 0 && msg[0] != "" {
		message = msg[0]
	}

	return &Error{
		Message:          message,
		ErrCode:          InvalidArgument,
		FieldViolations:  violations,
		ErrorDetailsFunc: addErrorDetails,
	}
}
