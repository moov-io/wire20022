package CustomerCreditTransfer

import (
	"reflect"

	pacs008 "github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer_pacs_008_001_08"
	fedwire "github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type PurposeOfPaymentType string
type RemittanceDeliveryMethod string

const (
	CashWithdrawal    PurposeOfPaymentType = "CASH" // Cash Withdrawal
	GoodsAndServices  PurposeOfPaymentType = "GDSV" // Goods and Services
	LabourInsurance   PurposeOfPaymentType = "CASH" // Labour Insurance
	SupplierPayment   PurposeOfPaymentType = "SUPP" // Supplier Payment
	TradeSettlement   PurposeOfPaymentType = "TRAD" // Trade Settlement
	InvestmentPayment PurposeOfPaymentType = "IVPT" // Investment Payment
	PensionPayment    PurposeOfPaymentType = "PENS" // Pension Payment
	AlimonyPayment    PurposeOfPaymentType = "ALMY" // Alimony Payment
	INSCPayment       PurposeOfPaymentType = "INSC"
)

const (
	Fax                       RemittanceDeliveryMethod = "FAXI" //Fax
	ElectronicDataInterchange RemittanceDeliveryMethod = "EDIC" //Electronic Data Interchange (EDI)
	UniformResourceIdentifier RemittanceDeliveryMethod = "URID" //Uniform Resource Identifier (URI)
	PostalMail                RemittanceDeliveryMethod = "POST" //Postal Mail
	Email                     RemittanceDeliveryMethod = "EMAL" //Email
)

type TaxRecord struct {
	//is used by governments to track tax obligations and payments.
	TaxId string
	//tax type code
	TaxTypeCode string
	// Tax Period Year
	TaxPeriodYear      model.Date
	TaxperiodTimeFrame string
}
type RemittanceDetail struct {
	//unique reference number used to identify a remittance transaction.
	RemittanceId string
	//Specifies how the remittance information is delivered.
	Method RemittanceDeliveryMethod
	//Provides the email address where the remittance details should be sent.
	ElectronicAddress string
}
type RemittanceDocument struct {
	//refers to Unstructured Remittance Information in the ISO 20022 payment message standard
	UnstructuredRemitInfo string
	//Code or Proprietary :It is used to specify the method for identifying the type of a document or reference.
	CodeOrProprietary model.CodeOrProprietaryType
	//invoice number
	Number string
	//default value: current date
	RelatedDate model.Date
	// Tax detail
	TaxDetail TaxRecord
}
type ChargeInfo struct {
	Amount         model.CurrencyAndAmount
	BusinessIdCode string
}

/*********************************************************/
/** Internal functions  **/
/*********************************************************/

func PostalAddress241From(param model.PostalAddress) (pacs008.PostalAddress241, *model.ValidateError) {
	var Dbtr_PstlAdr pacs008.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		vErr := pacs008.Max70Text(param.StreetName).Validate()
		if vErr != nil {
			return pacs008.PostalAddress241{}, &model.ValidateError{
				ParamName: "StreetName",
				Message:   vErr.Error(),
			}
		}
		StrtNm := pacs008.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		vErr := pacs008.Max16Text(param.BuildingNumber).Validate()
		if vErr != nil {
			return pacs008.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingNumber",
				Message:   vErr.Error(),
			}
		}
		BldgNb := pacs008.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		vErr := pacs008.Max35Text(param.BuildingName).Validate()
		if vErr != nil {
			return pacs008.PostalAddress241{}, &model.ValidateError{
				ParamName: "BuildingName",
				Message:   vErr.Error(),
			}
		}
		BldgNm := pacs008.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		vErr := pacs008.Max70Text(param.Floor).Validate()
		if vErr != nil {
			return pacs008.PostalAddress241{}, &model.ValidateError{
				ParamName: "Floor",
				Message:   vErr.Error(),
			}
		}
		Floor := pacs008.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		vErr := pacs008.Max70Text(param.RoomNumber).Validate()
		if vErr != nil {
			return pacs008.PostalAddress241{}, &model.ValidateError{
				ParamName: "RoomNumber",
				Message:   vErr.Error(),
			}
		}
		Room := pacs008.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		vErr := pacs008.Max16Text(param.PostalCode).Validate()
		if vErr != nil {
			return pacs008.PostalAddress241{}, &model.ValidateError{
				ParamName: "PostalCode",
				Message:   vErr.Error(),
			}
		}
		PstCd := pacs008.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		vErr := pacs008.Max35Text(param.TownName).Validate()
		if vErr != nil {
			return pacs008.PostalAddress241{}, &model.ValidateError{
				ParamName: "TownName",
				Message:   vErr.Error(),
			}
		}
		TwnNm := pacs008.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		vErr := pacs008.Max35Text(param.Subdivision).Validate()
		if vErr != nil {
			return pacs008.PostalAddress241{}, &model.ValidateError{
				ParamName: "Subdivision",
				Message:   vErr.Error(),
			}
		}
		CtrySubDvsn := pacs008.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		vErr := pacs008.CountryCode(param.Country).Validate()
		if vErr != nil {
			return pacs008.PostalAddress241{}, &model.ValidateError{
				ParamName: "Country",
				Message:   vErr.Error(),
			}
		}
		Ctry := pacs008.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs008.PostalAddress241{}, nil
	}

	return Dbtr_PstlAdr, nil
}
func isEmptyPostalAddress241(address pacs008.PostalAddress241) bool {
	// Compare the struct with its zero value
	return address.StrtNm == nil &&
		address.BldgNb == nil &&
		address.BldgNm == nil &&
		address.Flr == nil &&
		address.Room == nil &&
		address.PstCd == nil &&
		address.TwnNm == nil &&
		address.CtrySubDvsn == nil &&
		address.Ctry == nil
}
func PostalAddress242From(param model.PostalAddress) (pacs008.PostalAddress242, *model.ValidateError) {
	var Dbtr_PstlAdr pacs008.PostalAddress242

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		vErr := pacs008.Max70Text(param.StreetName).Validate()
		if vErr != nil {
			return pacs008.PostalAddress242{}, &model.ValidateError{
				ParamName: "StreetName",
				Message:   vErr.Error(),
			}
		}
		StrtNm := pacs008.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		vErr := pacs008.Max16Text(param.BuildingNumber).Validate()
		if vErr != nil {
			return pacs008.PostalAddress242{}, &model.ValidateError{
				ParamName: "BuildingNumber",
				Message:   vErr.Error(),
			}
		}
		BldgNb := pacs008.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.TownName != "" {
		vErr := pacs008.Max35Text(param.TownName).Validate()
		if vErr != nil {
			return pacs008.PostalAddress242{}, &model.ValidateError{
				ParamName: "TownName",
				Message:   vErr.Error(),
			}
		}
		Dbtr_PstlAdr.TwnNm = pacs008.Max35Text(param.TownName)
		hasData = true
	}
	if param.BuildingName != "" {
		vErr := pacs008.Max35Text(param.BuildingName).Validate()
		if vErr != nil {
			return pacs008.PostalAddress242{}, &model.ValidateError{
				ParamName: "BuildingName",
				Message:   vErr.Error(),
			}
		}
		BldgNm := pacs008.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		vErr := pacs008.Max70Text(param.Floor).Validate()
		if vErr != nil {
			return pacs008.PostalAddress242{}, &model.ValidateError{
				ParamName: "Floor",
				Message:   vErr.Error(),
			}
		}
		Floor := pacs008.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		vErr := pacs008.Max70Text(param.RoomNumber).Validate()
		if vErr != nil {
			return pacs008.PostalAddress242{}, &model.ValidateError{
				ParamName: "RoomNumber",
				Message:   vErr.Error(),
			}
		}
		Room := pacs008.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		vErr := pacs008.Max16Text(param.PostalCode).Validate()
		if vErr != nil {
			return pacs008.PostalAddress242{}, &model.ValidateError{
				ParamName: "PostalCode",
				Message:   vErr.Error(),
			}
		}
		PstCd := pacs008.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.Subdivision != "" {
		vErr := pacs008.Max35Text(param.Subdivision).Validate()
		if vErr != nil {
			return pacs008.PostalAddress242{}, &model.ValidateError{
				ParamName: "Subdivision",
				Message:   vErr.Error(),
			}
		}
		CtrySubDvsn := pacs008.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		vErr := pacs008.CountryCode(param.Country).Validate()
		if vErr != nil {
			return pacs008.PostalAddress242{}, &model.ValidateError{
				ParamName: "Country",
				Message:   vErr.Error(),
			}
		}
		Dbtr_PstlAdr.Ctry = pacs008.CountryCode(param.Country)
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs008.PostalAddress242{}, nil
	}

	return Dbtr_PstlAdr, nil
}
func isEmptyPostalAddress242(address pacs008.PostalAddress242) bool {
	// Compare the struct with its zero value
	return address.StrtNm == nil &&
		address.BldgNb == nil &&
		address.BldgNm == nil &&
		address.Flr == nil &&
		address.TwnNm == "" &&
		address.Room == nil &&
		address.PstCd == nil &&
		address.Ctry == ""
}
func CashAccount38From(ibanId string, iban string, otherId string, other string) (pacs008.CashAccount38, *model.ValidateError) {
	if iban == "" && other == "" {
		return pacs008.CashAccount38{}, nil // Return empty struct if input is empty
	}
	var account pacs008.AccountIdentification4Choice
	if iban != "" {
		err := pacs008.IBAN2007Identifier(iban).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: ibanId,
				Message:   err.Error(),
			}
			return pacs008.CashAccount38{}, &vErr
		}
		_IBAN := pacs008.IBAN2007Identifier(iban)
		account.IBAN = &_IBAN
	}
	if other != "" {
		err := pacs008.Max34Text(other).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: otherId,
				Message:   err.Error(),
			}
			return pacs008.CashAccount38{}, &vErr
		}
		account_Othr := pacs008.GenericAccountIdentification1{}
		account.Othr = &account_Othr
		account.Othr.Id = pacs008.Max34Text(other)
	}
	return pacs008.CashAccount38{
		Id: account,
	}, nil
}
func ClearingSystemMemberIdentification21From(param model.PaymentSystemType, paymentSysMemberId string) (pacs008.ClearingSystemMemberIdentification21, *model.ValidateError) {
	var result pacs008.ClearingSystemMemberIdentification21
	var hasData bool // Flag to check if there's valid data

	if param != "" {
		vErr := pacs008.ExternalClearingSystemIdentification1Code(param).Validate()
		if vErr != nil {
			return pacs008.ClearingSystemMemberIdentification21{}, &model.ValidateError{
				ParamName: "PaymentSystemType",
				Message:   vErr.Error(),
			}
		}
		Cd := pacs008.ExternalClearingSystemIdentification1Code(param)
		result.ClrSysId = pacs008.ClearingSystemIdentification2Choice1{
			Cd: &Cd,
		}
		hasData = true
	}

	if paymentSysMemberId != "" {
		vErr := pacs008.Max35Text(paymentSysMemberId).Validate()
		if vErr != nil {
			return pacs008.ClearingSystemMemberIdentification21{}, &model.ValidateError{
				ParamName: "paymentSysMemberId",
				Message:   vErr.Error(),
			}
		}
		result.MmbId = pacs008.Max35Text(paymentSysMemberId)
		hasData = true
	}

	// If no valid data, return an empty struct
	if !hasData {
		return pacs008.ClearingSystemMemberIdentification21{}, nil
	}

	return result, nil
}
func RemittanceInformation161From(doc RemittanceDocument) (pacs008.RemittanceInformation161, *model.ValidateError) {
	var result pacs008.RemittanceInformation161
	var hasData bool // Flag to check if we have any meaningful data

	// Set UnstructuredRemitInfo if not empty
	if doc.UnstructuredRemitInfo != "" {
		vErr := pacs008.Max140Text(doc.UnstructuredRemitInfo).Validate()
		if vErr != nil {
			return pacs008.RemittanceInformation161{}, &model.ValidateError{
				ParamName: "UnstructuredRemitInfo",
				Message:   vErr.Error(),
			}
		}
		UnstructuredRemitInfo := pacs008.Max140Text(doc.UnstructuredRemitInfo)
		result.Ustrd = &UnstructuredRemitInfo
		hasData = true
	}

	// Prepare referred document information
	var RD_item pacs008.ReferredDocumentInformation71
	var hasRDData bool // Check if RD_item contains meaningful data
	var hasTaxData bool
	var hasTaxPrData bool

	if doc.CodeOrProprietary != "" {
		vErr := pacs008.DocumentType6Code(doc.CodeOrProprietary).Validate()
		if vErr != nil {
			return pacs008.RemittanceInformation161{}, &model.ValidateError{
				ParamName: "CodeOrProprietary",
				Message:   vErr.Error(),
			}
		}
		RD_item_Tp_Cd := pacs008.DocumentType6Code(doc.CodeOrProprietary)
		RD_item.Tp = &pacs008.ReferredDocumentType4{
			CdOrPrtry: pacs008.ReferredDocumentType3Choice{
				Cd: &RD_item_Tp_Cd,
			},
		}
		hasRDData = true
	}

	if doc.Number != "" {
		vErr := pacs008.Max35Text(doc.Number).Validate()
		if vErr != nil {
			return pacs008.RemittanceInformation161{}, &model.ValidateError{
				ParamName: "Number",
				Message:   vErr.Error(),
			}
		}
		RD_item_Nb := pacs008.Max35Text(doc.Number)
		RD_item.Nb = &RD_item_Nb
		hasRDData = true
	}

	if !isEmpty(doc.RelatedDate) {
		vErr := doc.RelatedDate.Date().Validate()
		if vErr != nil {
			return pacs008.RemittanceInformation161{}, &model.ValidateError{
				ParamName: "RelatedDate",
				Message:   vErr.Error(),
			}
		}
		RD_item_RltdDt := doc.RelatedDate.Date()
		RD_item.RltdDt = &RD_item_RltdDt
		hasRDData = true
	}

	var TaxRmt pacs008.TaxInformation7
	if doc.TaxDetail.TaxId != "" {
		vErr := pacs008.Max35Text(doc.TaxDetail.TaxId).Validate()
		if vErr != nil {
			vvErr := model.ValidateError{
				ParamName: "TaxId",
				Message:   vErr.Error(),
			}
			vvErr.InsertPath("TaxDetail")
			return pacs008.RemittanceInformation161{}, &vvErr
		}
		TaxId := pacs008.Max35Text(doc.TaxDetail.TaxId)
		TaxRmt_Cdtr := pacs008.TaxParty1{}
		TaxRmt.Cdtr = &TaxRmt_Cdtr
		TaxRmt.Cdtr.TaxId = &TaxId
		hasTaxData = true
	}
	var TaxRecode pacs008.TaxRecord2
	if doc.TaxDetail.TaxTypeCode != "" {
		vErr := pacs008.Max35Text(doc.TaxDetail.TaxTypeCode).Validate()
		if vErr != nil {
			vvErr := model.ValidateError{
				ParamName: "TaxTypeCode",
				Message:   vErr.Error(),
			}
			vvErr.InsertPath("TaxDetail")
			return pacs008.RemittanceInformation161{}, &vvErr
		}
		TaxRecode_Tp := pacs008.Max35Text(doc.TaxDetail.TaxTypeCode)
		TaxRecode.Tp = &TaxRecode_Tp
		hasTaxPrData = true
	}
	if !isEmpty(doc.TaxDetail.TaxPeriodYear) {
		vErr := doc.TaxDetail.TaxPeriodYear.Date().Validate()
		if vErr != nil {
			vvErr := model.ValidateError{
				ParamName: "TaxPeriodYear",
				Message:   vErr.Error(),
			}
			vvErr.InsertPath("TaxDetail")
			return pacs008.RemittanceInformation161{}, &vvErr
		}
		TaxRecode_Prd_Y := doc.TaxDetail.TaxPeriodYear.Date()
		if TaxRecode.Prd == nil {
			TaxRecode_Prd := pacs008.TaxPeriod2{}
			TaxRecode.Prd = &TaxRecode_Prd
		}
		TaxRecode.Prd.Yr = &TaxRecode_Prd_Y
		hasTaxPrData = true
	}
	if doc.TaxDetail.TaxperiodTimeFrame != "" {
		vErr := pacs008.TaxRecordPeriod1Code(doc.TaxDetail.TaxperiodTimeFrame).Validate()
		if vErr != nil {
			vvErr := model.ValidateError{
				ParamName: "TaxperiodTimeFrame",
				Message:   vErr.Error(),
			}
			vvErr.InsertPath("TaxDetail")
			return pacs008.RemittanceInformation161{}, &vvErr
		}
		TaxRecode_Prd_tp := pacs008.TaxRecordPeriod1Code(doc.TaxDetail.TaxperiodTimeFrame)
		if TaxRecode.Prd == nil {
			TaxRecode_Prd := pacs008.TaxPeriod2{}
			TaxRecode.Prd = &TaxRecode_Prd
		}
		TaxRecode.Prd.Tp = &TaxRecode_Prd_tp
		hasTaxPrData = true
	}
	if hasTaxPrData {
		TaxRmt.Rcrd = []*pacs008.TaxRecord2{
			&TaxRecode,
		}
		hasTaxPrData = true
	}

	SR_item := pacs008.StructuredRemittanceInformation161{}
	if hasRDData {
		SR_item.RfrdDocInf = []*pacs008.ReferredDocumentInformation71{
			&RD_item,
		}
	}
	if hasTaxData {
		SR_item.TaxRmt = &TaxRmt
	}

	// If RD_item has data, add it to structured remittance info
	if hasRDData || hasTaxData {
		result.Strd = []*pacs008.StructuredRemittanceInformation161{
			&SR_item,
		}
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs008.RemittanceInformation161{}, nil
	}

	return result, nil
}
func FinancialInstitutionIdentification181From(agent model.Agent) (pacs008.FinancialInstitutionIdentification181, *model.ValidateError) {
	var result pacs008.FinancialInstitutionIdentification181
	if agent.BusinessIdCode != "" {
		vErr := pacs008.BICFIDec2014Identifier(agent.BusinessIdCode).Validate()
		if vErr != nil {
			return pacs008.FinancialInstitutionIdentification181{}, &model.ValidateError{
				ParamName: "BusinessIdCode",
				Message:   vErr.Error(),
			}
		}
		_BICFI := pacs008.BICFIDec2014Identifier(agent.BusinessIdCode)
		result.BICFI = &_BICFI
	}
	if agent.PaymentSysCode != "" || agent.PaymentSysMemberId != "" {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := pacs008.ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		if agent.PaymentSysCode != "" {
			vErr := pacs008.ExternalClearingSystemIdentification1Code(agent.PaymentSysCode).Validate()
			if vErr != nil {
				return pacs008.FinancialInstitutionIdentification181{}, &model.ValidateError{
					ParamName: "PaymentSysCode",
					Message:   vErr.Error(),
				}
			}
			Cd := pacs008.ExternalClearingSystemIdentification1Code(agent.PaymentSysCode)
			result.ClrSysMmbId.ClrSysId = pacs008.ClearingSystemIdentification2Choice1{
				Cd: &Cd,
			}
		}
		if agent.PaymentSysMemberId != "" {
			vErr := pacs008.Max35Text(agent.PaymentSysMemberId).Validate()
			if vErr != nil {
				return pacs008.FinancialInstitutionIdentification181{}, &model.ValidateError{
					ParamName: "PaymentSysMemberId",
					Message:   vErr.Error(),
				}
			}
			result.ClrSysMmbId.MmbId = pacs008.Max35Text(agent.PaymentSysMemberId)
		}
	}
	if agent.BankName != "" {
		vErr := pacs008.Max140Text(agent.BankName).Validate()
		if vErr != nil {
			return pacs008.FinancialInstitutionIdentification181{}, &model.ValidateError{
				ParamName: "BankName",
				Message:   vErr.Error(),
			}
		}
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := pacs008.ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		_BKNM := pacs008.Max140Text(agent.BankName)
		result.Nm = &_BKNM
	}
	postalAddress, vErr := PostalAddress241From(agent.PostalAddress)
	if vErr != nil {
		vErr.InsertPath("PostalAddress")
		return pacs008.FinancialInstitutionIdentification181{}, vErr
	}
	if !isEmptyPostalAddress241(postalAddress) {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := pacs008.ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		result.PstlAdr = &postalAddress
	}
	return result, nil
}

