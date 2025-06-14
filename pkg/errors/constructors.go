package errors

import "fmt"

// NewValidationError creates a new ValidationError with the specified field and reason.
// Returns a concrete ValidationError type, not an error interface.
//
// Example:
//
//	err := NewValidationError("MessageId", "cannot be empty")
//	// err.Error() returns: validation failed for field "MessageId": cannot be empty
func NewValidationError(field, reason string) *ValidationError {
	return &ValidationError{
		Field:  field,
		Reason: reason,
	}
}

// NewValidationErrorWithCause creates a new ValidationError that wraps an underlying error.
// Returns a concrete ValidationError type, not an error interface.
//
// Example:
//
//	err := NewValidationErrorWithCause("Amount", "invalid format", strconv.ErrSyntax)
func NewValidationErrorWithCause(field, reason string, cause error) *ValidationError {
	return &ValidationError{
		Field:  field,
		Reason: reason,
		Err:    cause,
	}
}

// NewParseError creates a new ParseError for parsing operations.
// Returns a concrete ParseError type, not an error interface.
//
// Example:
//
//	err := NewParseError("XML unmarshal", "Document", xmlErr)
//	// err.Error() returns: XML unmarshal failed for Document: <xmlErr>
func NewParseError(operation, content string, cause error) *ParseError {
	return &ParseError{
		Operation: operation,
		Content:   content,
		Err:       cause,
	}
}

// NewParseErrorSimple creates a new ParseError without content description.
// Returns a concrete ParseError type, not an error interface.
//
// Example:
//
//	err := NewParseErrorSimple("JSON marshal", jsonErr)
//	// err.Error() returns: JSON marshal failed: <jsonErr>
func NewParseErrorSimple(operation string, cause error) *ParseError {
	return &ParseError{
		Operation: operation,
		Err:       cause,
	}
}

// NewFieldError creates a new FieldError for field access operations.
// Returns a concrete FieldError type, not an error interface.
//
// Example:
//
//	err := NewFieldError("Header.ID", "get", ErrFieldNotFound)
//	// err.Error() returns: field get Header.ID failed: field not found
func NewFieldError(path, operation string, cause error) *FieldError {
	return &FieldError{
		Path:      path,
		Operation: operation,
		Err:       cause,
	}
}

// NewRequiredFieldError creates a ValidationError for missing required fields.
// This is a convenience constructor that uses the ErrRequiredField sentinel.
//
// Example:
//
//	err := NewRequiredFieldError("MessageId")
//	// err.Error() returns: validation failed for field "MessageId": is required
func NewRequiredFieldError(field string) *ValidationError {
	return &ValidationError{
		Field:  field,
		Reason: "is required",
		Err:    ErrRequiredField,
	}
}

// NewInvalidFieldError creates a ValidationError for invalid field values.
// This is a convenience constructor that uses the ErrInvalidField sentinel.
//
// Example:
//
//	err := NewInvalidFieldError("Amount", "must be positive")
//	// err.Error() returns: validation failed for field "Amount": must be positive
func NewInvalidFieldError(field, reason string) *ValidationError {
	return &ValidationError{
		Field:  field,
		Reason: reason,
		Err:    ErrInvalidField,
	}
}

// WrapValidationError wraps an existing error as a ValidationError.
// This is useful when you have a generic error that you want to associate with a field.
//
// Example:
//
//	parseErr := json.Unmarshal(data, &obj)
//	err := WrapValidationError("RequestData", "invalid JSON format", parseErr)
func WrapValidationError(field, reason string, cause error) *ValidationError {
	return &ValidationError{
		Field:  field,
		Reason: reason,
		Err:    fmt.Errorf("%w", cause),
	}
}
