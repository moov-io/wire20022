package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBuildPostalAddressHelper tests the postal address helper builder
func TestBuildPostalAddressHelper(t *testing.T) {
	helper := BuildPostalAddressHelper()

	// Test that all fields are properly initialized
	assert.Equal(t, "Street Name", helper.StreetName.Title)
	assert.Equal(t, "Building Number", helper.BuildingNumber.Title)
	assert.Equal(t, "Building Name", helper.BuildingName.Title)
	assert.Equal(t, "Floor", helper.Floor.Title)
	assert.Equal(t, "Room Number", helper.RoomNumber.Title)
	assert.Equal(t, "Postal Code", helper.PostalCode.Title)
	assert.Equal(t, "Town Name", helper.TownName.Title)
	assert.Equal(t, "Country Sub Division", helper.Subdivision.Title)
	assert.Equal(t, "Country", helper.Country.Title)

	// Test type information is present
	assert.Contains(t, helper.StreetName.Type, "Max70Text")
	assert.Contains(t, helper.BuildingNumber.Type, "Max16Text")
	assert.Contains(t, helper.PostalCode.Type, "Max16Text")

	// Test documentation is provided
	assert.Contains(t, helper.StreetName.Documentation, "street or thoroughfare")
	assert.Contains(t, helper.BuildingNumber.Documentation, "position of a building")
	assert.Contains(t, helper.Country.Documentation, "Nation")
}

// TestBuildAgentHelper tests the agent helper builder
func TestBuildAgentHelper(t *testing.T) {
	helper := BuildAgentHelper()

	// Test required fields are initialized
	assert.Equal(t, "Member Identification", helper.PaymentSysMemberId.Title)
	assert.Equal(t, "BICFI", helper.BusinessIdCode.Title)
	assert.Equal(t, "Name", helper.BankName.Title)
	assert.Equal(t, "Clearing System Identification Code", helper.PaymentSysCode.Title)
	assert.Equal(t, "Other Type Id", helper.OtherTypeId.Title)

	// Test type constraints
	assert.Contains(t, helper.PaymentSysMemberId.Type, "RoutingNumber_FRS_1")
	assert.Contains(t, helper.BusinessIdCode.Type, "BICFIDec2014Identifier")
	assert.Contains(t, helper.BankName.Type, "Max140Text")

	// Test documentation content
	assert.Contains(t, helper.PaymentSysMemberId.Documentation, "clearing system")
	assert.Contains(t, helper.BusinessIdCode.Documentation, "Business identifier code")
	assert.Contains(t, helper.BankName.Documentation, "agent")

	// Test postal address field exists (but is not initialized by BuildAgentHelper)
	// Note: PostalAddress is a zero-value struct in AgentHelper
	assert.Empty(t, helper.PostalAddress.StreetName.Title)
}

// TestBuildSequenceRangeHelper tests the sequence range helper builder
func TestBuildSequenceRangeHelper(t *testing.T) {
	helper := BuildSequenceRangeHelper()

	// Test field initialization
	assert.Equal(t, "From Sequence", helper.FromSeq.Title)
	assert.Equal(t, "To Sequence", helper.ToSeq.Title)

	// Test type information
	assert.Contains(t, helper.FromSeq.Type, "SequenceNumber_FedwireFunds_1")
	assert.Contains(t, helper.ToSeq.Type, "SequenceNumber_FedwireFunds_1")

	// Test documentation
	assert.Contains(t, helper.FromSeq.Documentation, "sequence")
	assert.Contains(t, helper.ToSeq.Documentation, "sequence")
}

// TestBuildMessagePagenationHelper tests the pagination helper builder
func TestBuildMessagePagenationHelper(t *testing.T) {
	helper := BuildMessagePagenationHelper()

	// Test field initialization
	assert.Equal(t, "Page Number", helper.PageNumber.Title)
	assert.Equal(t, "Last Page Indicator", helper.LastPageIndicator.Title)

	// Test type information
	assert.Contains(t, helper.PageNumber.Type, "Max5NumericText")
	assert.Contains(t, helper.LastPageIndicator.Type, "boolean")

	// Test documentation
	assert.Contains(t, helper.PageNumber.Documentation, "Page number")
	assert.Contains(t, helper.LastPageIndicator.Documentation, "last page")
}