func PaymentTypeInformation281From(InstrumentPropCode model.InstrumentPropCodeType, SericeLevel string) (pacs008.PaymentTypeInformation281, *model.ValidateError) {
	var result pacs008.PaymentTypeInformation281
	if InstrumentPropCode != "" {
		vErr := pacs008.LocalInstrumentFedwireFunds1(InstrumentPropCode).Validate()
		if vErr != nil {
			return pacs008.PaymentTypeInformation281{}, &model.ValidateError{
				ParamName: "InstrumentPropCode",
				Message:   vErr.Error(),
			}
		}
		result.LclInstrm = pacs008.LocalInstrument2Choice1{}
		CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry := pacs008.LocalInstrumentFedwireFunds1(InstrumentPropCode)
		result.LclInstrm.Prtry = &CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry
	}
	if SericeLevel != "" {
		vErr := pacs008.ExternalServiceLevel1Code(SericeLevel).Validate()
		if vErr != nil {
			return pacs008.PaymentTypeInformation281{}, &model.ValidateError{
				ParamName: "SericeLevel",
				Message:   vErr.Error(),
			}
		}
		svclv := pacs008.ExternalServiceLevel1Code(SericeLevel)
		CdtTrfTxInf_PmtTpInf_SvcLvl := pacs008.ServiceLevel8Choice{
			Cd: &svclv,
		}
		result.SvcLvl = []*pacs008.ServiceLevel8Choice{
			&CdtTrfTxInf_PmtTpInf_SvcLvl,
		}
	}
	return result, nil
}
func RemittanceLocation71From(param RemittanceDetail) (pacs008.RemittanceLocation71, *model.ValidateError) {
	var result pacs008.RemittanceLocation71
	if param.RemittanceId != "" {
		vErr := pacs008.Max35Text(param.RemittanceId).Validate()
		if vErr != nil {
			return pacs008.RemittanceLocation71{}, &model.ValidateError{
				ParamName: "RemittanceId",
				Message:   vErr.Error(),
			}
		}
		_RemittanceId := pacs008.Max35Text(param.RemittanceId)
		result.RmtId = &_RemittanceId
	}
	var locationData pacs008.RemittanceLocationData11
	var hasLocationData = false
	if param.Method != "" {
		vErr := pacs008.RemittanceLocationMethod2Code(param.Method).Validate()
		if vErr != nil {
			return pacs008.RemittanceLocation71{}, &model.ValidateError{
				ParamName: "Method",
				Message:   vErr.Error(),
			}
		}
		locationData.Mtd = pacs008.RemittanceLocationMethod2Code(param.Method)
		hasLocationData = true
	}
	if param.ElectronicAddress != "" {
		vErr := pacs008.Max2048Text(param.ElectronicAddress).Validate()
		if vErr != nil {
			return pacs008.RemittanceLocation71{}, &model.ValidateError{
				ParamName: "ElectronicAddress",
				Message:   vErr.Error(),
			}
		}
		_ElectronicAddress := pacs008.Max2048Text(param.ElectronicAddress)
		locationData.ElctrncAdr = &_ElectronicAddress
		hasLocationData = true
	}
	if hasLocationData {
		result.RmtLctnDtls = []*pacs008.RemittanceLocationData11{
			&locationData,
		}
	}
	return result, nil
}

