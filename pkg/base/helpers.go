package base

import "github.com/moov-io/wire20022/pkg/models"

// HelperBuilder represents a function that creates an ElementHelper
type HelperBuilder func() models.ElementHelper

// CommonHelpers provides standard ElementHelper definitions used across message types
var CommonHelpers = map[string]HelperBuilder{
	"MessageId": func() models.ElementHelper {
		return models.ElementHelper{
			Title: "Message Identification",
			Type:  `IMAD_FedwireFunds_1`,
			Documentation: `Point to point reference assigned by the assigner to unambiguously identify the message.
Usage: The assigner has to be the submitter of the message or an entity 
that is authorized to represent the message submitter. This reference can be used by any party to identify the message.`,
		}
	},
	"CreatedDateTime": func() models.ElementHelper {
		return models.ElementHelper{
			Title: "Creation Date Time",
			Type:  `ISODateTime`,
			Documentation: `Date and time at which the message was created.
Business rule: CreatedDateTime must be the central processing date of the Federal Reserve Bank that received the message.`,
		}
	},
	"NumberOfTransactions": func() models.ElementHelper {
		return models.ElementHelper{
			Title: "Number Of Transactions",
			Type:  `Max15NumericText`,
			Documentation: `Number of individual transactions contained in the message.
Business rule: NumberOfTransactions must equal the actual number of transactions in the message.`,
		}
	},
	"SettlementMethod": func() models.ElementHelper {
		return models.ElementHelper{
			Title: "Settlement Method",
			Type:  `SettlementMethod1Code`,
			Documentation: `Method used to settle the payment instruction.
Valid values: CLRG (Clearing), INDA (Cover), INGA (Guaranteed)`,
		}
	},
	"CommonClearingSysCode": func() models.ElementHelper {
		return models.ElementHelper{
			Title: "Common Clearing System Code",
			Type:  `ClearingSystemIdentification2Choice`,
			Documentation: `Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
Valid values: USABA (US ABA Routing Number), CHIPS (Clearing House Interbank Payments System)`,
		}
	},
	"InstructingAgent": func() models.ElementHelper {
		return models.ElementHelper{
			Title:         "Instructing Agent",
			Type:          `BranchAndFinancialInstitutionIdentification5`,
			Documentation: `Agent that instructs the next party in the chain to carry out the (set of) instruction(s).`,
		}
	},
	"InstructedAgent": func() models.ElementHelper {
		return models.ElementHelper{
			Title:         "Instructed Agent",
			Type:          `BranchAndFinancialInstitutionIdentification5`,
			Documentation: `Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).`,
		}
	},
	"DebtorAgent": func() models.ElementHelper {
		return models.ElementHelper{
			Title:         "Debtor Agent",
			Type:          `BranchAndFinancialInstitutionIdentification5`,
			Documentation: `Financial institution servicing an account for the debtor.`,
		}
	},
	"CreditorAgent": func() models.ElementHelper {
		return models.ElementHelper{
			Title:         "Creditor Agent",
			Type:          `BranchAndFinancialInstitutionIdentification5`,
			Documentation: `Financial institution servicing an account for the creditor.`,
		}
	},
	"InitiatingPartyName": func() models.ElementHelper {
		return models.ElementHelper{
			Title:         "Initiating Party Name",
			Type:          `Max140Text`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		}
	},
}

// StandardMessageHelper provides common helper structure for basic message fields
type StandardMessageHelper struct {
	MessageId       models.ElementHelper `json:"messageId"`
	CreatedDateTime models.ElementHelper `json:"createdDateTime"`
}

// BuildStandardMessageHelper creates a helper with common message fields
func BuildStandardMessageHelper() StandardMessageHelper {
	return StandardMessageHelper{
		MessageId:       CommonHelpers["MessageId"](),
		CreatedDateTime: CommonHelpers["CreatedDateTime"](),
	}
}

// PaymentMessageHelper extends StandardMessageHelper with payment-specific fields
type PaymentMessageHelper struct {
	StandardMessageHelper
	NumberOfTransactions  models.ElementHelper `json:"numberOfTransactions"`
	SettlementMethod      models.ElementHelper `json:"settlementMethod"`
	CommonClearingSysCode models.ElementHelper `json:"commonClearingSysCode"`
}

// BuildPaymentMessageHelper creates a helper with common payment fields
func BuildPaymentMessageHelper() PaymentMessageHelper {
	return PaymentMessageHelper{
		StandardMessageHelper: BuildStandardMessageHelper(),
		NumberOfTransactions:  CommonHelpers["NumberOfTransactions"](),
		SettlementMethod:      CommonHelpers["SettlementMethod"](),
		CommonClearingSysCode: CommonHelpers["CommonClearingSysCode"](),
	}
}

// AgentHelper provides helper structure for agent-related fields
type AgentHelper struct {
	InstructingAgent models.ElementHelper `json:"instructingAgent"`
	InstructedAgent  models.ElementHelper `json:"instructedAgent"`
	DebtorAgent      models.ElementHelper `json:"debtorAgent"`
	CreditorAgent    models.ElementHelper `json:"creditorAgent"`
}

// BuildAgentHelper creates a helper with common agent fields
func BuildAgentHelper() AgentHelper {
	return AgentHelper{
		InstructingAgent: CommonHelpers["InstructingAgent"](),
		InstructedAgent:  CommonHelpers["InstructedAgent"](),
		DebtorAgent:      CommonHelpers["DebtorAgent"](),
		CreditorAgent:    CommonHelpers["CreditorAgent"](),
	}
}

// AddressHelper provides helper structure for address-related fields
type AddressHelper struct {
	StreetName     models.ElementHelper `json:"streetName"`
	BuildingNumber models.ElementHelper `json:"buildingNumber"`
	RoomNumber     models.ElementHelper `json:"roomNumber"`
	PostalCode     models.ElementHelper `json:"postalCode"`
	TownName       models.ElementHelper `json:"townName"`
	Subdivision    models.ElementHelper `json:"subdivision"`
	Country        models.ElementHelper `json:"country"`
}

// BuildAddressHelper creates a helper with common address fields
func BuildAddressHelper() AddressHelper {
	return AddressHelper{
		StreetName: models.ElementHelper{
			Title:         "Street Name",
			Type:          `Max70Text`,
			Documentation: `Name of a street or thoroughfare.`,
		},
		BuildingNumber: models.ElementHelper{
			Title:         "Building Number",
			Type:          `Max16Text`,
			Documentation: `Number that identifies the position of a building on a street.`,
		},
		RoomNumber: models.ElementHelper{
			Title:         "Room Number",
			Type:          `Max70Text`,
			Documentation: `Number or identifier of a room, suite, or apartment within a building.`,
		},
		PostalCode: models.ElementHelper{
			Title:         "Postal Code",
			Type:          `Max16Text`,
			Documentation: `Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.`,
		},
		TownName: models.ElementHelper{
			Title:         "Town Name",
			Type:          `Max35Text`,
			Documentation: `Name of a built-up area, with defined boundaries, and a local government.`,
		},
		Subdivision: models.ElementHelper{
			Title:         "Country Subdivision",
			Type:          `Max35Text`,
			Documentation: `Identifies a subdivision of a country such as state, region, county.`,
		},
		Country: models.ElementHelper{
			Title:         "Country",
			Type:          `CountryCode`,
			Documentation: `Nation with its own government, occupying a particular territory.`,
		},
	}
}