// TestBuildNumberAndSumOfTransactionsHelper tests the transaction summary helper builder
func TestBuildNumberAndSumOfTransactionsHelper(t *testing.T) {
	helper := BuildNumberAndSumOfTransactionsHelper()

	// Test field initialization
	assert.Equal(t, "Number Of Entries", helper.NumberOfEntries.Title)
	assert.Equal(t, "Sum", helper.Sum.Title)

	// Test type information
	assert.Contains(t, helper.NumberOfEntries.Type, "Max15NumericText")
	assert.Contains(t, helper.Sum.Type, "DecimalNumber")

	// Test documentation
	assert.Contains(t, helper.NumberOfEntries.Documentation, "entries")
	assert.Contains(t, helper.Sum.Documentation, "Total of all individual entries")
}

// TestBuildCurrencyAndAmountHelper tests the currency amount helper builder
func TestBuildCurrencyAndAmountHelper(t *testing.T) {
	helper := BuildCurrencyAndAmountHelper()

	// Test field initialization
	assert.Equal(t, "Currency", helper.Currency.Title)
	assert.Equal(t, "Amount", helper.Amount.Title)

	// Test type information
	assert.Contains(t, helper.Currency.Type, "ActiveOrHistoricCurrencyCode")
	assert.Contains(t, helper.Amount.Type, "DecimalNumber")

	// Test documentation
	assert.Contains(t, helper.Currency.Documentation, "currency")
	assert.Contains(t, helper.Amount.Documentation, "fractional digits")
}

// TestBuildEntryDetailHelper tests the entry detail helper builder
func TestBuildEntryDetailHelper(t *testing.T) {
	helper := BuildEntryDetailHelper()

	// Test basic fields are properly initialized
	assert.Equal(t, "Message Identification", helper.MessageId.Title)
	assert.Equal(t, "Instruction Identification", helper.InstructionId.Title)
	assert.Equal(t, "UETR", helper.UniqueTransactionReference.Title)
	assert.Equal(t, "Clearing System Reference", helper.ClearingSystemRef.Title)
	assert.Equal(t, "Local Instrument Choice", helper.LocalInstrumentChoice.Title)

	// Test nested agent fields are properly initialized
	assert.Equal(t, "Member Identification", helper.InstructingAgent.PaymentSysMemberId.Title)
	assert.Equal(t, "Member Identification", helper.InstructedAgent.PaymentSysMemberId.Title)

	// Test type information
	assert.Contains(t, helper.MessageId.Type, "IMAD_FedwireFunds_1")
	assert.Contains(t, helper.InstructionId.Type, "Max35Text")
	assert.Contains(t, helper.UniqueTransactionReference.Type, "UUIDv4Identifier")
}

// TestBuildEntryHelper tests the entry helper builder
func TestBuildEntryHelper(t *testing.T) {
	helper := BuildEntryHelper()

	// Test required fields
	assert.Equal(t, "Credit Debit Indicator", helper.CreditDebitIndicator.Title)
	assert.Equal(t, "Status", helper.Status.Title)
	assert.Equal(t, "Bank Transaction Code", helper.BankTransactionCode.Title)
	assert.Equal(t, "Message Name Identification", helper.MessageNameId.Title)

	// Test nested structures are initialized
	require.NotNil(t, helper.Amount)
	require.NotNil(t, helper.EntryDetails)

	// Verify amount helper structure
	assert.Equal(t, "Currency", helper.Amount.Currency.Title)
	assert.Equal(t, "Amount", helper.Amount.Amount.Title)

	// Test type information
	assert.Contains(t, helper.CreditDebitIndicator.Type, "CdtDbtInd")
	assert.Contains(t, helper.Status.Type, "ReportStatus")
	assert.Contains(t, helper.BankTransactionCode.Type, "TransactionStatusCode")
	assert.Contains(t, helper.MessageNameId.Type, "MessageNameIdentification_FRS_1")
}

