package base

import (
	"encoding/xml"
	"reflect"

	"github.com/moov-io/wire20022/pkg/models"
)

// VersionedDocumentFactory provides generic factory functionality
// for creating versioned document factories with type safety
type VersionedDocumentFactory[T models.ISODocument, V comparable] struct {
	versionMap   map[string]V
	namespaceMap map[V]string
	factories    map[V]func() T
}

// NewVersionedDocumentFactory creates a new versioned document factory
func NewVersionedDocumentFactory[T models.ISODocument, V comparable]() *VersionedDocumentFactory[T, V] {
	return &VersionedDocumentFactory[T, V]{
		versionMap:   make(map[string]V),
		namespaceMap: make(map[V]string),
		factories:    make(map[V]func() T),
	}
}

// RegisterVersion registers a version with its namespace and factory function
func (f *VersionedDocumentFactory[T, V]) RegisterVersion(
	namespace string,
	version V,
	factory func() T,
) {
	f.versionMap[namespace] = version
	f.namespaceMap[version] = namespace
	f.factories[version] = factory
}

// BuildNameSpaceModelMap creates the namespace-to-factory map used by models.DocumentFrom
func (f *VersionedDocumentFactory[T, V]) BuildNameSpaceModelMap() map[string]models.DocumentFactory {
	result := make(map[string]models.DocumentFactory)

	for namespace, version := range f.versionMap {
		if factory, exists := f.factories[version]; exists {
			if space, hasSpace := f.namespaceMap[version]; hasSpace {
				// Capture variables for closure
				capturedFactory := factory
				capturedSpace := space

				result[namespace] = func() models.ISODocument {
					doc := capturedFactory()
					// Set XMLName using type assertion
					if xmlNameSetter, ok := any(doc).(XMLNameSetter); ok {
						xmlNameSetter.SetXMLName(xml.Name{Space: capturedSpace, Local: "Document"})
					} else {
						// Fallback using reflection
						setXMLNameByReflection(doc, capturedSpace)
					}
					return doc
				}
			}
		}
	}

	return result
}

// GetVersionMap returns the namespace to version mapping
func (f *VersionedDocumentFactory[T, V]) GetVersionMap() map[string]V {
	result := make(map[string]V)
	for k, v := range f.versionMap {
		result[k] = v
	}
	return result
}

// GetNamespaceMap returns the version to namespace mapping
func (f *VersionedDocumentFactory[T, V]) GetNamespaceMap() map[V]string {
	result := make(map[V]string)
	for k, v := range f.namespaceMap {
		result[k] = v
	}
	return result
}

// XMLNameSetter interface for documents that can set their XMLName
type XMLNameSetter interface {
	SetXMLName(name xml.Name)
}

// setXMLNameByReflection sets XMLName using reflection as fallback
func setXMLNameByReflection(doc any, namespace string) {
	docValue := reflect.ValueOf(doc)
	if docValue.Kind() == reflect.Ptr {
		docValue = docValue.Elem()
	}

	xmlNameField := docValue.FieldByName("XMLName")
	if xmlNameField.IsValid() && xmlNameField.CanSet() {
		xmlName := xml.Name{Space: namespace, Local: "Document"}
		xmlNameField.Set(reflect.ValueOf(xmlName))
	}
}

// FactoryRegistration represents a single version registration for building factories
type FactoryRegistration[T models.ISODocument, V comparable] struct {
	Namespace string
	Version   V
	Factory   func() T
}

// BuildFactoryFromRegistrations creates a factory from a slice of registrations
func BuildFactoryFromRegistrations[T models.ISODocument, V comparable](
	registrations []FactoryRegistration[T, V],
) *VersionedDocumentFactory[T, V] {
	factory := NewVersionedDocumentFactory[T, V]()

	for _, reg := range registrations {
		factory.RegisterVersion(reg.Namespace, reg.Version, reg.Factory)
	}

	return factory
}
