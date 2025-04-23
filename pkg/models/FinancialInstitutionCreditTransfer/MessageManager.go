package FinancialInstitutionCreditTransfer

import (
	"reflect"

	pacs009 "github.com/moov-io/fedwire20022/gen/FinancialInstitutionCreditTransfer_pacs_009_001_08"
	model "github.com/moov-io/wire20022/pkg/models"
)

type InstrumentType string

const (
	BankDrawdownRequest            InstrumentType = "DRRB"
	BankDrawdownTransfer           InstrumentType = "BTRD"
	CoreBankTransfer               InstrumentType = "BTRC"
	CoreCoverPayment               InstrumentType = "COVC"
	CoreCustomerTransfer           InstrumentType = "CTRC"
	CustomerDrawdownRequest        InstrumentType = "DRRC"
	CustomerDrawdownTransfer       InstrumentType = "CTRD"
	SpecialAccountBankTransfer     InstrumentType = "BTRS"
	SpecialAccountCoverPayment     InstrumentType = "COVS"
	SpecialAccountCustomerTransfer InstrumentType = "CTRS"
)

type CreditTransferTransaction struct {
	//Party that owes an amount of money to the (ultimate) creditor.
	Debtor model.FiniancialInstitutionId
	//Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount string
	//Financial institution servicing an account for the debtor.
	DebtorAgent model.FiniancialInstitutionId
	//Financial institution servicing an account for the creditor.
	CreditorAgent model.FiniancialInstitutionId
	//Party to which an amount of money is due.
	Creditor model.FiniancialInstitutionId
	//Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount string
	//Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation string
	//Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount model.CurrencyAndAmount
}

func PostalAddress241From(param model.PostalAddress) pacs009.PostalAddress241 {
	var Dbtr_PstlAdr pacs009.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := pacs009.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := pacs009.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := pacs009.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := pacs009.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := pacs009.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := pacs009.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		TwnNm := pacs009.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := pacs009.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Ctry := pacs009.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs009.PostalAddress241{}
	}

	return Dbtr_PstlAdr
}
func CreditTransferTransaction371From(param CreditTransferTransaction) pacs009.CreditTransferTransaction371 {
	var result pacs009.CreditTransferTransaction371
	if !isEmpty(param.Debtor) {
		var agent pacs009.PartyIdentification1352
		if param.Debtor.Name != "" {
			Nm := pacs009.Max140Text(param.Debtor.Name)
			agent.Nm = &Nm
		}
		if !isEmpty(param.Debtor.Address) {
			PstlAdr := PostalAddress241From(param.Debtor.Address)
			agent.PstlAdr = &PstlAdr
		}

		var finialialId pacs009.FinancialInstitutionIdentification181
		if param.Debtor.BusinessId != "" {
			BICFI := pacs009.BICFIDec2014Identifier(param.Debtor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if !isEmpty(agent) {
			result.Dbtr = agent
		}
	}
	if param.DebtorAccount != "" {
		IBAN := pacs009.IBAN2007Identifier(param.DebtorAccount)
		DbtrAcct := pacs009.CashAccount38{
			Id: pacs009.AccountIdentification4Choice{
				IBAN: &IBAN,
			},
		}
		result.DbtrAcct = &DbtrAcct
	}
	if !isEmpty(param.DebtorAgent) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if param.DebtorAgent.BusinessId != "" {
			BICFI := pacs009.BICFIDec2014Identifier(param.DebtorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if param.DebtorAgent.ClearingSystemId != "" {
			Cd := pacs009.ExternalClearingSystemIdentification1Code(param.DebtorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if param.DebtorAgent.Name != "" {
			Nm := pacs009.Max140Text(param.DebtorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(param.DebtorAgent.Address) {
			PstlAdr := PostalAddress241From(param.DebtorAgent.Address)
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(agent) {
			result.DbtrAgt = agent
		}
	}
	if !isEmpty(param.CreditorAgent) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification63
		var finialialId pacs009.FinancialInstitutionIdentification181
		if param.CreditorAgent.BusinessId != "" {
			BICFI := pacs009.BICFIDec2014Identifier(param.CreditorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if param.CreditorAgent.ClearingSystemId != "" {
			Cd := pacs009.ExternalClearingSystemIdentification1Code(param.CreditorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if param.CreditorAgent.Name != "" {
			Nm := pacs009.Max140Text(param.CreditorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(param.CreditorAgent.Address) {
			PstlAdr := PostalAddress241From(param.CreditorAgent.Address)
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(agent) {
			result.CdtrAgt = agent
		}
	}
	if !isEmpty(param.Creditor) {
		var agent pacs009.PartyIdentification1352
		if param.Creditor.Name != "" {
			Nm := pacs009.Max140Text(param.Creditor.Name)
			agent.Nm = &Nm
		}
		if !isEmpty(param.Creditor.Address) {
			PstlAdr := PostalAddress241From(param.Creditor.Address)
			agent.PstlAdr = &PstlAdr
		}

		var finialialId pacs009.FinancialInstitutionIdentification181
		if param.Creditor.BusinessId != "" {
			BICFI := pacs009.BICFIDec2014Identifier(param.Creditor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if !isEmpty(agent) {
			result.Cdtr = agent
		}
	}
	if param.CreditorAccount != "" {
		IBAN := pacs009.IBAN2007Identifier(param.CreditorAccount)
		CdtrAcct := pacs009.CashAccount38{
			Id: pacs009.AccountIdentification4Choice{
				IBAN: &IBAN,
			},
		}
		result.CdtrAcct = &CdtrAcct
	}
	if param.RemittanceInformation != "" {
		Ustrd := pacs009.Max140Text(param.RemittanceInformation)
		RmtInf := pacs009.RemittanceInformation161{
			Ustrd: &Ustrd,
		}
		result.RmtInf = &RmtInf
	}
	if !isEmpty(param.InstructedAmount) {
		InstdAmt := pacs009.ActiveOrHistoricCurrencyAndAmount{
			Value: pacs009.ActiveOrHistoricCurrencyAndAmountSimpleType(param.InstructedAmount.Amount),
			Ccy:   pacs009.ActiveOrHistoricCurrencyCode(param.InstructedAmount.Currency),
		}
		result.InstdAmt = &InstdAmt
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
