package Archive

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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
func SetElement(item any, path string, value any) error {
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
	if !field.CanSet() {
		return fmt.Errorf("field %s cannot be set (may be unexported)", last)
	}

	val := reflect.ValueOf(value)
	fmt.Println("val.Type() =", val.Type(), ", field.Type() =", field.Type())
	// Convert value if necessary
	if val.Type().ConvertibleTo(field.Type()) {
		field.Set(val.Convert(field.Type()))
	} else if val.Type().Kind() == reflect.String && field.Type().Kind() == reflect.String {
		// Both underlying are string but different named types,
		// convert value manually using string conversion
		strVal := val.Interface().(string)
		convertedVal := reflect.ValueOf(strVal).Convert(field.Type())
		field.Set(convertedVal)
	} else {
		return fmt.Errorf("cannot convert value to field type %s", field.Type())
	}

	return nil
}

func CopyDocumentValueToMessage(from IOSDocument, fromPah string, to any, toPath string) error {
	if from == nil || fromPah == "" || toPath == "" {
		return fmt.Errorf("invalid input")
	}
	_, value := GetElement(from, fromPah)
	if value == nil {
		return fmt.Errorf("failed to get element: %s", fromPah)
	}

	err := SetElement(to, toPath, value)
	if err != nil {
		return fmt.Errorf("failed to set element: %w", err)
	}

	return nil
}
