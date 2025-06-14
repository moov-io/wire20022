// Package errors provides domain-specific error types for wire20022 operations.
// These error types follow Go standard library conventions and support
// error wrapping, unwrapping, and comparison using errors.Is() and errors.As().
package errors

import (
	"errors"
	"fmt"
	"reflect"
)

// Sentinel errors for common validation failures.
// These can be compared using errors.Is().
var (
	ErrInvalidField     = errors.New("invalid field")
	ErrRequiredField    = errors.New("required field missing")
	ErrInvalidXML       = errors.New("invalid XML format")
	ErrInvalidJSON      = errors.New("invalid JSON format")
	ErrUnknownNamespace = errors.New("unknown XML namespace")
	ErrInvalidVersion   = errors.New("invalid version")
	ErrFieldNotFound    = errors.New("field not found")
	ErrIndexOutOfBounds = errors.New("array index out of bounds")
)

// ValidationError represents a field validation failure.
// It provides details about which field failed validation and why.
type ValidationError struct {
	Field  string // The field that failed validation
	Reason string // Human-readable reason for the failure
	Err    error  // Underlying error, if any
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("validation failed for field %q: %s: %v", e.Field, e.Reason, e.Err)
	}
	return fmt.Sprintf("validation failed for field %q: %s", e.Field, e.Reason)
}

// Unwrap returns the underlying error for error unwrapping.
func (e *ValidationError) Unwrap() error {
	return e.Err
}

// Is enables comparison with sentinel errors using errors.Is().
func (e *ValidationError) Is(target error) bool {
	if target == ErrInvalidField || target == ErrRequiredField {
		return true
	}
	return false
}

// ParseError represents an XML or JSON parsing failure.
// It provides context about what was being parsed and the underlying error.
type ParseError struct {
	Operation string // The parsing operation that failed (e.g., "XML unmarshal", "JSON marshal")
	Content   string // Brief description of content being parsed
	Err       error  // Underlying parsing error
}

// Error implements the error interface.
func (e *ParseError) Error() string {
	if e.Content != "" {
		return fmt.Sprintf("%s failed for %s: %v", e.Operation, e.Content, e.Err)
	}
	return fmt.Sprintf("%s failed: %v", e.Operation, e.Err)
}

// Unwrap returns the underlying error for error unwrapping.
func (e *ParseError) Unwrap() error {
	return e.Err
}

// Is enables comparison with sentinel errors using errors.Is().
func (e *ParseError) Is(target error) bool {
	if target == ErrInvalidXML || target == ErrInvalidJSON {
		return true
	}
	return false
}

// FieldError represents an error accessing or setting a field value.
// It provides context about the field path and operation.
type FieldError struct {
	Path      string // The field path that caused the error
	Operation string // The operation being performed (e.g., "get", "set")
	Err       error  // Underlying error
}

// Error implements the error interface.
func (e *FieldError) Error() string {
	return fmt.Sprintf("field %s %s failed: %v", e.Operation, e.Path, e.Err)
}

// Unwrap returns the underlying error for error unwrapping.
func (e *FieldError) Unwrap() error {
	return e.Err
}

// Is enables comparison with sentinel errors using errors.Is().
func (e *FieldError) Is(target error) bool {
	if target == ErrFieldNotFound || target == ErrIndexOutOfBounds {
		return true
	}
	return false
}

// Helper functions for internal error handling

// NewInternalError creates an error for internal programming issues.
// Use this for conditions that indicate bugs in the library code itself.
//
// Example usage:
//
//	if factoryMap == nil {
//	    return NewInternalError("factory map cannot be nil")
//	}
func NewInternalError(msg string) error {
	return fmt.Errorf("internal error: %s", msg)
}

// NewConfigurationError creates an error for configuration or setup issues.
// Use this when the library is misconfigured or initialized incorrectly.
//
// Example usage:
//
//	if len(supportedVersions) == 0 {
//	    return NewConfigurationError("no supported versions configured")
//	}
func NewConfigurationError(msg string) error {
	return fmt.Errorf("configuration error: %s", msg)
}

// ValidateCondition checks a condition and returns an error if it's false.
// This replaces traditional assertion patterns while maintaining error returns.
//
// Example usage:
//
//	if err := ValidateCondition(len(segments) > 0, "segments slice cannot be empty"); err != nil {
//	    return fmt.Errorf("failed to process path: %w", err)
//	}
func ValidateCondition(condition bool, msg string) error {
	if !condition {
		return NewInternalError(msg)
	}
	return nil
}

// ValidateNotNil checks that a value is not nil and returns an error if it is.
// This is a common pattern for preventing nil pointer dereferences.
// Properly handles both interface{} nil and typed nil pointers.
//
// Example usage:
//
//	if err := ValidateNotNil(factory, "factory"); err != nil {
//	    return fmt.Errorf("failed to create document: %w", err)
//	}
func ValidateNotNil(value interface{}, name string) error {
	if value == nil {
		return NewInternalError(fmt.Sprintf("%s cannot be nil", name))
	}

	// Use reflection to check for typed nil pointers/slices/maps/channels/functions
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		if v.IsNil() {
			return NewInternalError(fmt.Sprintf("%s cannot be nil", name))
		}
	case reflect.Invalid, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128, reflect.Array, reflect.String, reflect.Struct,
		reflect.UnsafePointer:
		// These types cannot be nil, no additional checking needed
	}

	return nil
}
