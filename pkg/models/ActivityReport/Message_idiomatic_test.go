package ActivityReport

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
			CreatedDateTime: time.Now(),
		},
		MessageId: models.CAMTReportType("AR1234567890123456789012"),
		Pagenation: models.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		},
		ReportId:             "RPT123",
		ReportCreateDateTime: time.Now(),
		TotalEntries:         "5",
		TotalCreditEntries: models.NumberAndSumOfTransactions{
			NumberOfEntries: "2",
			Sum:             1000.00,
		},
		TotalDebitEntries: models.NumberAndSumOfTransactions{
			NumberOfEntries: "3",
			Sum:             1500.00,
		},
		AccountEnhancement: &AccountEnhancementFields{
			AccountOtherId: "ACC123456789",
		},
	}

	// Test WriteXML
	var buf bytes.Buffer
	err := model.WriteXML(&buf, CAMT_052_001_05)
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
	require.Contains(t, buf.String(), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	require.Contains(t, buf.String(), "AR1234567890123456789012")

	// Test ReadXML with the generated XML
	var readModel MessageModel
	reader := strings.NewReader(buf.String())
	err = readModel.ReadXML(reader)
	require.NoError(t, err)
	require.Equal(t, model.MessageId, readModel.MessageId)
	require.Equal(t, model.ReportId, readModel.ReportId)
}

// TestParseXMLIdiomatic tests the ParseXML function
func TestParseXMLIdiomatic(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.05">
  <BkToCstmrAcctRpt>
    <GrpHdr>
      <MsgId>TEST123456789</MsgId>
      <CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
      <MsgPgntn>
        <PgNb>1</PgNb>
        <LastPgInd>true</LastPgInd>
      </MsgPgntn>
    </GrpHdr>
    <Rpt>
      <Id>RPT001</Id>
      <CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
      <Bal>
        <Tp>
          <CdOrPrtry>
            <Cd>OPBD</Cd>
          </CdOrPrtry>
        </Tp>
        <CdtDbtInd>CRDT</CdtDbtInd>
        <Amt Ccy="USD">1000.00</Amt>
      </Bal>
    </Rpt>
  </BkToCstmrAcctRpt>
</Document>`

	model, err := ParseXML([]byte(xmlData))
	require.NoError(t, err)
	require.NotNil(t, model)
	require.Equal(t, models.CAMTReportType("TEST123456789"), model.MessageId)
	require.Equal(t, models.ReportType("RPT001"), model.ReportId)
}

// TestInitializeVersionFields tests field initialization
func TestInitializeVersionFields(t *testing.T) {
	model := MessageModel{}
	require.Nil(t, model.AccountEnhancement)

	model.InitializeVersionFields()
	require.NotNil(t, model.AccountEnhancement)
}

// TestValidateFields tests the validation functions
func TestValidateFields(t *testing.T) {
	// Test valid model
	validModel := MessageModel{
		MessageHeader: base.MessageHeader{
			CreatedDateTime: time.Now(),
		},
		MessageId: models.CAMTReportType("VALID123456789"),
		Pagenation: models.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		},
		ReportId:             "RPT123",
		ReportCreateDateTime: time.Now(),
	}

	err := CheckRequiredFields(validModel)
	require.NoError(t, err)

	// Test invalid model - missing MessageId
	invalidModel := validModel
	invalidModel.MessageId = models.CAMTReportType("")
	err = CheckRequiredFields(invalidModel)
	require.Error(t, err)
	require.Contains(t, err.Error(), "MessageId")
}

// TestDocumentWithFunction tests the DocumentWith function
func TestDocumentWithFunction(t *testing.T) {
	model := MessageModel{
		MessageHeader: base.MessageHeader{
			CreatedDateTime: time.Now(),
		},
		MessageId: models.CAMTReportType("DOC123456789"),
		Pagenation: models.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		},
		ReportId:             "RPT123",
		ReportCreateDateTime: time.Now(),
	}

	// Initialize version fields to avoid access issues
	model.InitializeVersionFields()

	doc, err := DocumentWith(model, CAMT_052_001_05)
	require.NoError(t, err)
	require.NotNil(t, doc)
}

// TestAccountEnhancementFieldsValidation tests the nested field validation
func TestAccountEnhancementFieldsValidation(t *testing.T) {
	fields := &AccountEnhancementFields{
		AccountOtherId: "ACC123456789",
	}

	err := fields.Validate()
	require.NoError(t, err)
}

// TestUnmarshalJSONIdiomatic tests JSON unmarshaling
func TestUnmarshalJSONIdiomatic(t *testing.T) {
	jsonData := `{
		"messageId": "JSON123456789",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"pagenation": {
			"pageNumber": "1",
			"lastPageNumber": "1"
		},
		"reportId": "RPT001",
		"reportCreateDateTime": "2024-01-01T10:00:00Z",
		"totalEntries": "5",
		"accountOtherId": "ACC987654321"
	}`

	var model MessageModel
	err := model.UnmarshalJSON([]byte(jsonData))
	require.NoError(t, err)
	require.Equal(t, models.CAMTReportType("JSON123456789"), model.MessageId)
	require.NotNil(t, model.AccountEnhancement)
}