func PartyIdentification1352From(Nm string, PstlAdr model.PostalAddress) (pacs008.PartyIdentification1352, *model.ValidateError) {
	var result pacs008.PartyIdentification1352
	if Nm != "" {
		vErr := pacs008.Max140Text(Nm).Validate()
		if vErr != nil {
			return pacs008.PartyIdentification1352{}, &model.ValidateError{
				ParamName: "Name",
				Message:   vErr.Error(),
			}
		}
		_nm := pacs008.Max140Text(Nm)
		result.Nm = &_nm
	}
	_PstlAdr, vErr := PostalAddress241From(PstlAdr)
	if vErr != nil {
		return pacs008.PartyIdentification1352{}, vErr
	}
	if !isEmptyPostalAddress241(_PstlAdr) {
		result.PstlAdr = &_PstlAdr
	}
	return result, nil
}

func PartyIdentification1351From(Nm string, PstlAdr model.PostalAddress) (pacs008.PartyIdentification1351, *model.ValidateError) {
	var result pacs008.PartyIdentification1351
	if Nm != "" {
		vErr := pacs008.Max140Text(Nm).Validate()
		if vErr != nil {
			return pacs008.PartyIdentification1351{}, &model.ValidateError{
				ParamName: "Name",
				Message:   vErr.Error(),
			}
		}
		_nm := pacs008.Max140Text(Nm)
		result.Nm = &_nm
	}
	_PstlAdr, vErr := PostalAddress242From(PstlAdr)
	if vErr != nil {
		return pacs008.PartyIdentification1351{}, &model.ValidateError{
			ParamName: "PostalAddress",
			Message:   vErr.Error(),
		}
	}
	if !isEmptyPostalAddress242(_PstlAdr) {
		result.PstlAdr = &_PstlAdr
	}
	return result, nil
}

