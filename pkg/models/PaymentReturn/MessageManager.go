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

func BranchAndFinancialInstitutionIdentification62From(p model.Agent) pacs004.BranchAndFinancialInstitutionIdentification62 {
	var result pacs004.BranchAndFinancialInstitutionIdentification62
	var FinInstnId pacs004.FinancialInstitutionIdentification182
	var ClrSysMmbId pacs004.ClearingSystemMemberIdentification22
	if p.PaymentSysCode != "" {
		Cd := pacs004.ExternalClearingSystemIdentification1CodeFixed(p.PaymentSysCode)
		ClrSysMmbId.ClrSysId = pacs004.ClearingSystemIdentification2Choice2{
			Cd: &Cd,
		}
	}
	if p.PaymentSysMemberId != "" {
		ClrSysMmbId.MmbId = pacs004.RoutingNumberFRS1(p.PaymentSysMemberId)
	}
	if !isEmpty(ClrSysMmbId) {
		FinInstnId.ClrSysMmbId = ClrSysMmbId
	}
	if !isEmpty(FinInstnId) {
		result.FinInstnId = FinInstnId
	}
	return result
}
func PaymentReturnReason61From(p Reason) pacs004.PaymentReturnReason61 {
	var result pacs004.PaymentReturnReason61
	if p.Reason != "" {
		Cd := pacs004.ExternalReturnReason1Code(p.Reason)
		result.Rsn = pacs004.ReturnReason5Choice1{
			Cd: &Cd,
		}
	}
	var AddtlInf []*pacs004.Max105Text
	if p.AdditionalRequestData != "" {
		AddtlInfItem := pacs004.Max105Text(p.AdditionalRequestData)
		AddtlInf = append(AddtlInf, &AddtlInfItem)
	}
	if !isEmpty(AddtlInf) {
		result.AddtlInf = AddtlInf
	}
	return result
}
func PostalAddress241From(param model.PostalAddress) pacs004.PostalAddress241 {
	var Dbtr_PstlAdr pacs004.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := pacs004.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := pacs004.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := pacs004.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := pacs004.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := pacs004.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := pacs004.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		TwnNm := pacs004.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := pacs004.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Ctry := pacs004.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs004.PostalAddress241{}
	}

	return Dbtr_PstlAdr
}
func PartyIdentification1352From(p Party) pacs004.PartyIdentification1352 {
	var result pacs004.PartyIdentification1352
	if p.Name != "" {
		Nm := pacs004.Max140Text(p.Name)
		result.Nm = &Nm
	}
	if !isEmpty(p.Address) {
		PstlAdr := PostalAddress241From(p.Address)
		result.PstlAdr = &PstlAdr
	}
	return result
}
func TransactionParties81From(p ReturnChain) pacs004.TransactionParties81 {
	var result pacs004.TransactionParties81
	if !isEmpty(p.Debtor) {
		Pty := PartyIdentification1352From(p.Debtor)
		result.Dbtr = pacs004.Party40Choice2{
			Pty: &Pty,
		}
	}
	if p.DebtorOtherTypeId != "" {
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
			Cd := pacs004.ExternalClearingSystemIdentification1Code(p.DebtorAgent.PaymentSysCode)
			ClrSysMmbId.ClrSysId = pacs004.ClearingSystemIdentification2Choice1{
				Cd: &Cd,
			}
		}
		if p.DebtorAgent.PaymentSysMemberId != "" {
			ClrSysMmbId.MmbId = pacs004.Max35Text(p.DebtorAgent.PaymentSysMemberId)
		}
		if !isEmpty(ClrSysMmbId) {
			FinInstnId.ClrSysMmbId = &ClrSysMmbId
		}
		if p.DebtorAgent.BankName != "" {
			Nm := pacs004.Max140Text(p.DebtorAgent.BankName)
			FinInstnId.Nm = &Nm
		}
		if !isEmpty(p.DebtorAgent.PostalAddress) {
			PstlAdr := PostalAddress241From(p.DebtorAgent.PostalAddress)
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
			Cd := pacs004.ExternalClearingSystemIdentification1Code(p.CreditorAgent.PaymentSysCode)
			ClrSysMmbId.ClrSysId = pacs004.ClearingSystemIdentification2Choice1{
				Cd: &Cd,
			}
		}
		if p.CreditorAgent.PaymentSysMemberId != "" {
			ClrSysMmbId.MmbId = pacs004.Max35Text(p.CreditorAgent.PaymentSysMemberId)
		}
		if !isEmpty(ClrSysMmbId) {
			FinInstnId.ClrSysMmbId = &ClrSysMmbId
		}
		if p.CreditorAgent.BankName != "" {
			Nm := pacs004.Max140Text(p.CreditorAgent.BankName)
			FinInstnId.Nm = &Nm
		}
		if !isEmpty(p.CreditorAgent.PostalAddress) {
			PstlAdr := PostalAddress241From(p.CreditorAgent.PostalAddress)
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
		Pty := PartyIdentification1352From(p.Creditor)
		result.Cdtr = pacs004.Party40Choice2{
			Pty: &Pty,
		}
	}
	if p.CreditorAccountOtherTypeId != "" {
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
	return result
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
