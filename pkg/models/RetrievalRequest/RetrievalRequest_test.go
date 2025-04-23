package RetrievalRequest

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestMessageRetrieval_Scenario1_Step1_admi_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250301011104238MRSc1Step1MsgId"
	message.data.CreatedDateTime = time.Now()
	message.data.RequestType = model.RequestSent
	message.data.BusinessDate = model.FromTime(time.Now())
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.FileReference = "20250310B1QDRCQR000001"
	message.data.RecipientId = "B1QDRCQR"
	message.data.RecipientIssuer = "NA"

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("MessageRetrieval_Scenario1_Step1_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "MessageRetrieval_Scenario1_Step1_admi.006")
	genterated := filepath.Join("generated", "MessageRetrieval_Scenario1_Step1_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestMessageRetrieval_Scenario2_Step1_admi_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250301011104238MRSc2Step1MsgId"
	message.data.CreatedDateTime = time.Now()
	message.data.RequestType = model.RequestSent
	message.data.BusinessDate = model.FromTime(time.Now())
	message.data.SequenceRange = model.SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000003",
	}
	message.data.RecipientId = "B1QDRCQR"
	message.data.RecipientIssuer = "NA"

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("MessageRetrieval_Scenario2_Step1_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "MessageRetrieval_Scenario2_Step1_admi.006")
	genterated := filepath.Join("generated", "MessageRetrieval_Scenario2_Step1_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
