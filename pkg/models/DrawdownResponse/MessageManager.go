package DrawdownResponse

import "github.com/moov-io/wire20022/pkg/models"

type TransactionInfoAndStatus struct {
	//Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId string
	//Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId string
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUniqueId string
	//Specifies the status of a transaction, in a coded form.
	TransactionStatus models.TransactionStatusCode
	//Provides detailed information on the status reason.
	StatusReasonInfoCode models.StatusReasonInformationCode
}
