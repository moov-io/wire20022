package document

import (
	"encoding/xml"
	"fmt"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
)

type MessageInterface interface {
	ValidateRequiredFields() *model.ValidateError

	CreateDocument() *model.ValidateError

    GetDataModel() interface{}
	GetDocument() interface{}
    GetHelper() interface{}

    GetXML() ([]byte, error)
    GetJSON() ([]byte, error)
}

func CreateMessage(message interface{}) (MessageInterface, error) {
    switch message.(type) {
    case *AccountReportingRequest.Message:
        msg, err := AccountReportingRequest.NewMessage("")
        return &msg, err
    default:
        return nil, fmt.Errorf("unsupported message class")
    }
}
func CreateMessageFrom(xmlData []byte, message interface{}) (MessageInterface, error) {
	switch message.(type) {
    case *AccountReportingRequest.Message:
        msg, err := AccountReportingRequest.NewMessage("")
        if err != nil {
            return nil, fmt.Errorf("failed to create AccountReportingRequest.Message: %w", err)
        }
        // Pass a pointer to msg.doc for unmarshalling
        if len(xmlData) > 0 {
            if err := xml.Unmarshal(xmlData, msg.GetDocument()); err != nil {
                return nil, fmt.Errorf("XML parse error: %w", err)
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
                return nil, fmt.Errorf("failed to create AccountReportingRequest.Message: %w", err)
            }
            msg.Data = *model
            cErr := msg.CreateDocument()
            if cErr != nil {
                return nil, fmt.Errorf("failed to create document: %w", cErr)
            }
            return &msg, nil
        } else {
            return nil, fmt.Errorf("data is not of type *AccountReportingRequest.MessageModel")
        }
    default:
        return nil, fmt.Errorf("unsupported message class")
    }
}