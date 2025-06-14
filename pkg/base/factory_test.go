package base

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/moov-io/wire20022/pkg/models"
)

// Mock document for testing
type TestDoc struct {
	XMLName xml.Name `xml:"Document"`
	Content string   `xml:"content"`
}

func (t *TestDoc) Validate() error {
	return nil
}

func (t *TestDoc) SetXMLName(name xml.Name) {
	t.XMLName = name
}

// Test version type
type FactoryTestVersion string

const (
	FactoryV1 FactoryTestVersion = "v1"
	FactoryV2 FactoryTestVersion = "v2"
)

func TestVersionedDocumentFactory(t *testing.T) {
	t.Run("NewVersionedDocumentFactory", func(t *testing.T) {
		factory := NewVersionedDocumentFactory[*TestDoc, FactoryTestVersion]()
		
		assert.NotNil(t, factory)
		assert.NotNil(t, factory.versionMap)
		assert.NotNil(t, factory.namespaceMap)
		assert.NotNil(t, factory.factories)
	})

	t.Run("RegisterVersion", func(t *testing.T) {
		factory := NewVersionedDocumentFactory[*TestDoc, FactoryTestVersion]()
		
		namespace := "test:namespace:v1"
		version := FactoryV1
		factoryFunc := func() *TestDoc {
			return &TestDoc{Content: "test"}
		}
		
		factory.RegisterVersion(namespace, version, factoryFunc)
		
		// Check internal mappings
		assert.Equal(t, version, factory.versionMap[namespace])
		assert.Equal(t, namespace, factory.namespaceMap[version])
		assert.NotNil(t, factory.factories[version])
	})

	t.Run("BuildNameSpaceModelMap", func(t *testing.T) {
		factory := NewVersionedDocumentFactory[*TestDoc, FactoryTestVersion]()
		
		namespace1 := "test:namespace:v1"
		namespace2 := "test:namespace:v2"
		
		factory.RegisterVersion(namespace1, FactoryV1, func() *TestDoc {
			return &TestDoc{Content: "v1"}
		})
		factory.RegisterVersion(namespace2, FactoryV2, func() *TestDoc {
			return &TestDoc{Content: "v2"}
		})
		
		nameSpaceMap := factory.BuildNameSpaceModelMap()
		
		require.Len(t, nameSpaceMap, 2)
		assert.Contains(t, nameSpaceMap, namespace1)
		assert.Contains(t, nameSpaceMap, namespace2)
		
		// Test that factories work
		doc1 := nameSpaceMap[namespace1]()
		assert.NotNil(t, doc1)
		
		doc2 := nameSpaceMap[namespace2]()
		assert.NotNil(t, doc2)
		
		// Verify XMLName is set correctly
		testDoc1, ok := doc1.(*TestDoc)
		require.True(t, ok)
		assert.Equal(t, namespace1, testDoc1.XMLName.Space)
		assert.Equal(t, "Document", testDoc1.XMLName.Local)
	})

	t.Run("GetVersionMap", func(t *testing.T) {
		factory := NewVersionedDocumentFactory[*TestDoc, FactoryTestVersion]()
		
		namespace := "test:namespace:v1"
		version := FactoryV1
		
		factory.RegisterVersion(namespace, version, func() *TestDoc {
			return &TestDoc{Content: "test"}
		})
		
		versionMap := factory.GetVersionMap()
		
		assert.Len(t, versionMap, 1)
		assert.Equal(t, version, versionMap[namespace])
		
		// Verify it's a copy (not the same map)
		versionMap["new:namespace"] = FactoryV2
		assert.NotEqual(t, len(factory.versionMap), len(versionMap))
	})

	t.Run("GetNamespaceMap", func(t *testing.T) {
		factory := NewVersionedDocumentFactory[*TestDoc, FactoryTestVersion]()
		
		namespace := "test:namespace:v1"
		version := FactoryV1
		
		factory.RegisterVersion(namespace, version, func() *TestDoc {
			return &TestDoc{Content: "test"}
		})
		
		namespaceMap := factory.GetNamespaceMap()
		
		assert.Len(t, namespaceMap, 1)
		assert.Equal(t, namespace, namespaceMap[version])
		
		// Verify it's a copy (not the same map)
		namespaceMap[FactoryV2] = "new:namespace"
		assert.NotEqual(t, len(factory.namespaceMap), len(namespaceMap))
	})
}

func TestXMLNameSetter(t *testing.T) {
	t.Run("SetXMLName interface", func(t *testing.T) {
		doc := &TestDoc{Content: "test"}
		
		// Test that TestDoc implements XMLNameSetter
		var setter XMLNameSetter = doc
		require.NotNil(t, setter)
		
		newName := xml.Name{Space: "new:namespace", Local: "NewDocument"}
		setter.SetXMLName(newName)
		
		assert.Equal(t, newName, doc.XMLName)
	})
}

