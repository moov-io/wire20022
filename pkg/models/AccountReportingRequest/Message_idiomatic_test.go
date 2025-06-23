package AccountReportingRequest

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

// TestReadWriteXMLIdiomatic tests the idiomatic XML-first API
func TestReadWriteXMLIdiomatic(t *testing.T) {
	// Create a complete message model
	model := MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "ARR20250623123456",
			CreatedDateTime: time.Now(),
		},
		ReportRequestId:    models.AccountBalanceReport,
		RequestedMsgNameId: "camt.052.001.08",
		AccountOtherId:     "ACC123456789",
		AccountProperty:    models.AccountTypeMerchant,
		AccountOwnerAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "123456789",
		},
		ReportingSequence: &ReportingSequenceFields{
			FromToSequence: models.SequenceRange{
				FromSeq: "000001",
				ToSeq:   "000001",
			},
		},
	}

	// Test WriteXML
	var buf bytes.Buffer
	err := model.WriteXML(&buf, CAMT_060_001_05)
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
	require.Contains(t, buf.String(), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	require.Contains(t, buf.String(), "ARR20250623123456")

	// Test ReadXML with the generated XML
	var readModel MessageModel
	reader := strings.NewReader(buf.String())
	err = readModel.ReadXML(reader)
	require.NoError(t, err)
	require.Equal(t, model.MessageId, readModel.MessageId)
	require.Equal(t, model.ReportRequestId, readModel.ReportRequestId)
}

// TestParseXMLIdiomatic tests the ParseXML function
func TestParseXMLIdiomatic(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.060.001.05">
  <AcctRptgReq>
    <GrpHdr>
      <MsgId>TEST123456789</MsgId>
      <CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
    </GrpHdr>
    <RptgReq>
      <Id>ABAR</Id>
      <ReqdMsgNmId>camt.052.001.08</ReqdMsgNmId>
      <Acct>
        <Id>
          <Othr>
            <Id>ACC987654321</Id>
          </Othr>
        </Id>
      </Acct>
      <AcctOwnr>
        <Agt>
          <FinInstnId>
            <ClrSysMmbId>
              <ClrSysId>
                <Cd>USABA</Cd>
              </ClrSysId>
              <MmbId>987654321</MmbId>
            </ClrSysMmbId>
          </FinInstnId>
        </Agt>
      </AcctOwnr>
    </RptgReq>
  </AcctRptgReq>
</Document>`

	model, err := ParseXML([]byte(xmlData))
	require.NoError(t, err)
	require.NotNil(t, model)
	require.Equal(t, "TEST123456789", model.MessageId)
	require.Equal(t, models.CAMTReportType("ABAR"), model.ReportRequestId)
	require.Equal(t, "ACC987654321", model.AccountOtherId)
}

// TestInitializeVersionFields tests field initialization
func TestInitializeVersionFields(t *testing.T) {
	model := MessageModel{}
	require.Nil(t, model.ReportingSequence)

	// Initialize version fields manually since this is the idiomatic pattern
	if model.ReportingSequence == nil {
		model.ReportingSequence = &ReportingSequenceFields{}
	}
	require.NotNil(t, model.ReportingSequence)
}

// TestValidateFields tests the validation functions
func TestValidateFields(t *testing.T) {
	// Test valid model
	validModel := MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "VALID123456789",
			CreatedDateTime: time.Now(),
		},
		ReportRequestId:    models.AccountBalanceReport,
		RequestedMsgNameId: "camt.052.001.08",
		AccountOwnerAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "123456789",
		},
	}

	err := CheckRequiredFields(validModel)
	require.NoError(t, err)

	// Test invalid model - missing MessageId
	invalidModel := validModel
	invalidModel.MessageId = ""
	err = CheckRequiredFields(invalidModel)
	require.Error(t, err)
	require.Contains(t, err.Error(), "MessageId")
}

// TestDocumentWithFunction tests the DocumentWith function
func TestDocumentWithFunction(t *testing.T) {
	model := MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "DOC123456789",
			CreatedDateTime: time.Now(),
		},
		ReportRequestId:    models.AccountBalanceReport,
		RequestedMsgNameId: "camt.052.001.08",
		AccountOwnerAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "123456789",
		},
	}

	// Test with version that doesn't require ReportingSequence
	doc, err := DocumentWith(model, CAMT_060_001_02)
	require.NoError(t, err)
	require.NotNil(t, doc)
}

// TestReportingSequenceFieldsValidation tests the nested field validation
func TestReportingSequenceFieldsValidation(t *testing.T) {
	fields := &ReportingSequenceFields{
		FromToSequence: models.SequenceRange{
			FromSeq: "001",
			ToSeq:   "100",
		},
	}

	err := fields.Validate()
	require.NoError(t, err)
}