func Charges71From(data ChargeInfo) (pacs008.Charges71, *model.ValidateError) {
	var result pacs008.Charges71
	if data.Amount.Amount != 0 || data.Amount.Currency != "" {
		err := fedwire.Amount(data.Amount.Amount).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "Amount",
				Message:   err.Error(),
			}
			vErr.InsertPath("Amount")
			return pacs008.Charges71{}, &vErr
		}
		err = pacs008.ActiveOrHistoricCurrencyCode(data.Amount.Currency).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "Currency",
				Message:   err.Error(),
			}
			vErr.InsertPath("Amount")
			return pacs008.Charges71{}, &vErr
		}
		result.Amt = pacs008.ActiveOrHistoricCurrencyAndAmount{
			Value: pacs008.ActiveOrHistoricCurrencyAndAmountSimpleType(data.Amount.Amount),
			Ccy:   pacs008.ActiveOrHistoricCurrencyCode(data.Amount.Currency),
		}
	}
	if data.BusinessIdCode != "" {
		err := pacs008.BICFIDec2014Identifier(data.BusinessIdCode).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "BusinessIdCode",
				Message:   err.Error(),
			}
			return pacs008.Charges71{}, &vErr
		}
		result.Agt = pacs008.BranchAndFinancialInstitutionIdentification61{}
		result.Agt.FinInstnId = pacs008.FinancialInstitutionIdentification181{}
		_BICFI := pacs008.BICFIDec2014Identifier(data.BusinessIdCode)
		result.Agt.FinInstnId.BICFI = &_BICFI
	}
	return result, nil
}

func BranchAndFinancialInstitutionIdentification61From(BICFI string) (pacs008.BranchAndFinancialInstitutionIdentification61, *model.ValidateError) {
	var result pacs008.BranchAndFinancialInstitutionIdentification61
	if BICFI != "" {
		vErr := pacs008.BICFIDec2014Identifier(BICFI).Validate()
		if vErr != nil {
			return pacs008.BranchAndFinancialInstitutionIdentification61{}, &model.ValidateError{
				ParamName: "BICFI",
				Message:   vErr.Error(),
			}
		}
		result.FinInstnId = pacs008.FinancialInstitutionIdentification181{}
		_BICFI := pacs008.BICFIDec2014Identifier(BICFI)
		result.FinInstnId.BICFI = &_BICFI
	}
	return result, nil
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
