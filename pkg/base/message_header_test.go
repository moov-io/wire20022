package base

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessageHeader(t *testing.T) {
	t.Run("MessageHeader creation", func(t *testing.T) {
		header := MessageHeader{
			MessageId:       "MSG001",
			CreatedDateTime: time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
		}

		assert.Equal(t, "MSG001", header.MessageId)
		assert.Equal(t, 2025, header.CreatedDateTime.Year())
	})

	t.Run("MessageHeader JSON serialization", func(t *testing.T) {
		header := MessageHeader{
			MessageId:       "MSG001",
			CreatedDateTime: time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
		}

		data, err := json.Marshal(header)
		require.NoError(t, err)

		var decoded MessageHeader
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, header.MessageId, decoded.MessageId)
		assert.True(t, header.CreatedDateTime.Equal(decoded.CreatedDateTime))
	})
}

func TestPaymentCore(t *testing.T) {
	t.Run("PaymentCore creation with embedded MessageHeader", func(t *testing.T) {
		payment := PaymentCore{
			MessageHeader: MessageHeader{
				MessageId:       "PAY001",
				CreatedDateTime: time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
			},
			NumberOfTransactions:  "1",
			SettlementMethod:      "CLRG",
			CommonClearingSysCode: "USABA",
		}

		// Test embedded field access
		assert.Equal(t, "PAY001", payment.MessageId)
		assert.Equal(t, "1", payment.NumberOfTransactions)
		assert.Equal(t, models.SettlementMethodType("CLRG"), payment.SettlementMethod)
		assert.Equal(t, models.CommonClearingSysCodeType("USABA"), payment.CommonClearingSysCode)
	})

	t.Run("PaymentCore JSON serialization with inline embedding", func(t *testing.T) {
		payment := PaymentCore{
			MessageHeader: MessageHeader{
				MessageId:       "PAY001",
				CreatedDateTime: time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
			},
			NumberOfTransactions:  "1",
			SettlementMethod:      "CLRG",
			CommonClearingSysCode: "USABA",
		}

		data, err := json.Marshal(payment)
		require.NoError(t, err)

		// Verify JSON contains inline fields
		jsonStr := string(data)
		assert.Contains(t, jsonStr, `"messageId":"PAY001"`)
		assert.Contains(t, jsonStr, `"numberOfTransactions":"1"`)
		assert.Contains(t, jsonStr, `"settlementMethod":"CLRG"`)

		var decoded PaymentCore
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, payment.MessageId, decoded.MessageId)
		assert.Equal(t, payment.NumberOfTransactions, decoded.NumberOfTransactions)
	})
}

func TestAgentPair(t *testing.T) {
	t.Run("AgentPair creation", func(t *testing.T) {
		agents := AgentPair{
			InstructingAgent: mockAgent("INST001"),
			InstructedAgent:  mockAgent("INST002"),
		}

		assert.NotNil(t, agents.InstructingAgent)
		assert.NotNil(t, agents.InstructedAgent)
	})

	t.Run("AgentPair JSON serialization", func(t *testing.T) {
		agents := AgentPair{
			InstructingAgent: mockAgent("INST001"),
			InstructedAgent:  mockAgent("INST002"),
		}

		data, err := json.Marshal(agents)
		require.NoError(t, err)

		var decoded AgentPair
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, agents.InstructingAgent, decoded.InstructingAgent)
		assert.Equal(t, agents.InstructedAgent, decoded.InstructedAgent)
	})
}

func TestDebtorCreditorPair(t *testing.T) {
	t.Run("DebtorCreditorPair creation", func(t *testing.T) {
		pair := DebtorCreditorPair{
			DebtorAgent:   mockAgent("DEBT001"),
			CreditorAgent: mockAgent("CRED001"),
		}

		assert.NotNil(t, pair.DebtorAgent)
		assert.NotNil(t, pair.CreditorAgent)
	})

	t.Run("DebtorCreditorPair JSON serialization", func(t *testing.T) {
		pair := DebtorCreditorPair{
			DebtorAgent:   mockAgent("DEBT001"),
			CreditorAgent: mockAgent("CRED001"),
		}

		data, err := json.Marshal(pair)
		require.NoError(t, err)

		var decoded DebtorCreditorPair
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, pair.DebtorAgent, decoded.DebtorAgent)
		assert.Equal(t, pair.CreditorAgent, decoded.CreditorAgent)
	})
}

func TestPartyAddress(t *testing.T) {
	t.Run("PartyAddress creation", func(t *testing.T) {
		address := PartyAddress{
			StreetName:     "123 Main St",
			BuildingNumber: "Suite 100",
			RoomNumber:     "Room 5",
			PostalCode:     "12345",
			TownName:       "Anytown",
			Subdivision:    "ST",
			Country:        "US",
		}

		assert.Equal(t, "123 Main St", address.StreetName)
		assert.Equal(t, "Suite 100", address.BuildingNumber)
		assert.Equal(t, "Room 5", address.RoomNumber)
		assert.Equal(t, "12345", address.PostalCode)
		assert.Equal(t, "Anytown", address.TownName)
		assert.Equal(t, "ST", address.Subdivision)
		assert.Equal(t, "US", address.Country)
	})

	t.Run("PartyAddress JSON serialization", func(t *testing.T) {
		address := PartyAddress{
			StreetName:     "123 Main St",
			BuildingNumber: "Suite 100",
			PostalCode:     "12345",
			TownName:       "Anytown",
			Country:        "US",
		}

		data, err := json.Marshal(address)
		require.NoError(t, err)

		var decoded PartyAddress
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, address, decoded)
	})
}

func TestParty(t *testing.T) {
	t.Run("Party creation with embedded address", func(t *testing.T) {
		party := Party{
			Name: "Test Corporation",
			Address: PartyAddress{
				StreetName: "123 Business Ave",
				TownName:   "Corporate City",
				Country:    "US",
			},
		}

		assert.Equal(t, "Test Corporation", party.Name)
		assert.Equal(t, "123 Business Ave", party.Address.StreetName)
		assert.Equal(t, "Corporate City", party.Address.TownName)
		assert.Equal(t, "US", party.Address.Country)
	})

	t.Run("Party JSON serialization", func(t *testing.T) {
		party := Party{
			Name: "Test Corporation",
			Address: PartyAddress{
				StreetName: "123 Business Ave",
				TownName:   "Corporate City",
				Country:    "US",
			},
		}

		data, err := json.Marshal(party)
		require.NoError(t, err)

		var decoded Party
		err = json.Unmarshal(data, &decoded)
		require.NoError(t, err)

		assert.Equal(t, party, decoded)
	})
}

func TestInitiatingParty(t *testing.T) {
	t.Run("InitiatingParty with embedded Party", func(t *testing.T) {
		initiator := InitiatingParty{
			Party: Party{
				Name: "Initiating Bank",
				Address: PartyAddress{
					StreetName: "456 Finance St",
					TownName:   "Banking City",
					Country:    "US",
				},
			},
		}

		// Test embedded field access
		assert.Equal(t, "Initiating Bank", initiator.Name)
		assert.Equal(t, "456 Finance St", initiator.Address.StreetName)
	})
}

// Helper function to create mock Agent for testing
func mockAgent(id string) models.Agent {
	return models.Agent{
		BusinessIdCode:     id,
		PaymentSysCode:     "USABA",
		PaymentSysMemberId: id + "_MEMBER",
		BankName:           "Test Bank " + id,
		PostalAddress: models.PostalAddress{
			StreetName: "123 Test St",
			TownName:   "Test City",
			Country:    "US",
		},
		OtherTypeId: id + "_OTHER",
	}
}