// TestBuildPartyIdentifyHelper tests the party identification helper builder
func TestBuildPartyIdentifyHelper(t *testing.T) {
	helper := BuildPartyIdentifyHelper()

	// Test required fields
	assert.Equal(t, "Name", helper.Name.Title)

	// Test nested address structure is initialized
	require.NotNil(t, helper.Address)
	assert.Equal(t, "Street Name", helper.Address.StreetName.Title)
	assert.Equal(t, "Building Number", helper.Address.BuildingNumber.Title)
	assert.Equal(t, "Country", helper.Address.Country.Title)

	// Test type information
	assert.Contains(t, helper.Name.Type, "Max140Text")
	assert.Contains(t, helper.Name.Documentation, "party")
}

// TestBuildNumberAndStatusOfTransactionsHelper tests the number and status helper builder
func TestBuildNumberAndStatusOfTransactionsHelper(t *testing.T) {
	helper := BuildNumberAndStatusOfTransactionsHelper()

	// Test required fields
	assert.Equal(t, "Number Of Entries", helper.NumberOfEntries.Title)
	assert.Equal(t, "Status", helper.Status.Title)

	// Test type information
	assert.Contains(t, helper.NumberOfEntries.Type, "Max15NumericText")
	assert.Contains(t, helper.Status.Type, "TransactionStatusCode")

	// Test documentation
	assert.Contains(t, helper.NumberOfEntries.Documentation, "entries")
	assert.Contains(t, helper.Status.Documentation, "transaction")
}

// TestBuildFiniancialInstitutionIdHelper tests the financial institution helper builder
func TestBuildFiniancialInstitutionIdHelper(t *testing.T) {
	helper := BuildFiniancialInstitutionIdHelper()

	// Test required fields
	assert.Equal(t, "BICFI", helper.BusinessId.Title)
	assert.Equal(t, "Clearing System Identification Code", helper.ClearingSystemId.Title)
	assert.Equal(t, "Member Identification", helper.ClearintSystemMemberId.Title)
	assert.Equal(t, "Name", helper.Name.Title)

	// Test nested address structure
	require.NotNil(t, helper.Address)
	assert.Equal(t, "Street Name", helper.Address.StreetName.Title)

	// Test type information
	assert.Contains(t, helper.BusinessId.Type, "BICFIDec2014Identifier")
	assert.Contains(t, helper.ClearingSystemId.Type, "PaymentSystemType")
	assert.Contains(t, helper.ClearintSystemMemberId.Type, "RoutingNumber_FRS_1")
	assert.Contains(t, helper.Name.Type, "Max140Text")

	// Test documentation
	assert.Contains(t, helper.BusinessId.Documentation, "Business identifier code")
	assert.Contains(t, helper.ClearingSystemId.Documentation, "clearing system")
	assert.Contains(t, helper.Name.Documentation, "agent")
}

// TestElementHelperStructure tests the basic ElementHelper structure
func TestElementHelperStructure(t *testing.T) {
	element := ElementHelper{
		Title:         "Test Title",
		Rules:         "Test Rules",
		Type:          "Test Type",
		Documentation: "Test Documentation",
	}

	assert.Equal(t, "Test Title", element.Title)
	assert.Equal(t, "Test Rules", element.Rules)
	assert.Equal(t, "Test Type", element.Type)
	assert.Equal(t, "Test Documentation", element.Documentation)
}

