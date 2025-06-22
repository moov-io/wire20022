package base

import (
	"time"

	"github.com/moov-io/wire20022/pkg/models"
)

// MessageHeader contains common fields found in all ISO 20022 messages
type MessageHeader struct {
	MessageId       string    `json:"messageId"`
	CreatedDateTime time.Time `json:"createdDateTime"`
}

// PaymentCore extends MessageHeader with fields common to payment-related messages
type PaymentCore struct {
	MessageHeader
	NumberOfTransactions  string                           `json:"numberOfTransactions"`
	SettlementMethod      models.SettlementMethodType      `json:"settlementMethod"`
	CommonClearingSysCode models.CommonClearingSysCodeType `json:"commonClearingSysCode"`
}

// AgentPair represents the common pattern of agent pairs in messages
type AgentPair struct {
	InstructingAgent models.Agent `json:"instructingAgent"`
	InstructedAgent  models.Agent `json:"instructedAgent"`
}

// DebtorCreditorPair represents the common debtor/creditor agent pattern
type DebtorCreditorPair struct {
	DebtorAgent   models.Agent `json:"debtorAgent"`
	CreditorAgent models.Agent `json:"creditorAgent"`
}

// PartyAddress represents the common postal address pattern
type PartyAddress struct {
	StreetName     string `json:"streetName"`
	BuildingNumber string `json:"buildingNumber"`
	RoomNumber     string `json:"roomNumber"`
	PostalCode     string `json:"postalCode"`
	TownName       string `json:"townName"`
	Subdivision    string `json:"subdivision"`
	Country        string `json:"country"`
}

// Party represents a common party structure with name and address
type Party struct {
	Name    string       `json:"name"`
	Address PartyAddress `json:"address"`
}

// InitiatingParty embeds Party for the common initiating party pattern
type InitiatingParty struct {
	Party
}
