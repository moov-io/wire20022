package document

import (
	"encoding/xml"
	"fmt"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
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
	switch message.(type) {
	case *AccountReportingRequest.Message:
		msg, err := AccountReportingRequest.NewMessage("")
		if err != nil {
			return nil, err
		}
		return &msg, nil
	default:
		return nil, fmt.Errorf("unsupported message class")
	}
}
func CreateMessageFrom(xmlData []byte, message interface{}) (MessageInterface, error) {
	switch message.(type) {
	case *AccountReportingRequest.Message:
		msg, err := AccountReportingRequest.NewMessage("")
		if err != nil {
			return nil, err
		}
		if len(xmlData) > 0 {
			if err := xml.Unmarshal(xmlData, msg.GetDocument()); err != nil {
				return nil, err
			}
		}
		return &msg, nil

	default:
		return nil, fmt.Errorf("unsupported message class")
	}
}
func CreateMessageWith(dataModel interface{}, message interface{}) (MessageInterface, error) {
	switch message.(type) {
	case *AccountReportingRequest.Message:
		if model, ok := dataModel.(*AccountReportingRequest.MessageModel); ok {
			msg, err := AccountReportingRequest.NewMessage("")
			if err != nil {
				return nil, err
			}
			msg.Data = *model
			cErr := msg.CreateDocument()
			if cErr != nil {
				return nil, cErr
			}
			return &msg, nil
		} else {
			return nil, fmt.Errorf("data is not of type *AccountReportingRequest.MessageModel")
		}
	default:
		return nil, fmt.Errorf("unsupported message class")
	}
}

func GenerateXML(dataModel interface{}, message interface{}) ([]byte, error) {
	switch message.(type) {
	case *AccountReportingRequest.Message:
		msg, err := CreateMessageWith(dataModel, message)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := msg.(*AccountReportingRequest.Message); ok {
			vErr := msgSt.CreateDocument()
			if vErr != nil {
				return nil, vErr
			}
			xmlData, err := xml.MarshalIndent(msg.GetDocument(), "", "  ")
			if err != nil {
				return nil, err
			}
			return xmlData, nil
		} else {
			return nil, fmt.Errorf("failed to cast message to AccountReportingRequest.Message")
		}

	default:
		return nil, fmt.Errorf("unsupported message class")
	}
}

func ParseXML(xmlData []byte, message interface{}) (MessageInterface, error) {
	switch message.(type) {
	case *AccountReportingRequest.Message:
		if len(xmlData) == 0 {
			return nil, fmt.Errorf("XML data is empty")
		}
		msg, err := CreateMessageFrom(xmlData, message)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := msg.(*AccountReportingRequest.Message); ok {
			vErr := msgSt.CreateMessageModel()
			if vErr != nil {
				return nil, vErr
			}
		}
		return msg, nil
	default:
		return nil, fmt.Errorf("unsupported message class")
	}
}

func RequireFieldCheck(dataModel interface{}, message interface{}) (bool, error) {
	switch message.(type) {
	case *AccountReportingRequest.Message:
		_, err := CreateMessageWith(dataModel, message)
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, fmt.Errorf("unsupported message class")
	}
}
func Validate(xmlData []byte, message interface{}) (bool, error) {
	switch message.(type) {
	case *AccountReportingRequest.Message:
		if len(xmlData) == 0 {
			return false, fmt.Errorf("XML data is empty")
		}
		msg, err := CreateMessageFrom(xmlData, message)
		if err != nil {
			return false, err
		}
		if msgSt, ok := msg.(*AccountReportingRequest.Message); ok {
			vErr := msgSt.CreateMessageModel()
			if vErr != nil {
				return false, vErr
			}
			vaE := msgSt.Doc.Validate()
			if vaE != nil {
				return false, vaE
			}
			return true, nil
		}
		return false, fmt.Errorf("failed to cast message to AccountReportingRequest.Message")
	default:
		return false, fmt.Errorf("unsupported message class")
	}
}
