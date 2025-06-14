package errors

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinValidationErrors(t *testing.T) {
	t.Run("no errors returns nil", func(t *testing.T) {
		err := JoinValidationErrors()
		assert.NoError(t, err)
	})

	t.Run("single error returns that error", func(t *testing.T) {
		originalErr := NewValidationError("MessageId", "cannot be empty")
		err := JoinValidationErrors(originalErr)
		assert.Equal(t, originalErr, err)
	})

	t.Run("multiple errors are joined", func(t *testing.T) {
		err1 := NewValidationError("MessageId", "cannot be empty")
		err2 := NewValidationError("Amount", "must be positive")

		joinedErr := JoinValidationErrors(err1, err2)
		assert.Error(t, joinedErr)

		// Test that both errors are present using errors.Is
		assert.True(t, errors.Is(joinedErr, err1))
		assert.True(t, errors.Is(joinedErr, err2))

		// Test that error message contains both
		errMsg := joinedErr.Error()
		assert.Contains(t, errMsg, "MessageId")
		assert.Contains(t, errMsg, "Amount")
	})

	t.Run("joined errors can be unwrapped", func(t *testing.T) {
		err1 := NewValidationError("Field1", "error 1")
		err2 := NewValidationError("Field2", "error 2")
		err3 := NewValidationError("Field3", "error 3")

		joinedErr := JoinValidationErrors(err1, err2, err3)

		// Test individual error extraction
		var validationErr1 *ValidationError

		assert.True(t, errors.As(joinedErr, &validationErr1))
		// Note: errors.As only finds the first matching error in a chain
		// For multiple errors of the same type, we need to check differently

		// Verify all original errors are findable
		assert.True(t, errors.Is(joinedErr, err1))
		assert.True(t, errors.Is(joinedErr, err2))
		assert.True(t, errors.Is(joinedErr, err3))
	})
}

func TestValidationErrorCollector(t *testing.T) {
	t.Run("new collector has no errors", func(t *testing.T) {
		collector := NewValidationErrorCollector()

		assert.False(t, collector.HasErrors())
		assert.Equal(t, 0, collector.Count())
		assert.NoError(t, collector.Error())
		assert.Nil(t, collector.Errors())
	})

	t.Run("add single error", func(t *testing.T) {
		collector := NewValidationErrorCollector()
		err := NewValidationError("MessageId", "cannot be empty")

		collector.Add(err)

		assert.True(t, collector.HasErrors())
		assert.Equal(t, 1, collector.Count())
		assert.Equal(t, err, collector.Error())

		errs := collector.Errors()
		assert.Len(t, errs, 1)
		assert.Equal(t, err, errs[0])
	})

	t.Run("add multiple errors", func(t *testing.T) {
		collector := NewValidationErrorCollector()
		err1 := NewValidationError("MessageId", "cannot be empty")
		err2 := NewValidationError("Amount", "must be positive")

		collector.Add(err1)
		collector.Add(err2)

		assert.True(t, collector.HasErrors())
		assert.Equal(t, 2, collector.Count())

		joinedErr := collector.Error()
		assert.Error(t, joinedErr)
		assert.True(t, errors.Is(joinedErr, err1))
		assert.True(t, errors.Is(joinedErr, err2))

		errs := collector.Errors()
		assert.Len(t, errs, 2)
		assert.Contains(t, errs, err1)
		assert.Contains(t, errs, err2)
	})

	t.Run("add nil error is ignored", func(t *testing.T) {
		collector := NewValidationErrorCollector()

		collector.Add(nil)

		assert.False(t, collector.HasErrors())
		assert.Equal(t, 0, collector.Count())
		assert.NoError(t, collector.Error())
	})

	t.Run("convenience methods work correctly", func(t *testing.T) {
		collector := NewValidationErrorCollector()

		collector.AddFieldError("Field1", "custom reason")
		collector.AddRequiredField("Field2")
		collector.AddInvalidField("Field3", "must be numeric")

		assert.True(t, collector.HasErrors())
		assert.Equal(t, 3, collector.Count())

		joinedErr := collector.Error()
		errMsg := joinedErr.Error()

		assert.Contains(t, errMsg, "Field1")
		assert.Contains(t, errMsg, "custom reason")
		assert.Contains(t, errMsg, "Field2")
		assert.Contains(t, errMsg, "is required")
		assert.Contains(t, errMsg, "Field3")
		assert.Contains(t, errMsg, "must be numeric")
	})

	t.Run("errors slice is a copy", func(t *testing.T) {
		collector := NewValidationErrorCollector()
		err1 := NewValidationError("Field1", "error 1")

		collector.Add(err1)
		errs := collector.Errors()

		// Modify the returned slice
		errs[0] = NewValidationError("Modified", "modified")

		// Original collector should not be affected
		originalErrs := collector.Errors()
		assert.Equal(t, err1, originalErrs[0])
		assert.NotEqual(t, "Modified", originalErrs[0].(*ValidationError).Field)
	})
}

func TestValidationErrorCollectorRealWorldUsage(t *testing.T) {
	// Simulate a real validation scenario
	type MessageModel struct {
		MessageId string
		Amount    float64
		Currency  string
		ToAddress string
	}

	validateMessage := func(model MessageModel) error {
		collector := NewValidationErrorCollector()

		if model.MessageId == "" {
			collector.AddRequiredField("MessageId")
		}
		if strings.Contains(model.MessageId, " ") {
			collector.AddInvalidField("MessageId", "cannot contain spaces")
		}
		if model.Amount <= 0 {
			collector.AddInvalidField("Amount", "must be positive")
		}
		if model.Currency == "" {
			collector.AddRequiredField("Currency")
		}
		if len(model.Currency) != 3 {
			collector.AddInvalidField("Currency", "must be 3 characters (ISO code)")
		}
		if model.ToAddress == "" {
			collector.AddRequiredField("ToAddress")
		}

		return collector.Error()
	}

	t.Run("valid message has no errors", func(t *testing.T) {
		model := MessageModel{
			MessageId: "MSG123",
			Amount:    100.00,
			Currency:  "USD",
			ToAddress: "123 Main St",
		}

		err := validateMessage(model)
		assert.NoError(t, err)
	})

	t.Run("invalid message returns multiple errors", func(t *testing.T) {
		model := MessageModel{
			MessageId: "MSG 123", // Contains space
			Amount:    -100.00,   // Negative
			Currency:  "US",      // Too short
			ToAddress: "",        // Empty
		}

		err := validateMessage(model)
		assert.Error(t, err)

		errMsg := err.Error()
		// Check that all validation errors are present
		assert.Contains(t, errMsg, "MessageId")
		assert.Contains(t, errMsg, "cannot contain spaces")
		assert.Contains(t, errMsg, "Amount")
		assert.Contains(t, errMsg, "must be positive")
		assert.Contains(t, errMsg, "Currency")
		assert.Contains(t, errMsg, "must be 3 characters")
		assert.Contains(t, errMsg, "ToAddress")
		assert.Contains(t, errMsg, "is required")
	})

	t.Run("can extract individual validation errors", func(t *testing.T) {
		model := MessageModel{
			MessageId: "",     // Missing
			Amount:    -50.00, // Negative
		}

		err := validateMessage(model)
		assert.Error(t, err)

		// Can check for specific error types
		var validationErr *ValidationError
		assert.True(t, errors.As(err, &validationErr))

		// Can check for specific sentinel errors
		assert.True(t, errors.Is(err, ErrRequiredField))
		assert.True(t, errors.Is(err, ErrInvalidField))
	})
}
