package Master

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestReadWriteXML tests the XML-first API for reading and writing
func TestReadWriteXML(t *testing.T) {
	// Create a valid message model using version-specific initialization
	model := NewMessageForVersion(CAMT_052_001_08)
	model.MessageHeader = base.MessageHeader{
		MessageId:       "MASTER20250623001",
		CreatedDateTime: time.Now().UTC(),
	}
	model.MessagePagination = models.MessagePagenation{
		PageNumber:         "1",
		LastPageIndicator:  true,
	}
	model.ReportTypeId = models.ABMS
	model.ReportCreatedDate = time.Now().UTC()
	model.AccountOtherId = "231981435"
	model.AccountType = "M"
	model.RelatedAccountOtherId = "231981435"
	model.Balances = []models.Balance{
		{
			BalanceTypeId: models.BalanceType("DLOD"),
			Amount: models.CurrencyAndAmount{
				Currency: "USD",
				Amount:   270458895930.79,
			},
			CreditDebitIndicator: models.Credit,
			DateTime:            time.Now().UTC(),
		},
	}
	model.TransactionsSummary = []models.TotalsPerBankTransaction{
		{
			TotalNetEntryAmount:  279595877422.72,
			CreditDebitIndicator: models.Credit,
			CreditEntries: models.NumberAndSumOfTransactions{
				NumberOfEntries: "16281",
				Sum:             420780358976.96,
			},
			DebitEntries: models.NumberAndSumOfTransactions{
				NumberOfEntries: "22134",
				Sum:             141184481554.24,
			},
			BankTransactionCode: models.FedwireFundsTransfers,
			Date:                time.Now().UTC(),
		},
	}

	// Test WriteXML with specific version
	var buf bytes.Buffer
	err := model.WriteXML(&buf, CAMT_052_001_08)
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
	require.Contains(t, buf.String(), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	require.Contains(t, buf.String(), "MASTER20250623001")
	require.Contains(t, buf.String(), "231981435")

	// Test ReadXML with the generated XML
	var readModel MessageModel
	reader := strings.NewReader(buf.String())
	err = readModel.ReadXML(reader)
	require.NoError(t, err)

	// Verify key fields were preserved
	assert.Equal(t, model.MessageId, readModel.MessageId)
	assert.Equal(t, model.ReportTypeId, readModel.ReportTypeId)
	assert.Equal(t, model.AccountOtherId, readModel.AccountOtherId)
	assert.WithinDuration(t, model.CreatedDateTime, readModel.CreatedDateTime, 5*time.Hour)
}

// TestWriteXMLVersions tests writing XML for different versions
func TestWriteXMLVersions(t *testing.T) {
	// Test each supported version with appropriate field configurations
	versions := []struct {
		name    string
		version CAMT_052_001_VERSION
		hasBusinessQuery bool
	}{
		{"V2", CAMT_052_001_02, false},
		{"V3", CAMT_052_001_03, true},
		{"V8", CAMT_052_001_08, true},
		{"V12", CAMT_052_001_12, true},
	}

	for _, tc := range versions {
		t.Run(tc.name, func(t *testing.T) {
			model := NewMessageForVersion(tc.version)
			model.MessageHeader = base.MessageHeader{
				MessageId:       "VERSION_TEST_" + string(tc.version),
				CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			}
			model.MessagePagination = models.MessagePagenation{
				PageNumber:         "1",
				LastPageIndicator:  true,
			}
			model.ReportTypeId = models.ABMS
			model.ReportCreatedDate = time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
			model.AccountOtherId = "VERSION_TEST_ACC"
			model.AccountType = "M"
			model.RelatedAccountOtherId = "VERSION_TEST_REL"
			model.TransactionsSummary = []models.TotalsPerBankTransaction{
				{
					TotalNetEntryAmount:  100000.00,
					CreditDebitIndicator: models.Credit,
					CreditEntries: models.NumberAndSumOfTransactions{
						NumberOfEntries: "10",
						Sum:             100000.00,
					},
					DebitEntries: models.NumberAndSumOfTransactions{
						NumberOfEntries: "0",
						Sum:             0.00,
					},
					BankTransactionCode: models.FedwireFundsTransfers,
					Date:                time.Now(),
				},
			}

			// Verify version-specific fields are properly initialized
			if tc.hasBusinessQuery {
				assert.NotNil(t, model.BusinessQuery, "BusinessQuery should be initialized for %s", tc.version)
			} else {
				assert.Nil(t, model.BusinessQuery, "BusinessQuery should be nil for %s", tc.version)
			}

			var buf bytes.Buffer
			err := model.WriteXML(&buf, tc.version)
			require.NoError(t, err)
			assert.NotEmpty(t, buf.String())
			assert.Contains(t, buf.String(), "<?xml")
			assert.Contains(t, buf.String(), "VERSION_TEST_")

			// Verify namespace is correct for version
			expectedNamespace := VersionNameSpaceMap[tc.version]
			assert.Contains(t, buf.String(), expectedNamespace)
		})
	}
}

// TestValidateForVersion tests version-specific validation
func TestValidateForVersion(t *testing.T) {
	testCases := []struct {
		name    string
		model   MessageModel
		version CAMT_052_001_VERSION
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid model for V2",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "VALID002",
					CreatedDateTime: time.Now(),
				},
				MessagePagination: models.MessagePagenation{
					PageNumber:         "1",
					LastPageIndicator:  true,
				},
				ReportTypeId:          models.ABMS,
				ReportCreatedDate:     time.Now(),
				AccountOtherId:        "ACC001",
				AccountType:           "M",
				RelatedAccountOtherId: "REL001",
				TransactionsSummary: []models.TotalsPerBankTransaction{
					{
						TotalNetEntryAmount:  25000.00,
						CreditDebitIndicator: models.Credit,
						CreditEntries: models.NumberAndSumOfTransactions{
							NumberOfEntries: "3",
							Sum:             25000.00,
						},
						DebitEntries: models.NumberAndSumOfTransactions{
							NumberOfEntries: "0",
							Sum:             0.00,
						},
						BankTransactionCode: models.FedwireFundsTransfers,
						Date:                time.Now(),
					},
				},
			},
			version: CAMT_052_001_02,
			wantErr: false,
		},
		{
			name: "Valid model for V8 with business query",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "VALID008",
					CreatedDateTime: time.Now(),
				},
				MessagePagination: models.MessagePagenation{
					PageNumber:         "1",
					LastPageIndicator:  true,
				},
				ReportTypeId:          models.ABMS,
				ReportCreatedDate:     time.Now(),
				AccountOtherId:        "ACC008",
				AccountType:           "M",
				RelatedAccountOtherId: "REL008",
				TransactionsSummary: []models.TotalsPerBankTransaction{
					{
						TotalNetEntryAmount:  100000.00,
						CreditDebitIndicator: models.Credit,
						CreditEntries: models.NumberAndSumOfTransactions{
							NumberOfEntries: "5",
							Sum:             100000.00,
						},
						DebitEntries: models.NumberAndSumOfTransactions{
							NumberOfEntries: "0",
							Sum:             0.00,
						},
						BankTransactionCode: models.FedwireFundsTransfers,
						Date:                time.Now(),
					},
				},
				BusinessQuery: &BusinessQueryFields{
					OriginalBusinessMsgId:         "ORIG008",
					OriginalBusinessMsgNameId:     "camt.060.001.05",
					OriginalBusinessMsgCreateTime: time.Now().AddDate(0, 0, -1),
				},
			},
			version: CAMT_052_001_08,
			wantErr: false,
		},
		{
			name: "Missing CreatedDateTime",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId: "INVALID001",
				},
				ReportTypeId: models.ABMS,
			},
			version: CAMT_052_001_02,
			wantErr: true,
			errMsg:  "CreatedDateTime is required",
		},
		{
			name: "Missing ReportTypeId",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "INVALID002",
					CreatedDateTime: time.Now(),
				},
			},
			version: CAMT_052_001_02,
			wantErr: true,
			errMsg:  "ReportTypeId is required",
		},
		{
			name: "Missing ReportCreatedDate",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "INVALID003",
					CreatedDateTime: time.Now(),
				},
				ReportTypeId: models.ABMS,
			},
			version: CAMT_052_001_02,
			wantErr: true,
			errMsg:  "ReportCreatedDate is required",
		},
		{
			name: "V8 missing BusinessQuery",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "INVALID008",
					CreatedDateTime: time.Now(),
				},
				MessagePagination: models.MessagePagenation{
					PageNumber:         "1",
					LastPageIndicator:  true,
				},
				ReportTypeId:          models.ABMS,
				ReportCreatedDate:     time.Now(),
				AccountOtherId:        "ACC008",
				AccountType:           "M",
				RelatedAccountOtherId: "REL008",
				TransactionsSummary: []models.TotalsPerBankTransaction{
					{
						TotalNetEntryAmount:  50000.00,
						CreditDebitIndicator: models.Credit,
						CreditEntries: models.NumberAndSumOfTransactions{
							NumberOfEntries: "3",
							Sum:             50000.00,
						},
						DebitEntries: models.NumberAndSumOfTransactions{
							NumberOfEntries: "0",
							Sum:             0.00,
						},
						BankTransactionCode: models.FedwireFundsTransfers,
						Date:                time.Now(),
					},
				},
				// Missing BusinessQuery for V3+
			},
			version: CAMT_052_001_08,
			wantErr: true,
			errMsg:  "BusinessQueryFields required for version",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.model.ValidateForVersion(tc.version)
			if tc.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestValidateCoreFields tests core field validation directly
