package FedwireFundsPaymentStatus_002_001_10

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	FedwireFundsPaymentStatus "github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus_pacs_002_001_10"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Pacs002 struct {
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
	EffectiveInterbankSettlementDate time.Time
	//Provides detailed information on the status reason.
	StatusReasonInformation string
	//Further details on the status reason.
	ReasonAdditionalInfo string
	//Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent model.Agent
	//Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent model.Agent
}
type Pacs002Message struct {
	model Pacs002
	doc   FedwireFundsPaymentStatus.Document
}

func NewPacs002Message() Pacs002Message {
	return Pacs002Message{
		model: Pacs002{},
	}
}
func (msg *Pacs002Message) CreateDocument() {
	msg.doc = FedwireFundsPaymentStatus.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10",
			Local: "Document",
		},
	}
	var FIToFIPmtStsRpt FedwireFundsPaymentStatus.FIToFIPaymentStatusReportV10
	var GrpHdr FedwireFundsPaymentStatus.GroupHeader911
	if msg.model.MessageId != "" {
		GrpHdr.MsgId = FedwireFundsPaymentStatus.Max35Text(msg.model.MessageId)
	}
	if !isEmpty(msg.model.CreatedDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.model.CreatedDateTime)
	}
	if !isEmpty(GrpHdr) {
		FIToFIPmtStsRpt.GrpHdr = GrpHdr
	}
	var TxInfAndSts FedwireFundsPaymentStatus.PaymentTransaction1101
	var OrgnlGrpInf FedwireFundsPaymentStatus.OriginalGroupInformation291
	if msg.model.OriginalMessageId != "" {
		OrgnlGrpInf.OrgnlMsgId = FedwireFundsPaymentStatus.IMADFedwireFunds1(msg.model.OriginalMessageId)
	}
	if msg.model.OriginalMessageNameId != "" {
		OrgnlGrpInf.OrgnlMsgNmId = FedwireFundsPaymentStatus.MessageNameIdentificationFRS1(msg.model.OriginalMessageNameId)
	}
	if !isEmpty(msg.model.OriginalMessageCreateTime) {
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(msg.model.OriginalMessageCreateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		TxInfAndSts.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.model.OriginalUETR != "" {
		TxInfAndSts.OrgnlUETR = FedwireFundsPaymentStatus.UUIDv4Identifier(msg.model.OriginalUETR)
	}
	if msg.model.TransactionStatus != "" {
		TxInfAndSts.TxSts = FedwireFundsPaymentStatus.ExternalPaymentTransactionStatus1Code(msg.model.TransactionStatus)
	}
	if !isEmpty(msg.model.AcceptanceDateTime) {
		AccptncDtTm := fedwire.ISODateTime(msg.model.AcceptanceDateTime)
		TxInfAndSts.AccptncDtTm = &AccptncDtTm
	}
	if !isEmpty(msg.model.EffectiveInterbankSettlementDate) {
		Dt := fedwire.ISODate(civil.DateOf(msg.model.EffectiveInterbankSettlementDate))
		FctvIntrBkSttlmDt := FedwireFundsPaymentStatus.DateAndDateTime2Choice1{
			Dt: &Dt,
		}
		TxInfAndSts.FctvIntrBkSttlmDt = &FctvIntrBkSttlmDt
	}
	if msg.model.StatusReasonInformation != "" {
		Prtry := FedwireFundsPaymentStatus.Max35Text(msg.model.StatusReasonInformation)
		var StsRsnInf []*FedwireFundsPaymentStatus.StatusReasonInformation121
		reson := FedwireFundsPaymentStatus.StatusReasonInformation121{
			Rsn: FedwireFundsPaymentStatus.StatusReason6Choice1{
				Prtry: &Prtry,
			},
		}
		if msg.model.ReasonAdditionalInfo != "" {
			reson.AddtlInf = FedwireFundsPaymentStatus.Max105Text(msg.model.ReasonAdditionalInfo)
		}
		StsRsnInf = append(StsRsnInf, &reson)
		if !isEmpty(StsRsnInf) {
			TxInfAndSts.StsRsnInf = StsRsnInf
		}
	}

	if !isEmpty(msg.model.InstructingAgent) {
		Cd := FedwireFundsPaymentStatus.ExternalClearingSystemIdentification1CodeFixed(msg.model.InstructingAgent.PaymentSysCode)
		TxInfAndSts.InstgAgt = FedwireFundsPaymentStatus.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: FedwireFundsPaymentStatus.FinancialInstitutionIdentification181{
				ClrSysMmbId: FedwireFundsPaymentStatus.ClearingSystemMemberIdentification21{
					ClrSysId: FedwireFundsPaymentStatus.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: FedwireFundsPaymentStatus.Max35TextFixed(msg.model.InstructingAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(msg.model.InstructedAgent) {
		Cd := FedwireFundsPaymentStatus.ExternalClearingSystemIdentification1CodeFixed(msg.model.InstructedAgent.PaymentSysCode)
		TxInfAndSts.InstdAgt = FedwireFundsPaymentStatus.BranchAndFinancialInstitutionIdentification62{
			FinInstnId: FedwireFundsPaymentStatus.FinancialInstitutionIdentification182{
				ClrSysMmbId: FedwireFundsPaymentStatus.ClearingSystemMemberIdentification22{
					ClrSysId: FedwireFundsPaymentStatus.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: FedwireFundsPaymentStatus.RoutingNumberFRS1(msg.model.InstructedAgent.PaymentSysMemberId),
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
func (msg *Pacs002Message) GetXML() ([]byte, error) {
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

	// Convert back to []byte
	return []byte(xmlString), nil
	// return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *Pacs002Message) GetJson() ([]byte, error) {
	return json.MarshalIndent(msg.doc.FIToFIPmtStsRpt, "", " ")
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10"`, "")

	return cleanXML
}
