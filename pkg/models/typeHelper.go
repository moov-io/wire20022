package model

type ElementHelper struct {
	Title         string
	Rules         string
	Type          string
	Documentation string
}

type PostalAddressHelper struct {
	StreetName     ElementHelper
	BuildingNumber ElementHelper
	BuildingName   ElementHelper
	Floor          ElementHelper
	RoomNumber     ElementHelper
	PostalCode     ElementHelper
	TownName       ElementHelper
	Subdivision    ElementHelper
	Country        ElementHelper
}

func BuildPostalAddressHelper() PostalAddressHelper {
	return PostalAddressHelper{
		StreetName: ElementHelper{
			Title:         "Street Name",
			Rules:         "",
			Type:          `Max70Text (based on string) minLength: 1 maxLength: 70`,
			Documentation: `Name of a street or thoroughfare.`,
		},
		BuildingNumber: ElementHelper{
			Title:         "Building Number",
			Rules:         "",
			Type:          `Max16Text (based on string) minLength: 1 maxLength: 16`,
			Documentation: `Number that identifies the position of a building on a street.`,
		},
		BuildingName: ElementHelper{
			Title:         "Building Name",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Name of the building or house.`,
		},
		Floor: ElementHelper{
			Title:         "Floor",
			Rules:         "",
			Type:          `Max70Text (based on string) minLength: 1 maxLength: 70`,
			Documentation: `Floor or storey within a building.`,
		},
		RoomNumber: ElementHelper{
			Title:         "Room Number",
			Rules:         "",
			Type:          `Max70Text (based on string) minLength: 1 maxLength: 70`,
			Documentation: `Building room number.`,
		},
		PostalCode: ElementHelper{
			Title:         "Postal Code",
			Rules:         "",
			Type:          `Max16Text (based on string) minLength: 1 maxLength: 16`,
			Documentation: `Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.`,
		},
		TownName: ElementHelper{
			Title:         "Town Name",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Name of a built-up area, with defined boundaries, and a local government.`,
		},
		Subdivision: ElementHelper{
			Title:         "Country Sub Division",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identifies a subdivision of a country such as state, region, county.`,
		},
		Country: ElementHelper{
			Title:         "Country",
			Rules:         "",
			Type:          `CountryCode (based on string) pattern: [A-Z]{2,2}`,
			Documentation: `Nation with its own government.`,
		},
	}
}

type AgentHelper struct {
	BusinessIdCode     ElementHelper
	PaymentSysCode     ElementHelper
	PaymentSysMemberId ElementHelper
	BankName           ElementHelper
	PostalAddress      PostalAddressHelper
}

func BuildAgentHelper() AgentHelper {
	return AgentHelper{
		BusinessIdCode: ElementHelper{
			Title:         "BICFI",
			Rules:         "",
			Type:          `BICFIDec2014Identifier (based on string) pattern: [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1} identificationScheme: SWIFT; BICIdentifier`,
			Documentation: `Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".`,
		},
		PaymentSysCode: ElementHelper{
			Title:         "Clearing System Identification Code",
			Rules:         "",
			Type:          `PaymentSystemType(PaymentSysUSABA, PaymentSysCHIPS, PaymentSysSEPA, PaymentSysRTGS, PaymentSysSWIFT, PaymentSysBACS)`,
			Documentation: `Identification of a clearing system, in a coded form as published in an external list.`,
		},
		PaymentSysMemberId: ElementHelper{
			Title:         "Member Identification",
			Rules:         "",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification of a member of a clearing system.`,
		},
		BankName: ElementHelper{
			Title:         "Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which an agent is known and which is usually used to identify that agent.`,
		},
	}
}

type SequenceRangeHelper struct {
	FromSeq ElementHelper
	ToSeq   ElementHelper
}

func BuildSequenceRangeHelper() SequenceRangeHelper {
	return SequenceRangeHelper{
		FromSeq: ElementHelper{
			Title:         "From Sequence",
			Rules:         "Sequence number in From Sequence must be less than or equal to sequence number in To Sequence.",
			Type:          `SequenceNumber_FedwireFunds_1 (based on decimal) minInclusive: 000001 maxInclusive: 999999 pattern: [0-9]{6,6}`,
			Documentation: `Start sequence of the range.`,
		},
		ToSeq: ElementHelper{
			Title:         "To Sequence",
			Rules:         "Sequence number in From Sequence must be less than or equal to sequence number in To Sequence.",
			Type:          `SequenceNumber_FedwireFunds_1 (based on decimal) minInclusive: 000001 maxInclusive: 999999 pattern: [0-9]{6,6}`,
			Documentation: `End sequence of the range.`,
		},
	}
}

