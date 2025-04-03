package DrawdownResponse_014_001_07

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	DrawdownResponse "github.com/moov-io/fedwire20022/gen/DrawdownResponse_pain_014_001_07"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Pain014 struct {
	//A unique identifier (IMADFedwireFunds1) assigned to the message.
	MessageId string
	//The creation date and time (ISODateTime) of the message.
	CreateDatetime time.Time
	//In the Fedwire Funds Service, this is a person or entity that requests a drawdown.
	InitiatingParty model.PartyIdentify
	//Financial institution servicing an account for the debtor.
	DebtorAgent model.Agent
	// /Financial institution servicing an account for the creditor.
	CreditorAgent model.Agent
	//This should be the Message Identification of the original drawdown request (pain.013) message to which this drawdown request response message relates.
	OriginalMessageId string
	//Specifies the original message name identifier to which the message refers.
	OriginalMessageNameId string
	//Date and time at which the original message was created.
	OriginalCreationDateTime time.Time
	//Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInfoId string
	//Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus TransactionInfoAndStatus
}
type Pain014Message struct {
	model Pain014
	doc   DrawdownResponse.Document
}

func NewPain014Message() Pain014Message {
	return Pain014Message{
		model: Pain014{},
	}
}
func (msg *Pain014Message) CreateDocument() {
	msg.doc = DrawdownResponse.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.07",
			Local: "Document",
		},
	}
	var CdtrPmtActvtnReqStsRpt DrawdownResponse.CreditorPaymentActivationRequestStatusReportV07

	var GrpHdr DrawdownResponse.GroupHeader871
	if msg.model.MessageId != "" {
		GrpHdr.MsgId = DrawdownResponse.Max35Text(msg.model.MessageId)
	}
	if !isEmpty(msg.model.CreateDatetime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.model.CreateDatetime)
	}
	if !isEmpty(msg.model.InitiatingParty) {
		GrpHdr.InitgPty = PartyIdentification1351From(msg.model.InitiatingParty)
	}
	if !isEmpty(msg.model.DebtorAgent) {
		Cd := DrawdownResponse.ExternalClearingSystemIdentification1CodeFixed(msg.model.DebtorAgent.PaymentSysCode)
		DbtrAgt := DrawdownResponse.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: DrawdownResponse.FinancialInstitutionIdentification181{
				ClrSysMmbId: DrawdownResponse.ClearingSystemMemberIdentification21{
					ClrSysId: DrawdownResponse.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: DrawdownResponse.RoutingNumberFRS1(msg.model.DebtorAgent.PaymentSysMemberId),
				},
			},
		}
		if !isEmpty(DbtrAgt) {
			GrpHdr.DbtrAgt = DbtrAgt
		}
	}
	if !isEmpty(msg.model.CreditorAgent) {
		Cd := DrawdownResponse.ExternalClearingSystemIdentification1CodeFixed(msg.model.CreditorAgent.PaymentSysCode)
		CdtrAgt := DrawdownResponse.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: DrawdownResponse.FinancialInstitutionIdentification181{
				ClrSysMmbId: DrawdownResponse.ClearingSystemMemberIdentification21{
					ClrSysId: DrawdownResponse.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: DrawdownResponse.RoutingNumberFRS1(msg.model.CreditorAgent.PaymentSysMemberId),
				},
			},
		}
		if !isEmpty(CdtrAgt) {
			GrpHdr.CdtrAgt = CdtrAgt
		}
	}
	if !isEmpty(GrpHdr) {
		CdtrPmtActvtnReqStsRpt.GrpHdr = GrpHdr
	}

	var OrgnlGrpInfAndSts DrawdownResponse.OriginalGroupInformation301
	if msg.model.OriginalMessageId != "" {
		OrgnlGrpInfAndSts.OrgnlMsgId = DrawdownResponse.IMADFedwireFunds1(msg.model.OriginalMessageId)
	}
	if msg.model.OriginalMessageNameId != "" {
		OrgnlGrpInfAndSts.OrgnlMsgNmId = DrawdownResponse.MessageNameIdentificationFRS1(msg.model.OriginalMessageNameId)
	}
	if !isEmpty(msg.model.OriginalCreationDateTime) {
		OrgnlGrpInfAndSts.OrgnlCreDtTm = fedwire.ISODateTime(msg.model.OriginalCreationDateTime)
	}
	if !isEmpty(OrgnlGrpInfAndSts) {
		CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts = OrgnlGrpInfAndSts
	}

	var OrgnlPmtInfAndSts DrawdownResponse.OriginalPaymentInstruction311
	if msg.model.OriginalPaymentInfoId != "" {
		OrgnlPmtInfAndSts.OrgnlPmtInfId = DrawdownResponse.Max35Text(msg.model.OriginalPaymentInfoId)
	}
	if !isEmpty(msg.model.TransactionInformationAndStatus) {
		OrgnlPmtInfAndSts.TxInfAndSts = PaymentTransaction1041From(msg.model.TransactionInformationAndStatus)
	}

	if !isEmpty(OrgnlPmtInfAndSts) {
		CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts = OrgnlPmtInfAndSts
	}
	if !isEmpty(CdtrPmtActvtnReqStsRpt) {
		msg.doc.CdtrPmtActvtnReqStsRpt = CdtrPmtActvtnReqStsRpt
	}
}
func (msg *Pain014Message) GetXML() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(msg.doc, "", "\t")
	if err != nil {
		return nil, err
	}

	// Convert byte slice to string for manipulation
	xmlString := string(xmlData)

	// Keep the xmlns only in the <Document> tag, remove from others
	xmlString = removeExtraXMLNS(xmlString)

	// Regular expression to match scientific notation (e.g., 9.93229443e+06)
	re := regexp.MustCompile(`>(\d+\.\d+(?:e[+-]?\d+)?|\d+e[+-]?\d+)<`)

	// Replace matched numbers with properly formatted ones
	xmlString = re.ReplaceAllStringFunc(xmlString, func(match string) string {
		// Extract the number inside the tags
		numberStr := strings.Trim(match, "<>")

		// Convert to float
		number, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			return match // Return the original string if conversion fails
		}

		// Format it as a standard decimal number with 2 decimal places
		return fmt.Sprintf(">%.2f<", number)
	})

	// Convert back to []byte
	return []byte(xmlString), nil
	// return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *Pain014Message) GetJson() ([]byte, error) {
	return json.MarshalIndent(msg.doc.CdtrPmtActvtnReqStsRpt, "", " ")
}

func removeExtraXMLNS(xmlStr string) string {
	// Find the first occurrence of <Document ...> (keep this)
	docStart := strings.Index(xmlStr, "<Document")
	if docStart == -1 {
		return xmlStr // Return original if <Document> not found
	}

	// Find the end of the <Document> opening tag
	docEnd := strings.Index(xmlStr[docStart:], ">")
	if docEnd == -1 {
		return xmlStr
	}
	docEnd += docStart // Adjust index

	// Remove all occurrences of xmlns="..." except in <Document>
	cleanXML := xmlStr[:docEnd+1] + // Keep <Document> with its xmlns
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:pain.014.001.07"`, "")

	return cleanXML
}
