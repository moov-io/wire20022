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

func ReferredDocumentInformation71From(m RemittanceDocument) (pain013.ReferredDocumentInformation71, *model.ValidateError) {
	var result pain013.ReferredDocumentInformation71
	if m.CodeOrProprietary != "" {
		err := pain013.DocumentType6Code(m.CodeOrProprietary).Validate()
		if err != nil {
			return pain013.ReferredDocumentInformation71{}, &model.ValidateError{
				ParamName: "CodeOrProprietary",
				Message:   err.Error(),
			}
		}
		Cd := pain013.DocumentType6Code(m.CodeOrProprietary)
		Tp := pain013.ReferredDocumentType4{
			CdOrPrtry: pain013.ReferredDocumentType3Choice{
				Cd: &Cd,
			},
		}
		result.Tp = &Tp
	}
	if m.Number != "" {
		err := pain013.Max35Text(m.Number).Validate()
		if err != nil {
			return pain013.ReferredDocumentInformation71{}, &model.ValidateError{
				ParamName: "Number",
				Message:   err.Error(),
			}
		}
		Nb := pain013.Max35Text(m.Number)
		result.Nb = &Nb
	}
	if !isEmpty(m.RelatedDate) {
		err := m.RelatedDate.Date().Validate()
		if err != nil {
			return pain013.ReferredDocumentInformation71{}, &model.ValidateError{
				ParamName: "RelatedDate",
				Message:   err.Error(),
			}
		}
		RltdDt := m.RelatedDate.Date()
		result.RltdDt = &RltdDt
	}
	return result, nil
}
func ReferredDocumentInformation71To(m pain013.ReferredDocumentInformation71) RemittanceDocument {
	var result RemittanceDocument
	if !isEmpty(m.Tp) {
		if !isEmpty(m.Tp.CdOrPrtry) {
			if !isEmpty(m.Tp.CdOrPrtry.Cd) {
				result.CodeOrProprietary = model.CodeOrProprietaryType(*m.Tp.CdOrPrtry.Cd)
			}
		}
	}
	if !isEmpty(m.Nb) {
		result.Number = string(*m.Nb)
	}
	if !isEmpty(m.RltdDt) {
		result.RelatedDate = model.FromDate(*m.RltdDt)
	}
	return result
}
func CreditTransferTransaction351From(m CreditTransferTransaction) (pain013.CreditTransferTransaction351, *model.ValidateError) {
	var result pain013.CreditTransferTransaction351
	var PmtId pain013.PaymentIdentification61
	if m.PaymentInstructionId != "" {
		err := pain013.Max35Text(m.PaymentInstructionId).Validate()
		if err != nil {
			return pain013.CreditTransferTransaction351{}, &model.ValidateError{
				ParamName: "PaymentInstructionId",
				Message:   err.Error(),
			}
		}
		InstrId := pain013.Max35Text(m.PaymentInstructionId)
		PmtId.InstrId = &InstrId
	}
	if m.PaymentEndToEndId != "" {
		err := pain013.Max35Text(m.PaymentEndToEndId).Validate()
		if err != nil {
			return pain013.CreditTransferTransaction351{}, &model.ValidateError{
				ParamName: "PaymentEndToEndId",
				Message:   err.Error(),
			}
		}
		PmtId.EndToEndId = pain013.Max35Text(m.PaymentEndToEndId)
	}
	if m.PaymentUniqueId != "" {
		err := pain013.UUIDv4Identifier(m.PaymentUniqueId).Validate()
		if err != nil {
			return pain013.CreditTransferTransaction351{}, &model.ValidateError{
				ParamName: "PaymentUniqueId",
				Message:   err.Error(),
			}
		}
		PmtId.UETR = pain013.UUIDv4Identifier(m.PaymentUniqueId)
	}
	if !isEmpty(PmtId) {
		result.PmtId = PmtId
	}
	var PmtTpInf pain013.PaymentTypeInformation261
	if m.PayRequestType != "" {
		err := pain013.LocalInstrumentFedwireFunds1(m.PayRequestType).Validate()
		if err != nil {
			return pain013.CreditTransferTransaction351{}, &model.ValidateError{
				ParamName: "PayRequestType",
				Message:   err.Error(),
			}
		}
		Prtry := pain013.LocalInstrumentFedwireFunds1(m.PayRequestType)
		PmtTpInf.LclInstrm = pain013.LocalInstrument2Choice1{
			Prtry: &Prtry,
		}
	}
	if m.PayCategoryType != "" {
		err := pain013.ExternalCategoryPurpose1Code(m.PayCategoryType).Validate()
		if err != nil {
			return pain013.CreditTransferTransaction351{}, &model.ValidateError{
				ParamName: "PayCategoryType",
				Message:   err.Error(),
			}
		}
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
		err := pain013.ActiveCurrencyAndAmountFedwire1SimpleType(m.Amount.Amount).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "Amount",
				Message:   err.Error(),
			}
			vErr.InsertPath("Amount")
			return pain013.CreditTransferTransaction351{}, &vErr
		}
		err = pain013.ActiveCurrencyCodeFixed(m.Amount.Currency).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "Currency",
				Message:   err.Error(),
			}
			vErr.InsertPath("Amount")
			return pain013.CreditTransferTransaction351{}, &vErr
		}
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
		err := pain013.ExternalClearingSystemIdentification1CodeFixed(m.CreditorAgent.PaymentSysCode).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
			vErr.InsertPath("CreditorAgent")
			return pain013.CreditTransferTransaction351{}, &vErr
		}
		err = pain013.RoutingNumberFRS1(m.CreditorAgent.PaymentSysMemberId).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysMemberId",
				Message:   err.Error(),
			}
			vErr.InsertPath("CreditorAgent")
			return pain013.CreditTransferTransaction351{}, &vErr
		}
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
		Cdtr, vErr := PartyIdentification1352From(m.Creditor)
		if vErr != nil {
			vErr.InsertPath("Creditor")
			return pain013.CreditTransferTransaction351{}, vErr
		}
		if !isEmpty(Cdtr) {
			result.Cdtr = Cdtr
		}
	}
	if m.CrediorAccountOtherId != "" {
		err := pain013.Max34Text(m.CrediorAccountOtherId).Validate()
		if err != nil {
			return pain013.CreditTransferTransaction351{}, &model.ValidateError{
				ParamName: "CrediorAccountOtherId",
				Message:   err.Error(),
			}
		}
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
		err := pain013.Max140Text(m.RemittanceInformation).Validate()
		if err != nil {
			return pain013.CreditTransferTransaction351{}, &model.ValidateError{
				ParamName: "RemittanceInformation",
				Message:   err.Error(),
			}
		}
		Ustrd := pain013.Max140Text(m.RemittanceInformation)
		RmtInf.Ustrd = &Ustrd
	}
	if !isEmpty(m.document) {
		var Strd []*pain013.StructuredRemittanceInformation161
		var RfrdDocInf []*pain013.ReferredDocumentInformation71
		doc, vErr := ReferredDocumentInformation71From(m.document)
		if vErr != nil {
			vErr.InsertPath("document")
			return pain013.CreditTransferTransaction351{}, vErr
		}
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
	return result, nil
}
func CreditTransferTransaction351To(m pain013.CreditTransferTransaction351) CreditTransferTransaction {
	var result CreditTransferTransaction
	if !isEmpty(m.PmtId) {
		if !isEmpty(m.PmtId.InstrId) {
			result.PaymentInstructionId = string(*m.PmtId.InstrId)
		}
		if !isEmpty(m.PmtId.EndToEndId) {
			result.PaymentEndToEndId = string(m.PmtId.EndToEndId)
		}
		if !isEmpty(m.PmtId.UETR) {
			result.PaymentUniqueId = string(m.PmtId.UETR)
		}
	}
	if !isEmpty(m.PmtTpInf) {
		if !isEmpty(m.PmtTpInf.LclInstrm) {
			if !isEmpty(m.PmtTpInf.LclInstrm.Prtry) {
				result.PayRequestType = PaymentRequestType(*m.PmtTpInf.LclInstrm.Prtry)
			}
		}
		if !isEmpty(m.PmtTpInf.CtgyPurp) {
			if !isEmpty(m.PmtTpInf.CtgyPurp.Cd) {
				result.PayCategoryType = PaymentRequestType(*m.PmtTpInf.CtgyPurp.Cd)
			}
		}
	}
	if !isEmpty(m.Amt) {
		if !isEmpty(m.Amt.InstdAmt) {
			result.Amount = model.CurrencyAndAmount{
				Currency: string(m.Amt.InstdAmt.Ccy),
				Amount:   float64(m.Amt.InstdAmt.Value),
			}
		}
	}
	if !isEmpty(m.ChrgBr) {
		result.ChargeBearer = ChargeBearerType(m.ChrgBr)
	}
	if !isEmpty(m.CdtrAgt) {
		if !isEmpty(m.CdtrAgt.FinInstnId) {
			if !isEmpty(m.CdtrAgt.FinInstnId.ClrSysMmbId) {
				if !isEmpty(m.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId) {
					if !isEmpty(m.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
						result.CreditorAgent.PaymentSysCode = model.PaymentSystemType(*m.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
					}
					if !isEmpty(m.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId) {
						result.CreditorAgent.PaymentSysMemberId = string(m.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId)
					}
				}
			}
		}
	}
	if !isEmpty(m.Cdtr) {
		Cdtr := PartyIdentification1352To(m.Cdtr)
		result.Creditor = Cdtr
	}
	if !isEmpty(m.CdtrAcct) {
		if !isEmpty(m.CdtrAcct.Id) {
			if !isEmpty(m.CdtrAcct.Id.Othr) {
				if !isEmpty(m.CdtrAcct.Id.Othr.Id) {
					result.CrediorAccountOtherId = string(m.CdtrAcct.Id.Othr.Id)
				}
			}
		}
	}
	if !isEmpty(m.RmtInf) {
		if !isEmpty(m.RmtInf.Ustrd) {
			result.RemittanceInformation = string(*m.RmtInf.Ustrd)
		}
		if !isEmpty(m.RmtInf.Strd) {
			if len(m.RmtInf.Strd) > 0 {
				for _, strd := range m.RmtInf.Strd {
					if !isEmpty(strd.RfrdDocInf) {
						for _, doc := range strd.RfrdDocInf {
							docInfo := ReferredDocumentInformation71To(*doc)
							result.document = docInfo
						}
					}
				}
			}
		}
	}
	return result
}
func PostalAddress241From(a model.PostalAddress) (pain013.PostalAddress241, *model.ValidateError) {
	var result pain013.PostalAddress241
	if a.StreetName != "" {
		err := pain013.Max70Text(a.StreetName).Validate()
		if err != nil {
			return pain013.PostalAddress241{}, &model.ValidateError{
				ParamName: "StreetName",
				Message:   err.Error(),
			}
		}
		StrtNm := pain013.Max70Text(a.StreetName)
		result.StrtNm = &StrtNm
	}
	if a.BuildingNumber != "" {
		err := pain013.Max16Text(a.BuildingNumber).Validate()
		if err != nil {
			return pain013.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingNumber",
				Message:   err.Error(),
			}
		}
		BldgNb := pain013.Max16Text(a.BuildingNumber)
		result.BldgNb = &BldgNb
	}
	if a.RoomNumber != "" {
		err := pain013.Max70Text(a.RoomNumber).Validate()
		if err != nil {
			return pain013.PostalAddress241{}, &model.ValidateError{
				ParamName: "RoomNumber",
				Message:   err.Error(),
			}
		}
		Room := pain013.Max70Text(a.RoomNumber)
		result.Room = &Room
	}
	if a.PostalCode != "" {
		err := pain013.Max16Text(a.PostalCode).Validate()
		if err != nil {
			return pain013.PostalAddress241{}, &model.ValidateError{
				ParamName: "PostalCode",
				Message:   err.Error(),
			}
		}
		PstCd := pain013.Max16Text(a.PostalCode)
		result.PstCd = &PstCd
	}
	if a.TownName != "" {
		err := pain013.Max35Text(a.TownName).Validate()
		if err != nil {
			return pain013.PostalAddress241{}, &model.ValidateError{
				ParamName: "TownName",
				Message:   err.Error(),
			}
		}
		result.TwnNm = pain013.Max35Text(a.TownName)
	}
	if a.Subdivision != "" {
		err := pain013.Max35Text(a.Subdivision).Validate()
		if err != nil {
			return pain013.PostalAddress241{}, &model.ValidateError{
				ParamName: "Subdivision",
				Message:   err.Error(),
			}
		}
		CtrySubDvsn := pain013.Max35Text(a.Subdivision)
		result.CtrySubDvsn = &CtrySubDvsn
	}
	if a.Country != "" {
		err := pain013.CountryCode(a.Country).Validate()
		if err != nil {
			return pain013.PostalAddress241{}, &model.ValidateError{
				ParamName: "Country",
				Message:   err.Error(),
			}
		}
		result.Ctry = pain013.CountryCode(a.Country)
	}
	return result, nil
}
func PostalAddress241To(a pain013.PostalAddress241) model.PostalAddress {
	var result model.PostalAddress
	if !isEmpty(a.StrtNm) {
		result.StreetName = string(*a.StrtNm)
	}
	if !isEmpty(a.BldgNb) {
		result.BuildingNumber = string(*a.BldgNb)
	}
	if !isEmpty(a.Room) {
		result.RoomNumber = string(*a.Room)
	}
	if !isEmpty(a.PstCd) {
		result.PostalCode = string(*a.PstCd)
	}
	if !isEmpty(a.TwnNm) {
		result.TownName = string(a.TwnNm)
	}
	if !isEmpty(a.CtrySubDvsn) {
		result.Subdivision = string(*a.CtrySubDvsn)
	}
	if !isEmpty(a.Ctry) {
		result.Country = string(a.Ctry)
	}
	return result
}
func PostalAddress242From(a model.PostalAddress) (pain013.PostalAddress242, *model.ValidateError) {
	var result pain013.PostalAddress242
	if a.StreetName != "" {
		err := pain013.Max70Text(a.StreetName).Validate()
		if err != nil {
			return pain013.PostalAddress242{}, &model.ValidateError{
				ParamName: "StreetName",
				Message:   err.Error(),
			}
		}
		StrtNm := pain013.Max70Text(a.StreetName)
		result.StrtNm = &StrtNm
	}
	if a.BuildingNumber != "" {
		err := pain013.Max16Text(a.BuildingNumber).Validate()
		if err != nil {
			return pain013.PostalAddress242{}, &model.ValidateError{
				ParamName: "BuildingNumber",
				Message:   err.Error(),
			}
		}
		BldgNb := pain013.Max16Text(a.BuildingNumber)
		result.BldgNb = &BldgNb
	}
	if a.RoomNumber != "" {
		err := pain013.Max70Text(a.RoomNumber).Validate()
		if err != nil {
			return pain013.PostalAddress242{}, &model.ValidateError{
				ParamName: "RoomNumber",
				Message:   err.Error(),
			}
		}
		Room := pain013.Max70Text(a.RoomNumber)
		result.Room = &Room
	}
	if a.PostalCode != "" {
		err := pain013.Max16Text(a.PostalCode).Validate()
		if err != nil {
			return pain013.PostalAddress242{}, &model.ValidateError{
				ParamName: "PostalCode",
				Message:   err.Error(),
			}
		}
		PstCd := pain013.Max16Text(a.PostalCode)
		result.PstCd = &PstCd
	}
	if a.TownName != "" {
		err := pain013.Max35Text(a.TownName).Validate()
		if err != nil {
			return pain013.PostalAddress242{}, &model.ValidateError{
				ParamName: "TownName",
				Message:   err.Error(),
			}
		}
		TwnNm := pain013.Max35Text(a.TownName)
		result.TwnNm = &TwnNm
	}
	if a.Subdivision != "" {
		err := pain013.Max35Text(a.Subdivision).Validate()
		if err != nil {
			return pain013.PostalAddress242{}, &model.ValidateError{
				ParamName: "Subdivision",
				Message:   err.Error(),
			}
		}
		CtrySubDvsn := pain013.Max35Text(a.Subdivision)
		result.CtrySubDvsn = &CtrySubDvsn
	}
	if a.Country != "" {
		err := pain013.CountryCode(a.Country).Validate()
		if err != nil {
			return pain013.PostalAddress242{}, &model.ValidateError{
				ParamName: "Country",
				Message:   err.Error(),
			}
		}
		Ctry := pain013.CountryCode(a.Country)
		result.Ctry = &Ctry
	}
	return result, nil
}
func PostalAddress242To(a pain013.PostalAddress242) model.PostalAddress {
	var result model.PostalAddress
	if !isEmpty(a.StrtNm) {
		result.StreetName = string(*a.StrtNm)
	}
	if !isEmpty(a.BldgNb) {
		result.BuildingNumber = string(*a.BldgNb)
	}
	if !isEmpty(a.Room) {
		result.RoomNumber = string(*a.Room)
	}
	if !isEmpty(a.PstCd) {
		result.PostalCode = string(*a.PstCd)
	}
	if !isEmpty(a.TwnNm) {
		result.TownName = string(*a.TwnNm)
	}
	if !isEmpty(a.CtrySubDvsn) {
		result.Subdivision = string(*a.CtrySubDvsn)
	}
	if !isEmpty(a.Ctry) {
		result.Country = string(*a.Ctry)
	}
	return result
}
func PartyIdentification1351From(p model.PartyIdentify) (pain013.PartyIdentification1351, *model.ValidateError) {
	var result pain013.PartyIdentification1351
	if p.Name != "" {
		err := pain013.Max140Text(p.Name).Validate()
		if err != nil {
			return pain013.PartyIdentification1351{}, &model.ValidateError{
				ParamName: "Name",
				Message:   err.Error(),
			}
		}
		Nm := pain013.Max140Text(p.Name)
		result.Nm = &Nm
	}
	if !isEmpty(p.Address) {
		PstlAdr, vErr := PostalAddress241From(p.Address)
		if vErr != nil {
			vErr.InsertPath("Address")
			return pain013.PartyIdentification1351{}, vErr
		}
		if !isEmpty(PstlAdr) {
			result.PstlAdr = &PstlAdr
		}
	}
	return result, nil
}
func PartyIdentification1351To(p pain013.PartyIdentification1351) model.PartyIdentify {
	var result model.PartyIdentify
	if !isEmpty(p.Nm) {
		result.Name = string(*p.Nm)
	}
	if !isEmpty(p.PstlAdr) {
		result.Address = PostalAddress241To(*p.PstlAdr)
	}
	return result
}
func PartyIdentification1352From(p model.PartyIdentify) (pain013.PartyIdentification1352, *model.ValidateError) {
	var result pain013.PartyIdentification1352
	if p.Name != "" {
		err := pain013.Max140Text(p.Name).Validate()
		if err != nil {
			return pain013.PartyIdentification1352{}, &model.ValidateError{
				ParamName: "Name",
				Message:   err.Error(),
			}
		}
		Nm := pain013.Max140Text(p.Name)
		result.Nm = &Nm
	}
	if !isEmpty(p.Address) {
		PstlAdr, vErr := PostalAddress242From(p.Address)
		if vErr != nil {
			vErr.InsertPath("Address")
			return pain013.PartyIdentification1352{}, vErr
		}
		if !isEmpty(PstlAdr) {
			result.PstlAdr = &PstlAdr
		}
	}
	return result, nil
}
func PartyIdentification1352To(p pain013.PartyIdentification1352) model.PartyIdentify {
	var result model.PartyIdentify
	if !isEmpty(p.Nm) {
		result.Name = string(*p.Nm)
	}
	if !isEmpty(p.PstlAdr) {
		result.Address = PostalAddress242To(*p.PstlAdr)
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