type MessagePagenationHelper struct {
	PageNumber        ElementHelper
	LastPageIndicator ElementHelper
}

func BuildMessagePagenationHelper() MessagePagenationHelper {
	return MessagePagenationHelper{
		PageNumber: ElementHelper{
			Title:         "Page Number",
			Rules:         "",
			Type:          `Max5NumericText (based on string) pattern: [0-9]{1,5}`,
			Documentation: `Page number.`,
		},
		LastPageIndicator: ElementHelper{
			Title:         "Last Page Indicator",
			Rules:         "",
			Type:          `YesNoIndicator (based on boolean) meaningWhenFalse: No meaningWhenTrue: Yes`,
			Documentation: `Indicates the last page.`,
		},
	}
}

type NumberAndSumOfTransactionsHelper struct {
	NumberOfEntries ElementHelper
	Sum             ElementHelper
}

func BuildNumberAndSumOfTransactionsHelper() NumberAndSumOfTransactionsHelper {
	return NumberAndSumOfTransactionsHelper{
		NumberOfEntries: ElementHelper{
			Title:         "Number Of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) pattern: [0-9]{1,15}`,
			Documentation: `Number of individual entries included in the report.`,
		},
		Sum: ElementHelper{
			Title:         "Sum",
			Rules:         "",
			Type:          `DecimalNumber (based on decimal) totalDigits: 18 fractionDigits: 17`,
			Documentation: `Total of all individual entries included in the report.`,
		},
	}
}

type CurrencyAndAmountHelper struct {
	Currency ElementHelper
	Amount   ElementHelper
}

func BuildCurrencyAndAmountHelper() CurrencyAndAmountHelper {
	return CurrencyAndAmountHelper{
		Currency: ElementHelper{
			Title:         "Currency",
			Rules:         "",
			Type:          `ActiveOrHistoricCurrencyCode (based on string) pattern: [A-Z]{3,3}`,
			Documentation: `Medium of exchange of currency.`,
		},
		Amount: ElementHelper{
			Title:         "Amount",
			Rules:         "",
			Type:          `DecimalNumber (based on decimal) totalDigits: 18 fractionDigits: 17`,
			Documentation: `The number of fractional digits (or minor unit of currency) must comply with ISO 4217. Note: The decimal separator is a dot.`,
		},
	}
}

type EntryDetailHelper struct {
	MessageId                  ElementHelper
	InstructionId              ElementHelper
	UniqueTransactionReference ElementHelper
	ClearingSystemRef          ElementHelper
	InstructingAgent           AgentHelper
	InstructedAgent            AgentHelper
	LocalInstrumentChoice      ElementHelper
	RelatedDatesProprietary    ElementHelper
	RelatedDateTime            ElementHelper
}

func BuildEntryDetailHelper() EntryDetailHelper {
	return EntryDetailHelper{
		MessageId: ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `IMAD_FedwireFunds_1 (based on string) minLength: 22 maxLength: 22 pattern: [0-9]{8}[A-Z0-9]{8}[0-9]{6}`,
			Documentation: `Point to point reference, as assigned by the instructing party of the underlying message.`,
		},
		InstructionId: ElementHelper{
			Title:         "Instruction Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction. Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.`,
		},
		UniqueTransactionReference: ElementHelper{
			Title:         "UETR",
			Rules:         "",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide an end-to-end reference of a payment transaction.`,
		},
		ClearingSystemRef: ElementHelper{
			Title:         "Clearing System Reference",
			Rules:         "",
			Type:          `OMAD_FedwireFunds_1 (based on string) pattern: [A-Z0-9]{34,34}`,
			Documentation: `Unique reference, as assigned by a clearing system, to unambiguously identify the instruction. Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.`,
		},
		InstructingAgent: BuildAgentHelper(),
		InstructedAgent:  BuildAgentHelper(),
		LocalInstrumentChoice: ElementHelper{
			Title:         "Local Instrument Choice",
			Rules:         "",
			Type:          `InstrumentPropCodeType(InstrumentCTRC, InstrumentDD ...)`,
			Documentation: `User community specific instrument. Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.`,
		},
		RelatedDatesProprietary: ElementHelper{
			Title:         "Related Dates Proprietary",
			Rules:         "",
			Type:          `WorkingDayType`,
			Documentation: `Specifies the type of date.`,
		},
		RelatedDateTime: ElementHelper{
			Title:         "Related Dates Proprietary",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date in ISO format.`,
		},
	}
}