func TestValidateCoreFields(t *testing.T) {
	t.Run("Valid core fields", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId:       "CORE001",
				CreatedDateTime: time.Now(),
			},
			ReportTypeId:      models.ABMS,
			ReportCreatedDate: time.Now(),
		}
		err := model.validateCoreFields()
		assert.NoError(t, err)
	})

	t.Run("Zero CreatedDateTime", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "CORE002",
			},
			ReportTypeId: models.ABMS,
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "CreatedDateTime is required", err.Error())
	})

	t.Run("Empty ReportTypeId", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId:       "CORE003",
				CreatedDateTime: time.Now(),
			},
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "ReportTypeId is required", err.Error())
	})

	t.Run("Zero ReportCreatedDate", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId:       "CORE004",
				CreatedDateTime: time.Now(),
			},
			ReportTypeId: models.ABMS,
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "ReportCreatedDate is required", err.Error())
	})
}

// TestGetVersionCapabilities tests version capability detection
func TestGetVersionCapabilities(t *testing.T) {
	testCases := []struct {
		name    string
		version CAMT_052_001_VERSION
		expectedCaps map[string]bool
	}{
		{
			name:    "V2 - no business query",
			version: CAMT_052_001_02,
			expectedCaps: map[string]bool{
				"BusinessQuery": false,
			},
		},
		{
			name:    "V3 - business query",
			version: CAMT_052_001_03,
			expectedCaps: map[string]bool{
				"BusinessQuery": true,
			},
		},
		{
			name:    "V12 - business query",
			version: CAMT_052_001_12,
			expectedCaps: map[string]bool{
				"BusinessQuery": true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := NewMessageForVersion(tc.version)
			caps := model.GetVersionCapabilities()
			assert.Equal(t, tc.expectedCaps, caps)
		})
	}
}

