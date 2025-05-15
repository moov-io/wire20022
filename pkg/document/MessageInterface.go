package document

import (
	"encoding/xml"
	"errors"
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
	"github.com/moov-io/wire20022/pkg/models/FedwireFundsAcknowledgement"
	"github.com/moov-io/wire20022/pkg/models/FedwireFundsBroadcast"
	"github.com/moov-io/wire20022/pkg/models/FedwireFundsPaymentStatus"
	"github.com/moov-io/wire20022/pkg/models/FedwireFundsSystemResponse"
	"github.com/moov-io/wire20022/pkg/models/FinancialInstitutionCreditTransfer"
	"github.com/moov-io/wire20022/pkg/models/InvestRequest"
	"github.com/moov-io/wire20022/pkg/models/InvestResponse"
	"github.com/moov-io/wire20022/pkg/models/Master"
	"github.com/moov-io/wire20022/pkg/models/PaymentReturn"
	"github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest"
	"github.com/moov-io/wire20022/pkg/models/RetrievalRequest"
	"github.com/moov-io/wire20022/pkg/models/ReturnRequest"
	"github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse"
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
	case *FedwireFundsAcknowledgement.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return FedwireFundsAcknowledgement.NewMessage(path)
		})
	case *FedwireFundsBroadcast.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return FedwireFundsBroadcast.NewMessage(path)
		})
	case *FedwireFundsPaymentStatus.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return FedwireFundsPaymentStatus.NewMessage(path)
		})
	case *FedwireFundsSystemResponse.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return FedwireFundsSystemResponse.NewMessage(path)
		})
	case *FinancialInstitutionCreditTransfer.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return FinancialInstitutionCreditTransfer.NewMessage(path)
		})
	case *InvestRequest.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return InvestRequest.NewMessage(path)
		})
	case *InvestResponse.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return InvestResponse.NewMessage(path)
		})
	case *Master.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return Master.NewMessage(path)
		})
	case *PaymentReturn.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return PaymentReturn.NewMessage(path)
		})
	case *PaymentStatusRequest.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return PaymentStatusRequest.NewMessage(path)
		})
	case *RetrievalRequest.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return RetrievalRequest.NewMessage(path)
		})
	case *ReturnRequest.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return ReturnRequest.NewMessage(path)
		})
	case *ReturnRequestResponse.Message:
		return createNewMessage(func(path string) (MessageInterface, error) {
			return ReturnRequestResponse.NewMessage(path)
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
	case *FedwireFundsAcknowledgement.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return FedwireFundsAcknowledgement.NewMessage("")
		}
	case *FedwireFundsBroadcast.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return FedwireFundsBroadcast.NewMessage("")
		}
	case *FedwireFundsPaymentStatus.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return FedwireFundsPaymentStatus.NewMessage("")
		}
	case *FedwireFundsSystemResponse.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return FedwireFundsSystemResponse.NewMessage("")
		}
	case *FinancialInstitutionCreditTransfer.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return FinancialInstitutionCreditTransfer.NewMessage("")
		}
	case *InvestRequest.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return InvestRequest.NewMessage("")
		}
	case *InvestResponse.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return InvestResponse.NewMessage("")
		}
	case *Master.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return Master.NewMessage("")
		}
	case *PaymentReturn.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return PaymentReturn.NewMessage("")
		}
	case *PaymentStatusRequest.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return PaymentStatusRequest.NewMessage("")
		}
	case *RetrievalRequest.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return RetrievalRequest.NewMessage("")
		}
	case *ReturnRequest.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return ReturnRequest.NewMessage("")
		}
	case *ReturnRequestResponse.Message:
		newMessageFunc = func() (MessageInterface, error) {
			return ReturnRequestResponse.NewMessage("")
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
			casted, ok := msg.(*AccountReportingRequest.Message)
			if !ok {
				return errors.New("msg is not of type *AccountReportingRequest.Message")
			}
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
			casted, ok := msg.(*ActivityReport.Message)
			if !ok {
				return errors.New("msg is not of type *ActivityReport.Message")
			}
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
			casted, ok := msg.(*BusinessApplicationHeader.Message)
			if !ok {
				return errors.New("msg is not of type *BusinessApplicationHeader.Message")
			}
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
			casted, ok := msg.(*ConnectionCheck.Message)
			if !ok {
				return errors.New("msg is not of type *ConnectionCheck.Message")
			}
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
			casted, ok := msg.(*CustomerCreditTransfer.Message)
			if !ok {
				return errors.New("msg is not of type *CustomerCreditTransfer.Message")
			}
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
			casted, ok := msg.(*DrawdownRequest.Message)
			if !ok {
				return errors.New("msg is not of type *DrawdownRequest.Message")
			}
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
			casted, ok := msg.(*DrawdownResponse.Message)
			if !ok {
				return errors.New("msg is not of type *DrawdownResponse.Message")
			}
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
			casted, ok := msg.(*EndpointDetailsReport.Message)
			if !ok {
				return errors.New("msg is not of type *EndpointDetailsReport.Message")
			}
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
			casted, ok := msg.(*EndpointGapReport.Message)
			if !ok {
				return errors.New("msg is not of type *EndpointGapReport.Message")
			}
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
			casted, ok := msg.(*EndpointTotalsReport.Message)
			if !ok {
				return errors.New("msg is not of type *EndpointTotalsReport.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *FedwireFundsAcknowledgement.Message:
		model, ok := dataModel.(*FedwireFundsAcknowledgement.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *FedwireFundsAcknowledgement.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return FedwireFundsAcknowledgement.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*FedwireFundsAcknowledgement.Message)
			if !ok {
				return errors.New("msg is not of type *FedwireFundsAcknowledgement.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *FedwireFundsBroadcast.Message:
		model, ok := dataModel.(*FedwireFundsBroadcast.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *FedwireFundsBroadcast.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return FedwireFundsBroadcast.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*FedwireFundsBroadcast.Message)
			if !ok {
				return errors.New("msg is not of type *FedwireFundsBroadcast.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *FedwireFundsPaymentStatus.Message:
		model, ok := dataModel.(*FedwireFundsPaymentStatus.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *FedwireFundsPaymentStatus.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return FedwireFundsPaymentStatus.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*FedwireFundsPaymentStatus.Message)
			if !ok {
				return errors.New("msg is not of type *FedwireFundsPaymentStatus.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *FedwireFundsSystemResponse.Message:
		model, ok := dataModel.(*FedwireFundsSystemResponse.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *FedwireFundsSystemResponse.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return FedwireFundsSystemResponse.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*FedwireFundsSystemResponse.Message)
			if !ok {
				return errors.New("msg is not of type *FedwireFundsSystemResponse.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *FinancialInstitutionCreditTransfer.Message:
		model, ok := dataModel.(*FinancialInstitutionCreditTransfer.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *FinancialInstitutionCreditTransfer.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return FinancialInstitutionCreditTransfer.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*FinancialInstitutionCreditTransfer.Message)
			if !ok {
				return errors.New("msg is not of type *FinancialInstitutionCreditTransfer.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *InvestRequest.Message:
		model, ok := dataModel.(*InvestRequest.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *InvestRequest.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return InvestRequest.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*InvestRequest.Message)
			if !ok {
				return errors.New("msg is not of type *InvestRequest.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *InvestResponse.Message:
		model, ok := dataModel.(*InvestResponse.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *InvestResponse.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return InvestResponse.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*InvestResponse.Message)
			if !ok {
				return errors.New("msg is not of type *InvestResponse.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *Master.Message:
		model, ok := dataModel.(*Master.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *Master.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return Master.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*Master.Message)
			if !ok {
				return errors.New("msg is not of type *Master.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *PaymentReturn.Message:
		model, ok := dataModel.(*PaymentReturn.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *PaymentReturn.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return PaymentReturn.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*PaymentReturn.Message)
			if !ok {
				return errors.New("msg is not of type *PaymentReturn.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *PaymentStatusRequest.Message:
		model, ok := dataModel.(*PaymentStatusRequest.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *PaymentStatusRequest.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return PaymentStatusRequest.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*PaymentStatusRequest.Message)
			if !ok {
				return errors.New("msg is not of type *PaymentStatusRequest.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *RetrievalRequest.Message:
		model, ok := dataModel.(*RetrievalRequest.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *RetrievalRequest.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return RetrievalRequest.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*RetrievalRequest.Message)
			if !ok {
				return errors.New("msg is not of type *RetrievalRequest.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *ReturnRequest.Message:
		model, ok := dataModel.(*ReturnRequest.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *ReturnRequest.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return ReturnRequest.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*ReturnRequest.Message)
			if !ok {
				return errors.New("msg is not of type *ReturnRequest.Message")
			}
			casted.Data = *model
			return casted.CreateDocument()
		}
	case *ReturnRequestResponse.Message:
		model, ok := dataModel.(*ReturnRequestResponse.MessageModel)
		if !ok {
			return nil, fmt.Errorf("expected *ReturnRequestResponse.MessageModel, got %T", dataModel)
		}
		create = func() (MessageInterface, error) {
			return ReturnRequestResponse.NewMessage("")
		}
		assign = func(msg MessageInterface) error {
			casted, ok := msg.(*ReturnRequestResponse.Message)
			if !ok {
				return errors.New("msg is not of type *ReturnRequestResponse.Message")
			}
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
	case *FedwireFundsAcknowledgement.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*FedwireFundsAcknowledgement.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to FedwireFundsAcknowledgement.Message")
		}
	case *FedwireFundsBroadcast.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*FedwireFundsBroadcast.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to FedwireFundsBroadcast.Message")
		}
	case *FedwireFundsPaymentStatus.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*FedwireFundsPaymentStatus.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to FedwireFundsPaymentStatus.Message")
		}
	case *FedwireFundsSystemResponse.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*FedwireFundsSystemResponse.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to FedwireFundsSystemResponse.Message")
		}
	case *FinancialInstitutionCreditTransfer.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*FinancialInstitutionCreditTransfer.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to FinancialInstitutionCreditTransfer.Message")
		}
	case *InvestRequest.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*InvestRequest.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to InvestRequest.Message")
		}
	case *InvestResponse.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*InvestResponse.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to InvestResponse.Message")
		}
	case *Master.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*Master.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to Master.Message")
		}
	case *PaymentReturn.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*PaymentReturn.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to PaymentReturn.Message")
		}
	case *PaymentStatusRequest.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*PaymentStatusRequest.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to PaymentStatusRequest.Message")
		}
	case *RetrievalRequest.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*RetrievalRequest.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to RetrievalRequest.Message")
		}
	case *ReturnRequest.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*ReturnRequest.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to ReturnRequest.Message")
		}
	case *ReturnRequestResponse.Message:
		createdMsg, err := CreateMessageWith(dataModel, msg)
		if err != nil {
			return nil, err
		}
		if msgSt, ok := createdMsg.(*ReturnRequestResponse.Message); ok {
			return generateXMLForMessage(msgSt)
		} else {
			return nil, fmt.Errorf("failed to cast message to ReturnRequestResponse.Message")
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
		case *FedwireFundsAcknowledgement.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *FedwireFundsBroadcast.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *FedwireFundsPaymentStatus.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *FedwireFundsSystemResponse.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *FinancialInstitutionCreditTransfer.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *InvestRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *InvestResponse.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *Master.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *PaymentReturn.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *PaymentStatusRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *RetrievalRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *ReturnRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		case *ReturnRequestResponse.Message:
			if err := m.CreateMessageModel(); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported message type for CreateMessageModel: %T", m)
		}
		return parsedMsg, nil
	}
	casted, ok := message.(MessageInterface)
	if !ok {
		// Handle the error appropriately: log, return, or panic with context
		return nil, errors.New("msg is not of type MessageInterface")
	}
	return processMessage(casted)
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
		case *FedwireFundsAcknowledgement.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *FedwireFundsBroadcast.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *FedwireFundsPaymentStatus.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *FedwireFundsSystemResponse.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *FinancialInstitutionCreditTransfer.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *InvestRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *InvestResponse.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *Master.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *PaymentReturn.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *PaymentStatusRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *RetrievalRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *ReturnRequest.Message:
			if err := m.CreateMessageModel(); err != nil {
				return false, err
			}
			if vaErr := m.Doc.Validate(); vaErr != nil {
				return false, vaErr
			}
		case *ReturnRequestResponse.Message:
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
	casted, ok := message.(MessageInterface)
	if !ok {
		// Handle the error appropriately: log, return, or panic with context
		return false, errors.New("msg is not of type MessageInterface")
	}
	return validateMessage(casted)
}

func isNil(err error) bool {
	return err == nil || reflect.ValueOf(err).IsNil()
}
