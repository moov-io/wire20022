package Archive

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Match struct {
	SrcPath string
	DstPath string
}
type ElementMap struct {
	SrcElement string
	DstElement string
}

type Document struct {
	Attrs []xml.Attr `xml:",any,attr"`
}
type ISODocument interface {
	Validate() error
}
type DocumentFactory func() ISODocument

func DocumentFrom(data []byte, factoryMap map[string]DocumentFactory) (ISODocument, string, error) {
	var root Document
	if err := xml.Unmarshal(data, &root); err != nil {
		return nil, "", fmt.Errorf("XML decode error: %w", err)
	}

	var xmlns string
	for _, attr := range root.Attrs {
		if attr.Name.Local == "xmlns" && attr.Name.Space == "" {
			xmlns = attr.Value
			break
		}
	}

	if xmlns == "" {
		return nil, "", fmt.Errorf("no xmlns found")
	}

	// Lookup model factory
	factory, ok := factoryMap[xmlns]
	if !ok {
		return nil, "", fmt.Errorf("unknown namespace: %s", xmlns)
	}

	// Instantiate and unmarshal into actual model
	doc := factory()
	if err := xml.Unmarshal(data, doc); err != nil {
		return nil, "", fmt.Errorf("XML unmarshal to model failed: %w", err)
	}

	return doc, xmlns, nil
}
func GetElement(item any, path string) (reflect.Type, any) {
	if item == nil || path == "" {
		return nil, nil
	}

	v := reflect.ValueOf(item)

	// Dereference pointer
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Make sure we're starting from a struct
	if v.Kind() != reflect.Struct {
		return nil, nil
	}

	// Walk the path
	segments := strings.Split(path, ".")
	for _, segment := range segments {
		// Check if the segment is an array or slice access
		// e.g., "RptgReq[0]"
		re := regexp.MustCompile(`^(\w+)\[(\d+)\]$`)
		matches := re.FindStringSubmatch(segment)
		if matches != nil {
			fieldName := matches[1]
			index, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil, nil // Invalid index
			}
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			// Get the field by name
			if isReflectValueNil(v) {
				return nil, nil // Field is nil
			}
			v = v.FieldByName(fieldName)
			if !v.IsValid() || (v.Kind() != reflect.Slice && v.Kind() != reflect.Array) {
				return nil, nil // Field not found or not a slice/array
			}

			// Check if the index is within bounds
			if index < 0 || index >= v.Len() {
				return nil, nil // Index out of bounds
			}

			// Access the element at the specified index
			v = v.Index(index)
		} else {
			// Regular field access
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() != reflect.Struct {
				return nil, nil
			}
			if isReflectValueNil(v) {
				return nil, fmt.Errorf("field %s is nil", segment)
			}
			v = v.FieldByName(segment)
			if !v.IsValid() {
				return nil, nil // Field not found
			}
		}
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Type(), v.Interface()
}
func SetElementToModel(item any, path string, value any) error {
	return SetElementToDocument(item, path, value)
}
func SetElementToDocument(item any, path string, value any) error {
	if item == nil || path == "" {
		return fmt.Errorf("invalid input")
	}
	v := reflect.ValueOf(item)
	// item must be a pointer to a struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("item must be a pointer to a struct")
	}
	// Dereference to struct
	v = v.Elem()
	segments := strings.Split(path, ".")
	for i := 0; i < len(segments)-1; i++ {
		re := regexp.MustCompile(`^(\w+)\[(\d+)\]$`)
		matches := re.FindStringSubmatch(segments[i])
		if matches != nil {
			fieldName := matches[1]
			index, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil // Invalid index
			}
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if isReflectValueNil(v) {
				return nil // Field is nil
			}
			v = v.FieldByName(fieldName)
			if !v.IsValid() || (v.Kind() != reflect.Slice && v.Kind() != reflect.Array) {
				return nil // Field not found or not a slice/array
			}
			if isEmpty(v) {
				newArray := reflect.New(v.Type()).Elem()
				v.Set(newArray)
			}
			if index >= v.Len() {
				elementType := v.Type().Elem()

				// Handle pointer types
				if elementType.Kind() == reflect.Ptr {
					elementType = elementType.Elem()
				}

				// Ensure the underlying type is a struct
				if elementType.Kind() == reflect.Struct {
					// Create a new slice with the required length
					newSlice := reflect.MakeSlice(v.Type(), index+1, index+1)
					reflect.Copy(newSlice, v)

					// Initialize new elements in the slice
					for i := v.Len(); i <= index; i++ {
						newStruct := reflect.New(elementType).Elem()
						if v.Type().Elem().Kind() == reflect.Ptr {
							// If the slice holds pointers, set a pointer to the new struct
							newSlice.Index(i).Set(newStruct.Addr())
						} else {
							// Otherwise, set the struct directly
							newSlice.Index(i).Set(newStruct)
						}
					}

					// Replace the old slice with the new slice
					v.Set(newSlice)
				} else {
					return fmt.Errorf("element type is not a struct or pointer to a struct")
				}
			}
			if index < v.Len() {
				v = v.Index(index)
			}
		} else {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if isReflectValueNil(v) {
				return fmt.Errorf("field %s is nil", segments[i])
			}
			field := v.FieldByName(segments[i])
			if !field.IsValid() {
				return fmt.Errorf("field %s not found", segments[i])
			}
			if !field.IsValid() {
				return fmt.Errorf("field %s not found", segments[i])
			}
			if field.Kind() == reflect.Ptr {
				if field.IsNil() {
					field.Set(reflect.New(field.Type().Elem()))
				}
				field = field.Elem()
			}
			// Move deeper
			if field.Kind() != reflect.Struct {
				return fmt.Errorf("field %s is not a struct", segments[i])
			}

			v = field
		}
	}
	// Now set the last field
	last := segments[len(segments)-1]
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if isReflectValueNil(v) {
		return fmt.Errorf("field %s is nil", last)
	}
	field := v.FieldByName(last)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", last)
	}
	if !field.CanSet() {
		return fmt.Errorf("field %s cannot be set (may be unexported)", last)
	}

	// Convert value if necessary
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		field = field.Elem()
	}
	err := setValue(field, value)
	if err != nil {
		return err
	}
	return nil
}
func setValue(v reflect.Value, value any) error {
	if !v.CanSet() {
		return fmt.Errorf("cannot set value")
	}
	val := reflect.ValueOf(value)
	if val.Type().ConvertibleTo(v.Type()) {
		v.Set(val.Convert(v.Type()))
		if hasValidateMethod(v) {
			method := v.MethodByName("Validate")
			if method.IsValid() && method.Type().NumIn() == 0 && method.Type().NumOut() == 1 {
				// Call the Validate method
				results := method.Call(nil) //nolint:forbidigo
				if len(results) == 1 && !results[0].IsNil() {
					validationErr, ok := results[0].Interface().(error)
					if ok {
						return validationErr
					}
					return fmt.Errorf("%v", results[0].Interface()) // Fallback for non-error types
				}
			}
		}
	} else if val.Type().Kind() == reflect.String && v.Type().Kind() == reflect.String {
		if strVal, ok := val.Interface().(string); ok {
			convertedVal := reflect.ValueOf(strVal).Convert(v.Type())
			v.Set(convertedVal)
			if hasValidateMethod(v) {
				method := v.MethodByName("Validate")
				if method.IsValid() && method.Type().NumIn() == 0 && method.Type().NumOut() == 1 {
					// Call the Validate method
					results := method.Call(nil) //nolint:forbidigo
					if len(results) == 1 && !results[0].IsNil() {
						validationErr, ok := results[0].Interface().(error)
						if ok {
							return validationErr
						}
						return fmt.Errorf("%v", results[0].Interface()) // Fallback for non-error types
					}
				}
			}
		} else {
			return fmt.Errorf("value is not a string, cannot convert to field type %s", v.Type())
		}
	} else {
		return fmt.Errorf("cannot convert value to field type %s", v.Type())
	}

	return nil
}

