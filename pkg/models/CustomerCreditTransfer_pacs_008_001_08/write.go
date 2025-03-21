package CustomerCreditTransfer_pacs_008_001_08

import (
	"encoding/xml"
)

func encodeIfNotNil(e *xml.Encoder, value any, elementName string) error {
	if value == nil || value == "" {
		return nil
	}
	return e.EncodeElement(value, xml.StartElement{Name: xml.Name{Local: elementName}})
}
func (v Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Define namespace attribute
	namespaceAttr := xml.Attr{
		Name:  xml.Name{Local: "xmlns:urn"},
		Value: "urn2:iso:std:iso:20022:tech:xsd:pacs.008.001.08",
	}

	// Start Document element with namespace
	start.Name.Local = "urn2:Document"
	start.Attr = append(start.Attr, namespaceAttr)

	// Encode the start element
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Encode FIToFICstmrCdtTrf element with correct namespace
	if err := e.EncodeElement(v.FIToFICstmrCdtTrf, xml.StartElement{
		Name: xml.Name{Space: "urn2", Local: "FIToFICstmrCdtTrf"},
	}); err != nil {
		return err
	}

	// Encode closing Document element
	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v AccountIdentification4Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.IBAN, "urn2:IBAN")
	encodeIfNotNil(e, v.Othr, "urn2:Othr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v AccountSchemeName1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ActiveCurrencyAndAmountFedwire1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Value, "urn2:Value")
	encodeIfNotNil(e, v.Ccy, "urn2:Ccy")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ActiveOrHistoricCurrencyAndAmount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Value, "urn2:Value")
	encodeIfNotNil(e, v.Ccy, "urn2:Ccy")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification61) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.FinInstnId, "urn2:FinInstnId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification62) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.FinInstnId, "urn2:FinInstnId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification63) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.FinInstnId, "urn2:FinInstnId")
	encodeIfNotNil(e, v.BrnchId, "urn2:BrnchId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchData31) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Id, "urn2:Id")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CashAccount38) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Id, "urn2:Id")
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Ccy, "urn2:Ccy")
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	encodeIfNotNil(e, v.Prxy, "urn2:Prxy")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CashAccountType2Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CategoryPurpose1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Charges71) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	encodeIfNotNil(e, v.Agt, "urn2:Agt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemIdentification2Choice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemIdentification2Choice2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemIdentification3Choice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification21) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.ClrSysId, "urn2:ClrSysId")
	encodeIfNotNil(e, v.MmbId, "urn2:MmbId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification22) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.ClrSysId, "urn2:ClrSysId")
	encodeIfNotNil(e, v.MmbId, "urn2:MmbId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CreditTransferTransaction391) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.PmtId, "urn2:PmtId")
	encodeIfNotNil(e, v.PmtTpInf, "urn2:PmtTpInf")
	encodeIfNotNil(e, v.IntrBkSttlmAmt, "urn2:IntrBkSttlmAmt")
	encodeIfNotNil(e, v.IntrBkSttlmDt, "urn2:IntrBkSttlmDt")
	encodeIfNotNil(e, v.AccptncDtTm, "urn2:AccptncDtTm")
	encodeIfNotNil(e, v.InstdAmt, "urn2:InstdAmt")
	encodeIfNotNil(e, v.XchgRate, "urn2:XchgRate")
	encodeIfNotNil(e, v.ChrgBr, "urn2:ChrgBr")
	encodeIfNotNil(e, v.ChrgsInf, "urn2:ChrgsInf")
	encodeIfNotNil(e, v.PrvsInstgAgt1, "urn2:PrvsInstgAgt1")
	encodeIfNotNil(e, v.PrvsInstgAgt1Acct, "urn2:PrvsInstgAgt1Acct")
	encodeIfNotNil(e, v.PrvsInstgAgt2, "urn2:PrvsInstgAgt2")
	encodeIfNotNil(e, v.PrvsInstgAgt2Acct, "urn2:PrvsInstgAgt2Acct")
	encodeIfNotNil(e, v.PrvsInstgAgt3, "urn2:PrvsInstgAgt3")
	encodeIfNotNil(e, v.PrvsInstgAgt3Acct, "urn2:PrvsInstgAgt3Acct")
	encodeIfNotNil(e, v.InstgAgt, "urn2:InstgAgt")
	encodeIfNotNil(e, v.InstdAgt, "urn2:InstdAgt")
	encodeIfNotNil(e, v.IntrmyAgt1, "urn2:IntrmyAgt1")
	encodeIfNotNil(e, v.IntrmyAgt1Acct, "urn2:IntrmyAgt1Acct")
	encodeIfNotNil(e, v.IntrmyAgt2, "urn2:IntrmyAgt2")
	encodeIfNotNil(e, v.IntrmyAgt2Acct, "urn2:IntrmyAgt2Acct")
	encodeIfNotNil(e, v.IntrmyAgt3, "urn2:IntrmyAgt3")
	encodeIfNotNil(e, v.IntrmyAgt3Acct, "urn2:IntrmyAgt3Acct")
	encodeIfNotNil(e, v.UltmtDbtr, "urn2:UltmtDbtr")
	encodeIfNotNil(e, v.InitgPty, "urn2:InitgPty")
	encodeIfNotNil(e, v.Dbtr, "urn2:Dbtr")
	encodeIfNotNil(e, v.DbtrAcct, "urn2:DbtrAcct")
	encodeIfNotNil(e, v.DbtrAgt, "urn2:DbtrAgt")
	encodeIfNotNil(e, v.DbtrAgtAcct, "urn2:DbtrAgtAcct")
	encodeIfNotNil(e, v.CdtrAgt, "urn2:CdtrAgt")
	encodeIfNotNil(e, v.CdtrAgtAcct, "urn2:CdtrAgtAcct")
	encodeIfNotNil(e, v.Cdtr, "urn2:Cdtr")
	encodeIfNotNil(e, v.CdtrAcct, "urn2:CdtrAcct")
	encodeIfNotNil(e, v.UltmtCdtr, "urn2:UltmtCdtr")
	encodeIfNotNil(e, v.InstrForCdtrAgt, "urn2:InstrForCdtrAgt")
	encodeIfNotNil(e, v.Purp, "urn2:Purp")
	encodeIfNotNil(e, v.RgltryRptg, "urn2:RgltryRptg")
	encodeIfNotNil(e, v.RltdRmtInf, "urn2:RltdRmtInf")
	encodeIfNotNil(e, v.RmtInf, "urn2:RmtInf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CreditorReferenceInformation2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Ref, "urn2:Ref")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CreditorReferenceType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CreditorReferenceType2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DateAndPlaceOfBirth1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.BirthDt, "urn2:BirthDt")
	encodeIfNotNil(e, v.PrvcOfBirth, "urn2:PrvcOfBirth")
	encodeIfNotNil(e, v.CityOfBirth, "urn2:CityOfBirth")
	encodeIfNotNil(e, v.CtryOfBirth, "urn2:CtryOfBirth")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DatePeriod2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.FrDt, "urn2:FrDt")
	encodeIfNotNil(e, v.ToDt, "urn2:ToDt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DiscountAmountAndType1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DiscountAmountType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentAdjustment1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	encodeIfNotNil(e, v.CdtDbtInd, "urn2:CdtDbtInd")
	encodeIfNotNil(e, v.Rsn, "urn2:Rsn")
	encodeIfNotNil(e, v.AddtlInf, "urn2:AddtlInf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentLineIdentification1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Nb, "urn2:Nb")
	encodeIfNotNil(e, v.RltdDt, "urn2:RltdDt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentLineInformation11) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Id, "urn2:Id")
	encodeIfNotNil(e, v.Desc, "urn2:Desc")
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentLineType1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentLineType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FIToFICustomerCreditTransferV08) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.GrpHdr, "urn2:GrpHdr")
	encodeIfNotNil(e, v.CdtTrfTxInf, "urn2:CdtTrfTxInf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification181) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.BICFI, "urn2:BICFI")
	encodeIfNotNil(e, v.ClrSysMmbId, "urn2:ClrSysMmbId")
	encodeIfNotNil(e, v.LEI, "urn2:LEI")
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	encodeIfNotNil(e, v.PstlAdr, "urn2:PstlAdr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification182) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.ClrSysMmbId, "urn2:ClrSysMmbId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Garnishment31) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Grnshee, "urn2:Grnshee")
	encodeIfNotNil(e, v.GrnshmtAdmstr, "urn2:GrnshmtAdmstr")
	encodeIfNotNil(e, v.RefNb, "urn2:RefNb")
	encodeIfNotNil(e, v.Dt, "urn2:Dt")
	encodeIfNotNil(e, v.RmtdAmt, "urn2:RmtdAmt")
	encodeIfNotNil(e, v.FmlyMdclInsrncInd, "urn2:FmlyMdclInsrncInd")
	encodeIfNotNil(e, v.MplyeeTermntnInd, "urn2:MplyeeTermntnInd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GarnishmentType1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GarnishmentType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GenericAccountIdentification1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Id, "urn2:Id")
	encodeIfNotNil(e, v.SchmeNm, "urn2:SchmeNm")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GenericOrganisationIdentification1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Id, "urn2:Id")
	encodeIfNotNil(e, v.SchmeNm, "urn2:SchmeNm")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GenericPersonIdentification1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Id, "urn2:Id")
	encodeIfNotNil(e, v.SchmeNm, "urn2:SchmeNm")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GroupHeader931) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.MsgId, "urn2:MsgId")
	encodeIfNotNil(e, v.CreDtTm, "urn2:CreDtTm")
	encodeIfNotNil(e, v.NbOfTxs, "urn2:NbOfTxs")
	encodeIfNotNil(e, v.SttlmInf, "urn2:SttlmInf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v InstructionForCreditorAgent1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.InstrInf, "urn2:InstrInf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v LocalInstrument2Choice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v NameAndAddress161) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	encodeIfNotNil(e, v.Adr, "urn2:Adr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OrganisationIdentification291) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.AnyBIC, "urn2:AnyBIC")
	encodeIfNotNil(e, v.LEI, "urn2:LEI")
	encodeIfNotNil(e, v.Othr, "urn2:Othr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OrganisationIdentificationSchemeName1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party38Choice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.OrgId, "urn2:OrgId")
	encodeIfNotNil(e, v.PrvtId, "urn2:PrvtId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PartyIdentification1351) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	encodeIfNotNil(e, v.PstlAdr, "urn2:PstlAdr")
	encodeIfNotNil(e, v.Id, "urn2:Id")
	encodeIfNotNil(e, v.CtryOfRes, "urn2:CtryOfRes")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PartyIdentification1352) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	encodeIfNotNil(e, v.PstlAdr, "urn2:PstlAdr")
	encodeIfNotNil(e, v.Id, "urn2:Id")
	encodeIfNotNil(e, v.CtryOfRes, "urn2:CtryOfRes")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentIdentification71) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.InstrId, "urn2:InstrId")
	encodeIfNotNil(e, v.EndToEndId, "urn2:EndToEndId")
	encodeIfNotNil(e, v.TxId, "urn2:TxId")
	encodeIfNotNil(e, v.UETR, "urn2:UETR")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTypeInformation281) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.InstrPrty, "urn2:InstrPrty")
	encodeIfNotNil(e, v.SvcLvl, "urn2:SvcLvl")
	encodeIfNotNil(e, v.LclInstrm, "urn2:LclInstrm")
	encodeIfNotNil(e, v.CtgyPurp, "urn2:CtgyPurp")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PersonIdentification131) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.DtAndPlcOfBirth, "urn2:DtAndPlcOfBirth")
	encodeIfNotNil(e, v.Othr, "urn2:Othr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PersonIdentificationSchemeName1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PostalAddress241) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Dept, "urn2:Dept")
	encodeIfNotNil(e, v.SubDept, "urn2:SubDept")
	encodeIfNotNil(e, v.StrtNm, "urn2:StrtNm")
	encodeIfNotNil(e, v.BldgNb, "urn2:BldgNb")
	encodeIfNotNil(e, v.Flr, "urn2:Flr")
	encodeIfNotNil(e, v.PstBx, "urn2:PstBx")
	encodeIfNotNil(e, v.Room, "urn2:Room")
	encodeIfNotNil(e, v.PstCd, "urn2:PstCd")
	encodeIfNotNil(e, v.TwnNm, "urn2:TwnNm")
	encodeIfNotNil(e, v.TwnLctnNm, "urn2:TwnLctnNm")
	encodeIfNotNil(e, v.DstrctNm, "urn2:DstrctNm")
	encodeIfNotNil(e, v.CtrySubDvsn, "urn2:CtrySubDvsn")
	encodeIfNotNil(e, v.Ctry, "urn2:Ctry")
	encodeIfNotNil(e, v.AdrLine, "urn2:AdrLine")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PostalAddress242) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Dept, "urn2:Dept")
	encodeIfNotNil(e, v.SubDept, "urn2:SubDept")
	encodeIfNotNil(e, v.StrtNm, "urn2:StrtNm")
	encodeIfNotNil(e, v.BldgNb, "urn2:BldgNb")
	encodeIfNotNil(e, v.Flr, "urn2:Flr")
	encodeIfNotNil(e, v.PstBx, "urn2:PstBx")
	encodeIfNotNil(e, v.Room, "urn2:Room")
	encodeIfNotNil(e, v.PstCd, "urn2:PstCd")
	encodeIfNotNil(e, v.TwnNm, "urn2:TwnNm")
	encodeIfNotNil(e, v.TwnLctnNm, "urn2:TwnLctnNm")
	encodeIfNotNil(e, v.DstrctNm, "urn2:DstrctNm")
	encodeIfNotNil(e, v.CtrySubDvsn, "urn2:CtrySubDvsn")
	encodeIfNotNil(e, v.Ctry, "urn2:Ctry")
	encodeIfNotNil(e, v.AdrLine, "urn2:AdrLine")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ProxyAccountIdentification1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Id, "urn2:Id")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ProxyAccountType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Purpose2Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ReferredDocumentInformation71) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Nb, "urn2:Nb")
	encodeIfNotNil(e, v.RltdDt, "urn2:RltdDt")
	encodeIfNotNil(e, v.LineDtls, "urn2:LineDtls")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ReferredDocumentType3Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ReferredDocumentType4) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RegulatoryAuthority2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	encodeIfNotNil(e, v.Ctry, "urn2:Ctry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RegulatoryReporting3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.DbtCdtRptgInd, "urn2:DbtCdtRptgInd")
	encodeIfNotNil(e, v.Authrty, "urn2:Authrty")
	encodeIfNotNil(e, v.Dtls, "urn2:Dtls")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RemittanceAmount2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.DuePyblAmt, "urn2:DuePyblAmt")
	encodeIfNotNil(e, v.DscntApldAmt, "urn2:DscntApldAmt")
	encodeIfNotNil(e, v.CdtNoteAmt, "urn2:CdtNoteAmt")
	encodeIfNotNil(e, v.TaxAmt, "urn2:TaxAmt")
	encodeIfNotNil(e, v.AdjstmntAmtAndRsn, "urn2:AdjstmntAmtAndRsn")
	encodeIfNotNil(e, v.RmtdAmt, "urn2:RmtdAmt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RemittanceAmount3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.DuePyblAmt, "urn2:DuePyblAmt")
	encodeIfNotNil(e, v.DscntApldAmt, "urn2:DscntApldAmt")
	encodeIfNotNil(e, v.CdtNoteAmt, "urn2:CdtNoteAmt")
	encodeIfNotNil(e, v.TaxAmt, "urn2:TaxAmt")
	encodeIfNotNil(e, v.AdjstmntAmtAndRsn, "urn2:AdjstmntAmtAndRsn")
	encodeIfNotNil(e, v.RmtdAmt, "urn2:RmtdAmt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RemittanceInformation161) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Ustrd, "urn2:Ustrd")
	encodeIfNotNil(e, v.Strd, "urn2:Strd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RemittanceLocation71) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.RmtId, "urn2:RmtId")
	encodeIfNotNil(e, v.RmtLctnDtls, "urn2:RmtLctnDtls")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RemittanceLocationData11) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Mtd, "urn2:Mtd")
	encodeIfNotNil(e, v.ElctrncAdr, "urn2:ElctrncAdr")
	encodeIfNotNil(e, v.PstlAdr, "urn2:PstlAdr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ServiceLevel8Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SettlementInstruction71) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.SttlmMtd, "urn2:SttlmMtd")
	encodeIfNotNil(e, v.ClrSys, "urn2:ClrSys")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v StructuredRegulatoryReporting3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Dt, "urn2:Dt")
	encodeIfNotNil(e, v.Ctry, "urn2:Ctry")
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	encodeIfNotNil(e, v.Inf, "urn2:Inf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v StructuredRemittanceInformation161) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.RfrdDocInf, "urn2:RfrdDocInf")
	encodeIfNotNil(e, v.RfrdDocAmt, "urn2:RfrdDocAmt")
	encodeIfNotNil(e, v.CdtrRefInf, "urn2:CdtrRefInf")
	encodeIfNotNil(e, v.Invcr, "urn2:Invcr")
	encodeIfNotNil(e, v.Invcee, "urn2:Invcee")
	encodeIfNotNil(e, v.TaxRmt, "urn2:TaxRmt")
	encodeIfNotNil(e, v.GrnshmtRmt, "urn2:GrnshmtRmt")
	encodeIfNotNil(e, v.AddtlRmtInf, "urn2:AddtlRmtInf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxAmount2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Rate, "urn2:Rate")
	encodeIfNotNil(e, v.TaxblBaseAmt, "urn2:TaxblBaseAmt")
	encodeIfNotNil(e, v.TtlAmt, "urn2:TtlAmt")
	encodeIfNotNil(e, v.Dtls, "urn2:Dtls")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxAmountAndType1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxAmountType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxAuthorisation1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Titl, "urn2:Titl")
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxInformation7) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cdtr, "urn2:Cdtr")
	encodeIfNotNil(e, v.Dbtr, "urn2:Dbtr")
	encodeIfNotNil(e, v.UltmtDbtr, "urn2:UltmtDbtr")
	encodeIfNotNil(e, v.AdmstnZone, "urn2:AdmstnZone")
	encodeIfNotNil(e, v.RefNb, "urn2:RefNb")
	encodeIfNotNil(e, v.Mtd, "urn2:Mtd")
	encodeIfNotNil(e, v.TtlTaxblBaseAmt, "urn2:TtlTaxblBaseAmt")
	encodeIfNotNil(e, v.TtlTaxAmt, "urn2:TtlTaxAmt")
	encodeIfNotNil(e, v.Dt, "urn2:Dt")
	encodeIfNotNil(e, v.SeqNb, "urn2:SeqNb")
	encodeIfNotNil(e, v.Rcrd, "urn2:Rcrd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxParty1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.TaxId, "urn2:TaxId")
	encodeIfNotNil(e, v.RegnId, "urn2:RegnId")
	encodeIfNotNil(e, v.TaxTp, "urn2:TaxTp")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxParty2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.TaxId, "urn2:TaxId")
	encodeIfNotNil(e, v.RegnId, "urn2:RegnId")
	encodeIfNotNil(e, v.TaxTp, "urn2:TaxTp")
	encodeIfNotNil(e, v.Authstn, "urn2:Authstn")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxPeriod2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Yr, "urn2:Yr")
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.FrToDt, "urn2:FrToDt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxRecord2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Ctgy, "urn2:Ctgy")
	encodeIfNotNil(e, v.CtgyDtls, "urn2:CtgyDtls")
	encodeIfNotNil(e, v.DbtrSts, "urn2:DbtrSts")
	encodeIfNotNil(e, v.CertId, "urn2:CertId")
	encodeIfNotNil(e, v.FrmsCd, "urn2:FrmsCd")
	encodeIfNotNil(e, v.Prd, "urn2:Prd")
	encodeIfNotNil(e, v.TaxAmt, "urn2:TaxAmt")
	encodeIfNotNil(e, v.AddtlInf, "urn2:AddtlInf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxRecordDetails2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Prd, "urn2:Prd")
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
