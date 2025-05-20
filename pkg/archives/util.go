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

type Document struct {
	Attrs []xml.Attr `xml:",any,attr"`
}
type IOSDocument interface {
	Validate() error
}
type DocumentFactory func() IOSDocument

func DocumentFrom(data []byte, factoryMap map[string]DocumentFactory) (IOSDocument, error) {
	var root Document
	if err := xml.Unmarshal(data, &root); err != nil {
		return nil, fmt.Errorf("XML decode error: %w", err)
	}

	var xmlns string
	for _, attr := range root.Attrs {
		if attr.Name.Local == "xmlns" && attr.Name.Space == "" {
			xmlns = attr.Value
			break
		}
	}

	if xmlns == "" {
		return nil, fmt.Errorf("no xmlns found")
	}

	// Lookup model factory
	factory, ok := factoryMap[xmlns]
	if !ok {
		return nil, fmt.Errorf("unknown namespace: %s", xmlns)
	}

	// Instantiate and unmarshal into actual model
	doc := factory()
	if err := xml.Unmarshal(data, doc); err != nil {
		return nil, fmt.Errorf("XML unmarshal to model failed: %w", err)
	}

	return doc, nil
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

	// Walk the path up to the second-to-last field
	for i := 0; i < len(segments)-1; i++ {
		field := v.FieldByName(segments[i])
		if !field.IsValid() {
			return fmt.Errorf("field %s not found", segments[i])
		}

		// If pointer, initialize if nil
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

	// Now set the last field
	last := segments[len(segments)-1]
	field := v.FieldByName(last)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", last)
	}
	err := setValue(field, value)
	if err != nil {
		return fmt.Errorf("cannot convert value to field type %s", field.Type())
	}

	return nil
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
				if elementType.Kind() == reflect.Struct {
					newSlice := reflect.MakeSlice(v.Type(), index+1, index+1)
					reflect.Copy(newSlice, v)
					newStruct := reflect.New(elementType).Elem()
					newSlice.Index(index).Set(newStruct)
					v.Set(newSlice)
				} else {
					return fmt.Errorf("element type is not a struct")
				}
			}
			v = v.Index(index)
		} else {
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
		return fmt.Errorf("cannot convert value to field type %s", field.Type())
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
	} else if val.Type().Kind() == reflect.String && v.Type().Kind() == reflect.String {
		if strVal, ok := val.Interface().(string); ok {
			convertedVal := reflect.ValueOf(strVal).Convert(v.Type())
			v.Set(convertedVal)
		} else {
			return fmt.Errorf("value is not a string, cannot convert to field type %s", v.Type())
		}
	} else {
		return fmt.Errorf("cannot convert value to field type %s", v.Type())
	}

	return nil
}
func CopyDocumentValueToMessage(from IOSDocument, fromPah string, to any, toPath string) {
	if from == nil || fromPah == "" || toPath == "" {
		return
	}
	_, value := GetElement(from, fromPah)
	if value == nil {
		return
	}

	err := SetElementToModel(to, toPath, value)
	if err != nil {
		return
	}
}

func CopyMessageValueToDocument(from any, fromPath string, to IOSDocument, toPath string) {
	if from == nil || fromPath == "" || toPath == "" {
		return
	}
	_, value := GetElement(from, fromPath)
	if isEmpty(value) {
		return
	}

	err := SetElementToDocument(to, toPath, value)
	if err != nil {
		return
	}
}

func isReflectValueNil(v reflect.Value) bool {
	// First check if the reflect.Value is valid
	if !v.IsValid() {
		return true
	}

	// Then check if the kind supports nil and if it's actually nil
	switch v.Kind() {
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
