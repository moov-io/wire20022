package document

import (
	"encoding/json"
	"encoding/xml"

	"github.com/moov-io/wire20022/pkg/utils"
)

// Document interface for ISO 20022
type Iso20022Document interface {
	// Validate will be process validation check of document
	Validate() error

	// NameSpace check will return xmlns of document
	NameSpace() string

	// GetXmlName returns xml name of document
	GetXmlName() *xml.Name

	// GetAttrs returns attributes of document
	GetAttrs() []xml.Attr

	// InspectMessage returns message
	InspectMessage() Iso20022Message
}

// Element interface for ISO 20022
type Iso20022Message interface {
	// Validate will be process validation check of document
	Validate() error
}

type constructorFunc func() Iso20022Message

var (
	messageConstructor = map[string]constructorFunc{
		// utils.DocumentARRcamt06000105NameSpace: func() Iso20022Message {
		// 	return &AccountReportingRequest_camt_060_001_05.Document{}
		// },
		// utils.DocumentARcamt05200108NameSpace: func() Iso20022Message {
		// 	return &ActivityReport_camt_052_001_08.Document{}
		// },
		// utils.DocumentBAHead00100103NameSpace: func() Iso20022Message {
		// 	return &BusinessApplicationHeader_head_001_001_03.AppHdr{}
		// },
		// utils.DocumentCCA00400102NameSpace: func() Iso20022Message {
		// 	return &ConnectionCheck_admi_004_001_02.Document{}
		// },
		// utils.DocumentCCTPacs00800108NameSpace: func() Iso20022Message {
		// 	return &CustomerCreditTransfer_pacs_008_001_08.Document{}
		// },
	}
)

type documentDummy struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
}

func (dummy documentDummy) NameSpace() string {
	for _, attr := range dummy.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func NewDocument(space string) (doc Iso20022Document, err error) {
	constractor := messageConstructor[space]
	if constractor == nil {
		return nil, utils.NewErrUnsupportedNameSpace()
	}

	return &Iso20022DocumentObject{
		Message: constractor(),
	}, nil
}

// ParseIso20022Document will return a interface of ISO 20022 document after pass buffer
func ParseIso20022Document(buf []byte) (Iso20022Document, error) {
	docformat := utils.GetDocumentFormat(buf)
	if docformat == utils.DocumentTypeUnknown {
		return nil, utils.NewErrInvalidFileType()
	}

	var dummy documentDummy
	var err error

	if docformat == utils.DocumentTypeXml {
		err = xml.Unmarshal(buf, &dummy)
	} else {
		err = json.Unmarshal(buf, &dummy)
	}
	if err != nil {
		return nil, err
	}

	namespace := dummy.NameSpace()
	if namespace == "" {
		return nil, utils.NewErrOmittedNameSpace()
	}

	constractor := messageConstructor[namespace]
	if constractor == nil {
		return nil, utils.NewErrUnsupportedNameSpace()
	}

	doc := &Iso20022DocumentObject{
		Message: constractor(),
	}

	if docformat == utils.DocumentTypeXml {
		err = xml.Unmarshal(buf, doc)
	} else {
		err = json.Unmarshal(buf, doc)
	}
	if err != nil {
		return nil, err
	}

	return doc, nil
}

type Iso20022DocumentObject struct {
	XMLName xml.Name
	Attrs   []xml.Attr      `xml:",any,attr,omitempty" json:",omitempty"`
	Message Iso20022Message `xml:",any"`
}

func (doc Iso20022DocumentObject) Validate() error {
	if len(doc.NameSpace()) == 0 {
		return utils.Validate(&doc)
	}

	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() == attr.Value {
			return utils.Validate(&doc)
		}
	}

	return utils.NewErrInvalidNameSpace()
}

func (doc Iso20022DocumentObject) NameSpace() string {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (doc *Iso20022DocumentObject) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *Iso20022DocumentObject) GetAttrs() []xml.Attr {
	return doc.Attrs
}

func (doc *Iso20022DocumentObject) InspectMessage() Iso20022Message {
	return doc.Message
}

func (doc Iso20022DocumentObject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	a := struct {
		XMLName xml.Name
		Attrs   []xml.Attr      `xml:",any,attr,omitempty" json:",omitempty"`
		Message Iso20022Message `xml:",any"`
	}(doc)

	updatingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&a, start)
}

func updatingStartElement(start *xml.StartElement, attrs []xml.Attr, name xml.Name) {
	for _, attr := range attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			name.Space = ""
		}
	}
	if len(name.Local) > 0 {
		start.Name.Local = name.Local
	}
	start.Name.Space = name.Space
}
