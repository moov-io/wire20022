package InvestRequest

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
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, InvestigationType, UnderlyingData, Requestor, Responder")
}
func generateRequreFields(msg Message) Message {
	if msg.data.MessageId == "" {
		msg.data.MessageId = "20250310QMGFT015000901"
	}
	if msg.data.InvestigationType == "" {
		msg.data.InvestigationType = "UTAP"
	}
	if isEmpty(msg.data.UnderlyingData) {
		msg.data.UnderlyingData = Underlying{
			OriginalMessageId:        "20250310B1QDRCQR000001",
			OriginalMessageNameId:    "pacs.008.001.08",
			OriginalCreationDateTime: time.Now(),
			OriginalInstructionId:    "Scenario01InstrId001",
			OriginalEndToEndId:       "Scenario01EtoEId001",
			OriginalUETR:             "8a562c67-ca16-48ba-b074-65581be6f011",
			OriginalInterbankSettlementAmount: model.CurrencyAndAmount{
				Amount:   510000.74,
				Currency: "USD",
			},
			OriginalInterbankSettlementDate: model.FromTime(time.Now()),
		}
	}
	if isEmpty(msg.data.Requestor) {
		msg.data.Requestor = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	if isEmpty(msg.data.Responder) {
		msg.data.Responder = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	return msg
}
func TestInvestRequestFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "Investigations_Scenario1_Step2_camt.110")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250310QMGFT015000901", string(message.doc.InvstgtnReq.InvstgtnReq.MsgId))
	require.Equal(t, "20250310B1QDRCQR000001", string(message.doc.InvstgtnReq.InvstgtnReq.Undrlyg.IntrBk.OrgnlGrpInf.OrgnlMsgId))
	require.Equal(t, "pacs.008.001.08", string(message.doc.InvstgtnReq.InvstgtnReq.Undrlyg.IntrBk.OrgnlGrpInf.OrgnlMsgNmId))
	require.Equal(t, "Scenario01InstrId001", string(*message.doc.InvstgtnReq.InvstgtnReq.Undrlyg.IntrBk.OrgnlInstrId))
	require.Equal(t, "Scenario01EtoEId001", string(*message.doc.InvstgtnReq.InvstgtnReq.Undrlyg.IntrBk.OrgnlEndToEndId))
	require.Equal(t, "8a562c67-ca16-48ba-b074-65581be6f011", string(message.doc.InvstgtnReq.InvstgtnReq.Undrlyg.IntrBk.OrgnlUETR))
	require.Equal(t, "USABA", string(*message.doc.InvstgtnReq.InvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "011104238", string(message.doc.InvstgtnReq.InvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "IN14", string(*message.doc.InvstgtnReq.InvstgtnData[0].Rsn.Cd))
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

func TestInvestRequestValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"Invalid MessageId",
			Message{data: MessageModel{MessageId: INVALID_OTHER_ID}},
			"error occur at MessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid InvestigationType",
			Message{data: MessageModel{InvestigationType: INVALID_MESSAGE_NAME_ID}},
			"error occur at InvestigationType: sabcd-123-001-12 fails validation with length 16 <= required maxLength 4",
		},
		{
			"Invalid UnderlyingData - OriginalMessageId",
			Message{data: MessageModel{UnderlyingData: Underlying{
				OriginalMessageId:        INVALID_OTHER_ID,
				OriginalMessageNameId:    "pacs.008.001.08",
				OriginalCreationDateTime: time.Now(),
				OriginalInstructionId:    "Scenario01InstrId001",
				OriginalEndToEndId:       "Scenario01EtoEId001",
				OriginalUETR:             "8a562c67-ca16-48ba-b074-65581be6f011",
				OriginalInterbankSettlementAmount: model.CurrencyAndAmount{
					Amount:   510000.74,
					Currency: "USD",
				},
				OriginalInterbankSettlementDate: model.FromTime(time.Now()),
			}}},
			"error occur at UnderlyingData.OriginalMessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid Requestor",
			Message{data: MessageModel{Requestor: model.Agent{
				PaymentSysCode:     model.PaymentSysUSABA,
				PaymentSysMemberId: INVALID_MESSAGE_NAME_ID,
			}}},
			"error occur at Requestor.PaymentSysMemberId: sabcd-123-001-12 fails validation with pattern [0-9]{9,9}",
		},
		{
			"Invalid Responder",
			Message{data: MessageModel{Responder: model.Agent{
				PaymentSysCode:     model.PaymentSysUSABA,
				PaymentSysMemberId: INVALID_MESSAGE_NAME_ID,
			}}},
			"error occur at Responder.PaymentSysMemberId: sabcd-123-001-12 fails validation with pattern [0-9]{9,9}",
		},
		{
			"Invalid InvestReason",
			Message{data: MessageModel{InvestReason: InvestigationReason{
				Reason: "IN14",
			}}},
			"error occur at InvestReason.Reason: sabcd-123-001-12 fails validation with pattern [0-9]{9,9}",
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
func TestInvestigations_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFT015000901"
	message.data.InvestigationType = "UTAP"
	message.data.UnderlyingData = Underlying{
		OriginalMessageId:        "20250310B1QDRCQR000001",
		OriginalMessageNameId:    "pacs.008.001.08",
		OriginalCreationDateTime: time.Now(),
		OriginalInstructionId:    "Scenario01InstrId001",
		OriginalEndToEndId:       "Scenario01EtoEId001",
		OriginalUETR:             "8a562c67-ca16-48ba-b074-65581be6f011",
		OriginalInterbankSettlementAmount: model.CurrencyAndAmount{
			Amount:   510000.74,
			Currency: "USD",
		},
		OriginalInterbankSettlementDate: model.FromTime(time.Now()),
	}
	message.data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.InvestReason = InvestigationReason{
		Reason: "IN14",
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario1_Step2_camt.110")
	genterated := filepath.Join("generated", "Investigations_Scenario1_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario2_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFT015000902"
	message.data.InvestigationType = "OTHR"
	message.data.UnderlyingData = Underlying{
		OriginalMessageId:        "20250310B1QDRCQR000002",
		OriginalMessageNameId:    "pacs.008.001.08",
		OriginalCreationDateTime: time.Now(),
		OriginalInstructionId:    "Scenario01InstrId001",
		OriginalEndToEndId:       "Scenario01EtoEId001",
		OriginalUETR:             "8a562c67-ca16-48ba-b074-65581be6f011",
		OriginalInterbankSettlementAmount: model.CurrencyAndAmount{
			Amount:   510000.74,
			Currency: "USD",
		},
		OriginalInterbankSettlementDate: model.FromTime(time.Now()),
	}
	message.data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.InvestReason = InvestigationReason{
		Reason:                "PDUP",
		AdditionalRequestData: "Payment seems duplicate from previously received payment with IMAD 20250310B1QDRCQR000001.",
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario2_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario2_Step2_camt.110")
	genterated := filepath.Join("generated", "Investigations_Scenario2_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario3_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFT015000903"
	message.data.InvestigationType = "RQFI"
	message.data.UnderlyingData = Underlying{
		OriginalMessageId:        "20250310B1QDRCQR000007",
		OriginalMessageNameId:    "pacs.008.001.08",
		OriginalCreationDateTime: time.Now(),
		OriginalInstructionId:    "Scenario01InstrId001",
		OriginalEndToEndId:       "Scenario01EtoEId001",
		OriginalUETR:             "8a562c67-ca16-48ba-b074-65581be6f011",
		OriginalInterbankSettlementAmount: model.CurrencyAndAmount{
			Amount:   510000.74,
			Currency: "USD",
		},
		OriginalInterbankSettlementDate: model.FromTime(time.Now()),
	}
	message.data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.InvestReason = InvestigationReason{
		Reason: "MS01",
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario3_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario3_Step2_camt.110")
	genterated := filepath.Join("generated", "Investigations_Scenario3_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