type EntryHelper struct {
	Amount               CurrencyAndAmountHelper
	CreditDebitIndicator ElementHelper
	Status               ElementHelper
	BankTransactionCode  ElementHelper
	MessageNameId        ElementHelper
	EntryDetails         EntryDetailHelper
}

func BuildEntryHelper() EntryHelper {
	return EntryHelper{
		Amount: BuildCurrencyAndAmountHelper(),
		CreditDebitIndicator: ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `CdtDbtInd(Credit, Debit)`,
			Documentation: `Indicates whether the transaction is a credit or a debit transaction.`,
		},
		Status: ElementHelper{
			Title:         "Status",
			Rules:         "",
			Type:          `ReportStatus(Book, Pending, Received, Settled)`,
			Documentation: `Status of an entry on the books of the account servicer.`,
		},
		BankTransactionCode: ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `TransactionStatusCode(MessagesInProcess, MessagesIntercepted ...)`,
			Documentation: `Set of elements used to fully identify the type of underlying transaction resulting in an entry.`,
		},
		MessageNameId: ElementHelper{
			Title:         "Message Name Identification",
			Rules:         "",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `Specifies the message name identifier of the message that will be used to provide additional details.`,
		},
		EntryDetails: BuildEntryDetailHelper(),
	}
}

type PartyIdentifyHelper struct {
	Name    ElementHelper
	Address PostalAddressHelper
}

func BuildPartyIdentifyHelper() PartyIdentifyHelper {
	return PartyIdentifyHelper{
		Name: ElementHelper{
			Title:         "Name",
			Rules:         "If BIC is not present, then Name must be present. Postal address information may be required under applicable law. Even when not required, it is strongly recommended to include this information to the extent possible.",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		},
		Address: BuildPostalAddressHelper(),
	}
}

type NumberAndStatusOfTransactionsHelper struct {
	NumberOfEntries ElementHelper
	Status          ElementHelper
}

func BuildNumberAndStatusOfTransactionsHelper() NumberAndStatusOfTransactionsHelper {
	return NumberAndStatusOfTransactionsHelper{
		NumberOfEntries: ElementHelper{
			Title:         "Number Of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) pattern: [0-9]{1,15}`,
			Documentation: `Number of individual entries for the bank transaction code.`,
		},
		Status: ElementHelper{
			Title:         "Status",
			Rules:         "",
			Type:          `TransactionStatusCode(MessagesInProcess, MessagesIntercepted...)`,
			Documentation: `Proprietary bank transaction code to identify the underlying transaction.`,
		},
	}
}

type FiniancialInstitutionIdHelper struct {
	BusinessId             ElementHelper
	ClearingSystemId       ElementHelper
	ClearintSystemMemberId ElementHelper
	Name                   ElementHelper
	Address                PostalAddressHelper
}

func BuildFiniancialInstitutionIdHelper() FiniancialInstitutionIdHelper {
	return FiniancialInstitutionIdHelper{
		BusinessId: ElementHelper{
			Title:         "BICFI",
			Rules:         "",
			Type:          `BICFIDec2014Identifier (based on string) pattern: [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1} identificationScheme: SWIFT; BICIdentifier`,
			Documentation: `Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".`,
		},
		ClearingSystemId: ElementHelper{
			Title:         "Clearing System Identification Code",
			Rules:         "",
			Type:          `PaymentSystemType(PaymentSysUSABA, PaymentSysCHIPS, PaymentSysSEPA, PaymentSysRTGS, PaymentSysSWIFT, PaymentSysBACS)`,
			Documentation: `Identification of a clearing system, in a coded form as published in an external list.`,
		},
		ClearintSystemMemberId: ElementHelper{
			Title:         "Member Identification",
			Rules:         "",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification of a member of a clearing system.`,
		},
		Name: ElementHelper{
			Title:         "Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which an agent is known and which is usually used to identify that agent.`,
		},
		Address: BuildPostalAddressHelper(),
	}
}
