package DrawdownResponse

import (
	"reflect"

	pain014 "github.com/moov-io/fedwire20022/gen/DrawdownResponse_pain_014_001_07"
	model "github.com/moov-io/wire20022/pkg/models"
)

type StatusReasonInformationCode string

const (
	InsufficientFunds         StatusReasonInformationCode = "AM04" // The account does not have enough balance to process the transaction.
	DuplicateTransaction      StatusReasonInformationCode = "AM05" // The transaction appears to be a duplicate of another payment.
	WrongAccount              StatusReasonInformationCode = "AM09" // The account number provided is incorrect or does not exist.
	CreditorBankNotRegistered StatusReasonInformationCode = "CNOR" // The creditor’s bank is unknown or not part of the payment system.
	DebtorBankNotRegistered   StatusReasonInformationCode = "DNOR" // The debtor’s bank is unknown or not part of the payment system.
	InvalidFileFormat         StatusReasonInformationCode = "FF01" // The file submitted does not match the expected format.
	InvalidBIC                StatusReasonInformationCode = "RC01" // The Bank Identifier Code (BIC) provided is incorrect.
	MissingDebtorInfo         StatusReasonInformationCode = "RR01" // Mandatory information about the debtor is missing.
	MissingCreditorInfo       StatusReasonInformationCode = "RR02" // Mandatory information about the creditor is missing.
	CutOffTimeExceeded        StatusReasonInformationCode = "SL01" // The transaction was submitted after the cut-off time for processing.
)

type TransactionInfoAndStatus struct {
	//Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId string
	//Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId string
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUniqueId string
	//Specifies the status of a transaction, in a coded form.
	TransactionStatus model.TransactionStatusCode
	//Provides detailed information on the status reason.
	StatusReasonInfoCode StatusReasonInformationCode
}

