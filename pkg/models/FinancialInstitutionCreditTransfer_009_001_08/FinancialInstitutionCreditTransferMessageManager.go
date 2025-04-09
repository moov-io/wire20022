package FinancialInstitutionCreditTransfer_009_001_08

import (
	"reflect"

	FinancialInstitutionCreditTransfer "github.com/moov-io/fedwire20022/gen/FinancialInstitutionCreditTransfer_pacs_009_001_08"
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

func PostalAddress241From(param model.PostalAddress) FinancialInstitutionCreditTransfer.PostalAddress241 {
	var Dbtr_PstlAdr FinancialInstitutionCreditTransfer.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := FinancialInstitutionCreditTransfer.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := FinancialInstitutionCreditTransfer.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := FinancialInstitutionCreditTransfer.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := FinancialInstitutionCreditTransfer.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := FinancialInstitutionCreditTransfer.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := FinancialInstitutionCreditTransfer.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		TwnNm := FinancialInstitutionCreditTransfer.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := FinancialInstitutionCreditTransfer.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Ctry := FinancialInstitutionCreditTransfer.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return FinancialInstitutionCreditTransfer.PostalAddress241{}
	}

	return Dbtr_PstlAdr
}
func CreditTransferTransaction371From(param CreditTransferTransaction) FinancialInstitutionCreditTransfer.CreditTransferTransaction371 {
	var result FinancialInstitutionCreditTransfer.CreditTransferTransaction371
	if !isEmpty(param.Debtor) {
		var agent FinancialInstitutionCreditTransfer.PartyIdentification1352
		if param.Debtor.Name != "" {
			Nm := FinancialInstitutionCreditTransfer.Max140Text(param.Debtor.Name)
			agent.Nm = &Nm
		}
		if !isEmpty(param.Debtor.Address) {
			PstlAdr := PostalAddress241From(param.Debtor.Address)
			agent.PstlAdr = &PstlAdr
		}

		var finialialId FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification181
		if param.Debtor.BusinessId != "" {
			BICFI := FinancialInstitutionCreditTransfer.BICFIDec2014Identifier(param.Debtor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if !isEmpty(agent) {
			result.Dbtr = agent
		}
	}
	if param.DebtorAccount != "" {
		IBAN := FinancialInstitutionCreditTransfer.IBAN2007Identifier(param.DebtorAccount)
		DbtrAcct := FinancialInstitutionCreditTransfer.CashAccount38{
			Id: FinancialInstitutionCreditTransfer.AccountIdentification4Choice{
				IBAN: &IBAN,
			},
		}
		result.DbtrAcct = &DbtrAcct
	}
	if !isEmpty(param.DebtorAgent) {
		var agent FinancialInstitutionCreditTransfer.BranchAndFinancialInstitutionIdentification61
		var finialialId FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification181
		if param.DebtorAgent.BusinessId != "" {
			BICFI := FinancialInstitutionCreditTransfer.BICFIDec2014Identifier(param.DebtorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if param.DebtorAgent.ClearingSystemId != "" {
			Cd := FinancialInstitutionCreditTransfer.ExternalClearingSystemIdentification1Code(param.DebtorAgent.ClearingSystemId)
			ClrSysMmbId := FinancialInstitutionCreditTransfer.ClearingSystemMemberIdentification21{
				ClrSysId: FinancialInstitutionCreditTransfer.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if param.DebtorAgent.Name != "" {
			Nm := FinancialInstitutionCreditTransfer.Max140Text(param.DebtorAgent.Name)
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
		var agent FinancialInstitutionCreditTransfer.BranchAndFinancialInstitutionIdentification63
		var finialialId FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification181
		if param.CreditorAgent.BusinessId != "" {
			BICFI := FinancialInstitutionCreditTransfer.BICFIDec2014Identifier(param.CreditorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if param.CreditorAgent.ClearingSystemId != "" {
			Cd := FinancialInstitutionCreditTransfer.ExternalClearingSystemIdentification1Code(param.CreditorAgent.ClearingSystemId)
			ClrSysMmbId := FinancialInstitutionCreditTransfer.ClearingSystemMemberIdentification21{
				ClrSysId: FinancialInstitutionCreditTransfer.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if param.CreditorAgent.Name != "" {
			Nm := FinancialInstitutionCreditTransfer.Max140Text(param.CreditorAgent.Name)
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
		var agent FinancialInstitutionCreditTransfer.PartyIdentification1352
		if param.Creditor.Name != "" {
			Nm := FinancialInstitutionCreditTransfer.Max140Text(param.Creditor.Name)
			agent.Nm = &Nm
		}
		if !isEmpty(param.Creditor.Address) {
			PstlAdr := PostalAddress241From(param.Creditor.Address)
			agent.PstlAdr = &PstlAdr
		}

		var finialialId FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification181
		if param.Creditor.BusinessId != "" {
			BICFI := FinancialInstitutionCreditTransfer.BICFIDec2014Identifier(param.Creditor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if !isEmpty(agent) {
			result.Cdtr = agent
		}
	}
	if param.CreditorAccount != "" {
		IBAN := FinancialInstitutionCreditTransfer.IBAN2007Identifier(param.CreditorAccount)
		CdtrAcct := FinancialInstitutionCreditTransfer.CashAccount38{
			Id: FinancialInstitutionCreditTransfer.AccountIdentification4Choice{
				IBAN: &IBAN,
			},
		}
		result.CdtrAcct = &CdtrAcct
	}
	if param.RemittanceInformation != "" {
		Ustrd := FinancialInstitutionCreditTransfer.Max140Text(param.RemittanceInformation)
		RmtInf := FinancialInstitutionCreditTransfer.RemittanceInformation161{
			Ustrd: &Ustrd,
		}
		result.RmtInf = &RmtInf
	}
	if !isEmpty(param.InstructedAmount) {
		InstdAmt := FinancialInstitutionCreditTransfer.ActiveOrHistoricCurrencyAndAmount{
			Value: FinancialInstitutionCreditTransfer.ActiveOrHistoricCurrencyAndAmountSimpleType(param.InstructedAmount.Amount),
			Ccy:   FinancialInstitutionCreditTransfer.ActiveOrHistoricCurrencyCode(param.InstructedAmount.Currency),
		}
		result.InstdAmt = &InstdAmt
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
