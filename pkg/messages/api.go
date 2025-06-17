// Package messages provides type-safe processors for ISO 20022 message types.
package messages

// Message processors for each ISO 20022 message type
// Each processor provides:
// - CreateDocument: Convert JSON model to XML document
// - ValidateDocument: Validate JSON model and create document
// - Validate: Validate required fields
// - ConvertXMLToModel: Parse XML to typed model
// - GetHelp: Get field documentation

// Payment messages (pacs)
// - CustomerCreditTransfer: pacs.008 - Customer credit transfer
// - PaymentReturn: pacs.004 - Payment return
// - PaymentStatusRequest: pacs.028 - Payment status request
// - FedwireFundsPaymentStatus: pacs.002 - Payment status report

// Cash management messages (camt)
// - AccountReportingRequest: camt.060 - Account reporting request
// - ActivityReport: camt.086 - Bank services billing statement
// - EndpointDetailsReport: camt.090 - Request for member profile
// - EndpointGapReport: camt.087 - Request for duplicate
// - EndpointTotalsReport: camt.089 - Request to cancel payment
// - ReturnRequestResponse: camt.029 - Resolution of investigation
// - Master: camt.052 - Bank to customer account report

// Payment initiation messages (pain)
// - DrawdownRequest: pain.013 - Creditor payment activation request
// - DrawdownResponse: pain.014 - Creditor payment activation request status report

// Administrative messages (admi)
// - ConnectionCheck: admi.001 - Static data request
// - FedwireFundsAcknowledgement: admi.004 - System event acknowledgement
// - FedwireFundsSystemResponse: admi.011 - System event notification
