package DrawdownRequest_013_001_07

import (
	"reflect"

	"cloud.google.com/go/civil"
	DrawdownRequest "github.com/moov-io/fedwire20022/gen/DrawdownRequest_pain_013_001_07"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type PaymentMethod string
type PaymentRequestType string
type ChargeBearerType string
type CodeOrProprietaryType string

const (
	CreditTransform PaymentMethod = "TRF"
)
const (
	DrawDownRequestCredit PaymentRequestType = "DRRC"
	DrawDownRequestDebit  PaymentRequestType = "DRRB"
	IntraCompanyPayment   PaymentRequestType = "INTC"
)
const (
	ChargeBearerSLEV ChargeBearerType = "SLEV" // Sender Pays All Charges
	ChargeBearerRECV ChargeBearerType = "RECV" // Receiver Pays All Charges
	ChargeBearerSHAR ChargeBearerType = "SHAR" // Shared Charges
)
const (
	CodeCINV CodeOrProprietaryType = "CINV" // Invoice
	CodeCREQ CodeOrProprietaryType = "CREQ" // Credit Request
	CodeCNTR CodeOrProprietaryType = "CNTR" // Credit Note
	CodeDBTR CodeOrProprietaryType = "DBTR" // Debtor
	CodeCRED CodeOrProprietaryType = "CRED" // Credit
	CodeSCT  CodeOrProprietaryType = "SCT"  // SEPA Credit Transfer
	CodePAYM CodeOrProprietaryType = "PAYM" // Payment Message
	CodeRTGS CodeOrProprietaryType = "RTGS" // Real-Time Gross Settlement
	CodeRCLS CodeOrProprietaryType = "RCLS" // Reversal
	CodeRFF  CodeOrProprietaryType = "RFF"  // Reference
	CodeCMCN CodeOrProprietaryType = "CMCN" // Reference
)

type RemittanceDocument struct {
	//Code or Proprietary :It is used to specify the method for identifying the type of a document or reference.
	CodeOrProprietary CodeOrProprietaryType
	//invoice number
	Number string
	//default value: current date
	RelatedDate civil.Date
}

type CreditTransferTransaction struct {
	//Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	PaymentInstructionId string
	//Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	PaymentEndToEndId string
	//Universally unique identifier to provide an end-to-end reference of a payment transaction.
	PaymentUniqueId string
	//Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	PayRequestType PaymentRequestType
	//Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	PayCategoryType PaymentRequestType
	//Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount model.CurrencyAndAmount
	//Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer ChargeBearerType
	// /Financial institution servicing an account for the creditor.
	CreditorAgent model.Agent
	//This is the party whose account will be credited by the creditor agent if the drawdown request is honored.
	Creditor model.PartyIdentify
	//Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CrediorAccountOtherId string
	//Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation string
	document              RemittanceDocument
}

func ReferredDocumentInformation71From(m RemittanceDocument) DrawdownRequest.ReferredDocumentInformation71 {
	var result DrawdownRequest.ReferredDocumentInformation71
	if m.CodeOrProprietary != "" {
		Cd := DrawdownRequest.DocumentType6Code(m.CodeOrProprietary)
		Tp := DrawdownRequest.ReferredDocumentType4{
			CdOrPrtry: DrawdownRequest.ReferredDocumentType3Choice{
				Cd: &Cd,
			},
		}
		result.Tp = &Tp
	}
	if m.Number != "" {
		Nb := DrawdownRequest.Max35Text(m.Number)
		result.Nb = &Nb
	}
	if !isEmpty(m.RelatedDate) {
		RltdDt := fedwire.ISODate(m.RelatedDate)
		result.RltdDt = &RltdDt
	}
	return result
}
func CreditTransferTransaction351From(m CreditTransferTransaction) DrawdownRequest.CreditTransferTransaction351 {
	var result DrawdownRequest.CreditTransferTransaction351
	var PmtId DrawdownRequest.PaymentIdentification61
	if m.PaymentInstructionId != "" {
		InstrId := DrawdownRequest.Max35Text(m.PaymentInstructionId)
		PmtId.InstrId = &InstrId
	}
	if m.PaymentEndToEndId != "" {
		PmtId.EndToEndId = DrawdownRequest.Max35Text(m.PaymentEndToEndId)
	}
	if m.PaymentUniqueId != "" {
		PmtId.UETR = DrawdownRequest.UUIDv4Identifier(m.PaymentUniqueId)
	}
	if !isEmpty(PmtId) {
		result.PmtId = PmtId
	}
	var PmtTpInf DrawdownRequest.PaymentTypeInformation261
	if m.PayRequestType != "" {
		Prtry := DrawdownRequest.LocalInstrumentFedwireFunds1(m.PayRequestType)
		PmtTpInf.LclInstrm = DrawdownRequest.LocalInstrument2Choice1{
			Prtry: &Prtry,
		}
	}
	if m.PayCategoryType != "" {
		Cd := DrawdownRequest.ExternalCategoryPurpose1Code(m.PayCategoryType)
		CtgyPurp := DrawdownRequest.CategoryPurpose1Choice{
			Cd: &Cd,
		}
		PmtTpInf.CtgyPurp = &CtgyPurp
	}
	if !isEmpty(PmtTpInf) {
		result.PmtTpInf = PmtTpInf
	}
	if !isEmpty(m.Amount) {
		InstdAmt := DrawdownRequest.ActiveCurrencyAndAmountFedwire1{
			Value: DrawdownRequest.ActiveCurrencyAndAmountFedwire1SimpleType(m.Amount.Amount),
			Ccy:   DrawdownRequest.ActiveCurrencyCodeFixed(m.Amount.Currency),
		}
		result.Amt = DrawdownRequest.AmountType4Choice1{
			InstdAmt: &InstdAmt,
		}
	}
	if m.ChargeBearer != "" {
		result.ChrgBr = DrawdownRequest.ChargeBearerType1Code(m.ChargeBearer)
	}
	if !isEmpty(m.CreditorAgent) {
		Cd := DrawdownRequest.ExternalClearingSystemIdentification1CodeFixed(m.CreditorAgent.PaymentSysCode)
		CdtrAgt := DrawdownRequest.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: DrawdownRequest.FinancialInstitutionIdentification181{
				ClrSysMmbId: DrawdownRequest.ClearingSystemMemberIdentification21{
					ClrSysId: DrawdownRequest.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: DrawdownRequest.RoutingNumberFRS1(m.CreditorAgent.PaymentSysMemberId),
				},
			},
		}
		if !isEmpty(CdtrAgt) {
			result.CdtrAgt = CdtrAgt
		}
	}
	if !isEmpty(m.Creditor) {
		Cdtr := PartyIdentification1352From(m.Creditor)
		if !isEmpty(Cdtr) {
			result.Cdtr = Cdtr
		}
	}
	if m.CrediorAccountOtherId != "" {
		Othr := DrawdownRequest.GenericAccountIdentification1{
			Id: DrawdownRequest.Max34Text(m.CrediorAccountOtherId),
		}
		cashAcc := DrawdownRequest.CashAccount38{
			Id: DrawdownRequest.AccountIdentification4Choice{
				Othr: &Othr,
			},
		}
		result.CdtrAcct = &cashAcc
	}
	var RmtInf DrawdownRequest.RemittanceInformation161
	if m.RemittanceInformation != "" {
		Ustrd := DrawdownRequest.Max140Text(m.RemittanceInformation)
		RmtInf.Ustrd = &Ustrd
	}
	if !isEmpty(m.document) {
		var Strd []*DrawdownRequest.StructuredRemittanceInformation161
		var RfrdDocInf []*DrawdownRequest.ReferredDocumentInformation71
		doc := ReferredDocumentInformation71From(m.document)
		RfrdDocInf = append(RfrdDocInf, &doc)
		remitDoc := DrawdownRequest.StructuredRemittanceInformation161{
			RfrdDocInf: RfrdDocInf,
		}
		Strd = append(Strd, &remitDoc)
		RmtInf.Strd = Strd
	}
	if !isEmpty(RmtInf) {
		result.RmtInf = &RmtInf
	}
	return result
}
func PostalAddress241From(a model.PostalAddress) DrawdownRequest.PostalAddress241 {
	var result DrawdownRequest.PostalAddress241
	if a.StreetName != "" {
		StrtNm := DrawdownRequest.Max70Text(a.StreetName)
		result.StrtNm = &StrtNm
	}
	if a.BuildingNumber != "" {
		BldgNb := DrawdownRequest.Max16Text(a.BuildingNumber)
		result.BldgNb = &BldgNb
	}
	if a.RoomNumber != "" {
		Room := DrawdownRequest.Max70Text(a.RoomNumber)
		result.Room = &Room
	}
	if a.PostalCode != "" {
		PstCd := DrawdownRequest.Max16Text(a.PostalCode)
		result.PstCd = &PstCd
	}
	if a.TownName != "" {
		result.TwnNm = DrawdownRequest.Max35Text(a.TownName)
	}
	if a.Subdivision != "" {
		CtrySubDvsn := DrawdownRequest.Max35Text(a.Subdivision)
		result.CtrySubDvsn = &CtrySubDvsn
	}
	if a.Country != "" {
		result.Ctry = DrawdownRequest.CountryCode(a.Country)
	}
	return result
}
func PostalAddress242From(a model.PostalAddress) DrawdownRequest.PostalAddress242 {
	var result DrawdownRequest.PostalAddress242
	if a.StreetName != "" {
		StrtNm := DrawdownRequest.Max70Text(a.StreetName)
		result.StrtNm = &StrtNm
	}
	if a.BuildingNumber != "" {
		BldgNb := DrawdownRequest.Max16Text(a.BuildingNumber)
		result.BldgNb = &BldgNb
	}
	if a.RoomNumber != "" {
		Room := DrawdownRequest.Max70Text(a.RoomNumber)
		result.Room = &Room
	}
	if a.PostalCode != "" {
		PstCd := DrawdownRequest.Max16Text(a.PostalCode)
		result.PstCd = &PstCd
	}
	if a.TownName != "" {
		TwnNm := DrawdownRequest.Max35Text(a.TownName)
		result.TwnNm = &TwnNm
	}
	if a.Subdivision != "" {
		CtrySubDvsn := DrawdownRequest.Max35Text(a.Subdivision)
		result.CtrySubDvsn = &CtrySubDvsn
	}
	if a.Country != "" {
		Ctry := DrawdownRequest.CountryCode(a.Country)
		result.Ctry = &Ctry
	}
	return result
}

func PartyIdentification1351From(p model.PartyIdentify) DrawdownRequest.PartyIdentification1351 {
	var result DrawdownRequest.PartyIdentification1351
	if p.Name != "" {
		Nm := DrawdownRequest.Max140Text(p.Name)
		result.Nm = &Nm
	}
	if !isEmpty(p.Address) {
		PstlAdr := PostalAddress241From(p.Address)
		if !isEmpty(PstlAdr) {
			result.PstlAdr = &PstlAdr
		}
	}
	return result
}
func PartyIdentification1352From(p model.PartyIdentify) DrawdownRequest.PartyIdentification1352 {
	var result DrawdownRequest.PartyIdentification1352
	if p.Name != "" {
		Nm := DrawdownRequest.Max140Text(p.Name)
		result.Nm = &Nm
	}
	if !isEmpty(p.Address) {
		PstlAdr := PostalAddress242From(p.Address)
		if !isEmpty(PstlAdr) {
			result.PstlAdr = &PstlAdr
		}
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
