package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("MessageId", "cannot be empty")

	assert.Equal(t, "MessageId", err.Field)
	assert.Equal(t, "cannot be empty", err.Reason)
	assert.Nil(t, err.Err)

	expected := `validation failed for field "MessageId": cannot be empty`
	assert.Equal(t, expected, err.Error())
}

func TestNewValidationErrorWithCause(t *testing.T) {
	cause := errors.New("underlying cause")
	err := NewValidationErrorWithCause("Amount", "invalid format", cause)

	assert.Equal(t, "Amount", err.Field)
	assert.Equal(t, "invalid format", err.Reason)
	assert.Equal(t, cause, err.Err)

	expected := `validation failed for field "Amount": invalid format: underlying cause`
	assert.Equal(t, expected, err.Error())
}

func TestNewParseError(t *testing.T) {
	cause := errors.New("unexpected character")
	err := NewParseError("XML unmarshal", "Document", cause)

	assert.Equal(t, "XML unmarshal", err.Operation)
	assert.Equal(t, "Document", err.Content)
	assert.Equal(t, cause, err.Err)

	expected := "XML unmarshal failed for Document: unexpected character"
	assert.Equal(t, expected, err.Error())
}

func TestNewParseErrorSimple(t *testing.T) {
	cause := errors.New("invalid syntax")
	err := NewParseErrorSimple("JSON marshal", cause)

	assert.Equal(t, "JSON marshal", err.Operation)
	assert.Equal(t, "", err.Content)
	assert.Equal(t, cause, err.Err)

	expected := "JSON marshal failed: invalid syntax"
	assert.Equal(t, expected, err.Error())
}

func TestNewFieldError(t *testing.T) {
	err := NewFieldError("Header.ID", "get", ErrFieldNotFound)

	assert.Equal(t, "Header.ID", err.Path)
	assert.Equal(t, "get", err.Operation)
	assert.Equal(t, ErrFieldNotFound, err.Err)

	expected := "field get Header.ID failed: field not found"
	assert.Equal(t, expected, err.Error())
}

func TestNewRequiredFieldError(t *testing.T) {
	err := NewRequiredFieldError("MessageId")

	assert.Equal(t, "MessageId", err.Field)
	assert.Equal(t, "is required", err.Reason)
	assert.Equal(t, ErrRequiredField, err.Err)

	expected := `validation failed for field "MessageId": is required: required field missing`
	assert.Equal(t, expected, err.Error())

	// Test that it works with errors.Is
	assert.True(t, errors.Is(err, ErrRequiredField))
}

func TestNewInvalidFieldError(t *testing.T) {
	err := NewInvalidFieldError("Amount", "must be positive")

	assert.Equal(t, "Amount", err.Field)
	assert.Equal(t, "must be positive", err.Reason)
	assert.Equal(t, ErrInvalidField, err.Err)

	expected := `validation failed for field "Amount": must be positive: invalid field`
	assert.Equal(t, expected, err.Error())

	// Test that it works with errors.Is
	assert.True(t, errors.Is(err, ErrInvalidField))
}

func TestWrapValidationError(t *testing.T) {
	originalErr := errors.New("json: invalid character")
	err := WrapValidationError("RequestData", "invalid JSON format", originalErr)

	assert.Equal(t, "RequestData", err.Field)
	assert.Equal(t, "invalid JSON format", err.Reason)
	assert.NotNil(t, err.Err)

	// Test that the original error is wrapped and can be unwrapped
	assert.True(t, errors.Is(err, originalErr))

	expected := `validation failed for field "RequestData": invalid JSON format: json: invalid character`
	assert.Equal(t, expected, err.Error())
}

func TestConstructorReturnTypes(t *testing.T) {
	// Verify that constructors return concrete types, not error interfaces

	t.Run("NewValidationError returns *ValidationError", func(t *testing.T) {
		err := NewValidationError("test", "test")
		// Verify it's the concrete type by checking fields
		assert.Equal(t, "test", err.Field)
		assert.Equal(t, "test", err.Reason)
		// Verify it implements error interface
		var _ error = err
	})

	t.Run("NewParseError returns *ParseError", func(t *testing.T) {
		err := NewParseError("test", "test", errors.New("test"))
		// Verify it's the concrete type by checking fields
		assert.Equal(t, "test", err.Operation)
		assert.Equal(t, "test", err.Content)
		// Verify it implements error interface
		var _ error = err
	})

	t.Run("NewFieldError returns *FieldError", func(t *testing.T) {
		err := NewFieldError("test", "test", errors.New("test"))
		// Verify it's the concrete type by checking fields
		assert.Equal(t, "test", err.Path)
		assert.Equal(t, "test", err.Operation)
		// Verify it implements error interface
		var _ error = err
	})
}

func TestErrorChaining(t *testing.T) {
	// Test complex error chaining scenarios

	t.Run("validation error with parse error cause", func(t *testing.T) {
		parseErr := NewParseErrorSimple("JSON unmarshal", errors.New("invalid character"))
		validationErr := NewValidationErrorWithCause("RequestData", "malformed JSON", parseErr)

		// Should be able to check for both error types
		assert.True(t, errors.Is(validationErr, ErrInvalidJSON))

		// Should be able to extract the parse error
		var extractedParseErr *ParseError
		assert.True(t, errors.As(validationErr, &extractedParseErr))
		assert.Equal(t, "JSON unmarshal", extractedParseErr.Operation)
	})

	t.Run("field error with validation error cause", func(t *testing.T) {
		validationErr := NewInvalidFieldError("Amount", "must be numeric")
		fieldErr := NewFieldError("Transaction.Amount", "set", validationErr)

		// Should be able to check for both error types
		assert.True(t, errors.Is(fieldErr, ErrInvalidField))
		assert.True(t, errors.Is(fieldErr, ErrFieldNotFound)) // FieldError.Is() behavior

		// Should be able to extract the validation error
		var extractedValidationErr *ValidationError
		assert.True(t, errors.As(fieldErr, &extractedValidationErr))
		assert.Equal(t, "Amount", extractedValidationErr.Field)
	})
}
