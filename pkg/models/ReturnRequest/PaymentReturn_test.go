package ReturnRequest

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestRequireField(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	cErr := message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: AssignmentId, Assigner, Assignee, AssignmentCreateTime, CaseId, Creator, OriginalMessageId, OriginalMessageNameId, OriginalMessageCreateTime, OriginalUETR, OriginalInterbankSettlementAmount, OriginalInterbankSettlementDate, CancellationReason")
}
func generateRequreFields(msg Message) Message {
	if isEmpty(msg.Data.AssignmentId) {
		msg.Data.AssignmentId = "20250310B1QDRCQR000722"
	}
	if isEmpty(msg.Data.Assigner) {
		msg.Data.Assigner = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if isEmpty(msg.Data.Assignee) {
		msg.Data.Assignee = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	if msg.Data.AssignmentCreateTime.IsZero() {
		msg.Data.AssignmentCreateTime = time.Now()
	}
	if msg.Data.CaseId == "" {
		msg.Data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	}
	if isEmpty(msg.Data.Creator) {
		msg.Data.Creator = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if msg.Data.OriginalMessageId == "" {
		msg.Data.OriginalMessageId = "20250310B1QDRCQR000721"
	}
	if msg.Data.OriginalMessageNameId == "" {
		msg.Data.OriginalMessageNameId = "pacs.008.001.08"
	}
	if msg.Data.OriginalMessageCreateTime.IsZero() {
		msg.Data.OriginalMessageCreateTime = time.Now()
	}
	if msg.Data.OriginalUETR == "" {
		msg.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	}
	if isEmpty(msg.Data.OriginalInterbankSettlementAmount) {
		msg.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
			Amount:   151235.88,
			Currency: "USD",
		}
	}
	if isEmpty(msg.Data.OriginalInterbankSettlementDate) {
		msg.Data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	}
	if isEmpty(msg.Data.CancellationReason) {
		msg.Data.CancellationReason = Reason{
			Reason: "DUPL",
		}
	}
	return msg
}
func TestReturnRequestFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step2_camt.056")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250310B1QDRCQR000722", string(message.Doc.FIToFIPmtCxlReq.Assgnmt.Id))
	require.Equal(t, "011104238", string(message.Doc.FIToFIPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "USABA", string(*message.Doc.FIToFIPmtCxlReq.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "021040078", string(message.Doc.FIToFIPmtCxlReq.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "USABA", string(*message.Doc.FIToFIPmtCxlReq.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "20250310011104238Sc01Step1MsgIdDUPL", string(message.Doc.FIToFIPmtCxlReq.Case.Id))
	require.Equal(t, "011104238", string(message.Doc.FIToFIPmtCxlReq.Case.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "Bank A", string(*message.Doc.FIToFIPmtCxlReq.Case.Cretr.Agt.FinInstnId.Nm))
	require.Equal(t, "20250310B1QDRCQR000721", string(message.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlMsgId))
	require.Equal(t, "Scenario01InstrId001", string(*message.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlInstrId))
	require.Equal(t, "Scenario01EtoEId001", string(*message.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlEndToEndId))
	require.Equal(t, "Corporation A", string(*message.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf.Orgtr.Nm))
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"
const INVALID_TRCOUNT string = "123456789012345"
const INVALID_MESSAGE_ID string = "12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456"
const INVALID_OTHER_ID string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const INVALID_BUILD_NUM string = "12345678901234567"
const INVALID_POSTAL_CODE string = "12345678901234567"
const INVALID_COUNTRY_CODE string = "12345678"
const INVALID_MESSAGE_NAME_ID string = "sabcd-123-001-12"
const INVALID_PAY_SYSCODE model.PaymentSystemType = model.PaymentSystemType(INVALID_COUNT)

func TestReturnRequestValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"Invalid AssignmentId",
			Message{Data: MessageModel{AssignmentId: INVALID_OTHER_ID}},
			"error occur at AssignmentId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid Assigner",
			Message{Data: MessageModel{Assigner: model.Agent{
				PaymentSysCode:     INVALID_PAY_SYSCODE,
				PaymentSysMemberId: "011104238",
			}}},
			"error occur at Assigner.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid Assigner",
			Message{Data: MessageModel{Assignee: model.Agent{
				PaymentSysCode:     INVALID_PAY_SYSCODE,
				PaymentSysMemberId: "011104238",
			}}},
			"error occur at Assgne.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid CaseId",
			Message{Data: MessageModel{CaseId: INVALID_OTHER_ID}},
			"error occur at CaseId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid Creator",
			Message{Data: MessageModel{Creator: model.Agent{
				PaymentSysCode:     model.PaymentSysUSABA,
				PaymentSysMemberId: "011104238",
				BankName:           "Bank A",
				PostalAddress: model.PostalAddress{
					StreetName:     "Avenue A",
					BuildingNumber: "66",
					PostalCode:     INVALID_POSTAL_CODE,
					TownName:       "Lisle",
					Subdivision:    "IL",
					Country:        "US",
				},
			}}},
			"error occur at Creator.PostalAddress.PostalCode: 12345678901234567 fails validation with length 17 <= required maxLength 16",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			nMsg := generateRequreFields(tt.msg)
			msgErr := nMsg.CreateDocument()
			if msgErr != nil {
				require.Equal(t, tt.expectedErr, msgErr.Error())
			}
		})
	}
}
func TestFedwireFundsAcknowledgement_Scenario2_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.AssignmentId = "20250310B1QDRCQR000722"
	message.Data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.AssignmentCreateTime = time.Now()
	message.Data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.Data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.Data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.Data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.Data.CancellationReason = Reason{
		Originator:     "Corporation A",
		Reason:         "DUPL",
		AdditionalInfo: "Order cancelled. Ref:20250310B1QDRCQR000721",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step2_camt.056")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario2_Step2b_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.AssignmentId = "20250310B1QDRCQR000722"
	message.Data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.AssignmentCreateTime = time.Now()
	message.Data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.Data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.Data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.Data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.Data.CancellationReason = Reason{
		Originator:     "Corporation A",
		Reason:         "DUPL",
		AdditionalInfo: "Order cancelled. Ref:20250310B1QDRCQR000721",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step2b_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step2b_camt.056")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step2b_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario2_Step4_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.AssignmentId = "20250310B1QDRCQR000912"
	message.Data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.AssignmentCreateTime = time.Now()
	message.Data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.Data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.Data.OriginalMessageId = "20250310B1QDRCQR000002"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   510000.74,
		Currency: "USD",
	}
	message.Data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.Data.CancellationReason = Reason{
		Originator:     "Corporation A",
		Reason:         "DUPL",
		AdditionalInfo: "Payment is a duplicate. Please return payment.",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario2_Step4_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario2_Step4_camt.056")
	genterated := filepath.Join("generated", "Investigations_Scenario2_Step4_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.AssignmentId = "20250310B1QDRCQR000401"
	message.Data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.AssignmentCreateTime = time.Now()
	message.Data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.Data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.Data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1510000.74,
		Currency: "USD",
	}
	message.Data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.Data.CancellationReason = Reason{
		Originator:     "Corporation A",
		Reason:         "DUPL",
		AdditionalInfo: "Order cancelled. Ref:20250310B1QDRCQR000400",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario1_Step2_camt.056")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario1_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario2_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.AssignmentId = "20250310B1QDRCQR000421"
	message.Data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.AssignmentCreateTime = time.Now()
	message.Data.CaseId = "20250310011104238Sc02Step1MsgIdSVNR"
	message.Data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.Data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario02InstrId001"
	message.Data.OriginalEndToEndId = "Scenario02EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1234578.88,
		Currency: "USD",
	}
	message.Data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.Data.CancellationReason = Reason{
		Originator: "Corporation C",
		Reason:     "SVNR",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario2_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario2_Step2_camt.056")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario2_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario3_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.AssignmentId = "20250310B1QDRCQR000431"
	message.Data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.AssignmentCreateTime = time.Now()
	message.Data.CaseId = "20250310011104238Sc02Step1MsgIdSVNR"
	message.Data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.Data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario03InstrId001"
	message.Data.OriginalEndToEndId = "Scenario03EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   2234578.88,
		Currency: "USD",
	}
	message.Data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.Data.CancellationReason = Reason{
		Originator: "Corporation C",
		Reason:     "SVNR",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario3_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario3_Step2_camt.056")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario3_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario5_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.AssignmentId = "20250310B1QDRCQR000452"
	message.Data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.AssignmentCreateTime = time.Now()
	message.Data.CaseId = "20250310021307481Sc02Step1MsgIdCUST"
	message.Data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
		BankName:           "Corporation Z",
		PostalAddress: model.PostalAddress{
			StreetName: "Avenue Moliere 70",
			PostalCode: "1180",
			TownName:   "Brussels",
			Country:    "BE",
		},
	}
	message.Data.OriginalMessageId = "20250310B1QDRCQR000450"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario05InstrId001"
	message.Data.OriginalEndToEndId = "Scenario05EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   2234578.88,
		Currency: "USD",
	}
	message.Data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.Data.CancellationReason = Reason{
		Originator:     "Corporation Z",
		Reason:         "CUST",
		AdditionalInfo: "Goods ordered are on backfill.",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario5_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario5_Step2_camt.056")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario5_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
