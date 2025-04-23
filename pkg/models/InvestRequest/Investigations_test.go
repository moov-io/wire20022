package InvestRequest

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestInvestigations_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario1_Step2_camt.110")
	genterated := filepath.Join("generated", "Investigations_Scenario1_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario2_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario2_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario2_Step2_camt.110")
	genterated := filepath.Join("generated", "Investigations_Scenario2_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario3_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario3_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario3_Step2_camt.110")
	genterated := filepath.Join("generated", "Investigations_Scenario3_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
