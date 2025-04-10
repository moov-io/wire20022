package DrawdownRequest

import (
	"reflect"

	pain013 "github.com/moov-io/fedwire20022/gen/DrawdownRequest_pain_013_001_07"
	model "github.com/moov-io/wire20022/pkg/models"
)

type PaymentMethod string
type PaymentRequestType string
type ChargeBearerType string

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

type RemittanceDocument struct {
	//Code or Proprietary :It is used to specify the method for identifying the type of a document or reference.
	CodeOrProprietary model.CodeOrProprietaryType
	//invoice number
	Number string
	//default value: current date
	RelatedDate model.Date
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

func ReferredDocumentInformation71From(m RemittanceDocument) pain013.ReferredDocumentInformation71 {
	var result pain013.ReferredDocumentInformation71
	if m.CodeOrProprietary != "" {
		Cd := pain013.DocumentType6Code(m.CodeOrProprietary)
		Tp := pain013.ReferredDocumentType4{
			CdOrPrtry: pain013.ReferredDocumentType3Choice{
				Cd: &Cd,
			},
		}
		result.Tp = &Tp
	}
	if m.Number != "" {
		Nb := pain013.Max35Text(m.Number)
		result.Nb = &Nb
	}
	if !isEmpty(m.RelatedDate) {
		RltdDt := m.RelatedDate.Date()
		result.RltdDt = &RltdDt
	}
	return result
}
func CreditTransferTransaction351From(m CreditTransferTransaction) pain013.CreditTransferTransaction351 {
	var result pain013.CreditTransferTransaction351
	var PmtId pain013.PaymentIdentification61
	if m.PaymentInstructionId != "" {
		InstrId := pain013.Max35Text(m.PaymentInstructionId)
		PmtId.InstrId = &InstrId
	}
	if m.PaymentEndToEndId != "" {
		PmtId.EndToEndId = pain013.Max35Text(m.PaymentEndToEndId)
	}
	if m.PaymentUniqueId != "" {
		PmtId.UETR = pain013.UUIDv4Identifier(m.PaymentUniqueId)
	}
	if !isEmpty(PmtId) {
		result.PmtId = PmtId
	}
	var PmtTpInf pain013.PaymentTypeInformation261
	if m.PayRequestType != "" {
		Prtry := pain013.LocalInstrumentFedwireFunds1(m.PayRequestType)
		PmtTpInf.LclInstrm = pain013.LocalInstrument2Choice1{
			Prtry: &Prtry,
		}
	}
	if m.PayCategoryType != "" {
		Cd := pain013.ExternalCategoryPurpose1Code(m.PayCategoryType)
		CtgyPurp := pain013.CategoryPurpose1Choice{
			Cd: &Cd,
		}
		PmtTpInf.CtgyPurp = &CtgyPurp
	}
	if !isEmpty(PmtTpInf) {
		result.PmtTpInf = PmtTpInf
	}
	if !isEmpty(m.Amount) {
		InstdAmt := pain013.ActiveCurrencyAndAmountFedwire1{
			Value: pain013.ActiveCurrencyAndAmountFedwire1SimpleType(m.Amount.Amount),
			Ccy:   pain013.ActiveCurrencyCodeFixed(m.Amount.Currency),
		}
		result.Amt = pain013.AmountType4Choice1{
			InstdAmt: &InstdAmt,
		}
	}
	if m.ChargeBearer != "" {
		result.ChrgBr = pain013.ChargeBearerType1Code(m.ChargeBearer)
	}
	if !isEmpty(m.CreditorAgent) {
		Cd := pain013.ExternalClearingSystemIdentification1CodeFixed(m.CreditorAgent.PaymentSysCode)
		CdtrAgt := pain013.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pain013.FinancialInstitutionIdentification181{
				ClrSysMmbId: pain013.ClearingSystemMemberIdentification21{
					ClrSysId: pain013.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pain013.RoutingNumberFRS1(m.CreditorAgent.PaymentSysMemberId),
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
		Othr := pain013.GenericAccountIdentification1{
			Id: pain013.Max34Text(m.CrediorAccountOtherId),
		}
		cashAcc := pain013.CashAccount38{
			Id: pain013.AccountIdentification4Choice{
				Othr: &Othr,
			},
		}
		result.CdtrAcct = &cashAcc
	}
	var RmtInf pain013.RemittanceInformation161
	if m.RemittanceInformation != "" {
		Ustrd := pain013.Max140Text(m.RemittanceInformation)
		RmtInf.Ustrd = &Ustrd
	}
	if !isEmpty(m.document) {
		var Strd []*pain013.StructuredRemittanceInformation161
		var RfrdDocInf []*pain013.ReferredDocumentInformation71
		doc := ReferredDocumentInformation71From(m.document)
		RfrdDocInf = append(RfrdDocInf, &doc)
		remitDoc := pain013.StructuredRemittanceInformation161{
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
func PostalAddress241From(a model.PostalAddress) pain013.PostalAddress241 {
	var result pain013.PostalAddress241
	if a.StreetName != "" {
		StrtNm := pain013.Max70Text(a.StreetName)
		result.StrtNm = &StrtNm
	}
	if a.BuildingNumber != "" {
		BldgNb := pain013.Max16Text(a.BuildingNumber)
		result.BldgNb = &BldgNb
	}
	if a.RoomNumber != "" {
		Room := pain013.Max70Text(a.RoomNumber)
		result.Room = &Room
	}
	if a.PostalCode != "" {
		PstCd := pain013.Max16Text(a.PostalCode)
		result.PstCd = &PstCd
	}
	if a.TownName != "" {
		result.TwnNm = pain013.Max35Text(a.TownName)
	}
	if a.Subdivision != "" {
		CtrySubDvsn := pain013.Max35Text(a.Subdivision)
		result.CtrySubDvsn = &CtrySubDvsn
	}
	if a.Country != "" {
		result.Ctry = pain013.CountryCode(a.Country)
	}
	return result
}
func PostalAddress242From(a model.PostalAddress) pain013.PostalAddress242 {
	var result pain013.PostalAddress242
	if a.StreetName != "" {
		StrtNm := pain013.Max70Text(a.StreetName)
		result.StrtNm = &StrtNm
	}
	if a.BuildingNumber != "" {
		BldgNb := pain013.Max16Text(a.BuildingNumber)
		result.BldgNb = &BldgNb
	}
	if a.RoomNumber != "" {
		Room := pain013.Max70Text(a.RoomNumber)
		result.Room = &Room
	}
	if a.PostalCode != "" {
		PstCd := pain013.Max16Text(a.PostalCode)
		result.PstCd = &PstCd
	}
	if a.TownName != "" {
		TwnNm := pain013.Max35Text(a.TownName)
		result.TwnNm = &TwnNm
	}
	if a.Subdivision != "" {
		CtrySubDvsn := pain013.Max35Text(a.Subdivision)
		result.CtrySubDvsn = &CtrySubDvsn
	}
	if a.Country != "" {
		Ctry := pain013.CountryCode(a.Country)
		result.Ctry = &Ctry
	}
	return result
}

func PartyIdentification1351From(p model.PartyIdentify) pain013.PartyIdentification1351 {
	var result pain013.PartyIdentification1351
	if p.Name != "" {
		Nm := pain013.Max140Text(p.Name)
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
func PartyIdentification1352From(p model.PartyIdentify) pain013.PartyIdentification1352 {
	var result pain013.PartyIdentification1352
	if p.Name != "" {
		Nm := pain013.Max140Text(p.Name)
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
