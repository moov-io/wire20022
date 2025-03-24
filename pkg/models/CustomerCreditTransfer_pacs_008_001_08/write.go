package CustomerCreditTransfer_pacs_008_001_08

import (
	"encoding/json"
	"encoding/xml"
)

func isNotNil(value any) bool {
	if value == nil || value == "" {
		return false
	}
	return true
}

func encodeIfNotNil(e *xml.Encoder, value any, elementName string) error {
	if value == nil || value == "" {
		return nil
	}
	return e.EncodeElement(value, xml.StartElement{Name: xml.Name{Local: elementName}})
}
func encodeJsonIfNotNil(data map[string]any, value any, elementName string) {
	if value == nil || value == "" {
		return
	}
	data[elementName] = value
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

func (v Document) MarshalJSON() ([]byte, error) {
	// Define a map to structure the JSON output
	data := map[string]any{
		"Document": map[string]any{
			"xmlns:urn":              "urn2:iso:std:iso:20022:tech:xsd:pacs.008.001.08",
			"urn2:FIToFICstmrCdtTrf": v.FIToFICstmrCdtTrf,
		},
	}

	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v AccountIdentification4Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.IBAN, "urn2:IBAN")
	encodeIfNotNil(e, v.Othr, "urn2:Othr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v AccountIdentification4Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.IBAN, "urn2:IBAN")
	encodeJsonIfNotNil(data, v.Othr, "urn2:Othr")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v AccountSchemeName1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v AccountSchemeName1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ActiveCurrencyAndAmountFedwire1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Value, "urn2:Value")
	encodeIfNotNil(e, v.Ccy, "urn2:Ccy")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ActiveCurrencyAndAmountFedwire1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Value, "urn2:Value")
	encodeJsonIfNotNil(data, v.Ccy, "urn2:Ccy")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ActiveOrHistoricCurrencyAndAmount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Value, "urn2:Value")
	encodeIfNotNil(e, v.Ccy, "urn2:Ccy")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ActiveOrHistoricCurrencyAndAmount) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Value, "urn2:Value")
	encodeJsonIfNotNil(data, v.Ccy, "urn2:Ccy")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification61) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.FinInstnId, "urn2:FinInstnId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v BranchAndFinancialInstitutionIdentification61) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.FinInstnId, "urn2:FinInstnId")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification62) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.FinInstnId, "urn2:FinInstnId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v BranchAndFinancialInstitutionIdentification62) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.FinInstnId, "urn2:FinInstnId")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification63) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.FinInstnId, "urn2:FinInstnId")
	encodeIfNotNil(e, v.BrnchId, "urn2:BrnchId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v BranchAndFinancialInstitutionIdentification63) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.FinInstnId, "urn2:FinInstnId")
	encodeJsonIfNotNil(data, v.BrnchId, "urn2:BrnchId")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchData31) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Id, "urn2:Id")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v BranchData31) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Id, "urn2:Id")
	return json.Marshal(data)
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
func (v CashAccount38) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Id, "urn2:Id")
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.Ccy, "urn2:Ccy")
	encodeJsonIfNotNil(data, v.Nm, "urn2:Nm")
	encodeJsonIfNotNil(data, v.Prxy, "urn2:Prxy")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CashAccountType2Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v CashAccountType2Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CategoryPurpose1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v CategoryPurpose1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Charges71) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	encodeIfNotNil(e, v.Agt, "urn2:Agt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v Charges71) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Amt, "urn2:Amt")
	encodeJsonIfNotNil(data, v.Agt, "urn2:Agt")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemIdentification2Choice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ClearingSystemIdentification2Choice1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemIdentification2Choice2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ClearingSystemIdentification2Choice2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemIdentification3Choice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ClearingSystemIdentification3Choice1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification21) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.ClrSysId, "urn2:ClrSysId")
	encodeIfNotNil(e, v.MmbId, "urn2:MmbId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ClearingSystemMemberIdentification21) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.ClrSysId, "urn2:ClrSysId")
	encodeJsonIfNotNil(data, v.MmbId, "urn2:MmbId")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification22) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.ClrSysId, "urn2:ClrSysId")
	encodeIfNotNil(e, v.MmbId, "urn2:MmbId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ClearingSystemMemberIdentification22) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.ClrSysId, "urn2:ClrSysId")
	encodeJsonIfNotNil(data, v.MmbId, "urn2:MmbId")
	return json.Marshal(data)
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

func (v CreditTransferTransaction391) MarshalJSON() ([]byte, error) {
	e := make(map[string]any)
	encodeJsonIfNotNil(e, v.PmtId, "urn2:PmtId")
	encodeJsonIfNotNil(e, v.PmtTpInf, "urn2:PmtTpInf")
	encodeJsonIfNotNil(e, v.IntrBkSttlmAmt, "urn2:IntrBkSttlmAmt")
	encodeJsonIfNotNil(e, v.IntrBkSttlmDt, "urn2:IntrBkSttlmDt")
	encodeJsonIfNotNil(e, v.AccptncDtTm, "urn2:AccptncDtTm")
	encodeJsonIfNotNil(e, v.InstdAmt, "urn2:InstdAmt")
	encodeJsonIfNotNil(e, v.XchgRate, "urn2:XchgRate")
	encodeJsonIfNotNil(e, v.ChrgBr, "urn2:ChrgBr")
	encodeJsonIfNotNil(e, v.ChrgsInf, "urn2:ChrgsInf")
	encodeJsonIfNotNil(e, v.PrvsInstgAgt1, "urn2:PrvsInstgAgt1")
	encodeJsonIfNotNil(e, v.PrvsInstgAgt1Acct, "urn2:PrvsInstgAgt1Acct")
	encodeJsonIfNotNil(e, v.PrvsInstgAgt2, "urn2:PrvsInstgAgt2")
	encodeJsonIfNotNil(e, v.PrvsInstgAgt2Acct, "urn2:PrvsInstgAgt2Acct")
	encodeJsonIfNotNil(e, v.PrvsInstgAgt3, "urn2:PrvsInstgAgt3")
	encodeJsonIfNotNil(e, v.PrvsInstgAgt3Acct, "urn2:PrvsInstgAgt3Acct")
	encodeJsonIfNotNil(e, v.InstgAgt, "urn2:InstgAgt")
	encodeJsonIfNotNil(e, v.InstdAgt, "urn2:InstdAgt")
	encodeJsonIfNotNil(e, v.IntrmyAgt1, "urn2:IntrmyAgt1")
	encodeJsonIfNotNil(e, v.IntrmyAgt1Acct, "urn2:IntrmyAgt1Acct")
	encodeJsonIfNotNil(e, v.IntrmyAgt2, "urn2:IntrmyAgt2")
	encodeJsonIfNotNil(e, v.IntrmyAgt2Acct, "urn2:IntrmyAgt2Acct")
	encodeJsonIfNotNil(e, v.IntrmyAgt3, "urn2:IntrmyAgt3")
	encodeJsonIfNotNil(e, v.IntrmyAgt3Acct, "urn2:IntrmyAgt3Acct")
	encodeJsonIfNotNil(e, v.UltmtDbtr, "urn2:UltmtDbtr")
	encodeJsonIfNotNil(e, v.InitgPty, "urn2:InitgPty")
	encodeJsonIfNotNil(e, v.Dbtr, "urn2:Dbtr")
	encodeJsonIfNotNil(e, v.DbtrAcct, "urn2:DbtrAcct")
	encodeJsonIfNotNil(e, v.DbtrAgt, "urn2:DbtrAgt")
	encodeJsonIfNotNil(e, v.DbtrAgtAcct, "urn2:DbtrAgtAcct")
	encodeJsonIfNotNil(e, v.CdtrAgt, "urn2:CdtrAgt")
	encodeJsonIfNotNil(e, v.CdtrAgtAcct, "urn2:CdtrAgtAcct")
	encodeJsonIfNotNil(e, v.Cdtr, "urn2:Cdtr")
	encodeJsonIfNotNil(e, v.CdtrAcct, "urn2:CdtrAcct")
	encodeJsonIfNotNil(e, v.UltmtCdtr, "urn2:UltmtCdtr")
	encodeJsonIfNotNil(e, v.InstrForCdtrAgt, "urn2:InstrForCdtrAgt")
	encodeJsonIfNotNil(e, v.Purp, "urn2:Purp")
	encodeJsonIfNotNil(e, v.RgltryRptg, "urn2:RgltryRptg")
	encodeJsonIfNotNil(e, v.RltdRmtInf, "urn2:RltdRmtInf")
	encodeJsonIfNotNil(e, v.RmtInf, "urn2:RmtInf")
	return json.Marshal(e)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CreditorReferenceInformation2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Ref, "urn2:Ref")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v CreditorReferenceInformation2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.Ref, "urn2:Ref")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CreditorReferenceType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v CreditorReferenceType1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v CreditorReferenceType2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v CreditorReferenceType2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeJsonIfNotNil(data, v.Issr, "urn2:Issr")
	return json.Marshal(data)
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
func (v DateAndPlaceOfBirth1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.BirthDt, "urn2:BirthDt")
	encodeJsonIfNotNil(data, v.PrvcOfBirth, "urn2:PrvcOfBirth")
	encodeJsonIfNotNil(data, v.CityOfBirth, "urn2:CityOfBirth")
	encodeJsonIfNotNil(data, v.CtryOfBirth, "urn2:CtryOfBirth")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DatePeriod2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.FrDt, "urn2:FrDt")
	encodeIfNotNil(e, v.ToDt, "urn2:ToDt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v DatePeriod2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.FrDt, "urn2:FrDt")
	encodeJsonIfNotNil(data, v.ToDt, "urn2:ToDt")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DiscountAmountAndType1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v DiscountAmountAndType1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.Amt, "urn2:Amt")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DiscountAmountType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v DiscountAmountType1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
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
func (v DocumentAdjustment1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Amt, "urn2:Amt")
	encodeJsonIfNotNil(data, v.CdtDbtInd, "urn2:CdtDbtInd")
	encodeJsonIfNotNil(data, v.Rsn, "urn2:Rsn")
	encodeJsonIfNotNil(data, v.AddtlInf, "urn2:AddtlInf")
	return json.Marshal(data)
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
func (v DocumentLineIdentification1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.Nb, "urn2:Nb")
	encodeJsonIfNotNil(data, v.RltdDt, "urn2:RltdDt")
	return json.Marshal(data)
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
func (v DocumentLineInformation11) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Id, "urn2:Id")
	encodeJsonIfNotNil(data, v.Desc, "urn2:Desc")
	encodeJsonIfNotNil(data, v.Amt, "urn2:Amt")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentLineType1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v DocumentLineType1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeJsonIfNotNil(data, v.Issr, "urn2:Issr")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentLineType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v DocumentLineType1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FIToFICustomerCreditTransferV08) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.GrpHdr, "urn2:GrpHdr")
	encodeIfNotNil(e, v.CdtTrfTxInf, "urn2:CdtTrfTxInf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v FIToFICustomerCreditTransferV08) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.GrpHdr, "urn2:GrpHdr")
	encodeJsonIfNotNil(data, v.CdtTrfTxInf, "urn2:CdtTrfTxInf")
	return json.Marshal(data)
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
func (v FinancialInstitutionIdentification181) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.BICFI, "urn2:BICFI")
	encodeJsonIfNotNil(data, v.ClrSysMmbId, "urn2:ClrSysMmbId")
	encodeJsonIfNotNil(data, v.LEI, "urn2:LEI")
	encodeJsonIfNotNil(data, v.Nm, "urn2:Nm")
	encodeJsonIfNotNil(data, v.PstlAdr, "urn2:PstlAdr")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification182) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.ClrSysMmbId, "urn2:ClrSysMmbId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v FinancialInstitutionIdentification182) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.ClrSysMmbId, "urn2:ClrSysMmbId")
	return json.Marshal(data)
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
func (v Garnishment31) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Tp, "urn2:ClrSysMmbId")
	encodeJsonIfNotNil(data, v.Grnshee, "urn2:Grnshee")
	encodeJsonIfNotNil(data, v.GrnshmtAdmstr, "urn2:GrnshmtAdmstr")
	encodeJsonIfNotNil(data, v.RefNb, "urn2:RefNb")
	encodeJsonIfNotNil(data, v.Dt, "urn2:Dt")
	encodeJsonIfNotNil(data, v.RmtdAmt, "urn2:RmtdAmt")
	encodeJsonIfNotNil(data, v.FmlyMdclInsrncInd, "urn2:FmlyMdclInsrncInd")
	encodeJsonIfNotNil(data, v.MplyeeTermntnInd, "urn2:MplyeeTermntnInd")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GarnishmentType1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v GarnishmentType1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeJsonIfNotNil(data, v.Issr, "urn2:Issr")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GarnishmentType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v GarnishmentType1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
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
func (v GenericAccountIdentification1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Id, "urn2:Id")
	encodeJsonIfNotNil(data, v.SchmeNm, "urn2:SchmeNm")
	encodeJsonIfNotNil(data, v.Issr, "urn2:Issr")
	return json.Marshal(data)
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
func (v GenericOrganisationIdentification1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Id, "urn2:Id")
	encodeJsonIfNotNil(data, v.SchmeNm, "urn2:SchmeNm")
	encodeJsonIfNotNil(data, v.Issr, "urn2:Issr")
	return json.Marshal(data)
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
func (v GenericPersonIdentification1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Id, "urn2:Id")
	encodeJsonIfNotNil(data, v.SchmeNm, "urn2:SchmeNm")
	encodeJsonIfNotNil(data, v.Issr, "urn2:Issr")
	return json.Marshal(data)
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
func (v GroupHeader931) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.MsgId, "urn2:MsgId")
	encodeJsonIfNotNil(data, v.CreDtTm, "urn2:CreDtTm")
	encodeJsonIfNotNil(data, v.NbOfTxs, "urn2:NbOfTxs")
	encodeJsonIfNotNil(data, v.SttlmInf, "urn2:SttlmInf")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v InstructionForCreditorAgent1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.InstrInf, "urn2:InstrInf")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v InstructionForCreditorAgent1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.InstrInf, "urn2:InstrInf")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v LocalInstrument2Choice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v LocalInstrument2Choice1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v NameAndAddress161) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	encodeIfNotNil(e, v.Adr, "urn2:Adr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v NameAndAddress161) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Nm, "urn2:Nm")
	encodeJsonIfNotNil(data, v.Adr, "urn2:Adr")
	return json.Marshal(data)
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
func (v OrganisationIdentification291) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.AnyBIC, "urn2:AnyBIC")
	encodeJsonIfNotNil(data, v.LEI, "urn2:LEI")
	encodeJsonIfNotNil(data, v.Othr, "urn2:Othr")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OrganisationIdentificationSchemeName1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v OrganisationIdentificationSchemeName1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Party38Choice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.OrgId, "urn2:OrgId")
	encodeIfNotNil(e, v.PrvtId, "urn2:PrvtId")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v Party38Choice1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.OrgId, "urn2:OrgId")
	encodeJsonIfNotNil(data, v.PrvtId, "urn2:PrvtId")
	return json.Marshal(data)
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
func (v PartyIdentification1351) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Nm, "urn2:Nm")
	encodeJsonIfNotNil(data, v.PstlAdr, "urn2:PstlAdr")
	encodeJsonIfNotNil(data, v.Id, "urn2:Id")
	encodeJsonIfNotNil(data, v.CtryOfRes, "urn2:CtryOfRes")
	return json.Marshal(data)
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
func (v PartyIdentification1352) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Nm, "urn2:Nm")
	encodeJsonIfNotNil(data, v.PstlAdr, "urn2:PstlAdr")
	encodeJsonIfNotNil(data, v.Id, "urn2:Id")
	encodeJsonIfNotNil(data, v.CtryOfRes, "urn2:CtryOfRes")
	return json.Marshal(data)
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
func (v PaymentIdentification71) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.InstrId, "urn2:InstrId")
	encodeJsonIfNotNil(data, v.EndToEndId, "urn2:EndToEndId")
	encodeJsonIfNotNil(data, v.TxId, "urn2:TxId")
	encodeJsonIfNotNil(data, v.UETR, "urn2:UETR")
	return json.Marshal(data)
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
func (v PaymentTypeInformation281) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.InstrPrty, "urn2:InstrPrty")
	encodeJsonIfNotNil(data, v.SvcLvl, "urn2:SvcLvl")
	encodeJsonIfNotNil(data, v.LclInstrm, "urn2:LclInstrm")
	encodeJsonIfNotNil(data, v.CtgyPurp, "urn2:CtgyPurp")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PersonIdentification131) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.DtAndPlcOfBirth, "urn2:DtAndPlcOfBirth")
	encodeIfNotNil(e, v.Othr, "urn2:Othr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v PersonIdentification131) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.DtAndPlcOfBirth, "urn2:DtAndPlcOfBirth")
	encodeJsonIfNotNil(data, v.Othr, "urn2:Othr")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PersonIdentificationSchemeName1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v PersonIdentificationSchemeName1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
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
func (v PostalAddress241) MarshalJSON() ([]byte, error) {
	e := make(map[string]any)
	encodeJsonIfNotNil(e, v.Dept, "urn2:Dept")
	encodeJsonIfNotNil(e, v.SubDept, "urn2:SubDept")
	encodeJsonIfNotNil(e, v.StrtNm, "urn2:StrtNm")
	encodeJsonIfNotNil(e, v.BldgNb, "urn2:BldgNb")
	encodeJsonIfNotNil(e, v.Flr, "urn2:Flr")
	encodeJsonIfNotNil(e, v.PstBx, "urn2:PstBx")
	encodeJsonIfNotNil(e, v.Room, "urn2:Room")
	encodeJsonIfNotNil(e, v.PstCd, "urn2:PstCd")
	encodeJsonIfNotNil(e, v.TwnNm, "urn2:TwnNm")
	encodeJsonIfNotNil(e, v.TwnLctnNm, "urn2:TwnLctnNm")
	encodeJsonIfNotNil(e, v.DstrctNm, "urn2:DstrctNm")
	encodeJsonIfNotNil(e, v.CtrySubDvsn, "urn2:CtrySubDvsn")
	encodeJsonIfNotNil(e, v.Ctry, "urn2:Ctry")
	encodeJsonIfNotNil(e, v.AdrLine, "urn2:AdrLine")
	return json.Marshal(e)
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
func (v PostalAddress242) MarshalJSON() ([]byte, error) {
	e := make(map[string]any)
	encodeJsonIfNotNil(e, v.Dept, "urn2:Dept")
	encodeJsonIfNotNil(e, v.SubDept, "urn2:SubDept")
	encodeJsonIfNotNil(e, v.StrtNm, "urn2:StrtNm")
	encodeJsonIfNotNil(e, v.BldgNb, "urn2:BldgNb")
	encodeJsonIfNotNil(e, v.Flr, "urn2:Flr")
	encodeJsonIfNotNil(e, v.PstBx, "urn2:PstBx")
	encodeJsonIfNotNil(e, v.Room, "urn2:Room")
	encodeJsonIfNotNil(e, v.PstCd, "urn2:PstCd")
	encodeJsonIfNotNil(e, v.TwnNm, "urn2:TwnNm")
	encodeJsonIfNotNil(e, v.TwnLctnNm, "urn2:TwnLctnNm")
	encodeJsonIfNotNil(e, v.DstrctNm, "urn2:DstrctNm")
	encodeJsonIfNotNil(e, v.CtrySubDvsn, "urn2:CtrySubDvsn")
	encodeJsonIfNotNil(e, v.Ctry, "urn2:Ctry")
	encodeJsonIfNotNil(e, v.AdrLine, "urn2:AdrLine")
	return json.Marshal(e)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ProxyAccountIdentification1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Id, "urn2:Id")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ProxyAccountIdentification1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.Id, "urn2:Id")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ProxyAccountType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ProxyAccountType1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Purpose2Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v Purpose2Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
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
func (v ReferredDocumentInformation71) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.Nb, "urn2:Nb")
	encodeJsonIfNotNil(data, v.RltdDt, "urn2:RltdDt")
	encodeJsonIfNotNil(data, v.LineDtls, "urn2:LineDtls")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ReferredDocumentType3Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ReferredDocumentType3Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ReferredDocumentType4) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeIfNotNil(e, v.Issr, "urn2:Issr")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ReferredDocumentType4) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.CdOrPrtry, "urn2:CdOrPrtry")
	encodeJsonIfNotNil(data, v.Issr, "urn2:Issr")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RegulatoryAuthority2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	encodeIfNotNil(e, v.Ctry, "urn2:Ctry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v RegulatoryAuthority2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Nm, "urn2:Nm")
	encodeJsonIfNotNil(data, v.Ctry, "urn2:Ctry")
	return json.Marshal(data)
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
func (v RegulatoryReporting3) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.DbtCdtRptgInd, "urn2:DbtCdtRptgInd")
	encodeJsonIfNotNil(data, v.Authrty, "urn2:Authrty")
	encodeJsonIfNotNil(data, v.Dtls, "urn2:Dtls")
	return json.Marshal(data)
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
func (v RemittanceAmount2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.DuePyblAmt, "urn2:DuePyblAmt")
	encodeJsonIfNotNil(data, v.DscntApldAmt, "urn2:DscntApldAmt")
	encodeJsonIfNotNil(data, v.CdtNoteAmt, "urn2:CdtNoteAmt")
	encodeJsonIfNotNil(data, v.TaxAmt, "urn2:TaxAmt")
	encodeJsonIfNotNil(data, v.AdjstmntAmtAndRsn, "urn2:AdjstmntAmtAndRsn")
	encodeJsonIfNotNil(data, v.RmtdAmt, "urn2:RmtdAmt")
	return json.Marshal(data)
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
func (v RemittanceAmount3) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.DuePyblAmt, "urn2:DuePyblAmt")
	encodeJsonIfNotNil(data, v.DscntApldAmt, "urn2:DscntApldAmt")
	encodeJsonIfNotNil(data, v.CdtNoteAmt, "urn2:CdtNoteAmt")
	encodeJsonIfNotNil(data, v.TaxAmt, "urn2:TaxAmt")
	encodeJsonIfNotNil(data, v.AdjstmntAmtAndRsn, "urn2:AdjstmntAmtAndRsn")
	encodeJsonIfNotNil(data, v.RmtdAmt, "urn2:RmtdAmt")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RemittanceInformation161) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Ustrd, "urn2:Ustrd")
	encodeIfNotNil(e, v.Strd, "urn2:Strd")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v RemittanceInformation161) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Ustrd, "urn2:Ustrd")
	encodeJsonIfNotNil(data, v.Strd, "urn2:Strd")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RemittanceLocation71) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.RmtId, "urn2:RmtId")
	encodeIfNotNil(e, v.RmtLctnDtls, "urn2:RmtLctnDtls")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v RemittanceLocation71) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.RmtId, "urn2:RmtId")
	encodeJsonIfNotNil(data, v.RmtLctnDtls, "urn2:RmtLctnDtls")
	return json.Marshal(data)
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
func (v RemittanceLocationData11) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Mtd, "urn2:Mtd")
	encodeJsonIfNotNil(data, v.ElctrncAdr, "urn2:ElctrncAdr")
	encodeJsonIfNotNil(data, v.PstlAdr, "urn2:PstlAdr")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ServiceLevel8Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v ServiceLevel8Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SettlementInstruction71) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.SttlmMtd, "urn2:SttlmMtd")
	encodeIfNotNil(e, v.ClrSys, "urn2:ClrSys")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v SettlementInstruction71) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.SttlmMtd, "urn2:SttlmMtd")
	encodeJsonIfNotNil(data, v.ClrSys, "urn2:ClrSys")
	return json.Marshal(data)
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
func (v StructuredRegulatoryReporting3) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.Dt, "urn2:Dt")
	encodeJsonIfNotNil(data, v.Ctry, "urn2:Ctry")
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Amt, "urn2:Amt")
	encodeJsonIfNotNil(data, v.Inf, "urn2:Inf")
	return json.Marshal(data)
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
func (v StructuredRemittanceInformation161) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.RfrdDocInf, "urn2:RfrdDocInf")
	encodeJsonIfNotNil(data, v.RfrdDocAmt, "urn2:RfrdDocAmt")
	encodeJsonIfNotNil(data, v.CdtrRefInf, "urn2:CdtrRefInf")
	encodeJsonIfNotNil(data, v.Invcr, "urn2:Invcr")
	encodeJsonIfNotNil(data, v.Invcee, "urn2:Invcee")
	encodeJsonIfNotNil(data, v.TaxRmt, "urn2:TaxRmt")
	encodeJsonIfNotNil(data, v.GrnshmtRmt, "urn2:GrnshmtRmt")
	encodeJsonIfNotNil(data, v.AddtlRmtInf, "urn2:AddtlRmtInf")
	return json.Marshal(data)
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
func (v TaxAmount2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Rate, "urn2:Rate")
	encodeJsonIfNotNil(data, v.TaxblBaseAmt, "urn2:TaxblBaseAmt")
	encodeJsonIfNotNil(data, v.TtlAmt, "urn2:TtlAmt")
	encodeJsonIfNotNil(data, v.Dtls, "urn2:Dtls")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxAmountAndType1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Tp, "urn2:Tp")
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