func hasValidateMethod(v reflect.Value) bool {
	// Get the type of the value
	t := v.Type()

	// Check if the method "Validate" exists
	method, exists := t.MethodByName("Validate")
	if !exists {
		return false
	}

	// Ensure the method has the correct signature (e.g., no parameters and returns an error)
	if method.Type.NumIn() == 1 && method.Type.NumOut() == 1 && method.Type.Out(0) == reflect.TypeOf((*error)(nil)).Elem() {
		return true
	}

	return false
}
func CopyDocumentValueToMessage(from any, fromPah string, to any, toPath string) {
	if from == nil || fromPah == "" || toPath == "" {
		return
	}
	_, value := GetElement(from, fromPah)
	if isEmpty(value) {
		return
	}
	if value == nil {
		return
	}

	err := SetElementToModel(to, toPath, value)
	if err != nil {
		return
	}
}

func CopyMessageValueToDocument(from any, fromPath string, to ISODocument, toPath string) error {
	if from == nil || fromPath == "" || toPath == "" {
		return fmt.Errorf("invalid input")
	}
	_, value := GetElement(from, fromPath)
	if isEmpty(value) {
		return nil
	}

	err := SetElementToDocument(to, toPath, value)
	if err != nil {
		return fmt.Errorf("failed to set %s: %w", fromPath, err)
	}
	return nil
}