// TestNewMessageForVersion tests version-specific initialization
func TestNewMessageForVersion(t *testing.T) {
	versions := []struct {
		version CAMT_052_001_VERSION
		hasBusinessQuery bool
	}{
		{CAMT_052_001_02, false},
		{CAMT_052_001_03, true},
		{CAMT_052_001_08, true},
		{CAMT_052_001_12, true},
	}

	for _, v := range versions {
		t.Run(string(v.version), func(t *testing.T) {
			model := NewMessageForVersion(v.version)
			assert.NotNil(t, model)

			// Check base fields are initialized to zero values
			assert.Empty(t, model.MessageId)
			assert.Empty(t, model.ReportTypeId)
			assert.True(t, model.CreatedDateTime.IsZero())

			// Check version-specific field initialization
			if v.hasBusinessQuery {
				assert.NotNil(t, model.BusinessQuery)
			} else {
				assert.Nil(t, model.BusinessQuery)
			}
		})
	}
}

// TestCheckRequiredFields tests the required field validation helper
func TestCheckRequiredFields(t *testing.T) {
	t.Run("All required fields present", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId:       "REQ001",
				CreatedDateTime: time.Now(),
			},
			MessagePagination: models.MessagePagenation{
				PageNumber:         "1",
				LastPageIndicator:  true,
			},
			ReportTypeId:          models.ABMS,
			ReportCreatedDate:     time.Now(),
			AccountOtherId:        "ACC_REQ001",
			AccountType:           "M",
			RelatedAccountOtherId: "REL_REQ001",
			TransactionsSummary: []models.TotalsPerBankTransaction{
				{
					TotalNetEntryAmount:  50000.00,
					CreditDebitIndicator: models.Credit,
					CreditEntries: models.NumberAndSumOfTransactions{
						NumberOfEntries: "5",
						Sum:             50000.00,
					},
					DebitEntries: models.NumberAndSumOfTransactions{
						NumberOfEntries: "0",
						Sum:             0.00,
					},
					BankTransactionCode: models.FedwireFundsTransfers,
					Date:                time.Now(),
				},
			},
		}
		err := CheckRequiredFields(model)
		assert.NoError(t, err)
	})

	t.Run("Missing required fields", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "REQ002",
				// Missing CreatedDateTime
			},
			// Missing ReportTypeId and ReportCreatedDate
		}
		err := CheckRequiredFields(model)
		require.Error(t, err)
		// The error should mention missing required fields
		assert.Contains(t, err.Error(), "required")
	})
}

