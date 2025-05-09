package document

import (
	"encoding/xml"
	"fmt"
	"reflect"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
	"github.com/moov-io/wire20022/pkg/models/ActivityReport"
	"github.com/moov-io/wire20022/pkg/models/BusinessApplicationHeader"
	"github.com/moov-io/wire20022/pkg/models/ConnectionCheck"
	"github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
	"github.com/moov-io/wire20022/pkg/models/DrawdownRequest"
	"github.com/moov-io/wire20022/pkg/models/DrawdownResponse"
	"github.com/moov-io/wire20022/pkg/models/EndpointDetailsReport"
	"github.com/moov-io/wire20022/pkg/models/EndpointGapReport"
	"github.com/moov-io/wire20022/pkg/models/EndpointTotalsReport"
)

type MessageInterface interface {
	CreateDocument() *model.ValidateError
	CreateMessageModel() *model.ValidateError
	ValidateRequiredFields() *model.ValidateError

	GetDataModel() any
	GetDocument() any
	GetHelper() any
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
	case *BusinessApplicationHeader.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return BusinessApplicationHeader.NewMessage(path)
		})
	case *ConnectionCheck.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return ConnectionCheck.NewMessage(path)
		})
	case *CustomerCreditTransfer.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return CustomerCreditTransfer.NewMessage(path)
		})
	case *DrawdownRequest.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return DrawdownRequest.NewMessage(path)
		})
	case *DrawdownResponse.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return DrawdownResponse.NewMessage(path)
		})
	case *EndpointDetailsReport.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return EndpointDetailsReport.NewMessage(path)
		})
	case *EndpointGapReport.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return EndpointGapReport.NewMessage(path)
		})
	case *EndpointTotalsReport.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return EndpointTotalsReport.NewMessage(path)
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
	case *BusinessApplicationHeader.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return BusinessApplicationHeader.NewMessage("")
		}
	case *ConnectionCheck.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return ConnectionCheck.NewMessage("")
		}
	case *CustomerCreditTransfer.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return CustomerCreditTransfer.NewMessage("")
		}
	case *DrawdownRequest.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return DrawdownRequest.NewMessage("")
		}
	case *DrawdownResponse.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return DrawdownResponse.NewMessage("")
		}
	case *EndpointDetailsReport.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return EndpointDetailsReport.NewMessage("")
		}
	case *EndpointGapReport.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return EndpointGapReport.NewMessage("")
		}
	case *EndpointTotalsReport.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return EndpointTotalsReport.NewMessage("")
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
	case *BusinessApplicationHeader.Message:
		model, ok := dataModel.(*BusinessApplicationHeader.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *BusinessApplicationHeader.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return BusinessApplicationHeader.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*BusinessApplicationHeader.Message)
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *ConnectionCheck.Message:
		model, ok := dataModel.(*ConnectionCheck.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *ConnectionCheck.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return ConnectionCheck.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*ConnectionCheck.Message)
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *CustomerCreditTransfer.Message:
		model, ok := dataModel.(*CustomerCreditTransfer.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *CustomerCreditTransfer.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return CustomerCreditTransfer.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*CustomerCreditTransfer.Message)
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *DrawdownRequest.Message:
		model, ok := dataModel.(*DrawdownRequest.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *DrawdownRequest.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return DrawdownRequest.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*DrawdownRequest.Message)
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *DrawdownResponse.Message:
		model, ok := dataModel.(*DrawdownResponse.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *DrawdownResponse.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return DrawdownResponse.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*DrawdownResponse.Message)
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *EndpointDetailsReport.Message:
		model, ok := dataModel.(*EndpointDetailsReport.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *EndpointDetailsReport.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return EndpointDetailsReport.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*EndpointDetailsReport.Message)
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *EndpointGapReport.Message:
		model, ok := dataModel.(*EndpointGapReport.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *EndpointGapReport.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return EndpointGapReport.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*EndpointGapReport.Message)
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *EndpointTotalsReport.Message:
		model, ok := dataModel.(*EndpointTotalsReport.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *EndpointTotalsReport.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return EndpointTotalsReport.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted := msg.(*EndpointTotalsReport.Message)
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

func GenerateXML(dataModel interface{}, message interface{}) ([]byte, error) {
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
	case *BusinessApplicationHeader.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*BusinessApplicationHeader.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to BusinessApplicationHeader.Message")
		}
	case *ConnectionCheck.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*ConnectionCheck.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to ConnectionCheck.Message")
		}
	case *CustomerCreditTransfer.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*CustomerCreditTransfer.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to CustomerCreditTransfer.Message")
		}
	case *DrawdownRequest.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*DrawdownRequest.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to DrawdownRequest.Message")
		}
	case *DrawdownResponse.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*DrawdownResponse.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to DrawdownResponse.Message")
		}
	case *EndpointDetailsReport.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*EndpointDetailsReport.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to EndpointDetailsReport.Message")
		}
	case *EndpointGapReport.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*EndpointGapReport.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to EndpointGapReport.Message")
		}
	case *EndpointTotalsReport.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*EndpointTotalsReport.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to EndpointTotalsReport.Message")
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
		case *BusinessApplicationHeader.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *ConnectionCheck.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *CustomerCreditTransfer.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *DrawdownRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *DrawdownResponse.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *EndpointDetailsReport.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *EndpointGapReport.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *EndpointTotalsReport.Message:
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
		case *BusinessApplicationHeader.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *ConnectionCheck.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *CustomerCreditTransfer.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *DrawdownRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *DrawdownResponse.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *EndpointDetailsReport.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *EndpointGapReport.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *EndpointTotalsReport.Message:
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
