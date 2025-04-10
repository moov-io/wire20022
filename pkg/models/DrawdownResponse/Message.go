package DrawdownResponse

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	pain014 "github.com/moov-io/fedwire20022/gen/DrawdownResponse_pain_014_001_07"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:pain.014.001.07"

type MessageModel struct {
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
type Message struct {
	data MessageModel
	doc  pain014.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = pain014.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var CdtrPmtActvtnReqStsRpt pain014.CreditorPaymentActivationRequestStatusReportV07

	var GrpHdr pain014.GroupHeader871
	if msg.data.MessageId != "" {
		GrpHdr.MsgId = pain014.Max35Text(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreateDatetime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreateDatetime)
	}
	if !isEmpty(msg.data.InitiatingParty) {
		GrpHdr.InitgPty = PartyIdentification1351From(msg.data.InitiatingParty)
	}
	if !isEmpty(msg.data.DebtorAgent) {
		Cd := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.data.DebtorAgent.PaymentSysCode)
		DbtrAgt := pain014.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pain014.FinancialInstitutionIdentification181{
				ClrSysMmbId: pain014.ClearingSystemMemberIdentification21{
					ClrSysId: pain014.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pain014.RoutingNumberFRS1(msg.data.DebtorAgent.PaymentSysMemberId),
				},
			},
		}
		if !isEmpty(DbtrAgt) {
			GrpHdr.DbtrAgt = DbtrAgt
		}
	}
	if !isEmpty(msg.data.CreditorAgent) {
		Cd := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.data.CreditorAgent.PaymentSysCode)
		CdtrAgt := pain014.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pain014.FinancialInstitutionIdentification181{
				ClrSysMmbId: pain014.ClearingSystemMemberIdentification21{
					ClrSysId: pain014.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pain014.RoutingNumberFRS1(msg.data.CreditorAgent.PaymentSysMemberId),
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

	var OrgnlGrpInfAndSts pain014.OriginalGroupInformation301
	if msg.data.OriginalMessageId != "" {
		OrgnlGrpInfAndSts.OrgnlMsgId = pain014.IMADFedwireFunds1(msg.data.OriginalMessageId)
	}
	if msg.data.OriginalMessageNameId != "" {
		OrgnlGrpInfAndSts.OrgnlMsgNmId = pain014.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
	}
	if !isEmpty(msg.data.OriginalCreationDateTime) {
		OrgnlGrpInfAndSts.OrgnlCreDtTm = fedwire.ISODateTime(msg.data.OriginalCreationDateTime)
	}
	if !isEmpty(OrgnlGrpInfAndSts) {
		CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts = OrgnlGrpInfAndSts
	}

	var OrgnlPmtInfAndSts pain014.OriginalPaymentInstruction311
	if msg.data.OriginalPaymentInfoId != "" {
		OrgnlPmtInfAndSts.OrgnlPmtInfId = pain014.Max35Text(msg.data.OriginalPaymentInfoId)
	}
	if !isEmpty(msg.data.TransactionInformationAndStatus) {
		OrgnlPmtInfAndSts.TxInfAndSts = PaymentTransaction1041From(msg.data.TransactionInformationAndStatus)
	}

	if !isEmpty(OrgnlPmtInfAndSts) {
		CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts = OrgnlPmtInfAndSts
	}
	if !isEmpty(CdtrPmtActvtnReqStsRpt) {
		msg.doc.CdtrPmtActvtnReqStsRpt = CdtrPmtActvtnReqStsRpt
	}
}
func WriteXMLTo(filePath string, xml []byte) error {
	os.Mkdir("generated", 0755)
	xmlFileName := filepath.Join("generated", filePath)

	xmlString := string(xml)
	xmlString = removeExtraXMLNS(xmlString)
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

	re = regexp.MustCompile(`<(FrSeq|ToSeq)>(\d+)</(FrSeq|ToSeq)>`)

	// Replace numeric values with zero-padded format (6 digits)
	xmlString = re.ReplaceAllStringFunc(xmlString, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 4 {
			num := parts[2] // Extract number as string
			return fmt.Sprintf("<%s>%06s</%s>", parts[1], num, parts[3])
		}
		return match
	})

	return os.WriteFile(xmlFileName, []byte(xmlString), 0644)
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="`+XMLINS+`"`, "")

	return cleanXML
}
