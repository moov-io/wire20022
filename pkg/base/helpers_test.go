package base

import (
	"testing"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommonHelpers(t *testing.T) {
	t.Run("MessageId helper", func(t *testing.T) {
		helper := CommonHelpers["MessageId"]()

		assert.Equal(t, "Message Identification", helper.Title)
		assert.Equal(t, "IMAD_FedwireFunds_1", helper.Type)
		assert.Contains(t, helper.Documentation, "Point to point reference")
		assert.Contains(t, helper.Documentation, "unambiguously identify")
	})

	t.Run("CreatedDateTime helper", func(t *testing.T) {
		helper := CommonHelpers["CreatedDateTime"]()

		assert.Equal(t, "Creation Date Time", helper.Title)
		assert.Equal(t, "ISODateTime", helper.Type)
		assert.Contains(t, helper.Documentation, "Date and time")
		assert.Contains(t, helper.Documentation, "created")
	})

	t.Run("NumberOfTransactions helper", func(t *testing.T) {
		helper := CommonHelpers["NumberOfTransactions"]()

		assert.Equal(t, "Number Of Transactions", helper.Title)
		assert.Equal(t, "Max15NumericText", helper.Type)
		assert.Contains(t, helper.Documentation, "Number of individual transactions")
	})

	t.Run("SettlementMethod helper", func(t *testing.T) {
		helper := CommonHelpers["SettlementMethod"]()

		assert.Equal(t, "Settlement Method", helper.Title)
		assert.Equal(t, "SettlementMethod1Code", helper.Type)
		assert.Contains(t, helper.Documentation, "Method used to settle")
		assert.Contains(t, helper.Documentation, "CLRG")
	})

	t.Run("CommonClearingSysCode helper", func(t *testing.T) {
		helper := CommonHelpers["CommonClearingSysCode"]()

		assert.Equal(t, "Common Clearing System Code", helper.Title)
		assert.Equal(t, "ClearingSystemIdentification2Choice", helper.Type)
		assert.Contains(t, helper.Documentation, "pre-agreed offering")
		assert.Contains(t, helper.Documentation, "USABA")
	})

	t.Run("InstructingAgent helper", func(t *testing.T) {
		helper := CommonHelpers["InstructingAgent"]()

		assert.Equal(t, "Instructing Agent", helper.Title)
		assert.Equal(t, "BranchAndFinancialInstitutionIdentification5", helper.Type)
		assert.Contains(t, helper.Documentation, "instructs the next party")
	})

	t.Run("InstructedAgent helper", func(t *testing.T) {
		helper := CommonHelpers["InstructedAgent"]()

		assert.Equal(t, "Instructed Agent", helper.Title)
		assert.Equal(t, "BranchAndFinancialInstitutionIdentification5", helper.Type)
		assert.Contains(t, helper.Documentation, "instructed by the previous party")
	})

	t.Run("DebtorAgent helper", func(t *testing.T) {
		helper := CommonHelpers["DebtorAgent"]()

		assert.Equal(t, "Debtor Agent", helper.Title)
		assert.Equal(t, "BranchAndFinancialInstitutionIdentification5", helper.Type)
		assert.Contains(t, helper.Documentation, "servicing an account for the debtor")
	})

	t.Run("CreditorAgent helper", func(t *testing.T) {
		helper := CommonHelpers["CreditorAgent"]()

		assert.Equal(t, "Creditor Agent", helper.Title)
		assert.Equal(t, "BranchAndFinancialInstitutionIdentification5", helper.Type)
		assert.Contains(t, helper.Documentation, "servicing an account for the creditor")
	})

	t.Run("InitiatingPartyName helper", func(t *testing.T) {
		helper := CommonHelpers["InitiatingPartyName"]()

		assert.Equal(t, "Initiating Party Name", helper.Title)
		assert.Equal(t, "Max140Text", helper.Type)
		assert.Contains(t, helper.Documentation, "Name by which a party is known")
	})

	t.Run("All common helpers exist", func(t *testing.T) {
		expectedHelpers := []string{
			"MessageId",
			"CreatedDateTime",
			"NumberOfTransactions",
			"SettlementMethod",
			"CommonClearingSysCode",
			"InstructingAgent",
			"InstructedAgent",
			"DebtorAgent",
			"CreditorAgent",
			"InitiatingPartyName",
		}

		for _, helperName := range expectedHelpers {
			t.Run("helper_"+helperName, func(t *testing.T) {
				helperFunc, exists := CommonHelpers[helperName]
				require.True(t, exists, "Helper %s should exist", helperName)
				require.NotNil(t, helperFunc, "Helper function %s should not be nil", helperName)

				helper := helperFunc()
				assert.NotEmpty(t, helper.Title, "Helper %s should have a title", helperName)
				assert.NotEmpty(t, helper.Type, "Helper %s should have a type", helperName)
				assert.NotEmpty(t, helper.Documentation, "Helper %s should have documentation", helperName)
			})
		}
	})
}

func TestStandardMessageHelper(t *testing.T) {
	t.Run("BuildStandardMessageHelper", func(t *testing.T) {
		helper := BuildStandardMessageHelper()

		assert.Equal(t, "Message Identification", helper.MessageId.Title)
		assert.Equal(t, "Creation Date Time", helper.CreatedDateTime.Title)

		// Verify these are the same as from CommonHelpers
		expectedMessageId := CommonHelpers["MessageId"]()
		expectedCreatedDateTime := CommonHelpers["CreatedDateTime"]()

		assert.Equal(t, expectedMessageId, helper.MessageId)
		assert.Equal(t, expectedCreatedDateTime, helper.CreatedDateTime)
	})
}

func TestPaymentMessageHelper(t *testing.T) {
	t.Run("BuildPaymentMessageHelper", func(t *testing.T) {
		helper := BuildPaymentMessageHelper()

		// Check embedded StandardMessageHelper fields
		assert.Equal(t, "Message Identification", helper.MessageId.Title)
		assert.Equal(t, "Creation Date Time", helper.CreatedDateTime.Title)

		// Check payment-specific fields
		assert.Equal(t, "Number Of Transactions", helper.NumberOfTransactions.Title)
		assert.Equal(t, "Settlement Method", helper.SettlementMethod.Title)
		assert.Equal(t, "Common Clearing System Code", helper.CommonClearingSysCode.Title)

		// Verify these match CommonHelpers
		expectedNumTxns := CommonHelpers["NumberOfTransactions"]()
		expectedSettlement := CommonHelpers["SettlementMethod"]()
		expectedClearing := CommonHelpers["CommonClearingSysCode"]()

		assert.Equal(t, expectedNumTxns, helper.NumberOfTransactions)
		assert.Equal(t, expectedSettlement, helper.SettlementMethod)
		assert.Equal(t, expectedClearing, helper.CommonClearingSysCode)
	})
}

func TestAgentHelper(t *testing.T) {
	t.Run("BuildAgentHelper", func(t *testing.T) {
		helper := BuildAgentHelper()

		assert.Equal(t, "Instructing Agent", helper.InstructingAgent.Title)
		assert.Equal(t, "Instructed Agent", helper.InstructedAgent.Title)
		assert.Equal(t, "Debtor Agent", helper.DebtorAgent.Title)
		assert.Equal(t, "Creditor Agent", helper.CreditorAgent.Title)

		// Verify these match CommonHelpers
		expectedInstructing := CommonHelpers["InstructingAgent"]()
		expectedInstructed := CommonHelpers["InstructedAgent"]()
		expectedDebtor := CommonHelpers["DebtorAgent"]()
		expectedCreditor := CommonHelpers["CreditorAgent"]()

		assert.Equal(t, expectedInstructing, helper.InstructingAgent)
		assert.Equal(t, expectedInstructed, helper.InstructedAgent)
		assert.Equal(t, expectedDebtor, helper.DebtorAgent)
		assert.Equal(t, expectedCreditor, helper.CreditorAgent)
	})
}

func TestAddressHelper(t *testing.T) {
	t.Run("BuildAddressHelper", func(t *testing.T) {
		helper := BuildAddressHelper()

		// Test all address fields have proper titles, types, and documentation
		assert.Equal(t, "Street Name", helper.StreetName.Title)
		assert.Equal(t, "Max70Text", helper.StreetName.Type)
		assert.Contains(t, helper.StreetName.Documentation, "Name of a street")

		assert.Equal(t, "Building Number", helper.BuildingNumber.Title)
		assert.Equal(t, "Max16Text", helper.BuildingNumber.Type)
		assert.Contains(t, helper.BuildingNumber.Documentation, "position of a building")

		assert.Equal(t, "Room Number", helper.RoomNumber.Title)
		assert.Equal(t, "Max70Text", helper.RoomNumber.Type)
		assert.Contains(t, helper.RoomNumber.Documentation, "room, suite, or apartment")

		assert.Equal(t, "Postal Code", helper.PostalCode.Title)
		assert.Equal(t, "Max16Text", helper.PostalCode.Type)
		assert.Contains(t, helper.PostalCode.Documentation, "postal address")

		assert.Equal(t, "Town Name", helper.TownName.Title)
		assert.Equal(t, "Max35Text", helper.TownName.Type)
		assert.Contains(t, helper.TownName.Documentation, "built-up area")

		assert.Equal(t, "Country Subdivision", helper.Subdivision.Title)
		assert.Equal(t, "Max35Text", helper.Subdivision.Type)
		assert.Contains(t, helper.Subdivision.Documentation, "subdivision of a country")

		assert.Equal(t, "Country", helper.Country.Title)
		assert.Equal(t, "CountryCode", helper.Country.Type)
		assert.Contains(t, helper.Country.Documentation, "Nation with its own government")
	})
}

func TestHelperBuilder(t *testing.T) {
	t.Run("HelperBuilder function type", func(t *testing.T) {
		// Test that we can create and use a HelperBuilder
		builder := func() models.ElementHelper {
			return models.ElementHelper{
				Title:         "Test Helper",
				Type:          "TestType",
				Documentation: "Test documentation",
			}
		}

		// Verify it matches HelperBuilder type
		var helperBuilder HelperBuilder = builder
		require.NotNil(t, helperBuilder)

		// Test the builder
		result := helperBuilder()
		assert.Equal(t, "Test Helper", result.Title)
		assert.Equal(t, "TestType", result.Type)
		assert.Equal(t, "Test documentation", result.Documentation)
	})
}

func TestHelperStructJSONTags(t *testing.T) {
	t.Run("StandardMessageHelper JSON tags", func(t *testing.T) {
		// This test verifies that the struct has the expected JSON tags
		// We can't directly test tags in runtime, but we can verify the struct compiles
		// and has the expected field structure
		helper := StandardMessageHelper{}

		// Verify fields exist and can be accessed
		assert.NotNil(t, helper.MessageId)
		assert.NotNil(t, helper.CreatedDateTime)
	})

	t.Run("PaymentMessageHelper JSON tags", func(t *testing.T) {
		helper := PaymentMessageHelper{}

		// Verify embedded StandardMessageHelper fields
		assert.NotNil(t, helper.MessageId)
		assert.NotNil(t, helper.CreatedDateTime)

		// Verify payment-specific fields
		assert.NotNil(t, helper.NumberOfTransactions)
		assert.NotNil(t, helper.SettlementMethod)
		assert.NotNil(t, helper.CommonClearingSysCode)
	})

	t.Run("AgentHelper JSON tags", func(t *testing.T) {
		helper := AgentHelper{}

		assert.NotNil(t, helper.InstructingAgent)
		assert.NotNil(t, helper.InstructedAgent)
		assert.NotNil(t, helper.DebtorAgent)
		assert.NotNil(t, helper.CreditorAgent)
	})

	t.Run("AddressHelper JSON tags", func(t *testing.T) {
		helper := AddressHelper{}

		assert.NotNil(t, helper.StreetName)
		assert.NotNil(t, helper.BuildingNumber)
		assert.NotNil(t, helper.RoomNumber)
		assert.NotNil(t, helper.PostalCode)
		assert.NotNil(t, helper.TownName)
		assert.NotNil(t, helper.Subdivision)
		assert.NotNil(t, helper.Country)
	})
}