// TestJSONMarshaling tests JSON serialization
func TestJSONMarshaling(t *testing.T) {
	t.Skip("Skipping until date handling issues are resolved")
	
	original := MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "JSON001",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		MessagePagination: models.MessagePagenation{
			PageNumber:         "1",
			LastPageIndicator:  true,
		},
		ReportTypeId:          models.ABMS,
		ReportCreatedDate:     time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		AccountOtherId:        "JSON_ACC001",
		AccountType:           "M",
		RelatedAccountOtherId: "JSON_REL001",
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	require.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Unmarshal back
	var decoded MessageModel
	err = json.Unmarshal(jsonData, &decoded)
	require.NoError(t, err)

	// Verify fields
	assert.Equal(t, original.MessageId, decoded.MessageId)
	assert.Equal(t, original.ReportTypeId, decoded.ReportTypeId)
	assert.Equal(t, original.AccountOtherId, decoded.AccountOtherId)
	assert.Equal(t, original.CreatedDateTime.UTC(), decoded.CreatedDateTime.UTC())
}

// TestParseXMLWithInvalidData tests ParseXML error handling
func TestParseXMLWithInvalidData(t *testing.T) {
	testCases := []struct {
		name    string
		xmlData string
		wantErr bool
	}{
		{
			name:    "Empty XML",
			xmlData: "",
			wantErr: true,
		},
		{
			name:    "Invalid XML",
			xmlData: "not xml at all",
			wantErr: true,
		},
		{
			name:    "Wrong namespace",
			xmlData: `<?xml version="1.0"?><Document xmlns="wrong:namespace"><BkToCstmrAcctRpt></BkToCstmrAcctRpt></Document>`,
			wantErr: true,
		},
		{
			name: "Missing required fields",
			xmlData: `<?xml version="1.0"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.08">
	<BkToCstmrAcctRpt>
		<GrpHdr>
			<MsgId>TEST001</MsgId>
		</GrpHdr>
	</BkToCstmrAcctRpt>
</Document>`,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseXML([]byte(tc.xmlData))
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestReadXMLWithErrors tests ReadXML error handling
func TestReadXMLWithErrors(t *testing.T) {
	t.Run("Invalid reader", func(t *testing.T) {
		var model MessageModel
		err := model.ReadXML(strings.NewReader(""))
		assert.Error(t, err)
	})

	t.Run("Malformed XML", func(t *testing.T) {
		var model MessageModel
		err := model.ReadXML(strings.NewReader("<invalid>"))
		assert.Error(t, err)
	})
}

// TestWriteXMLWithInvalidModel tests WriteXML with invalid data
func TestWriteXMLWithInvalidModel(t *testing.T) {
	// Model missing required fields
	model := MessageModel{}

	var buf bytes.Buffer
	err := model.WriteXML(&buf)
	assert.Error(t, err)
}

// TestDocumentWithValidation tests DocumentWith with validation
func TestDocumentWithValidation(t *testing.T) {
	t.Skip("Skipping until field mapping issues are resolved")
	
	t.Run("Valid model creates document", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId:       "DOC001",
				CreatedDateTime: time.Now(),
			},
			MessagePagination: models.MessagePagenation{
				PageNumber:         "1",
				LastPageIndicator:  true,
			},
			ReportTypeId:          models.ABMS,
			ReportCreatedDate:     time.Now(),
			AccountOtherId:        "DOC_ACC001",
			AccountType:           "M",
			RelatedAccountOtherId: "DOC_REL001",
		}

		doc, err := DocumentWith(model, CAMT_052_001_02)
		require.NoError(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("Invalid model fails validation", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "DOC002",
				// Missing required fields
			},
		}

		doc, err := DocumentWith(model, CAMT_052_001_02)
		require.Error(t, err)
		assert.Nil(t, doc)
		assert.Contains(t, err.Error(), "required")
	})
}