func PostalAddress241From(a model.PostalAddress) (pain014.PostalAddress241, *model.ValidateError) {
	var result pain014.PostalAddress241
	if a.StreetName != "" {
		err := pain014.Max70Text(a.StreetName).Validate()
		if err != nil {
			return pain014.PostalAddress241{}, &model.ValidateError{
				ParamName: "StreetName",
				Message:   err.Error(),
			}
		}
		StrtNm := pain014.Max70Text(a.StreetName)
		result.StrtNm = &StrtNm
	}
	if a.BuildingNumber != "" {
		err := pain014.Max16Text(a.BuildingNumber).Validate()
		if err != nil {
			return pain014.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingNumber",
				Message:   err.Error(),
			}
		}
		BldgNb := pain014.Max16Text(a.BuildingNumber)
		result.BldgNb = &BldgNb
	}
	if a.RoomNumber != "" {
		err := pain014.Max70Text(a.RoomNumber).Validate()
		if err != nil {
			return pain014.PostalAddress241{}, &model.ValidateError{
				ParamName: "RoomNumber",
				Message:   err.Error(),
			}
		}
		Room := pain014.Max70Text(a.RoomNumber)
		result.Room = &Room
	}
	if a.PostalCode != "" {
		err := pain014.Max16Text(a.PostalCode).Validate()
		if err != nil {
			return pain014.PostalAddress241{}, &model.ValidateError{
				ParamName: "PostalCode",
				Message:   err.Error(),
			}
		}
		PstCd := pain014.Max16Text(a.PostalCode)
		result.PstCd = &PstCd
	}
	if a.TownName != "" {
		err := pain014.Max35Text(a.TownName).Validate()
		if err != nil {
			return pain014.PostalAddress241{}, &model.ValidateError{
				ParamName: "TownName",
				Message:   err.Error(),
			}
		}
		result.TwnNm = pain014.Max35Text(a.TownName)
	}
	if a.Subdivision != "" {
		err := pain014.Max35Text(a.Subdivision).Validate()
		if err != nil {
			return pain014.PostalAddress241{}, &model.ValidateError{
				ParamName: "Subdivision",
				Message:   err.Error(),
			}
		}
		CtrySubDvsn := pain014.Max35Text(a.Subdivision)
		result.CtrySubDvsn = &CtrySubDvsn
	}
	if a.Country != "" {
		err := pain014.CountryCode(a.Country).Validate()
		if err != nil {
			return pain014.PostalAddress241{}, &model.ValidateError{
				ParamName: "Country",
				Message:   err.Error(),
			}
		}
		result.Ctry = pain014.CountryCode(a.Country)
	}
	return result, nil
}
func PartyIdentification1351From(p model.PartyIdentify) (pain014.PartyIdentification1351, *model.ValidateError) {
	var result pain014.PartyIdentification1351
	if p.Name != "" {
		err := pain014.Max140Text(p.Name).Validate()
		if err != nil {
			return pain014.PartyIdentification1351{}, &model.ValidateError{
				ParamName: "Name",
				Message:   err.Error(),
			}
		}
		Nm := pain014.Max140Text(p.Name)
		result.Nm = &Nm
	}
	if !isEmpty(p.Address) {
		PstlAdr, vErr := PostalAddress241From(p.Address)
		if vErr != nil {
			vErr.InsertPath("Address")
			return pain014.PartyIdentification1351{}, vErr
		}
		if !isEmpty(PstlAdr) {
			result.PstlAdr = &PstlAdr
		}
	}
	return result, nil
}
func PaymentTransaction1041From(p TransactionInfoAndStatus) (pain014.PaymentTransaction1041, *model.ValidateError) {
	var result pain014.PaymentTransaction1041
	if p.OriginalInstructionId != "" {
		err := pain014.Max35Text(p.OriginalInstructionId).Validate()
		if err != nil {
			return pain014.PaymentTransaction1041{}, &model.ValidateError{
				ParamName: "OriginalInstructionId",
				Message:   err.Error(),
			}
		}
		OrgnlInstrId := pain014.Max35Text(p.OriginalInstructionId)
		result.OrgnlInstrId = &OrgnlInstrId
	}
	if p.OriginalEndToEndId != "" {
		err := pain014.Max35Text(p.OriginalEndToEndId).Validate()
		if err != nil {
			return pain014.PaymentTransaction1041{}, &model.ValidateError{
				ParamName: "OriginalEndToEndId",
				Message:   err.Error(),
			}
		}
		OrgnlEndToEndId := pain014.Max35Text(p.OriginalEndToEndId)
		result.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if p.OriginalUniqueId != "" {
		err := pain014.UUIDv4Identifier(p.OriginalUniqueId).Validate()
		if err != nil {
			return pain014.PaymentTransaction1041{}, &model.ValidateError{
				ParamName: "OriginalUniqueId",
				Message:   err.Error(),
			}
		}
		result.OrgnlUETR = pain014.UUIDv4Identifier(p.OriginalUniqueId)
	}
	if p.TransactionStatus != "" {
		err := pain014.ExternalPaymentTransactionStatus1Code(p.TransactionStatus).Validate()
		if err != nil {
			return pain014.PaymentTransaction1041{}, &model.ValidateError{
				ParamName: "TransactionStatus",
				Message:   err.Error(),
			}
		}
		result.TxSts = pain014.ExternalPaymentTransactionStatus1Code(p.TransactionStatus)
	}
	if p.StatusReasonInfoCode != "" {
		err := pain014.Max35Text(p.StatusReasonInfoCode).Validate()
		if err != nil {
			return pain014.PaymentTransaction1041{}, &model.ValidateError{
				ParamName: "StatusReasonInfoCode",
				Message:   err.Error(),
			}
		}
		Prtry := pain014.Max35Text(p.StatusReasonInfoCode)
		result.StsRsnInf = []*pain014.StatusReasonInformation121{
			{
				Rsn: pain014.StatusReason6Choice{
					Prtry: &Prtry,
				},
			},
		}
	}
	return result, nil
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
