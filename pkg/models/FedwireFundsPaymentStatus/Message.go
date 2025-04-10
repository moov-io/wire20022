package FedwireFundsPaymentStatus

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	pacs002 "github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus_pacs_002_001_10"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10"

type MessageModel struct {
	//Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	MessageId string
	//Date and time at which the message was created.
	CreatedDateTime time.Time
	//Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageId string
	//Specifies the original message name identifier to which the message refers.
	OriginalMessageNameId string
	//Date and time at which the original message was created.
	OriginalMessageCreateTime time.Time
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUETR string
	//Specifies the status of a transaction, in a coded form.
	TransactionStatus model.TransactionStatusCode
	//Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime time.Time
	//Date and time at which a transaction is completed and cleared, that is, payment is effected.
	EffectiveInterbankSettlementDate model.Date
	//Provides detailed information on the status reason.
	StatusReasonInformation string
	//Further details on the status reason.
	ReasonAdditionalInfo string
	//Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent model.Agent
	//Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent model.Agent
}
type Message struct {
	data MessageModel
	doc  pacs002.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = pacs002.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10",
			Local: "Document",
		},
	}
	var FIToFIPmtStsRpt pacs002.FIToFIPaymentStatusReportV10
	var GrpHdr pacs002.GroupHeader911
	if msg.data.MessageId != "" {
		GrpHdr.MsgId = pacs002.Max35Text(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
	if !isEmpty(GrpHdr) {
		FIToFIPmtStsRpt.GrpHdr = GrpHdr
	}
	var TxInfAndSts pacs002.PaymentTransaction1101
	var OrgnlGrpInf pacs002.OriginalGroupInformation291
	if msg.data.OriginalMessageId != "" {
		OrgnlGrpInf.OrgnlMsgId = pacs002.IMADFedwireFunds1(msg.data.OriginalMessageId)
	}
	if msg.data.OriginalMessageNameId != "" {
		OrgnlGrpInf.OrgnlMsgNmId = pacs002.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
	}
	if !isEmpty(msg.data.OriginalMessageCreateTime) {
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(msg.data.OriginalMessageCreateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		TxInfAndSts.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.data.OriginalUETR != "" {
		TxInfAndSts.OrgnlUETR = pacs002.UUIDv4Identifier(msg.data.OriginalUETR)
	}
	if msg.data.TransactionStatus != "" {
		TxInfAndSts.TxSts = pacs002.ExternalPaymentTransactionStatus1Code(msg.data.TransactionStatus)
	}
	if !isEmpty(msg.data.AcceptanceDateTime) {
		AccptncDtTm := fedwire.ISODateTime(msg.data.AcceptanceDateTime)
		TxInfAndSts.AccptncDtTm = &AccptncDtTm
	}
	if !isEmpty(msg.data.EffectiveInterbankSettlementDate) {
		Dt := msg.data.EffectiveInterbankSettlementDate.Date()
		FctvIntrBkSttlmDt := pacs002.DateAndDateTime2Choice1{
			Dt: &Dt,
		}
		TxInfAndSts.FctvIntrBkSttlmDt = &FctvIntrBkSttlmDt
	}
	if msg.data.StatusReasonInformation != "" {
		Prtry := pacs002.Max35Text(msg.data.StatusReasonInformation)
		var StsRsnInf []*pacs002.StatusReasonInformation121
		reson := pacs002.StatusReasonInformation121{
			Rsn: pacs002.StatusReason6Choice1{
				Prtry: &Prtry,
			},
		}
		if msg.data.ReasonAdditionalInfo != "" {
			reson.AddtlInf = pacs002.Max105Text(msg.data.ReasonAdditionalInfo)
		}
		StsRsnInf = append(StsRsnInf, &reson)
		if !isEmpty(StsRsnInf) {
			TxInfAndSts.StsRsnInf = StsRsnInf
		}
	}

	if !isEmpty(msg.data.InstructingAgent) {
		Cd := pacs002.ExternalClearingSystemIdentification1CodeFixed(msg.data.InstructingAgent.PaymentSysCode)
		TxInfAndSts.InstgAgt = pacs002.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pacs002.FinancialInstitutionIdentification181{
				ClrSysMmbId: pacs002.ClearingSystemMemberIdentification21{
					ClrSysId: pacs002.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pacs002.Max35TextFixed(msg.data.InstructingAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(msg.data.InstructedAgent) {
		Cd := pacs002.ExternalClearingSystemIdentification1CodeFixed(msg.data.InstructedAgent.PaymentSysCode)
		TxInfAndSts.InstdAgt = pacs002.BranchAndFinancialInstitutionIdentification62{
			FinInstnId: pacs002.FinancialInstitutionIdentification182{
				ClrSysMmbId: pacs002.ClearingSystemMemberIdentification22{
					ClrSysId: pacs002.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pacs002.RoutingNumberFRS1(msg.data.InstructedAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(TxInfAndSts) {
		FIToFIPmtStsRpt.TxInfAndSts = TxInfAndSts
	}
	if !isEmpty(FIToFIPmtStsRpt) {
		msg.doc.FIToFIPmtStsRpt = FIToFIPmtStsRpt
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
