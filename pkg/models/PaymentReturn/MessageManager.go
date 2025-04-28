package PaymentReturn

import (
	"reflect"

	pacs004 "github.com/moov-io/fedwire20022/gen/PaymentReturn_pacs_004_001_10"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Party struct {
	//Name by which a party is known and which is usually used to identify that party.
	Name string
	//Information that locates and identifies a specific address, as defined by postal services.
	Address model.PostalAddress
}
type Reason struct {
	Reason                string
	AdditionalRequestData string
}
type ReturnChain struct {
	//Party that owes an amount of money to the (ultimate) creditor.
	Debtor Party
	//Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorOtherTypeId string
	//Financial institution servicing an account for the debtor.
	DebtorAgent model.Agent
	//Financial institution servicing an account for the creditor.
	CreditorAgent model.Agent
	//Party to which an amount of money is due.
	Creditor Party
	//Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccountOtherTypeId string
}

func BranchAndFinancialInstitutionIdentification62From(p model.Agent) (pacs004.BranchAndFinancialInstitutionIdentification62, *model.ValidateError) {
	var result pacs004.BranchAndFinancialInstitutionIdentification62
	var FinInstnId pacs004.FinancialInstitutionIdentification182
	var ClrSysMmbId pacs004.ClearingSystemMemberIdentification22
	if p.PaymentSysCode != "" {
		err := pacs004.ExternalClearingSystemIdentification1CodeFixed(p.PaymentSysCode).Validate()
		if err != nil {
			return pacs004.BranchAndFinancialInstitutionIdentification62{}, &model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
		}
		Cd := pacs004.ExternalClearingSystemIdentification1CodeFixed(p.PaymentSysCode)
		ClrSysMmbId.ClrSysId = pacs004.ClearingSystemIdentification2Choice2{
			Cd: &Cd,
		}
	}
	if p.PaymentSysMemberId != "" {
		err := pacs004.RoutingNumberFRS1(p.PaymentSysMemberId).Validate()
		if err != nil {
			return pacs004.BranchAndFinancialInstitutionIdentification62{}, &model.ValidateError{
				ParamName: "PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		ClrSysMmbId.MmbId = pacs004.RoutingNumberFRS1(p.PaymentSysMemberId)
	}
	if !isEmpty(ClrSysMmbId) {
		FinInstnId.ClrSysMmbId = ClrSysMmbId
	}
	if !isEmpty(FinInstnId) {
		result.FinInstnId = FinInstnId
	}
	return result, nil
}
func PaymentReturnReason61From(p Reason) (pacs004.PaymentReturnReason61, *model.ValidateError) {
	var result pacs004.PaymentReturnReason61
	if p.Reason != "" {
		err := pacs004.ExternalReturnReason1Code(p.Reason).Validate()
		if err != nil {
			return pacs004.PaymentReturnReason61{}, &model.ValidateError{
				ParamName: "Reason",
				Message:   err.Error(),
			}
		}
		Cd := pacs004.ExternalReturnReason1Code(p.Reason)
		result.Rsn = pacs004.ReturnReason5Choice1{
			Cd: &Cd,
		}
	}
	var AddtlInf []*pacs004.Max105Text
	if p.AdditionalRequestData != "" {
		err := pacs004.Max105Text(p.AdditionalRequestData).Validate()
		if err != nil {
			return pacs004.PaymentReturnReason61{}, &model.ValidateError{
				ParamName: "AdditionalRequestData",
				Message:   err.Error(),
			}
		}
		AddtlInfItem := pacs004.Max105Text(p.AdditionalRequestData)
		AddtlInf = append(AddtlInf, &AddtlInfItem)
	}
	if !isEmpty(AddtlInf) {
		result.AddtlInf = AddtlInf
	}
	return result, nil
}
func PostalAddress241From(param model.PostalAddress) (pacs004.PostalAddress241, *model.ValidateError) {
	var Dbtr_PstlAdr pacs004.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		err := pacs004.Max70Text(param.StreetName).Validate()
		if err != nil {
			return pacs004.PostalAddress241{}, &model.ValidateError{
				ParamName: "StreetName",
				Message:   err.Error(),
			}
		}
		StrtNm := pacs004.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		err := pacs004.Max16Text(param.BuildingNumber).Validate()
		if err != nil {
			return pacs004.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingNumber",
				Message:   err.Error(),
			}
		}
		BldgNb := pacs004.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		err := pacs004.Max35Text(param.BuildingName).Validate()
		if err != nil {
			return pacs004.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingName",
				Message:   err.Error(),
			}
		}
		BldgNm := pacs004.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		err := pacs004.Max70Text(param.Floor).Validate()
		if err != nil {
			return pacs004.PostalAddress241{}, &model.ValidateError{
				ParamName: "Floor",
				Message:   err.Error(),
			}
		}
		Floor := pacs004.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		err := pacs004.Max70Text(param.RoomNumber).Validate()
		if err != nil {
			return pacs004.PostalAddress241{}, &model.ValidateError{
				ParamName: "RoomNumber",
				Message:   err.Error(),
			}
		}
		Room := pacs004.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		err := pacs004.Max16Text(param.PostalCode).Validate()
		if err != nil {
			return pacs004.PostalAddress241{}, &model.ValidateError{
				ParamName: "PostalCode",
				Message:   err.Error(),
			}
		}
		PstCd := pacs004.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		err := pacs004.Max35Text(param.TownName).Validate()
		if err != nil {
			return pacs004.PostalAddress241{}, &model.ValidateError{
				ParamName: "TownName",
				Message:   err.Error(),
			}
		}
		TwnNm := pacs004.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		err := pacs004.Max35Text(param.Subdivision).Validate()
		if err != nil {
			return pacs004.PostalAddress241{}, &model.ValidateError{
				ParamName: "Subdivision",
				Message:   err.Error(),
			}
		}
		CtrySubDvsn := pacs004.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		err := pacs004.CountryCode(param.Country).Validate()
		if err != nil {
			return pacs004.PostalAddress241{}, &model.ValidateError{
				ParamName: "Country",
				Message:   err.Error(),
			}
		}
		Ctry := pacs004.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs004.PostalAddress241{}, nil
	}

	return Dbtr_PstlAdr, nil
}
func PartyIdentification1352From(p Party) (pacs004.PartyIdentification1352, *model.ValidateError) {
	var result pacs004.PartyIdentification1352
	if p.Name != "" {
		err := pacs004.Max140Text(p.Name).Validate()
		if err != nil {
			return pacs004.PartyIdentification1352{}, &model.ValidateError{
				ParamName: "Name",
				Message:   err.Error(),
			}
		}
		Nm := pacs004.Max140Text(p.Name)
		result.Nm = &Nm
	}
	if !isEmpty(p.Address) {
		PstlAdr, err := PostalAddress241From(p.Address)
		if err != nil {
			err.InsertPath("Address")
			return pacs004.PartyIdentification1352{}, &model.ValidateError{
				ParamName: "Name",
				Message:   err.Error(),
			}
		}
		result.PstlAdr = &PstlAdr
	}
	return result, nil
}
func TransactionParties81From(p ReturnChain) (pacs004.TransactionParties81, *model.ValidateError) {
	var result pacs004.TransactionParties81
	if !isEmpty(p.Debtor) {
		Pty, err := PartyIdentification1352From(p.Debtor)
		if err != nil {
			err.InsertPath("Debtor")
			return pacs004.TransactionParties81{}, err
		}
		result.Dbtr = pacs004.Party40Choice2{
			Pty: &Pty,
		}
	}
	if p.DebtorOtherTypeId != "" {
		err := pacs004.Max34Text(p.DebtorOtherTypeId).Validate()
		if err != nil {
			return pacs004.TransactionParties81{}, &model.ValidateError{
				ParamName: "DebtorOtherTypeId",
				Message:   err.Error(),
			}
		}
		Othr := pacs004.GenericAccountIdentification1{
			Id: pacs004.Max34Text(p.DebtorOtherTypeId),
		}
		DbtrAcct := pacs004.CashAccount38{
			Id: pacs004.AccountIdentification4Choice{
				Othr: &Othr,
			},
		}
		result.DbtrAcct = &DbtrAcct
	}
	if !isEmpty(p.DebtorAgent) {
		var DbtrAgt pacs004.BranchAndFinancialInstitutionIdentification61
		var FinInstnId pacs004.FinancialInstitutionIdentification181
		var ClrSysMmbId pacs004.ClearingSystemMemberIdentification21
		if p.DebtorAgent.PaymentSysCode != "" {
			err := pacs004.ExternalClearingSystemIdentification1Code(p.DebtorAgent.PaymentSysCode).Validate()
			if err != nil {
				return pacs004.TransactionParties81{}, &model.ValidateError{
					ParamName: "DebtorAgent.PaymentSysCode",
					Message:   err.Error(),
				}
			}
			Cd := pacs004.ExternalClearingSystemIdentification1Code(p.DebtorAgent.PaymentSysCode)
			ClrSysMmbId.ClrSysId = pacs004.ClearingSystemIdentification2Choice1{
				Cd: &Cd,
			}
		}
		if p.DebtorAgent.PaymentSysMemberId != "" {
			err := pacs004.Max35Text(p.DebtorAgent.PaymentSysMemberId).Validate()
			if err != nil {
				return pacs004.TransactionParties81{}, &model.ValidateError{
					ParamName: "DebtorAgent.PaymentSysMemberId",
					Message:   err.Error(),
				}
			}
			ClrSysMmbId.MmbId = pacs004.Max35Text(p.DebtorAgent.PaymentSysMemberId)
		}
		if !isEmpty(ClrSysMmbId) {
			FinInstnId.ClrSysMmbId = &ClrSysMmbId
		}
		if p.DebtorAgent.BankName != "" {
			err := pacs004.Max140Text(p.DebtorAgent.BankName).Validate()
			if err != nil {
				return pacs004.TransactionParties81{}, &model.ValidateError{
					ParamName: "DebtorAgent.BankName",
					Message:   err.Error(),
				}
			}
			Nm := pacs004.Max140Text(p.DebtorAgent.BankName)
			FinInstnId.Nm = &Nm
		}
		if !isEmpty(p.DebtorAgent.PostalAddress) {
			PstlAdr, err := PostalAddress241From(p.DebtorAgent.PostalAddress)
			if err != nil {
				err.InsertPath("DebtorAgent.PostalAddress")
				return pacs004.TransactionParties81{}, err
			}
			FinInstnId.PstlAdr = &PstlAdr
		}
		if !isEmpty(FinInstnId) {
			DbtrAgt.FinInstnId = FinInstnId
		}
		if !isEmpty(DbtrAgt) {
			result.DbtrAgt = &DbtrAgt
		}
	}
	if !isEmpty(p.CreditorAgent) {
		var CdtrAgt pacs004.BranchAndFinancialInstitutionIdentification63
		var FinInstnId pacs004.FinancialInstitutionIdentification181
		var ClrSysMmbId pacs004.ClearingSystemMemberIdentification21
		if p.CreditorAgent.PaymentSysCode != "" {
			err := pacs004.ExternalClearingSystemIdentification1Code(p.CreditorAgent.PaymentSysCode).Validate()
			if err != nil {
				return pacs004.TransactionParties81{}, &model.ValidateError{
					ParamName: "CreditorAgent.PaymentSysCode",
					Message:   err.Error(),
				}
			}
			Cd := pacs004.ExternalClearingSystemIdentification1Code(p.CreditorAgent.PaymentSysCode)
			ClrSysMmbId.ClrSysId = pacs004.ClearingSystemIdentification2Choice1{
				Cd: &Cd,
			}
		}
		if p.CreditorAgent.PaymentSysMemberId != "" {
			err := pacs004.Max35Text(p.CreditorAgent.PaymentSysMemberId).Validate()
			if err != nil {
				return pacs004.TransactionParties81{}, &model.ValidateError{
					ParamName: "CreditorAgent.PaymentSysMemberId",
					Message:   err.Error(),
				}
			}
			ClrSysMmbId.MmbId = pacs004.Max35Text(p.CreditorAgent.PaymentSysMemberId)
		}
		if !isEmpty(ClrSysMmbId) {
			FinInstnId.ClrSysMmbId = &ClrSysMmbId
		}
		if p.CreditorAgent.BankName != "" {
			err := pacs004.Max140Text(p.CreditorAgent.BankName).Validate()
			if err != nil {
				return pacs004.TransactionParties81{}, &model.ValidateError{
					ParamName: "CreditorAgent.BankName",
					Message:   err.Error(),
				}
			}
			Nm := pacs004.Max140Text(p.CreditorAgent.BankName)
			FinInstnId.Nm = &Nm
		}
		if !isEmpty(p.CreditorAgent.PostalAddress) {
			PstlAdr, err := PostalAddress241From(p.CreditorAgent.PostalAddress)
			if err != nil {
				err.InsertPath("CreditorAgent.PostalAddress")
				return pacs004.TransactionParties81{}, err
			}
			FinInstnId.PstlAdr = &PstlAdr
		}
		if !isEmpty(FinInstnId) {
			CdtrAgt.FinInstnId = FinInstnId
		}
		if !isEmpty(CdtrAgt) {
			result.CdtrAgt = &CdtrAgt
		}
	}
	if !isEmpty(p.Creditor) {
		Pty, err := PartyIdentification1352From(p.Creditor)
		if err != nil {
			err.InsertPath("Creditor")
			return pacs004.TransactionParties81{}, err
		}
		result.Cdtr = pacs004.Party40Choice2{
			Pty: &Pty,
		}
	}
	if p.CreditorAccountOtherTypeId != "" {
		err := pacs004.Max34Text(p.CreditorAccountOtherTypeId).Validate()
			if err != nil {
				return pacs004.TransactionParties81{}, &model.ValidateError{
					ParamName: "CreditorAccountOtherTypeId",
					Message:   err.Error(),
				}
			}
		Othr := pacs004.GenericAccountIdentification1{
			Id: pacs004.Max34Text(p.CreditorAccountOtherTypeId),
		}
		CdtrAcct := pacs004.CashAccount38{
			Id: pacs004.AccountIdentification4Choice{
				Othr: &Othr,
			},
		}
		result.CdtrAcct = &CdtrAcct
	}
	return result, nil
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