// TestHelperFunctionsReturnNonNilStructures tests that all helper functions return properly initialized structures
func TestHelperFunctionsReturnNonNilStructures(t *testing.T) {
	// Test all build functions return non-nil structures
	postalHelper := BuildPostalAddressHelper()
	assert.NotEmpty(t, postalHelper.StreetName.Title)

	agentHelper := BuildAgentHelper()
	assert.NotEmpty(t, agentHelper.BusinessIdCode.Title)

	seqHelper := BuildSequenceRangeHelper()
	assert.NotEmpty(t, seqHelper.FromSeq.Title)

	paginationHelper := BuildMessagePagenationHelper()
	assert.NotEmpty(t, paginationHelper.PageNumber.Title)

	numSumHelper := BuildNumberAndSumOfTransactionsHelper()
	assert.NotEmpty(t, numSumHelper.NumberOfEntries.Title)

	currencyHelper := BuildCurrencyAndAmountHelper()
	assert.NotEmpty(t, currencyHelper.Currency.Title)

	entryDetailHelper := BuildEntryDetailHelper()
	assert.NotEmpty(t, entryDetailHelper.MessageId.Title)

	entryHelper := BuildEntryHelper()
	assert.NotEmpty(t, entryHelper.CreditDebitIndicator.Title)

	partyHelper := BuildPartyIdentifyHelper()
	assert.NotEmpty(t, partyHelper.Name.Title)

	numStatusHelper := BuildNumberAndStatusOfTransactionsHelper()
	assert.NotEmpty(t, numStatusHelper.NumberOfEntries.Title)

	finInstHelper := BuildFiniancialInstitutionIdHelper()
	assert.NotEmpty(t, finInstHelper.BusinessId.Title)
}

// TestNestedHelperStructuresAreProperlyCombined tests that nested helpers work correctly
func TestNestedHelperStructuresAreProperlyCombined(t *testing.T) {
	// Test that EntryDetailHelper properly incorporates AgentHelper structures
	entryDetailHelper := BuildEntryDetailHelper()

	// Verify instructing agent has complete agent structure
	assert.Equal(t, "Member Identification", entryDetailHelper.InstructingAgent.PaymentSysMemberId.Title)
	assert.Equal(t, "BICFI", entryDetailHelper.InstructingAgent.BusinessIdCode.Title)
	assert.Equal(t, "Name", entryDetailHelper.InstructingAgent.BankName.Title)

	// Verify instructed agent has complete agent structure
	assert.Equal(t, "Member Identification", entryDetailHelper.InstructedAgent.PaymentSysMemberId.Title)
	assert.Equal(t, "BICFI", entryDetailHelper.InstructedAgent.BusinessIdCode.Title)
	assert.Equal(t, "Name", entryDetailHelper.InstructedAgent.BankName.Title)

	// Test that PartyIdentifyHelper properly incorporates PostalAddressHelper
	partyHelper := BuildPartyIdentifyHelper()
	assert.Equal(t, "Street Name", partyHelper.Address.StreetName.Title)
	assert.Equal(t, "Building Number", partyHelper.Address.BuildingNumber.Title)
	assert.Equal(t, "Country", partyHelper.Address.Country.Title)

	// Test that FinancialInstitutionIdHelper properly incorporates PostalAddressHelper
	finInstHelper := BuildFiniancialInstitutionIdHelper()
	assert.Equal(t, "Street Name", finInstHelper.Address.StreetName.Title)
	assert.Equal(t, "Postal Code", finInstHelper.Address.PostalCode.Title)
	assert.Equal(t, "Town Name", finInstHelper.Address.TownName.Title)

	// Test that EntryHelper properly incorporates CurrencyAndAmountHelper and EntryDetailHelper
	entryHelper := BuildEntryHelper()
	assert.Equal(t, "Currency", entryHelper.Amount.Currency.Title)
	assert.Equal(t, "Amount", entryHelper.Amount.Amount.Title)

	// Verify the nested EntryDetailHelper structure within EntryHelper
	assert.Equal(t, "Message Identification", entryHelper.EntryDetails.MessageId.Title)
	assert.Equal(t, "Instruction Identification", entryHelper.EntryDetails.InstructionId.Title)
}

// BenchmarkBuildHelpers benchmarks the performance of helper building functions
func BenchmarkBuildHelpers(b *testing.B) {
	b.Run("BuildPostalAddressHelper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = BuildPostalAddressHelper()
		}
	})

	b.Run("BuildAgentHelper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = BuildAgentHelper()
		}
	})

	b.Run("BuildEntryDetailHelper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = BuildEntryDetailHelper()
		}
	})

	b.Run("BuildEntryHelper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = BuildEntryHelper()
		}
	})
}
