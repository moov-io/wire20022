package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {
	t.Run("Error method", func(t *testing.T) {
		err := &ValidationError{
			Field:  "MessageId",
			Reason: "cannot be empty",
		}
		expected := `validation failed for field "MessageId": cannot be empty`
		assert.Equal(t, expected, err.Error())
	})

	t.Run("Error method with underlying error", func(t *testing.T) {
		underlying := errors.New("string too long")
		err := &ValidationError{
			Field:  "MessageId",
			Reason: "invalid format",
			Err:    underlying,
		}
		expected := `validation failed for field "MessageId": invalid format: string too long`
		assert.Equal(t, expected, err.Error())
	})

	t.Run("Unwrap method", func(t *testing.T) {
		underlying := errors.New("underlying error")
		err := &ValidationError{
			Field:  "test",
			Reason: "test",
			Err:    underlying,
		}
		assert.Equal(t, underlying, err.Unwrap())
	})

	t.Run("Is method with sentinel errors", func(t *testing.T) {
		err := &ValidationError{Field: "test", Reason: "test"}

		assert.True(t, err.Is(ErrInvalidField))
		assert.True(t, err.Is(ErrRequiredField))
		assert.False(t, err.Is(ErrInvalidXML))
		assert.False(t, err.Is(errors.New("other error")))
	})

	t.Run("errors.Is compatibility", func(t *testing.T) {
		err := &ValidationError{Field: "test", Reason: "test"}

		assert.True(t, errors.Is(err, ErrInvalidField))
		assert.True(t, errors.Is(err, ErrRequiredField))
		assert.False(t, errors.Is(err, ErrInvalidXML))
	})

	t.Run("errors.As compatibility", func(t *testing.T) {
		err := &ValidationError{Field: "TestField", Reason: "test reason"}

		var validationErr *ValidationError
		assert.True(t, errors.As(err, &validationErr))
		assert.Equal(t, "TestField", validationErr.Field)
		assert.Equal(t, "test reason", validationErr.Reason)
	})
}

func TestParseError(t *testing.T) {
	t.Run("Error method with content", func(t *testing.T) {
		underlying := errors.New("unexpected character")
		err := &ParseError{
			Operation: "XML unmarshal",
			Content:   "Document",
			Err:       underlying,
		}
		expected := "XML unmarshal failed for Document: unexpected character"
		assert.Equal(t, expected, err.Error())
	})

	t.Run("Error method without content", func(t *testing.T) {
		underlying := errors.New("invalid syntax")
		err := &ParseError{
			Operation: "JSON marshal",
			Err:       underlying,
		}
		expected := "JSON marshal failed: invalid syntax"
		assert.Equal(t, expected, err.Error())
	})

	t.Run("Unwrap method", func(t *testing.T) {
		underlying := errors.New("underlying error")
		err := &ParseError{
			Operation: "test",
			Err:       underlying,
		}
		assert.Equal(t, underlying, err.Unwrap())
	})

	t.Run("Is method with sentinel errors", func(t *testing.T) {
		err := &ParseError{Operation: "test"}

		assert.True(t, err.Is(ErrInvalidXML))
		assert.True(t, err.Is(ErrInvalidJSON))
		assert.False(t, err.Is(ErrInvalidField))
		assert.False(t, err.Is(errors.New("other error")))
	})

	t.Run("errors.Is compatibility", func(t *testing.T) {
		err := &ParseError{Operation: "test"}

		assert.True(t, errors.Is(err, ErrInvalidXML))
		assert.True(t, errors.Is(err, ErrInvalidJSON))
		assert.False(t, errors.Is(err, ErrFieldNotFound))
	})

	t.Run("errors.As compatibility", func(t *testing.T) {
		err := &ParseError{Operation: "XML parse", Content: "test document"}

		var parseErr *ParseError
		assert.True(t, errors.As(err, &parseErr))
		assert.Equal(t, "XML parse", parseErr.Operation)
		assert.Equal(t, "test document", parseErr.Content)
	})
}

func TestFieldError(t *testing.T) {
	t.Run("Error method", func(t *testing.T) {
		underlying := errors.New("not found")
		err := &FieldError{
			Path:      "Header.ID",
			Operation: "get",
			Err:       underlying,
		}
		expected := "field get Header.ID failed: not found"
		assert.Equal(t, expected, err.Error())
	})

	t.Run("Unwrap method", func(t *testing.T) {
		underlying := errors.New("underlying error")
		err := &FieldError{
			Path:      "test",
			Operation: "set",
			Err:       underlying,
		}
		assert.Equal(t, underlying, err.Unwrap())
	})

	t.Run("Is method with sentinel errors", func(t *testing.T) {
		err := &FieldError{Path: "test", Operation: "get"}

		assert.True(t, err.Is(ErrFieldNotFound))
		assert.True(t, err.Is(ErrIndexOutOfBounds))
		assert.False(t, err.Is(ErrInvalidField))
		assert.False(t, err.Is(errors.New("other error")))
	})

	t.Run("errors.Is compatibility", func(t *testing.T) {
		err := &FieldError{Path: "test", Operation: "get"}

		assert.True(t, errors.Is(err, ErrFieldNotFound))
		assert.True(t, errors.Is(err, ErrIndexOutOfBounds))
		assert.False(t, errors.Is(err, ErrInvalidField))
	})

	t.Run("errors.As compatibility", func(t *testing.T) {
		err := &FieldError{Path: "Header.ID", Operation: "set"}

		var fieldErr *FieldError
		assert.True(t, errors.As(err, &fieldErr))
		assert.Equal(t, "Header.ID", fieldErr.Path)
		assert.Equal(t, "set", fieldErr.Operation)
	})
}

func TestSentinelErrors(t *testing.T) {
	t.Run("sentinel errors are distinct", func(t *testing.T) {
		sentinelErrors := []error{
			ErrInvalidField,
			ErrRequiredField,
			ErrInvalidXML,
			ErrInvalidJSON,
			ErrUnknownNamespace,
			ErrInvalidVersion,
			ErrFieldNotFound,
			ErrIndexOutOfBounds,
		}

		// Each sentinel error should be different from all others
		for i, err1 := range sentinelErrors {
			for j, err2 := range sentinelErrors {
				if i != j {
					assert.False(t, errors.Is(err1, err2),
						"sentinel errors should be distinct: %v should not equal %v", err1, err2)
				}
			}
		}
	})

	t.Run("sentinel errors have proper messages", func(t *testing.T) {
		// Verify error messages follow Go conventions (no capitals, no punctuation)
		testCases := map[error]string{
			ErrInvalidField:     "invalid field",
			ErrRequiredField:    "required field missing",
			ErrInvalidXML:       "invalid XML format",
			ErrInvalidJSON:      "invalid JSON format",
			ErrUnknownNamespace: "unknown XML namespace",
			ErrInvalidVersion:   "invalid version",
			ErrFieldNotFound:    "field not found",
			ErrIndexOutOfBounds: "array index out of bounds",
		}

		for err, expectedMsg := range testCases {
			assert.Equal(t, expectedMsg, err.Error())
		}
	})
}

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
