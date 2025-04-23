package ReturnRequest

import (
	"reflect"

	camt056 "github.com/moov-io/fedwire20022/gen/ReturnRequest_camt_056_001_08"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Reason struct {
	//Party that issues the cancellation request.
	Originator string
	//Specifies the reason for the cancellation.
	Reason string
	//Further details on the cancellation request reason.
	AdditionalInfo string
}

func Party40Choice2From(p model.Agent) camt056.Party40Choice2 {
	var result camt056.Party40Choice2
	var Agt camt056.BranchAndFinancialInstitutionIdentification62
	var FinInstnId camt056.FinancialInstitutionIdentification182
	var ClrSysMmbId camt056.ClearingSystemMemberIdentification22
	if p.PaymentSysCode != "" {
		Cd := camt056.ExternalClearingSystemIdentification1Code(p.PaymentSysCode)
		ClrSysMmbId.ClrSysId = camt056.ClearingSystemIdentification2Choice2{
			Cd: &Cd,
		}
	}
	if p.PaymentSysMemberId != "" {
		ClrSysMmbId.MmbId = camt056.Max35Text(p.PaymentSysMemberId)
	}
	if !isEmpty(ClrSysMmbId) {
		FinInstnId.ClrSysMmbId = &ClrSysMmbId
	}
	if p.BankName != "" {
		Nm := camt056.Max140Text(p.BankName)
		FinInstnId.Nm = &Nm
	}
	if !isEmpty(p.PostalAddress) {
		PstlAdr := PostalAddress241From(p.PostalAddress)
		FinInstnId.PstlAdr = &PstlAdr
	}
	if !isEmpty(FinInstnId) {
		Agt.FinInstnId = FinInstnId
	}
	if !isEmpty(Agt) {
		result.Agt = &Agt
	}
	return result
}
func Party40Choice1From(p model.Agent) camt056.Party40Choice1 {
	var result camt056.Party40Choice1
	var Agt camt056.BranchAndFinancialInstitutionIdentification61
	var FinInstnId camt056.FinancialInstitutionIdentification181
	var ClrSysMmbId camt056.ClearingSystemMemberIdentification21
	if p.PaymentSysCode != "" {
		Cd := camt056.ExternalClearingSystemIdentification1CodeFixed(p.PaymentSysCode)
		ClrSysMmbId.ClrSysId = camt056.ClearingSystemIdentification2Choice1{
			Cd: &Cd,
		}
	}
	if p.PaymentSysMemberId != "" {
		ClrSysMmbId.MmbId = camt056.RoutingNumberFRS1(p.PaymentSysMemberId)
	}
	if !isEmpty(ClrSysMmbId) {
		FinInstnId.ClrSysMmbId = ClrSysMmbId
	}
	if !isEmpty(FinInstnId) {
		Agt.FinInstnId = FinInstnId
	}
	if !isEmpty(Agt) {
		result.Agt = &Agt
	}
	return result
}
func PaymentCancellationReason51From(p Reason) camt056.PaymentCancellationReason51 {
	var result camt056.PaymentCancellationReason51
	if p.Originator != "" {
		Nm := camt056.Max140Text(p.Originator)
		Orgtr := camt056.PartyIdentification1352{
			Nm: &Nm,
		}
		result.Orgtr = &Orgtr
	}
	if p.Reason != "" {
		Cd := camt056.ExternalCancellationReason1Code(p.Reason)
		result.Rsn = camt056.CancellationReason33Choice1{
			Cd: &Cd,
		}
	}
	if p.AdditionalInfo != "" {
		var AddtlInf []*camt056.Max105Text
		info := camt056.Max105Text(p.AdditionalInfo)
		AddtlInf = append(AddtlInf, &info)
		result.AddtlInf = AddtlInf
	}
	return result
}
func PostalAddress241From(param model.PostalAddress) camt056.PostalAddress241 {
	var Dbtr_PstlAdr camt056.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := camt056.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := camt056.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := camt056.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := camt056.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := camt056.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := camt056.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		TwnNm := camt056.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := camt056.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Ctry := camt056.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return camt056.PostalAddress241{}
	}

	return Dbtr_PstlAdr
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