func (v TaxAmountAndType1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.Amt, "urn2:Amt")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxAmountType1Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Cd, "urn2:Cd")
	encodeIfNotNil(e, v.Prtry, "urn2:Prtry")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v TaxAmountType1Choice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Cd, "urn2:Cd")
	encodeJsonIfNotNil(data, v.Prtry, "urn2:Prtry")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxAuthorisation1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Titl, "urn2:Titl")
	encodeIfNotNil(e, v.Nm, "urn2:Nm")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v TaxAuthorisation1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Titl, "urn2:Titl")
	encodeJsonIfNotNil(data, v.Nm, "urn2:Nm")
	return json.Marshal(data)
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
func (v TaxInformation7) MarshalJSON() ([]byte, error) {
	e := make(map[string]any)
	encodeJsonIfNotNil(e, v.Cdtr, "urn2:Cdtr")
	encodeJsonIfNotNil(e, v.Dbtr, "urn2:Dbtr")
	encodeJsonIfNotNil(e, v.UltmtDbtr, "urn2:UltmtDbtr")
	encodeJsonIfNotNil(e, v.AdmstnZone, "urn2:AdmstnZone")
	encodeJsonIfNotNil(e, v.RefNb, "urn2:RefNb")
	encodeJsonIfNotNil(e, v.Mtd, "urn2:Mtd")
	encodeJsonIfNotNil(e, v.TtlTaxblBaseAmt, "urn2:TtlTaxblBaseAmt")
	encodeJsonIfNotNil(e, v.TtlTaxAmt, "urn2:TtlTaxAmt")
	encodeJsonIfNotNil(e, v.Dt, "urn2:Dt")
	encodeJsonIfNotNil(e, v.SeqNb, "urn2:SeqNb")
	encodeJsonIfNotNil(e, v.Rcrd, "urn2:Rcrd")
	return json.Marshal(e)
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
func (v TaxParty1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.TaxId, "urn2:TaxId")
	encodeJsonIfNotNil(data, v.RegnId, "urn2:RegnId")
	encodeJsonIfNotNil(data, v.TaxTp, "urn2:TaxTp")
	return json.Marshal(data)
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
func (v TaxParty2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.TaxId, "urn2:TaxId")
	encodeJsonIfNotNil(data, v.RegnId, "urn2:RegnId")
	encodeJsonIfNotNil(data, v.TaxTp, "urn2:TaxTp")
	encodeJsonIfNotNil(data, v.Authstn, "urn2:Authstn")
	return json.Marshal(data)
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
func (v TaxPeriod2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Yr, "urn2:Yr")
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.FrToDt, "urn2:FrToDt")
	return json.Marshal(data)
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
func (v TaxRecord2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Tp, "urn2:Tp")
	encodeJsonIfNotNil(data, v.Ctgy, "urn2:Ctgy")
	encodeJsonIfNotNil(data, v.CtgyDtls, "urn2:CtgyDtls")
	encodeJsonIfNotNil(data, v.DbtrSts, "urn2:DbtrSts")
	encodeJsonIfNotNil(data, v.CertId, "urn2:CertId")
	encodeJsonIfNotNil(data, v.FrmsCd, "urn2:FrmsCd")
	encodeJsonIfNotNil(data, v.Prd, "urn2:Prd")
	encodeJsonIfNotNil(data, v.TaxAmt, "urn2:TaxAmt")
	encodeJsonIfNotNil(data, v.AddtlInf, "urn2:AddtlInf")
	return json.Marshal(data)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TaxRecordDetails2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	encodeIfNotNil(e, v.Prd, "urn2:Prd")
	encodeIfNotNil(e, v.Amt, "urn2:Amt")
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
func (v TaxRecordDetails2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)
	encodeJsonIfNotNil(data, v.Prd, "urn2:Prd")
	encodeJsonIfNotNil(data, v.Amt, "urn2:Amt")
	return json.Marshal(data)
}