func RemakeMapping(from any, modelMap map[string]any, toModel bool) map[string]string {
	newMap := make(map[string]string)
	for k, v := range modelMap {
		switch v := v.(type) {
		case string:
			newMap[k] = v
		case map[string]string:
			src, dst := seperateKeyAndValue(k, ":")
			if src == "" || dst == "" {
				continue
			}
			targetPath := strings.TrimSpace(dst)
			if toModel {
				targetPath = strings.TrimSpace(src)
			}
			_, val := GetElement(from, targetPath)
			if val == nil {
				continue
			}

			// Check if val is an array or slice
			valValue := reflect.ValueOf(val)
			if valValue.Kind() != reflect.Array && valValue.Kind() != reflect.Slice {
				continue
			}
			// Get the length of the array or slice
			length := valValue.Len()

			// Iterate over the slice
			for i := 0; i < length; i++ {
				for k1, v1 := range v {
					newMap[fmt.Sprintf("%s[%d].%s", src, i, k1)] = fmt.Sprintf("%s[%d].%s", dst, i, v1)
				}
			}
		}
	}
	return newMap
}
func seperateKeyAndValue(src string, separate string) (string, string) {
	parts := strings.Split(src, separate)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	}
	return "", ""
}

func isReflectValueNil(v reflect.Value) bool {
	// First check if the reflect.Value is valid
	if !v.IsValid() {
		return true
	}

	// Then check if the kind supports nil and if it's actually nil
	switch v.Kind() {
	case reflect.Invalid, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128, reflect.Array, reflect.String, reflect.Struct,
		reflect.UnsafePointer:
		return false

	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
func isEmpty(val interface{}) bool {
	switch v := val.(type) {
	case nil:
		return true
	case string:
		return v == ""
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int() == 0
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(v).Uint() == 0
	case float32, float64:
		return reflect.ValueOf(v).Float() == 0
	case bool:
		return !v
	case time.Time:
		return v.IsZero()
	default:
		// Use reflect for fallback
		return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
	}
}

func WriteXMLTo(filePath string, data []byte) error {

	if ext := filepath.Ext(filePath); ext != ".xml" {
		return fmt.Errorf("invalid file extension %q, must be .xml", ext)
	}

	// Write file with atomic replacement
	tempFile := filePath + ".tmp"
	err := os.WriteFile(tempFile, data, 0600)
	if err != nil {
		return fmt.Errorf("temporary file write failed: %w", err)
	}

	// Atomic rename for crash safety
	if err := os.Rename(tempFile, filePath); err != nil {
		// Clean up temp file if rename fails
		if err := os.Remove(tempFile); err != nil && !os.IsNotExist(err) {
			log.Printf("failed to remove temp file %q: %v", tempFile, err)
		}
		return fmt.Errorf("file rename failed: %w", err)
	}

	return nil
}
func WriteXMLToGenerate(filePath string, data []byte) error {
	// Ensure directory exists with proper permissions
	if err := os.MkdirAll("generated", 0750); err != nil && !os.IsExist(err) {
		return fmt.Errorf("directory creation failed: %w", err)
	}

	// Construct full file path
	xmlFileName := filepath.Join("generated", filePath)

	// Validate file extension
	if ext := filepath.Ext(xmlFileName); ext != ".xml" {
		return fmt.Errorf("invalid file extension %q, must be .xml", ext)
	}

	// Write file with atomic replacement
	tempFile := xmlFileName + ".tmp"
	err := os.WriteFile(tempFile, data, 0600)
	if err != nil {
		return fmt.Errorf("temporary file write failed: %w", err)
	}

	// Atomic rename for crash safety
	if err := os.Rename(tempFile, xmlFileName); err != nil {
		// Clean up temp file if rename fails
		if err := os.Remove(tempFile); err != nil && !os.IsNotExist(err) {
			log.Printf("failed to remove temp file %q: %v", tempFile, err)
		}
		return fmt.Errorf("file rename failed: %w", err)
	}

	return nil
}
func ReadXMLFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", filename, err)
	}
	return data, nil
}