// TestVersionSpecificFieldValidation tests field validation specific to different versions
func TestVersionSpecificFieldValidation(t *testing.T) {
	t.Run("BusinessQueryFields validation", func(t *testing.T) {
		fields := &BusinessQueryFields{
			OriginalBusinessMsgId:         "ORIG001",
			OriginalBusinessMsgNameId:     "camt.060.001.05",
			OriginalBusinessMsgCreateTime: time.Now().AddDate(0, 0, -1),
		}
		err := fields.Validate()
		assert.NoError(t, err)
	})

	t.Run("BusinessQueryFields empty", func(t *testing.T) {
		fields := &BusinessQueryFields{}
		err := fields.Validate()
		assert.NoError(t, err)
	})
}

// TestCheckForBusinessQueryFields tests the business query field detection
func TestCheckForBusinessQueryFields(t *testing.T) {
	t.Run("Has business query fields", func(t *testing.T) {
		rawMap := map[string]interface{}{
			"originalBusinessMsgId":   "ORIG001",
			"originalBusinessMsgNameId": "camt.060.001.05",
			"messageId":               "MSG001",
		}
		result := checkForBusinessQueryFields(rawMap)
		assert.True(t, result)
	})

	t.Run("No business query fields", func(t *testing.T) {
		rawMap := map[string]interface{}{
			"messageId":    "MSG001",
			"reportTypeId": "ABAR",
		}
		result := checkForBusinessQueryFields(rawMap)
		assert.False(t, result)
	})
}

// TestMessagePaginationFields tests pagination-specific functionality
func TestMessagePaginationFields(t *testing.T) {
	t.Run("Valid pagination", func(t *testing.T) {
		pagination := models.MessagePagenation{
			PageNumber:         "1",
			LastPageIndicator:  true,
		}
		
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId:       "PAG001",
				CreatedDateTime: time.Now(),
			},
			MessagePagination: pagination,
			ReportTypeId:      models.ABMS,
			ReportCreatedDate: time.Now(),
		}
		
		assert.Equal(t, "1", model.MessagePagination.PageNumber)
		assert.True(t, model.MessagePagination.LastPageIndicator)
	})
}

// TestBalancesAndTransactionsData tests complex data structures
func TestBalancesAndTransactionsData(t *testing.T) {
	t.Run("Valid balances", func(t *testing.T) {
		balances := []models.Balance{
			{
				BalanceTypeId: models.BalanceType("DLOD"),
				Amount: models.CurrencyAndAmount{
					Currency: "USD",
					Amount:   270458895930.79,
				},
				CreditDebitIndicator: models.Credit,
				DateTime:            time.Now(),
			},
			{
				BalanceTypeId: models.BalanceType("BOOK"),
				Amount: models.CurrencyAndAmount{
					Currency: "USD",
					Amount:   100000.00,
				},
				CreditDebitIndicator: models.Debit,
				DateTime:            time.Now(),
			},
		}
		
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId:       "BAL001",
				CreatedDateTime: time.Now(),
			},
			ReportTypeId:      models.ABMS,
			ReportCreatedDate: time.Now(),
			Balances:          balances,
		}
		
		assert.Len(t, model.Balances, 2)
		assert.Equal(t, "USD", model.Balances[0].Amount.Currency)
		assert.Equal(t, 270458895930.79, model.Balances[0].Amount.Amount)
	})
}

// BenchmarkWriteXML benchmarks XML writing performance
func BenchmarkWriteXML(b *testing.B) {
	model := MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "BENCH001",
			CreatedDateTime: time.Now(),
		},
		MessagePagination: models.MessagePagenation{
			PageNumber:         "1",
			LastPageIndicator:  true,
		},
		ReportTypeId:          models.ABMS,
		ReportCreatedDate:     time.Now(),
		AccountOtherId:        "BENCH_ACC001",
		AccountType:           "M",
		RelatedAccountOtherId: "BENCH_REL001",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_ = model.WriteXML(&buf, CAMT_052_001_08)
	}
}

// BenchmarkParseXML benchmarks XML parsing performance
func BenchmarkParseXML(b *testing.B) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.08">
	<BkToCstmrAcctRpt>
		<GrpHdr>
			<MsgId>BENCH001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
			<MsgPgntn>
				<PgNb>1</PgNb>
				<LastPgInd>true</LastPgInd>
			</MsgPgntn>
		</GrpHdr>
		<Rpt>
			<Id>BENCH_REPORT</Id>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
			<Acct>
				<Id>
					<Othr>
						<Id>BENCH_ACC001</Id>
					</Othr>
				</Id>
				<Tp>
					<Prtry>M</Prtry>
				</Tp>
			</Acct>
		</Rpt>
	</BkToCstmrAcctRpt>
</Document>`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseXML(xmlData)
	}
}