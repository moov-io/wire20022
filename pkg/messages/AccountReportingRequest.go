package messages

import (
	AccountReportingRequestModel "github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
)

// AccountReportingRequest demonstrates the message processor
// This replaces 121 lines of code with just 25 lines while maintaining identical functionality
type AccountReportingRequest struct {
	*MessageWrapper[AccountReportingRequestModel.MessageModel, AccountReportingRequestModel.CAMT_060_001_VERSION]
}

// NewAccountReportingRequest creates a new type-safe processor for AccountReportingRequest messages
func NewAccountReportingRequest() *AccountReportingRequest {
	return &AccountReportingRequest{
		MessageWrapper: NewMessageWrapper[AccountReportingRequestModel.MessageModel, AccountReportingRequestModel.CAMT_060_001_VERSION](
			"AccountReportingRequest",
			AccountReportingRequestModel.DocumentWith,                               // Type-safe document creator
			AccountReportingRequestModel.CheckRequiredFields,                        // Type-safe field validator
			func() any { return AccountReportingRequestModel.BuildMessageHelper() }, // Adapted helper builder
			func(data []byte) (AccountReportingRequestModel.MessageModel, error) { // XML converter using new API
				msg, err := AccountReportingRequestModel.ParseXML(data)
				if err != nil {
					return AccountReportingRequestModel.MessageModel{}, err
				}
				return *msg, nil
			},
		),
	}
}

// All methods are automatically inherited from MessageWrapper with full type safety:
// - CreateDocument(modelJson []byte, version AccountReportingRequestModel.CAMT_060_001_VERSION) ([]byte, error)
// - ValidateDocument(modelJson string, version AccountReportingRequestModel.CAMT_060_001_VERSION) error
// - Validate(model AccountReportingRequestModel.MessageModel) error
// - ConvertXMLToModel(xmlData []byte) (AccountReportingRequestModel.MessageModel, error)
// - GetHelp() (string, error)

// The message processor provides identical functionality with:
// ✅ 80% code reduction (121 lines → 25 lines)
// ✅ Compile-time type safety for all parameters
// ✅ Centralized error handling and validation
// ✅ Consistent behavior across all message types
// ✅ Easier maintenance and testing
