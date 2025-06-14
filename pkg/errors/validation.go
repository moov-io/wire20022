package errors

import (
	"errors"
)

// JoinValidationErrors combines multiple validation errors into a single error.
// This is useful when validating multiple fields and you want to return all
// validation failures at once instead of stopping at the first error.
//
// Uses errors.Join() from Go 1.20+ for proper error chaining and unwrapping.
//
// Example:
//
//	var errs []error
//	if model.MessageId == "" {
//	    errs = append(errs, NewRequiredFieldError("MessageId"))
//	}
//	if model.Amount < 0 {
//	    errs = append(errs, NewInvalidFieldError("Amount", "must be positive"))
//	}
//	if len(errs) > 0 {
//	    return JoinValidationErrors(errs...)
//	}
func JoinValidationErrors(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return errors.Join(errs...)
}

// CollectValidationErrors provides a convenient way to collect multiple validation
// errors and return them as a joined error. Use this when you want to validate
// multiple fields and return all errors at once.
//
// Example:
//
//	collector := NewValidationErrorCollector()
//
//	if model.MessageId == "" {
//	    collector.Add(NewRequiredFieldError("MessageId"))
//	}
//	if model.Amount < 0 {
//	    collector.Add(NewInvalidFieldError("Amount", "must be positive"))
//	}
//
//	return collector.Error() // Returns nil if no errors, or joined error if any
type ValidationErrorCollector struct {
	errors []error
}

// NewValidationErrorCollector creates a new error collector for validation errors.
func NewValidationErrorCollector() *ValidationErrorCollector {
	return &ValidationErrorCollector{
		errors: make([]error, 0),
	}
}

// Add appends a validation error to the collector.
// Nil errors are ignored.
func (c *ValidationErrorCollector) Add(err error) {
	if err != nil {
		c.errors = append(c.errors, err)
	}
}

// AddFieldError is a convenience method to add a field validation error.
func (c *ValidationErrorCollector) AddFieldError(field, reason string) {
	c.Add(NewValidationError(field, reason))
}

// AddRequiredField is a convenience method to add a required field error.
func (c *ValidationErrorCollector) AddRequiredField(field string) {
	c.Add(NewRequiredFieldError(field))
}

// AddInvalidField is a convenience method to add an invalid field error.
func (c *ValidationErrorCollector) AddInvalidField(field, reason string) {
	c.Add(NewInvalidFieldError(field, reason))
}

// HasErrors returns true if any validation errors have been collected.
func (c *ValidationErrorCollector) HasErrors() bool {
	return len(c.errors) > 0
}

// Count returns the number of validation errors collected.
func (c *ValidationErrorCollector) Count() int {
	return len(c.errors)
}

// Error returns the collected errors as a single joined error.
// Returns nil if no errors were collected.
func (c *ValidationErrorCollector) Error() error {
	return JoinValidationErrors(c.errors...)
}

// Errors returns a copy of all collected errors.
func (c *ValidationErrorCollector) Errors() []error {
	if len(c.errors) == 0 {
		return nil
	}
	result := make([]error, len(c.errors))
	copy(result, c.errors)
	return result
}
