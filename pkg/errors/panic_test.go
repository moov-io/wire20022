package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalError(t *testing.T) {
	err := NewInternalError("factory map cannot be nil")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "internal error")
	assert.Contains(t, err.Error(), "factory map cannot be nil")
}

func TestNewConfigurationError(t *testing.T) {
	err := NewConfigurationError("no supported versions configured")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "configuration error")
	assert.Contains(t, err.Error(), "no supported versions configured")
}

func TestValidateCondition(t *testing.T) {
	t.Run("returns nil when condition is true", func(t *testing.T) {
		err := ValidateCondition(true, "this should not error")
		assert.NoError(t, err)
	})
	
	t.Run("returns error when condition is false", func(t *testing.T) {
		err := ValidateCondition(false, "condition failed")
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "internal error")
		assert.Contains(t, err.Error(), "condition failed")
	})
}

func TestValidateNotNil(t *testing.T) {
	t.Run("returns nil when value is not nil", func(t *testing.T) {
		value := "not nil"
		err := ValidateNotNil(value, "test value")
		assert.NoError(t, err)
	})
	
	t.Run("returns error when value is nil", func(t *testing.T) {
		var value *string = nil
		err := ValidateNotNil(value, "test value")
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "internal error")
		assert.Contains(t, err.Error(), "test value cannot be nil")
	})
	
	t.Run("works with interface{} nil", func(t *testing.T) {
		var value interface{} = nil
		err := ValidateNotNil(value, "interface value")
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "interface value cannot be nil")
	})
	
	t.Run("works with non-nil interface{}", func(t *testing.T) {
		var value interface{} = "something"
		err := ValidateNotNil(value, "interface value")
		assert.NoError(t, err)
	})
}

func TestErrorHandlingGuidelines(t *testing.T) {
	// This test documents the proper error handling patterns for wire20022
	
	t.Run("NEVER use panic - always return errors", func(t *testing.T) {
		// All these scenarios should return errors, never panic
		
		// ❌ DON'T: panic("factory map cannot be nil")
		// ✅ DO: 
		err1 := NewInternalError("factory map cannot be nil")
		assert.Error(t, err1)
		
		// ❌ DON'T: panic("unreachable code path") 
		// ✅ DO:
		err2 := NewInternalError("unexpected state in parser")
		assert.Error(t, err2)
		
		// ❌ DON'T: panic("invalid XML")
		// ✅ DO:
		err3 := NewParseError("XML decode", "document", NewInvalidFieldError("test", "invalid"))
		assert.Error(t, err3)
		
		// ❌ DON'T: panic("required field missing")
		// ✅ DO:
		err4 := NewRequiredFieldError("MessageId")
		assert.Error(t, err4)
	})
	
	t.Run("proper error types for different scenarios", func(t *testing.T) {
		// User input validation
		userErr := NewValidationError("Amount", "must be positive")
		assert.Error(t, userErr)
		
		// Data parsing issues
		parseErr := NewParseError("JSON unmarshal", "request body", NewInvalidFieldError("format", "invalid"))
		assert.Error(t, parseErr)
		
		// Field access problems
		fieldErr := NewFieldError("Header.ID", "get", ErrFieldNotFound)
		assert.Error(t, fieldErr)
		
		// Internal programming issues
		internalErr := NewInternalError("unexpected nil pointer")
		assert.Error(t, internalErr)
		
		// Configuration problems
		configErr := NewConfigurationError("missing required configuration")
		assert.Error(t, configErr)
	})
}

func TestReplacingPanicPatterns(t *testing.T) {
	// Examples of how to replace common panic patterns with error returns
	
	t.Run("nil pointer check replacement", func(t *testing.T) {
		// Old panic pattern:
		// if factory == nil { panic("factory cannot be nil") }
		
		// New error pattern:
		var factory interface{} = nil
		err := ValidateNotNil(factory, "factory")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "factory cannot be nil")
	})
	
	t.Run("condition assertion replacement", func(t *testing.T) {
		// Old panic pattern:
		// if len(segments) == 0 { panic("segments cannot be empty") }
		
		// New error pattern:
		segments := []string{}
		err := ValidateCondition(len(segments) > 0, "segments cannot be empty")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "segments cannot be empty")
	})
	
	t.Run("unreachable code replacement", func(t *testing.T) {
		// Old panic pattern:
		// default: panic("unreachable code")
		
		// New error pattern:
		handleUnexpectedCase := func(value string) error {
			switch value {
			case "valid1", "valid2":
				return nil
			default:
				return NewInternalError("unexpected enum value: " + value)
			}
		}
		
		err := handleUnexpectedCase("invalid")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected enum value")
	})
}

func TestLibraryErrorHandlingPrinciples(t *testing.T) {
	// Test that demonstrates the library's error handling principles
	
	t.Run("library functions never panic", func(t *testing.T) {
		// All library functions should return errors instead of panicking
		// This ensures applications can always recover gracefully
		
		// Example: Processing invalid input
		processInvalidInput := func(input string) error {
			if input == "" {
				return NewRequiredFieldError("input")
			}
			if input == "invalid" {
				return NewValidationError("input", "contains invalid characters")
			}
			// Internal error condition
			if input == "internal_error" {
				return NewInternalError("unexpected internal state")
			}
			return nil
		}
		
		// All these should return errors, never panic
		assert.Error(t, processInvalidInput(""))
		assert.Error(t, processInvalidInput("invalid"))
		assert.Error(t, processInvalidInput("internal_error"))
		assert.NoError(t, processInvalidInput("valid"))
	})
	
	t.Run("errors provide clear context", func(t *testing.T) {
		// All errors should provide clear context about what went wrong
		
		err1 := NewValidationError("MessageId", "cannot contain spaces")
		assert.Contains(t, err1.Error(), "MessageId")
		assert.Contains(t, err1.Error(), "cannot contain spaces")
		
		err2 := NewParseError("XML decode", "document header", NewInternalError("test"))
		assert.Contains(t, err2.Error(), "XML decode")
		assert.Contains(t, err2.Error(), "document header")
		
		err3 := NewInternalError("factory map is nil")
		assert.Contains(t, err3.Error(), "internal error")
		assert.Contains(t, err3.Error(), "factory map is nil")
	})
}