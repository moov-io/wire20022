package FinancialInstitutionCreditTransfer

import (
	"reflect"

	pacs009 "github.com/moov-io/fedwire20022/gen/FinancialInstitutionCreditTransfer_pacs_009_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
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

func PostalAddress241From(param model.PostalAddress) (pacs009.PostalAddress241, *model.ValidateError) {
	var Dbtr_PstlAdr pacs009.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		err := pacs009.Max70Text(param.StreetName).Validate()
		if err != nil {
			return pacs009.PostalAddress241{}, &model.ValidateError{
				ParamName: "StreetName",
				Message:   err.Error(),
			}
		}
		StrtNm := pacs009.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		err := pacs009.Max16Text(param.BuildingNumber).Validate()
		if err != nil {
			return pacs009.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingNumber",
				Message:   err.Error(),
			}
		}
		BldgNb := pacs009.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		err := pacs009.Max35Text(param.BuildingName).Validate()
		if err != nil {
			return pacs009.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingName",
				Message:   err.Error(),
			}
		}
		BldgNm := pacs009.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		err := pacs009.Max70Text(param.Floor).Validate()
		if err != nil {
			return pacs009.PostalAddress241{}, &model.ValidateError{
				ParamName: "Floor",
				Message:   err.Error(),
			}
		}
		Floor := pacs009.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		err := pacs009.Max70Text(param.RoomNumber).Validate()
		if err != nil {
			return pacs009.PostalAddress241{}, &model.ValidateError{
				ParamName: "RoomNumber",
				Message:   err.Error(),
			}
		}
		Room := pacs009.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		err := pacs009.Max16Text(param.PostalCode).Validate()
		if err != nil {
			return pacs009.PostalAddress241{}, &model.ValidateError{
				ParamName: "PostalCode",
				Message:   err.Error(),
			}
		}
		PstCd := pacs009.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		err := pacs009.Max35Text(param.TownName).Validate()
		if err != nil {
			return pacs009.PostalAddress241{}, &model.ValidateError{
				ParamName: "TownName",
				Message:   err.Error(),
			}
		}
		TwnNm := pacs009.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		err := pacs009.Max35Text(param.Subdivision).Validate()
		if err != nil {
			return pacs009.PostalAddress241{}, &model.ValidateError{
				ParamName: "Subdivision",
				Message:   err.Error(),
			}
		}
		CtrySubDvsn := pacs009.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		err := pacs009.CountryCode(param.Country).Validate()
		if err != nil {
			return pacs009.PostalAddress241{}, &model.ValidateError{
				ParamName: "Country",
				Message:   err.Error(),
			}
		}
		Ctry := pacs009.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs009.PostalAddress241{}, nil
	}

	return Dbtr_PstlAdr, nil
}
func PostalAddress241To(param pacs009.PostalAddress241) model.PostalAddress {
	var result model.PostalAddress
	if param.StrtNm != nil {
		result.StreetName = string(*param.StrtNm)
	}
	if param.BldgNb != nil {
		result.BuildingNumber = string(*param.BldgNb)
	}
	if param.BldgNm != nil {
		result.BuildingName = string(*param.BldgNm)
	}
	if param.Flr != nil {
		result.Floor = string(*param.Flr)
	}
	if param.Room != nil {
		result.RoomNumber = string(*param.Room)
	}
	if param.PstCd != nil {
		result.PostalCode = string(*param.PstCd)
	}
	if param.TwnNm != nil {
		result.TownName = string(*param.TwnNm)
	}
	if param.CtrySubDvsn != nil {
		result.Subdivision = string(*param.CtrySubDvsn)
	}
	if param.Ctry != nil {
		result.Country = string(*param.Ctry)
	}
	return result
}
func CreditTransferTransaction371From(param CreditTransferTransaction) (pacs009.CreditTransferTransaction371, *model.ValidateError) {
	var result pacs009.CreditTransferTransaction371
	if !isEmpty(param.Debtor) {
		var agent pacs009.PartyIdentification1352
		if param.Debtor.Name != "" {
			err := pacs009.Max140Text(param.Debtor.Name).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "Debtor.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(param.Debtor.Name)
			agent.Nm = &Nm
		}
		if !isEmpty(param.Debtor.Address) {
			PstlAdr, vErr := PostalAddress241From(param.Debtor.Address)
			if vErr != nil {
				vErr.InsertPath("Debtor.Address")
				return pacs009.CreditTransferTransaction371{}, vErr
			}
			agent.PstlAdr = &PstlAdr
		}

		var finialialId pacs009.FinancialInstitutionIdentification181
		if param.Debtor.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(param.Debtor.BusinessId).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "Debtor.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(param.Debtor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if !isEmpty(agent) {
			result.Dbtr = agent
		}
	}
	if param.DebtorAccount != "" {
		err := pacs009.IBAN2007Identifier(param.DebtorAccount).Validate()
		if err != nil {
			return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
				ParamName: "DebtorAccount",
				Message:   err.Error(),
			}
		}
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
			err := pacs009.BICFIDec2014Identifier(param.DebtorAgent.BusinessId).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "DebtorAgent.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(param.DebtorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if param.DebtorAgent.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(param.DebtorAgent.ClearingSystemId).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "DebtorAgent.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(param.DebtorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if param.DebtorAgent.Name != "" {
			err := pacs009.Max140Text(param.DebtorAgent.Name).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "DebtorAgent.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(param.DebtorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(param.DebtorAgent.Address) {
			PstlAdr, vErr := PostalAddress241From(param.DebtorAgent.Address)
			if vErr != nil {
				vErr.InsertPath("DebtorAgent.Address")
				return pacs009.CreditTransferTransaction371{}, vErr
			}
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
			err := pacs009.BICFIDec2014Identifier(param.CreditorAgent.BusinessId).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "CreditorAgent.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(param.CreditorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if param.CreditorAgent.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(param.CreditorAgent.ClearingSystemId).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "CreditorAgent.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(param.CreditorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if param.CreditorAgent.Name != "" {
			err := pacs009.Max140Text(param.CreditorAgent.Name).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "CreditorAgent.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(param.CreditorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(param.CreditorAgent.Address) {
			PstlAdr, vErr := PostalAddress241From(param.CreditorAgent.Address)
			if vErr != nil {
				vErr.InsertPath("CreditorAgent.Address")
				return pacs009.CreditTransferTransaction371{}, vErr
			}
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
			err := pacs009.Max140Text(param.Creditor.Name).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "Creditor.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(param.Creditor.Name)
			agent.Nm = &Nm
		}
		if !isEmpty(param.Creditor.Address) {
			PstlAdr, vErr := PostalAddress241From(param.Creditor.Address)
			if vErr != nil {
				vErr.InsertPath("Creditor.Address")
				return pacs009.CreditTransferTransaction371{}, vErr
			}
			agent.PstlAdr = &PstlAdr
		}

		var finialialId pacs009.FinancialInstitutionIdentification181
		if param.Creditor.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(param.Creditor.BusinessId).Validate()
			if err != nil {
				return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
					ParamName: "Creditor.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(param.Creditor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if !isEmpty(agent) {
			result.Cdtr = agent
		}
	}
	if param.CreditorAccount != "" {
		err := pacs009.IBAN2007Identifier(param.CreditorAccount).Validate()
		if err != nil {
			return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
				ParamName: "CreditorAccount",
				Message:   err.Error(),
			}
		}
		IBAN := pacs009.IBAN2007Identifier(param.CreditorAccount)
		CdtrAcct := pacs009.CashAccount38{
			Id: pacs009.AccountIdentification4Choice{
				IBAN: &IBAN,
			},
		}
		result.CdtrAcct = &CdtrAcct
	}
	if param.RemittanceInformation != "" {
		err := pacs009.Max140Text(param.RemittanceInformation).Validate()
		if err != nil {
			return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
				ParamName: "RemittanceInformation",
				Message:   err.Error(),
			}
		}
		Ustrd := pacs009.Max140Text(param.RemittanceInformation)
		RmtInf := pacs009.RemittanceInformation161{
			Ustrd: &Ustrd,
		}
		result.RmtInf = &RmtInf
	}
	if !isEmpty(param.InstructedAmount) {
		err := fedwire.Amount(param.InstructedAmount.Amount).Validate()
		if err != nil {
			return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
				ParamName: "InstructedAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = pacs009.ActiveOrHistoricCurrencyCode(param.InstructedAmount.Currency).Validate()
		if err != nil {
			return pacs009.CreditTransferTransaction371{}, &model.ValidateError{
				ParamName: "InstructedAmount.Currency",
				Message:   err.Error(),
			}
		}
		InstdAmt := pacs009.ActiveOrHistoricCurrencyAndAmount{
			Value: pacs009.ActiveOrHistoricCurrencyAndAmountSimpleType(param.InstructedAmount.Amount),
			Ccy:   pacs009.ActiveOrHistoricCurrencyCode(param.InstructedAmount.Currency),
		}
		result.InstdAmt = &InstdAmt
	}
	return result, nil
}
func CreditTransferTransaction371To(param pacs009.CreditTransferTransaction371) CreditTransferTransaction {
	var result CreditTransferTransaction
	if !isEmpty(param.Dbtr) {
		if !isEmpty(param.Dbtr.Nm) {
			result.Debtor.Name = string(*param.Dbtr.Nm)
		}
		if param.Dbtr.PstlAdr != nil {
			result.Debtor.Address = PostalAddress241To(*param.Dbtr.PstlAdr)
		}
	}
	if !isEmpty(param.DbtrAcct) {
		if param.DbtrAcct.Id.IBAN != nil {
			result.DebtorAccount = string(*param.DbtrAcct.Id.IBAN)
		}
	}
	if !isEmpty(param.DbtrAgt) {
		if !isEmpty(param.DbtrAgt.FinInstnId) {
			if !isEmpty(param.DbtrAgt.FinInstnId.BICFI) {
				result.DebtorAgent.BusinessId = string(*param.DbtrAgt.FinInstnId.BICFI)
			}
			if !isEmpty(param.DbtrAgt.FinInstnId.ClrSysMmbId) {
				if !isEmpty(param.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
					result.DebtorAgent.ClearingSystemId = model.PaymentSystemType(*param.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
				}
			}
			if !isEmpty(param.DbtrAgt.FinInstnId.Nm) {
				result.DebtorAgent.Name = string(*param.DbtrAgt.FinInstnId.Nm)
			}
			if param.DbtrAgt.FinInstnId.PstlAdr != nil {
				result.DebtorAgent.Address = PostalAddress241To(*param.DbtrAgt.FinInstnId.PstlAdr)
			}
		}
	}
	if !isEmpty(param.CdtrAgt) {
		if !isEmpty(param.CdtrAgt.FinInstnId) {
			if !isEmpty(param.CdtrAgt.FinInstnId.BICFI) {
				result.CreditorAgent.BusinessId = string(*param.CdtrAgt.FinInstnId.BICFI)
			}
			if !isEmpty(param.CdtrAgt.FinInstnId.ClrSysMmbId) {
				if !isEmpty(param.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
					result.CreditorAgent.ClearingSystemId = model.PaymentSystemType(*param.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
				}
			}
			if !isEmpty(param.CdtrAgt.FinInstnId.Nm) {
				result.CreditorAgent.Name = string(*param.CdtrAgt.FinInstnId.Nm)
			}
			if param.CdtrAgt.FinInstnId.PstlAdr != nil {
				result.CreditorAgent.Address = PostalAddress241To(*param.CdtrAgt.FinInstnId.PstlAdr)
			}
		}
	}
	if !isEmpty(param.Cdtr) {
		if !isEmpty(param.Cdtr.Nm) {
			result.Creditor.Name = string(*param.Cdtr.Nm)
		}
		if param.Cdtr.PstlAdr != nil {
			result.Creditor.Address = PostalAddress241To(*param.Cdtr.PstlAdr)
		}
	}
	if !isEmpty(param.CdtrAcct) {
		if param.CdtrAcct.Id.IBAN != nil {
			result.CreditorAccount = string(*param.CdtrAcct.Id.IBAN)
		}
	}
	if !isEmpty(param.RmtInf) {
		if !isEmpty(param.RmtInf.Ustrd) {
			result.RemittanceInformation = string(*param.RmtInf.Ustrd)
		}
	}
	if !isEmpty(param.InstdAmt) {
		result.InstructedAmount.Amount = float64(param.InstdAmt.Value)
		result.InstructedAmount.Currency = string(param.InstdAmt.Ccy)
	}

	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
