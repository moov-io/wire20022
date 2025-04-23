// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

/*
	Validate Function will run validate functions of the their data fields
	A lot of types have special Validate function (all of the implementation are in types.go file)
	The goal of this validate logic is to call the validate functions

	All of the structures in this package should have validate function
	If there is not special validate() function, should add dummy validate() function
	Example:
    	func (r AccountSwitchTerminationSwitchV01) Validate() error {
			return utils.Validate(&r)
		}

	With this logic, we will go validation check about special Iso20022Message
*/

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	DefaultValidateFunction = "Validate"
)

func getTypeName(value string) string {
	values := strings.Split(value, ".")
	if len(values) > 1 {
		values := strings.Split(values[1], " ")
		return values[0]
	} else {
		return values[0]
	}
}

func validateCallbackByValue(data reflect.Value) error {
	method := data.MethodByName(DefaultValidateFunction)
	if method.IsValid() {
		response := method.Call(nil)
		if len(response) > 0 {
			err := response[0]
			if !err.IsNil() {
				typeName := getTypeName(data.String())
				if len(typeName) > 0 {
					errValue, ok := err.Interface().(error)
					if !ok {
						return fmt.Errorf("failed to assert type as error: %v", err.Interface())
					}
					errStr := errValue.Error()
					if !strings.Contains(errStr, ")") {
						errStr = errStr + " (" + typeName + ")"
					} else {
						errStr = errStr[:len(errStr)-1] + ", " + typeName + ")"
					}
					return errors.New(errStr)
				}
				if errValue, ok := err.Interface().(error); ok {
					return errValue
				}
				return fmt.Errorf("type assertion failed for: %v", err.Interface())
			}
		}
	}
	return nil
}

// to validate interface
func Validate(r interface{}) error {
	fields := reflect.ValueOf(r).Elem()

	for i := 0; i < fields.NumField(); i++ {
		fieldData := fields.Field(i)

		// nolint:exhaustive // Reason: These cases are intentionally not handled here.
		switch kind := fieldData.Kind(); kind {
		case reflect.Slice:
			for j := 0; j < fieldData.Len(); j++ {
				if err := validateCallbackByValue(fieldData.Index(j)); err != nil {
					return err
				}
			}

		case reflect.Map:
			for _, key := range fieldData.MapKeys() {
				if err := validateCallbackByValue(fieldData.MapIndex(key)); err != nil {
					return err
				}
			}

		case reflect.Ptr:
			if !fieldData.IsNil() {
				if err := validateCallbackByValue(fieldData); err != nil {
					return err
				}
			}

		case reflect.Struct, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
			// Handle additional types
			if err := validateCallbackByValue(fieldData); err != nil {
				return err
			}

		default: // Catch any unhandled types
			// Log or handle unsupported types if necessary
			return fmt.Errorf("unsupported field type: %v", kind)
		}
	}

	return nil
}
