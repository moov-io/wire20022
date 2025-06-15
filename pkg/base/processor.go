package base

import (
	"errors"
	"fmt"
	"reflect"

	wirerrors "github.com/moov-io/wire20022/pkg/errors"
	"github.com/moov-io/wire20022/pkg/models"
)

// MessageProcessor provides generic message processing capabilities
// using type parameters to maintain type safety while reducing duplication
type MessageProcessor[M any, V comparable] struct {
	namespaceMap   map[string]models.DocumentFactory
	versionMap     map[string]V
	pathMaps       map[V]map[string]any
	requiredFields []string
}

// NewMessageProcessor creates a new generic message processor
func NewMessageProcessor[M any, V comparable](
	namespaceMap map[string]models.DocumentFactory,
	versionMap map[string]V,
	pathMaps map[V]map[string]any,
	requiredFields []string,
) *MessageProcessor[M, V] {
	return &MessageProcessor[M, V]{
		namespaceMap:   namespaceMap,
		versionMap:     versionMap,
		pathMaps:       pathMaps,
		requiredFields: requiredFields,
	}
}

// ProcessMessage handles the common pattern of converting XML to message model
func (p *MessageProcessor[M, V]) ProcessMessage(data []byte) (M, error) {
	var result M

	doc, xmlns, err := models.DocumentFrom(data, p.namespaceMap)
	if err != nil {
		return result, HandleDocumentCreationError(err)
	}

	version, exists := p.versionMap[xmlns]
	if !exists {
		return result, wirerrors.NewParseError("version lookup", xmlns,
			errors.New("unsupported namespace"))
	}

	pathMap, exists := p.pathMaps[version]
	if !exists {
		return result, wirerrors.NewParseError("path map lookup", fmt.Sprintf("%v", version),
			errors.New("missing path map for version"))
	}

	rePathMap := models.RemakeMapping(doc, pathMap, true)

	for sourcePath, targetPath := range rePathMap {
		models.CopyDocumentValueToMessage(doc, sourcePath, &result, targetPath)
	}

	// Validate required fields
	if err := p.ValidateRequiredFields(result); err != nil {
		return result, err
	}

	return result, nil
}

// CreateDocument handles the common pattern of converting message model to XML document
func (p *MessageProcessor[M, V]) CreateDocument(message M, version V) (models.ISODocument, error) {
	pathMap, exists := p.pathMaps[version]
	if !exists {
		return nil, wirerrors.NewValidationError("version", "unsupported version")
	}

	// Get document factory for this version
	var targetNamespace string
	for namespace, ver := range p.versionMap {
		if ver == version {
			targetNamespace = namespace
			break
		}
	}

	factory, exists := p.namespaceMap[targetNamespace]
	if !exists {
		return nil, wirerrors.NewValidationError("namespace", "missing factory for namespace")
	}

	doc := factory()
	rePathMap := models.RemakeMapping(message, pathMap, false)

	for targetPath, sourcePath := range rePathMap {
		if err := models.CopyMessageValueToDocument(&message, sourcePath, doc, targetPath); err != nil {
			return nil, HandleFieldCopyError(sourcePath, err)
		}
	}

	return doc, nil
}

// ValidateRequiredFields performs generic required field validation
func (p *MessageProcessor[M, V]) ValidateRequiredFields(model M) error {
	validator := &FieldValidator{requiredFields: p.requiredFields}
	return validator.ValidateRequired(model)
}

// FieldValidator provides generic field validation capabilities
type FieldValidator struct {
	requiredFields []string
}

// ValidateRequired checks that all required fields are present and non-empty
func (v *FieldValidator) ValidateRequired(model any) error {
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
	}

	for _, fieldName := range v.requiredFields {
		field := modelValue.FieldByName(fieldName)
		if !field.IsValid() || models.IsEmpty(field.Interface()) {
			return wirerrors.NewRequiredFieldError(fieldName)
		}
	}
	return nil
}

// Standard error handling functions to reduce duplication
func HandleDocumentCreationError(err error) error {
	return wirerrors.NewParseError("document creation", "XML data", err)
}

func HandleFieldCopyError(targetPath string, err error) error {
	return wirerrors.NewFieldError(targetPath, "copy", err)
}

func HandleVersionLookupError(xmlns string) error {
	return wirerrors.NewParseError("version lookup", xmlns, errors.New("unsupported namespace"))
}
