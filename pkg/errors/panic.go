package errors

import (
	"fmt"
	"reflect"
)

// Error handling guidelines for wire20022
//
// NEVER USE PANIC - ALWAYS RETURN ERRORS
//
// This library should never panic under any circumstances. All error conditions,
// including programming errors, should be handled by returning appropriate errors.
// This ensures that applications using this library can always recover gracefully.
//
// Error return guidelines:
//
// For user input issues:
//   - Use ValidationError for field validation failures
//   - Use ParseError for malformed XML/JSON data
//   - Use FieldError for field access issues
//
// For programming issues (that might traditionally use panic):
//   - Return error with clear message about the programming issue
//   - Use appropriate error type (usually ValidationError or custom error)
//   - Document the error condition in function comments
//
// Examples of proper error handling:
//
// ❌ DON'T: panic("factory map cannot be nil")
// ✅ DO: return fmt.Errorf("invalid configuration: factory map cannot be nil")
//
// ❌ DON'T: panic("unreachable code path") 
// ✅ DO: return fmt.Errorf("internal error: unexpected state in parser")
//
// ❌ DON'T: panic("invalid XML")
// ✅ DO: return NewParseError("XML decode", "document", err)
//
// ❌ DON'T: panic("required field missing")
// ✅ DO: return NewRequiredFieldError("MessageId")

// NewInternalError creates an error for internal programming issues that 
// traditionally might have used panic. This ensures the library never panics
// while still clearly indicating programming problems.
//
// Use this for conditions that indicate bugs in the library code itself,
// such as:
//   - Nil pointers that should never be nil
//   - Invalid internal state
//   - Unreachable code paths
//   - Configuration errors during initialization
//
// Example usage:
//   if factoryMap == nil {
//       return NewInternalError("factory map cannot be nil")
//   }
func NewInternalError(msg string) error {
	return fmt.Errorf("internal error: %s", msg)
}

// NewConfigurationError creates an error for configuration or setup issues.
// Use this when the library is misconfigured or initialized incorrectly.
//
// Example usage:
//   if len(supportedVersions) == 0 {
//       return NewConfigurationError("no supported versions configured")
//   }
func NewConfigurationError(msg string) error {
	return fmt.Errorf("configuration error: %s", msg)
}

// ValidateCondition checks a condition and returns an error if it's false.
// This replaces traditional assertion patterns while maintaining error returns.
//
// Example usage:
//   if err := ValidateCondition(len(segments) > 0, "segments slice cannot be empty"); err != nil {
//       return fmt.Errorf("failed to process path: %w", err)
//   }
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
//   if err := ValidateNotNil(factory, "factory"); err != nil {
//       return fmt.Errorf("failed to create document: %w", err)
//   }
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
	}
	
	return nil
}