func TestSetXMLNameByReflection(t *testing.T) {
	t.Run("setXMLNameByReflection with valid struct", func(t *testing.T) {
		doc := &TestDoc{Content: "test"}
		namespace := "reflection:namespace"
		
		setXMLNameByReflection(doc, namespace)
		
		assert.Equal(t, namespace, doc.XMLName.Space)
		assert.Equal(t, "Document", doc.XMLName.Local)
	})

	t.Run("setXMLNameByReflection with non-pointer", func(t *testing.T) {
		doc := TestDoc{Content: "test"}
		namespace := "reflection:namespace"
		
		// This should not panic, but also won't modify the struct
		setXMLNameByReflection(doc, namespace)
		
		// Original struct should be unchanged since it's not a pointer
		assert.Empty(t, doc.XMLName.Space)
	})

	t.Run("setXMLNameByReflection with struct without XMLName field", func(t *testing.T) {
		type NoXMLName struct {
			Content string
		}
		
		doc := &NoXMLName{Content: "test"}
		namespace := "reflection:namespace"
		
		// This should not panic even though there's no XMLName field
		setXMLNameByReflection(doc, namespace)
		
		// Nothing should happen
		assert.Equal(t, "test", doc.Content)
	})
}

func TestFactoryRegistration(t *testing.T) {
	t.Run("FactoryRegistration struct", func(t *testing.T) {
		reg := FactoryRegistration[*TestDoc, FactoryTestVersion]{
			Namespace: "test:namespace:v1",
			Version:   FactoryV1,
			Factory: func() *TestDoc {
				return &TestDoc{Content: "registered"}
			},
		}
		
		assert.Equal(t, "test:namespace:v1", reg.Namespace)
		assert.Equal(t, FactoryV1, reg.Version)
		assert.NotNil(t, reg.Factory)
		
		// Test factory function
		doc := reg.Factory()
		assert.Equal(t, "registered", doc.Content)
	})
}

func TestBuildFactoryFromRegistrations(t *testing.T) {
	t.Run("BuildFactoryFromRegistrations with multiple registrations", func(t *testing.T) {
		registrations := []FactoryRegistration[*TestDoc, FactoryTestVersion]{
			{
				Namespace: "test:namespace:v1",
				Version:   FactoryV1,
				Factory: func() *TestDoc {
					return &TestDoc{Content: "v1"}
				},
			},
			{
				Namespace: "test:namespace:v2", 
				Version:   FactoryV2,
				Factory: func() *TestDoc {
					return &TestDoc{Content: "v2"}
				},
			},
		}
		
		factory := BuildFactoryFromRegistrations(registrations)
		
		assert.NotNil(t, factory)
		
		// Verify both versions are registered
		versionMap := factory.GetVersionMap()
		assert.Len(t, versionMap, 2)
		assert.Equal(t, FactoryV1, versionMap["test:namespace:v1"])
		assert.Equal(t, FactoryV2, versionMap["test:namespace:v2"])
		
		// Verify namespace map
		namespaceMap := factory.GetNamespaceMap()
		assert.Len(t, namespaceMap, 2)
		assert.Equal(t, "test:namespace:v1", namespaceMap[FactoryV1])
		assert.Equal(t, "test:namespace:v2", namespaceMap[FactoryV2])
		
		// Test factory map works
		factoryMap := factory.BuildNameSpaceModelMap()
		assert.Len(t, factoryMap, 2)
		
		doc1 := factoryMap["test:namespace:v1"]()
		testDoc1, ok := doc1.(*TestDoc)
		require.True(t, ok)
		assert.Equal(t, "v1", testDoc1.Content)
		
		doc2 := factoryMap["test:namespace:v2"]()
		testDoc2, ok := doc2.(*TestDoc)
		require.True(t, ok)
		assert.Equal(t, "v2", testDoc2.Content)
	})

	t.Run("BuildFactoryFromRegistrations with empty slice", func(t *testing.T) {
		registrations := []FactoryRegistration[*TestDoc, FactoryTestVersion]{}
		
		factory := BuildFactoryFromRegistrations(registrations)
		
		assert.NotNil(t, factory)
		
		versionMap := factory.GetVersionMap()
		assert.Len(t, versionMap, 0)
		
		namespaceMap := factory.GetNamespaceMap()
		assert.Len(t, namespaceMap, 0)
		
		factoryMap := factory.BuildNameSpaceModelMap()
		assert.Len(t, factoryMap, 0)
	})
}

// Mock document that does NOT implement XMLNameSetter for testing reflection fallback
type NoSetterDoc struct {
	XMLName xml.Name `xml:"Document"`
	Content string   `xml:"content"`
}

func (n *NoSetterDoc) Validate() error {
	return nil
}

func TestBuildNameSpaceModelMapWithReflectionFallback(t *testing.T) {
	t.Run("Factory with document that doesn't implement XMLNameSetter", func(t *testing.T) {
		factory := NewVersionedDocumentFactory[models.ISODocument, FactoryTestVersion]()
		
		namespace := "test:namespace:v1"
		factory.RegisterVersion(namespace, FactoryV1, func() models.ISODocument {
			return &NoSetterDoc{Content: "test"}
		})
		
		factoryMap := factory.BuildNameSpaceModelMap()
		
		require.Len(t, factoryMap, 1)
		
		doc := factoryMap[namespace]()
		noSetterDoc, ok := doc.(*NoSetterDoc)
		require.True(t, ok)
		
		// Should have XMLName set via reflection
		assert.Equal(t, namespace, noSetterDoc.XMLName.Space)
		assert.Equal(t, "Document", noSetterDoc.XMLName.Local)
		assert.Equal(t, "test", noSetterDoc.Content)
	})
}