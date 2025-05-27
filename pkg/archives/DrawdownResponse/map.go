package DrawdownResponse

import (
	"time"

	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_01"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type Message struct {
	Doc  pain_014_001_01.Document
	Data MessageModel
}

func Convert() {
	msg := Message{}
	msg.Data.MessageId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId)
	msg.Data.CreateDatetime = time.Time(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CreDtTm)
	msg.Data.InitiatingParty.Name = string(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.Nm)
	msg.Data.InitiatingParty.Address.StreetName = string(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.StrtNm)
	msg.Data.InitiatingParty.Address.BuildingNumber = string(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.BldgNb)
	msg.Data.InitiatingParty.Address.PostalCode = string(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.PstCd)
	msg.Data.InitiatingParty.Address.TownName = string(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.TwnNm)
	msg.Data.InitiatingParty.Address.Subdivision = string(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn)
	msg.Data.InitiatingParty.Address.Country = string(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.Ctry)
	msg.Data.DebtorAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.DebtorAgent.PaymentSysMemberId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.CreditorAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.CreditorAgent.PaymentSysMemberId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.OriginalMessageId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId)
	msg.Data.OriginalMessageNameId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId)
	msg.Data.OriginalCreationDateTime = time.Time(*msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlCreDtTm)
	msg.Data.OriginalPaymentInfoId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].OrgnlPmtInfId)
	msg.Data.TransactionInformationAndStatus.OriginalInstructionId = string(*msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].OrgnlInstrId)
	msg.Data.TransactionInformationAndStatus.OriginalEndToEndId = string(*msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].OrgnlEndToEndId)
	msg.Data.TransactionInformationAndStatus.TransactionStatus = Archive.TransactionStatusCode(*msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].TxSts)
	msg.Data.TransactionInformationAndStatus.StatusReasonInfoCode = Archive.StatusReasonInformationCode(*msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].StsRsnInf[0].Rsn.Cd)
}

func PathMapV1() map[string]string {
	return PathMapV6()
}
func PathMapV2() map[string]string {
	return PathMapV6()
}
func PathMapV3() map[string]string {
	return PathMapV6()
}
func PathMapV4() map[string]string {
	return PathMapV6()
}
func PathMapV5() map[string]string {
	return PathMapV6()
}
func PathMapV6() map[string]string {
	return map[string]string{
		"CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId":                                            "MessageId",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.CreDtTm":                                          "CreateDatetime",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.Nm":                                      "InitiatingParty.Name",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.StrtNm":                          "InitiatingParty.Address.StreetName",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.BldgNb":                          "InitiatingParty.Address.BuildingNumber",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.PstCd":                           "InitiatingParty.Address.PostalCode",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.TwnNm":                           "InitiatingParty.Address.TownName",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn":                     "InitiatingParty.Address.Subdivision",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.Ctry":                            "InitiatingParty.Address.Country",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":       "DebtorAgent.PaymentSysCode",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":             "DebtorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":       "CreditorAgent.PaymentSysCode",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":             "CreditorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId":                            "OriginalMessageId",
		"CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId":                          "OriginalMessageNameId",
		"CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlCreDtTm":                          "OriginalCreationDateTime",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].OrgnlPmtInfId":                      "OriginalPaymentInfoId",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].OrgnlInstrId":        "TransactionInformationAndStatus.OriginalInstructionId",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].OrgnlEndToEndId":     "TransactionInformationAndStatus.OriginalEndToEndId",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].TxSts":               "TransactionInformationAndStatus.TransactionStatus",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].StsRsnInf[0].Rsn.Cd": "TransactionInformationAndStatus.StatusReasonInfoCode",
	}
}
func PathMapV7() map[string]string {
	return map[string]string{
		"CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId":                                            "MessageId",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.CreDtTm":                                          "CreateDatetime",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.Nm":                                      "InitiatingParty.Name",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.StrtNm":                          "InitiatingParty.Address.StreetName",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.BldgNb":                          "InitiatingParty.Address.BuildingNumber",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.Room":                            "InitiatingParty.Address.RoomNumber",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.PstCd":                           "InitiatingParty.Address.PostalCode",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.TwnNm":                           "InitiatingParty.Address.TownName",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn":                     "InitiatingParty.Address.Subdivision",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty.PstlAdr.Ctry":                            "InitiatingParty.Address.Country",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":       "DebtorAgent.PaymentSysCode",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":             "DebtorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":       "CreditorAgent.PaymentSysCode",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":             "CreditorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId":                            "OriginalMessageId",
		"CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId":                          "OriginalMessageNameId",
		"CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlCreDtTm":                          "OriginalCreationDateTime",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].OrgnlPmtInfId":                      "OriginalPaymentInfoId",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].OrgnlInstrId":        "TransactionInformationAndStatus.OriginalInstructionId",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].OrgnlEndToEndId":     "TransactionInformationAndStatus.OriginalEndToEndId",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].OrgnlUETR":           "TransactionInformationAndStatus.OriginalUniqueId",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].TxSts":               "TransactionInformationAndStatus.TransactionStatus",
		"CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts[0].TxInfAndSts[0].StsRsnInf[0].Rsn.Cd": "TransactionInformationAndStatus.StatusReasonInfoCode",
	}
}
func PathMapV8() map[string]string {
	return PathMapV7()
}
func PathMapV9() map[string]string {
	return PathMapV7()
}
func PathMapV10() map[string]string {
	return PathMapV7()
}
