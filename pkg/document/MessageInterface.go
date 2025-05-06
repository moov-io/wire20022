package document

import (
	"encoding/xml"
	"fmt"
	"reflect"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
	"github.com/moov-io/wire20022/pkg/models/ActivityReport"
)

type MessageInterface interface {
	CreateDocument() *model.ValidateError
	CreateMessageModel() *model.ValidateError
	ValidateRequiredFields() *model.ValidateError

	GetDataModel() interface{}
	GetDocument() interface{}
	GetHelper() interface{}
}

func CreateMessage(message interface{}) (MessageInterface, error) {
	createNewMessage := func(newMessageFunc func(string) (MessageInterface, error)) (MessageInterface, error) {
		msg, err := newMessageFunc("")
		if err != nil {
			return nil, err
		}
		return msg, nil
	}
	switch msg := message.(type) {
	case *AccountReportingRequest.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return AccountReportingRequest.NewMessage(path)
		})
	case *ActivityReport.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return ActivityReport.NewMessage(path)
		})
	default:
		return nil, fmt.Errorf("unsupported message class: %T", msg)
	}
}

func CreateMessageFrom(xmlData []byte, message interface{}) (MessageInterface, error) {
	var newMessageFunc func() (MessageInterface, error)

	switch message.(type) {
	case *AccountReportingRequest.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return AccountReportingRequest.NewMessage("")
		}
	case *ActivityReport.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return ActivityReport.NewMessage("")
		}
	default:
		return nil, fmt.Errorf("unsupported message class: %T", message)
	}

	msg, err := newMessageFunc()
	if err != nil {
		return nil, err
	}

	if len(xmlData) > 0 {
		if err := xml.Unmarshal(xmlData, msg.GetDocument()); err != nil {
			return nil, err
		}
	}
	return msg, nil
}
func CreateMessageWith(dataModel interface{}, message interface{}) (MessageInterface, error) {
	type buildFn func() (MessageInterface, error)

	var (
		create buildFn
		assign func(msg MessageInterface) error
	)

	switch m := message.(type) {
	case *AccountReportingRequest.Message:
		model, ok := dataModel.(*AccountReportingRequest.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *AccountReportingRequest.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return AccountReportingRequest.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*AccountReportingRequest.Message)
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *ActivityReport.Message:
		model, ok := dataModel.(*ActivityReport.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *ActivityReport.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return ActivityReport.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*ActivityReport.Message)
			casted.Data = *model
			return casted.CreateDocument()
		}
	default:
		return nil, fmt.Errorf("unsupported message class: %T", m)
	}
	msg, err := create()
	if err != nil {
		return nil, err
	}
	if err := assign(msg); err != nil && !isNil(err) {
		return nil, err
	}
	return msg, nil
}

func GenerateXML(dataModel interface{}, message MessageInterface{}) ([]byte, error) {
	// Helper function to generate XML from a message
	generateXMLForMessage := func(msg MessageInterface) ([]byte, error) {
		if err := msg.CreateDocument(); err != nil {
			return nil, err
		}
		xmlData, err := xml.MarshalIndent(msg.GetDocument(), "", "  ")
		if err != nil {
			return nil, err
		}
		return xmlData, nil
	}

	// Create message and generate XML based on type
	switch msg := message.(type) {
	case *AccountReportingRequest.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*AccountReportingRequest.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to AccountReportingRequest.Message")
		}

	case *ActivityReport.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*ActivityReport.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to ActivityReport.Message")
		}

	default:
		return nil, fmt.Errorf("unsupported message class")
	}
}

func ParseXML(xmlData []byte, message interface{}) (MessageInterface, error) {
	// Helper function to process the XML data for a message
	processMessage := func(msg MessageInterface) (MessageInterface, error) {
		if len(xmlData) == 0 {
			return nil, fmt.Errorf("XML data is empty")
		}
		parsedMsg, err := CreateMessageFrom(xmlData, msg)
		if err != nil {
			return nil, err
		}

		// Handle the message-specific CreateMessageModel step
		switch m := parsedMsg.(type) {
		case *AccountReportingRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *ActivityReport.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported message type for CreateMessageModel: %T", m)
		}
		return parsedMsg, nil
	}

	return processMessage(message.(MessageInterface))
}

func RequireFieldCheck(dataModel interface{}, message interface{}) (bool, error) {
	// Helper function to check the required field for a given message
	checkMessage := func() (bool, error) {
		_, err := CreateMessageWith(dataModel, message)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return checkMessage()
}

func Validate(xmlData []byte, message interface{}) (bool, error) {
	// Helper function to validate the message
	validateMessage := func(msg MessageInterface) (bool, error) {
		if len(xmlData) == 0 {
			return false, fmt.Errorf("XML data is empty")
		}
		parsedMsg, err := CreateMessageFrom(xmlData, msg)
		if err != nil {
			return false, err
		}

		// Perform type assertion and validation
		switch m := parsedMsg.(type) {
		case *AccountReportingRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *ActivityReport.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		default:
			return false, fmt.Errorf("unsupported message type for validation: %T", m)
		}
		return true, nil
	}

	return validateMessage(message.(MessageInterface))
}

func isNil(err error) bool {
	return err == nil || reflect.ValueOf(err).IsNil()
}
