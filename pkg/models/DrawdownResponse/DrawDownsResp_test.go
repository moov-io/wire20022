package DrawdownResponse

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestRequireField(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	cErr := message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreateDatetime, InitiatingParty, DebtorAgent, CreditorAgent, OriginalMessageId, OriginalMessageNameId, OriginalCreationDateTime, OriginalPaymentInfoId, TransactionInformationAndStatus")
}
func generateRequreFields(msg Message) Message {
	if msg.data.MessageId == "" {
		msg.data.MessageId = "20250310B1QDRCQR000602"
	}
	if msg.data.CreateDatetime.IsZero() {
		msg.data.CreateDatetime = time.Now()
	}
	if isEmpty(msg.data.InitiatingParty) {
		msg.data.InitiatingParty = model.PartyIdentify{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		}
	}
	if isEmpty(msg.data.DebtorAgent) {
		msg.data.DebtorAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	if isEmpty(msg.data.CreditorAgent) {
		msg.data.CreditorAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if msg.data.OriginalMessageId == "" {
		msg.data.OriginalMessageId = "20250310B1QDRCQR000601"
	}
	if msg.data.OriginalMessageNameId == "" {
		msg.data.OriginalMessageNameId = "pain.013.001.07"
	}
	if msg.data.OriginalCreationDateTime.IsZero() {
		msg.data.OriginalCreationDateTime = time.Now()
	}
	if msg.data.OriginalPaymentInfoId == "" {
		msg.data.OriginalPaymentInfoId = "20250310B1QDRCQR000601"
	}
	if isEmpty(msg.data.TransactionInformationAndStatus) {
		msg.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
			OriginalInstructionId: "Scenario01Step1InstrId001",
			OriginalEndToEndId:    "Scenario1EndToEndId001",
			OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
			TransactionStatus:     model.AcceptedTechnicalValidation,
		}
	}
	return msg
}
func TestDrawdownResponseFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step2_pain.014")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.doc.CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId), "20250310B1QDRCQR000602")
	require.Equal(t, string(*message.doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.Nm), "Corporation A")
	require.Equal(t, string(message.doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId), "021040078")
	require.Equal(t, string(*message.doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA")
	require.Equal(t, string(message.doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId), "20250310B1QDRCQR000601")
	require.Equal(t, string(message.doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId), "pain.013.001.07")
	require.Equal(t, string(message.doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.OrgnlPmtInfId), "20250310B1QDRCQR000601")
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"
const INVALID_TRCOUNT string = "123456789012345"
const INVALID_MESSAGE_ID string = "12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456"
const INVALID_OTHER_ID string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const INVALID_BUILD_NUM string = "12345678901234567"
const INVALID_POSTAL_CODE string = "12345678901234567"
const INVALID_COUNTRY_CODE string = "12345678"
const INVALID_MESSAGE_NAME_ID string = "sabcd-123-001-12"
const INVALID_PAY_SYSCODE model.PaymentSystemType = model.PaymentSystemType(INVALID_COUNT)

func TestDrawdownResponseValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"MessageId",
			Message{data: MessageModel{MessageId: INVALID_MESSAGE_ID}},
			"error occur at MessageId: 12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456 fails validation with length 66 <= required maxLength 35",
		},
		{
			"InitiatingParty - Postal",
			Message{data: MessageModel{InitiatingParty: model.PartyIdentify{
				Name: "Corporation A",
				Address: model.PostalAddress{
					StreetName:     "Avenue of the Fountains",
					BuildingNumber: "167565",
					RoomNumber:     "Suite D110",
					PostalCode:     INVALID_POSTAL_CODE,
					TownName:       "Fountain Hills",
					Subdivision:    "AZ",
					Country:        "US",
				},
			}}},
			"error occur at InitiatingParty.Address.PostalCode: 12345678901234567 fails validation with length 17 <= required maxLength 16",
		},
		{
			"InitiatingParty - Postal",
			Message{data: MessageModel{InitiatingParty: model.PartyIdentify{
				Name: "Corporation A",
				Address: model.PostalAddress{
					StreetName:     "Avenue of the Fountains",
					BuildingNumber: "167565",
					RoomNumber:     "Suite D110",
					PostalCode:     "85268",
					TownName:       "Fountain Hills",
					Subdivision:    "AZ",
					Country:        INVALID_COUNTRY_CODE,
				},
			}}},
			"error occur at InitiatingParty.Address.Country: 12345678 fails validation with pattern [A-Z]{2,2}",
		},
		{
			"DebtorAgent - PaymentSysCode",
			Message{data: MessageModel{DebtorAgent: model.Agent{
				PaymentSysCode:     INVALID_PAY_SYSCODE,
				PaymentSysMemberId: "021040078",
			}}},
			"error occur at DebtorAgent.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"DebtorAgent - PaymentSysMemberId",
			Message{data: MessageModel{DebtorAgent: model.Agent{
				PaymentSysCode:     model.PaymentSysUSABA,
				PaymentSysMemberId: INVALID_ACCOUNT_ID,
			}}},
			"error occur at DebtorAgent.PaymentSysMemberId: 123ABC789 fails validation with pattern [0-9]{9,9}",
		},
		{
			"OriginalMessageNameId",
			Message{data: MessageModel{OriginalMessageNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at OriginalMessageNameId: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			nMsg := generateRequreFields(tt.msg)
			msgErr := nMsg.CreateDocument()
			if msgErr != nil {
				require.Equal(t, tt.expectedErr, msgErr.Error())
			}
		})
	}
}
func TestDrawdowns_Scenario1_Step2_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000602"
	message.data.CreateDatetime = time.Now()
	message.data.InitiatingParty = model.PartyIdentify{
		Name: "Corporation A",
		Address: model.PostalAddress{
			StreetName:     "Avenue of the Fountains",
			BuildingNumber: "167565",
			RoomNumber:     "Suite D110",
			PostalCode:     "85268",
			TownName:       "Fountain Hills",
			Subdivision:    "AZ",
			Country:        "US",
		},
	}
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000601"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000601"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01Step1InstrId001",
		OriginalEndToEndId:    "Scenario1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario1_Step2_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step2_pain.014")
	genterated := filepath.Join("generated", "Drawdowns_Scenario1_Step2_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario2_Step2_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000612"
	message.data.CreateDatetime = time.Now()
	message.data.InitiatingParty = model.PartyIdentify{
		Name: "Corporation A",
		Address: model.PostalAddress{
			StreetName:     "Avenue of the Fountains",
			BuildingNumber: "167565",
			RoomNumber:     "Suite D110",
			PostalCode:     "85268",
			TownName:       "Fountain Hills",
			Subdivision:    "AZ",
			Country:        "US",
		},
	}
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000611"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000611"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario02Step1InstrId001",
		OriginalEndToEndId:    "Scenario2EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f068",
		TransactionStatus:     model.Rejected,
		StatusReasonInfoCode:  InsufficientFunds,
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario2_Step2_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario2_Step2_pain.014")
	genterated := filepath.Join("generated", "Drawdowns_Scenario2_Step2_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario3_Step2_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000622"
	message.data.CreateDatetime = time.Now()
	message.data.InitiatingParty = model.PartyIdentify{
		Name: "Bank A",
		Address: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000621"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000621"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario03Step1InstrId001",
		OriginalEndToEndId:    "Scenario3EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f070",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario3_Step2_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step2_pain.014")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step2_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step2_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000682"
	message.data.CreateDatetime = time.Now()
	message.data.InitiatingParty = model.PartyIdentify{
		Name: "Bank Aa",
		Address: model.PostalAddress{
			StreetName:     "Main Road",
			BuildingNumber: "3",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000681"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000681"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario04Step1InstrId001",
		OriginalEndToEndId:    "Scenario4EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f070",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario4_Step2_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step2_pain.014")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step2_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario5_Step2_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000632"
	message.data.CreateDatetime = time.Now()
	message.data.InitiatingParty = model.PartyIdentify{
		Name: "Corporation A",
		Address: model.PostalAddress{
			StreetName:     "Avenue of the Fountains",
			BuildingNumber: "167565",
			RoomNumber:     "Suite D110",
			PostalCode:     "85268",
			TownName:       "Fountain Hills",
			Subdivision:    "AZ",
			Country:        "US",
		},
	}
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000631"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000631"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario04Step1InstrId001",
		OriginalEndToEndId:    "Scenario4EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f258",
		TransactionStatus:     model.TransPending,
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario5_Step2_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario5_Step2_pain.014")
	genterated := filepath.Join("generated", "Drawdowns_Scenario5_Step2_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000712"
	message.data.CreateDatetime = time.Now()
	message.data.InitiatingParty = model.PartyIdentify{
		Name: "Corporation A",
		Address: model.PostalAddress{
			StreetName:     "Avenue of the Fountains",
			BuildingNumber: "167565",
			RoomNumber:     "Suite D110",
			PostalCode:     "85268",
			TownName:       "Fountain Hills",
			Subdivision:    "AZ",
			Country:        "US",
		},
	}
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000711"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000711"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01InstrId001",
		OriginalEndToEndId:    "Scenario01Step1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f078",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step2_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step2_pain.014")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step2_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2b_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000712"
	message.data.CreateDatetime = time.Now()
	message.data.InitiatingParty = model.PartyIdentify{
		Name: "Corporation A",
		Address: model.PostalAddress{
			StreetName:     "Avenue of the Fountains",
			BuildingNumber: "167565",
			RoomNumber:     "Suite D110",
			PostalCode:     "85268",
			TownName:       "Fountain Hills",
			Subdivision:    "AZ",
			Country:        "US",
		},
	}
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000711"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000711"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01InstrId001",
		OriginalEndToEndId:    "Scenario01Step1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f078",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step2b_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step2b_pain.014")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step2b_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
