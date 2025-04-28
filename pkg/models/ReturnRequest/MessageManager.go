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

func Party40Choice2From(p model.Agent) (camt056.Party40Choice2, *model.ValidateError) {
	var result camt056.Party40Choice2
	var Agt camt056.BranchAndFinancialInstitutionIdentification62
	var FinInstnId camt056.FinancialInstitutionIdentification182
	var ClrSysMmbId camt056.ClearingSystemMemberIdentification22
	if p.PaymentSysCode != "" {
		err := camt056.ExternalClearingSystemIdentification1Code(p.PaymentSysCode).Validate()
		if err != nil {
			return camt056.Party40Choice2{}, &model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
		}
		Cd := camt056.ExternalClearingSystemIdentification1Code(p.PaymentSysCode)
		ClrSysMmbId.ClrSysId = camt056.ClearingSystemIdentification2Choice2{
			Cd: &Cd,
		}
	}
	if p.PaymentSysMemberId != "" {
		err := camt056.Max35Text(p.PaymentSysMemberId).Validate()
		if err != nil {
			return camt056.Party40Choice2{}, &model.ValidateError{
				ParamName: "PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		ClrSysMmbId.MmbId = camt056.Max35Text(p.PaymentSysMemberId)
	}
	if !isEmpty(ClrSysMmbId) {
		FinInstnId.ClrSysMmbId = &ClrSysMmbId
	}
	if p.BankName != "" {
		err := camt056.Max140Text(p.BankName).Validate()
		if err != nil {
			return camt056.Party40Choice2{}, &model.ValidateError{
				ParamName: "BankName",
				Message:   err.Error(),
			}
		}
		Nm := camt056.Max140Text(p.BankName)
		FinInstnId.Nm = &Nm
	}
	if !isEmpty(p.PostalAddress) {
		PstlAdr, err := PostalAddress241From(p.PostalAddress)
		if err != nil {
			err.InsertPath("PostalAddress")
			return camt056.Party40Choice2{}, err
		}
		FinInstnId.PstlAdr = &PstlAdr
	}
	if !isEmpty(FinInstnId) {
		Agt.FinInstnId = FinInstnId
	}
	if !isEmpty(Agt) {
		result.Agt = &Agt
	}
	return result, nil
}
func Party40Choice1From(p model.Agent) (camt056.Party40Choice1, *model.ValidateError) {
	var result camt056.Party40Choice1
	var Agt camt056.BranchAndFinancialInstitutionIdentification61
	var FinInstnId camt056.FinancialInstitutionIdentification181
	var ClrSysMmbId camt056.ClearingSystemMemberIdentification21
	if p.PaymentSysCode != "" {
		err := camt056.ExternalClearingSystemIdentification1CodeFixed(p.PaymentSysCode).Validate()
		if err != nil {
			return camt056.Party40Choice1{}, &model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
		}
		Cd := camt056.ExternalClearingSystemIdentification1CodeFixed(p.PaymentSysCode)
		ClrSysMmbId.ClrSysId = camt056.ClearingSystemIdentification2Choice1{
			Cd: &Cd,
		}
	}
	if p.PaymentSysMemberId != "" {
		err := camt056.RoutingNumberFRS1(p.PaymentSysMemberId).Validate()
		if err != nil {
			return camt056.Party40Choice1{}, &model.ValidateError{
				ParamName: "PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
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
	return result, nil
}
func PaymentCancellationReason51From(p Reason) (camt056.PaymentCancellationReason51, *model.ValidateError) {
	var result camt056.PaymentCancellationReason51
	if p.Originator != "" {
		err := camt056.Max140Text(p.Originator).Validate()
		if err != nil {
			return camt056.PaymentCancellationReason51{}, &model.ValidateError{
				ParamName: "Originator",
				Message:   err.Error(),
			}
		}
		Nm := camt056.Max140Text(p.Originator)
		Orgtr := camt056.PartyIdentification1352{
			Nm: &Nm,
		}
		result.Orgtr = &Orgtr
	}
	if p.Reason != "" {
		err := camt056.ExternalCancellationReason1Code(p.Reason).Validate()
		if err != nil {
			return camt056.PaymentCancellationReason51{}, &model.ValidateError{
				ParamName: "Reason",
				Message:   err.Error(),
			}
		}
		Cd := camt056.ExternalCancellationReason1Code(p.Reason)
		result.Rsn = camt056.CancellationReason33Choice1{
			Cd: &Cd,
		}
	}
	if p.AdditionalInfo != "" {
		var AddtlInf []*camt056.Max105Text
		err := camt056.Max105Text(p.AdditionalInfo).Validate()
		if err != nil {
			return camt056.PaymentCancellationReason51{}, &model.ValidateError{
				ParamName: "AdditionalInfo",
				Message:   err.Error(),
			}
		}
		info := camt056.Max105Text(p.AdditionalInfo)
		AddtlInf = append(AddtlInf, &info)
		result.AddtlInf = AddtlInf
	}
	return result, nil
}
func PostalAddress241From(param model.PostalAddress) (camt056.PostalAddress241, *model.ValidateError) {
	var Dbtr_PstlAdr camt056.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		err := camt056.Max70Text(param.StreetName).Validate()
		if err != nil {
			return camt056.PostalAddress241{}, &model.ValidateError{
				ParamName: "StreetName",
				Message:   err.Error(),
			}
		}
		StrtNm := camt056.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		err := camt056.Max16Text(param.BuildingNumber).Validate()
		if err != nil {
			return camt056.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingNumber",
				Message:   err.Error(),
			}
		}
		BldgNb := camt056.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		err := camt056.Max35Text(param.BuildingName).Validate()
		if err != nil {
			return camt056.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingName",
				Message:   err.Error(),
			}
		}
		BldgNm := camt056.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		err := camt056.Max70Text(param.Floor).Validate()
		if err != nil {
			return camt056.PostalAddress241{}, &model.ValidateError{
				ParamName: "Floor",
				Message:   err.Error(),
			}
		}
		Floor := camt056.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		err := camt056.Max70Text(param.RoomNumber).Validate()
		if err != nil {
			return camt056.PostalAddress241{}, &model.ValidateError{
				ParamName: "RoomNumber",
				Message:   err.Error(),
			}
		}
		Room := camt056.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		err := camt056.Max16Text(param.PostalCode).Validate()
		if err != nil {
			return camt056.PostalAddress241{}, &model.ValidateError{
				ParamName: "PostalCode",
				Message:   err.Error(),
			}
		}
		PstCd := camt056.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		err := camt056.Max35Text(param.TownName).Validate()
		if err != nil {
			return camt056.PostalAddress241{}, &model.ValidateError{
				ParamName: "TownName",
				Message:   err.Error(),
			}
		}
		TwnNm := camt056.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		err := camt056.Max35Text(param.Subdivision).Validate()
		if err != nil {
			return camt056.PostalAddress241{}, &model.ValidateError{
				ParamName: "Subdivision",
				Message:   err.Error(),
			}
		}
		CtrySubDvsn := camt056.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		err := camt056.CountryCode(param.Country).Validate()
		if err != nil {
			return camt056.PostalAddress241{}, &model.ValidateError{
				ParamName: "Country",
				Message:   err.Error(),
			}
		}
		Ctry := camt056.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return camt056.PostalAddress241{}, nil
	}

	return Dbtr_PstlAdr, nil
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
