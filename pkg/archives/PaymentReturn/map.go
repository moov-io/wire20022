package PaymentReturn

import (
	"time"

	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_02"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type Message struct {
	Data MessageModel
	Doc  pacs_004_001_02.Document
}

func Convert() {
	msg := Message{}
	msg.Data.MessageId = string(msg.Doc.PmtRtr.GrpHdr.MsgId)
	msg.Data.CreatedDateTime = time.Time(msg.Doc.PmtRtr.GrpHdr.CreDtTm)
	msg.Data.NumberOfTransactions = string(msg.Doc.PmtRtr.GrpHdr.NbOfTxs)
	msg.Data.SettlementMethod = Archive.SettlementMethodType(msg.Doc.PmtRtr.GrpHdr.SttlmInf.SttlmMtd)
	msg.Data.ClearingSystem = Archive.CommonClearingSysCodeType(*msg.Doc.PmtRtr.GrpHdr.SttlmInf.ClrSys.Cd)
	msg.Data.OriginalMessageId = string(msg.Doc.PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgId)
	msg.Data.OriginalMessageNameId = string(msg.Doc.PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgNmId)
	msg.Data.OriginalCreationDateTime = time.Time(*msg.Doc.PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlCreDtTm)
	msg.Data.OriginalInstructionId = string(*msg.Doc.PmtRtr.TxInf[0].OrgnlInstrId)
	msg.Data.OriginalEndToEndId = string(*msg.Doc.PmtRtr.TxInf[0].OrgnlEndToEndId)
	// msg.Data.OriginalUETR = string(*msg.Doc.PmtRtr.TxInf[0].OrgnlUETR)
	msg.Data.ReturnedInterbankSettlementAmount.Amount = float64(msg.Doc.PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Value)
	msg.Data.ReturnedInterbankSettlementAmount.Currency = string(msg.Doc.PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Ccy)
	msg.Data.InterbankSettlementDate = *msg.Doc.PmtRtr.TxInf[0].IntrBkSttlmDt
	msg.Data.ReturnedInstructedAmount.Amount = float64(msg.Doc.PmtRtr.TxInf[0].RtrdInstdAmt.Value)
	msg.Data.ReturnedInstructedAmount.Currency = string(msg.Doc.PmtRtr.TxInf[0].RtrdInstdAmt.Ccy)
	msg.Data.ChargeBearer = Archive.ChargeBearerType(*msg.Doc.PmtRtr.TxInf[0].ChrgBr)
	msg.Data.InstructingAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.InstructingAgent.PaymentSysMemberId = string(msg.Doc.PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.InstructedAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.InstructedAgent.PaymentSysMemberId = string(msg.Doc.PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId)
	// msg.Data.RtrChain.Debtor.Name = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.Nm)
	// msg.Data.RtrChain.Debtor.Address.StreetName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.StrtNm)
	// msg.Data.RtrChain.Debtor.Address.BuildingNumber = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.BldgNb)
	// msg.Data.RtrChain.Debtor.Address.BuildingName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.BldgNm)
	// msg.Data.RtrChain.Debtor.Address.Floor = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Flr)
	// msg.Data.RtrChain.Debtor.Address.RoomNumber = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Room)
	// msg.Data.RtrChain.Debtor.Address.PostalCode = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.PstCd)
	// msg.Data.RtrChain.Debtor.Address.TownName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.TwnNm)
	// msg.Data.RtrChain.Debtor.Address.Subdivision = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.CtrySubDvsn)
	// msg.Data.RtrChain.Debtor.Address.Country = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Ctry)
	// msg.Data.RtrChain.Creditor.Name = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.Nm)
	// msg.Data.RtrChain.Creditor.Address.StreetName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.StrtNm)
	// msg.Data.RtrChain.Creditor.Address.BuildingNumber = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.BldgNb)
	// msg.Data.RtrChain.Creditor.Address.BuildingName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.BldgNm)
	// msg.Data.RtrChain.Creditor.Address.Floor = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Flr)
	// msg.Data.RtrChain.Creditor.Address.RoomNumber = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Room)
	// msg.Data.RtrChain.Creditor.Address.PostalCode = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.PstCd)
	// msg.Data.RtrChain.Creditor.Address.TownName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.TwnNm)
	// msg.Data.RtrChain.Creditor.Address.Subdivision = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.CtrySubDvsn)
	// msg.Data.RtrChain.Creditor.Address.Country = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Ctry)
	// msg.Data.RtrChain.DebtorOtherTypeId = string(msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAcct.Id.Othr.Id)
	// msg.Data.RtrChain.DebtorAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	// msg.Data.RtrChain.DebtorAgent.PaymentSysMemberId = string(msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId)
	// msg.Data.RtrChain.DebtorAgent.BankName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.Nm)
	// msg.Data.RtrChain.DebtorAgent.PostalAddress.StreetName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.StrtNm)
	// msg.Data.RtrChain.DebtorAgent.PostalAddress.BuildingNumber = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.BldgNb)
	// msg.Data.RtrChain.DebtorAgent.PostalAddress.BuildingName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.BldgNm)
	// msg.Data.RtrChain.DebtorAgent.PostalAddress.Floor = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Flr)
	// msg.Data.RtrChain.DebtorAgent.PostalAddress.RoomNumber = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Room)
	// msg.Data.RtrChain.DebtorAgent.PostalAddress.PostalCode = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.PstCd)
	// msg.Data.RtrChain.DebtorAgent.PostalAddress.TownName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.TwnNm)
	// msg.Data.RtrChain.DebtorAgent.PostalAddress.Subdivision = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn)
	// msg.Data.RtrChain.DebtorAgent.PostalAddress.Country = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Ctry)
	// msg.Data.RtrChain.CreditorAccountOtherTypeId = string(msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAcct.Id.Othr.Id)
	// msg.Data.RtrChain.CreditorAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	// msg.Data.RtrChain.CreditorAgent.PaymentSysMemberId = string(msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId)
	// msg.Data.RtrChain.CreditorAgent.BankName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.Nm)
	// msg.Data.RtrChain.CreditorAgent.PostalAddress.StreetName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.StrtNm)
	// msg.Data.RtrChain.CreditorAgent.PostalAddress.BuildingNumber = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.BldgNb)
	// msg.Data.RtrChain.CreditorAgent.PostalAddress.BuildingName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.BldgNm)
	// msg.Data.RtrChain.CreditorAgent.PostalAddress.Floor = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Flr)
	// msg.Data.RtrChain.CreditorAgent.PostalAddress.RoomNumber = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Room)
	// msg.Data.RtrChain.CreditorAgent.PostalAddress.PostalCode = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.PstCd)
	// msg.Data.RtrChain.CreditorAgent.PostalAddress.TownName = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.TwnNm)
	// msg.Data.RtrChain.CreditorAgent.PostalAddress.Subdivision = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn)
	// msg.Data.RtrChain.CreditorAgent.PostalAddress.Country = string(*msg.Doc.PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Ctry)
	msg.Data.ReturnReasonInformation.Reason = string(*msg.Doc.PmtRtr.TxInf[0].RtrRsnInf[0].Rsn.Cd)
	msg.Data.ReturnReasonInformation.AdditionalInfo = string(*msg.Doc.PmtRtr.TxInf[0].RtrRsnInf[0].AddtlInf[0])
	msg.Data.OriginalTransactionRef = Archive.InstrumentPropCodeType(*msg.Doc.PmtRtr.TxInf[0].OrgnlTxRef.PmtTpInf.LclInstrm.Prtry)
}

func PathMapV2() map[string]string {
	return PathMapV7()
}
func PathMapV3() map[string]string {
	return PathMapV7()
}
func PathMapV4() map[string]string {
	return PathMapV7()
}
func PathMapV5() map[string]string {
	return PathMapV7()
}
func PathMapV6() map[string]string {
	return PathMapV7()
}
func PathMapV7() map[string]string {
	return map[string]string{
		"PmtRtr.GrpHdr.MsgId":                                         "MessageId",
		"PmtRtr.GrpHdr.CreDtTm":                                       "CreatedDateTime",
		"PmtRtr.GrpHdr.NbOfTxs":                                       "NumberOfTransactions",
		"PmtRtr.GrpHdr.SttlmInf.SttlmMtd":                             "SettlementMethod",
		"PmtRtr.GrpHdr.SttlmInf.ClrSys.Cd":                            "ClearingSystem",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgId":                      "OriginalMessageId",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgNmId":                    "OriginalMessageNameId",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlCreDtTm":                    "OriginalCreationDateTime",
		"PmtRtr.TxInf[0].OrgnlInstrId":                                "OriginalInstructionId",
		"PmtRtr.TxInf[0].OrgnlEndToEndId":                             "OriginalEndToEndId",
		"PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Value":                    "ReturnedInterbankSettlementAmount.Amount",
		"PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Ccy":                      "ReturnedInterbankSettlementAmount.Currency",
		"PmtRtr.TxInf[0].IntrBkSttlmDt":                               "InterbankSettlementDate",
		"PmtRtr.TxInf[0].RtrdInstdAmt.Value":                          "ReturnedInstructedAmount.Amount",
		"PmtRtr.TxInf[0].RtrdInstdAmt.Ccy":                            "ReturnedInstructedAmount.Currency",
		"PmtRtr.TxInf[0].ChrgBr":                                      "ChargeBearer",
		"PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructingAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructingAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructedAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructedAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrRsnInf[0].Rsn.Cd":                         "ReturnReasonInformation.Reason",
		"PmtRtr.TxInf[0].RtrRsnInf[0].AddtlInf[0]":                    "ReturnReasonInformation.AdditionalInfo",
		"PmtRtr.TxInf[0].OrgnlTxRef.PmtTpInf.LclInstrm.Prtry":         "OriginalTransactionRef",
	}
}
func PathMapV8() map[string]string {
	return map[string]string{
		"PmtRtr.GrpHdr.MsgId":                                                 "MessageId",
		"PmtRtr.GrpHdr.CreDtTm":                                               "CreatedDateTime",
		"PmtRtr.GrpHdr.NbOfTxs":                                               "NumberOfTransactions",
		"PmtRtr.GrpHdr.SttlmInf.SttlmMtd":                                     "SettlementMethod",
		"PmtRtr.GrpHdr.SttlmInf.ClrSys.Cd":                                    "ClearingSystem",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgId":                              "OriginalMessageId",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgNmId":                            "OriginalMessageNameId",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlCreDtTm":                            "OriginalCreationDateTime",
		"PmtRtr.TxInf[0].OrgnlInstrId":                                        "OriginalInstructionId",
		"PmtRtr.TxInf[0].OrgnlEndToEndId":                                     "OriginalEndToEndId",
		"PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Value":                            "ReturnedInterbankSettlementAmount.Amount",
		"PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Ccy":                              "ReturnedInterbankSettlementAmount.Currency",
		"PmtRtr.TxInf[0].IntrBkSttlmDt":                                       "InterbankSettlementDate",
		"PmtRtr.TxInf[0].RtrdInstdAmt.Value":                                  "ReturnedInstructedAmount.Amount",
		"PmtRtr.TxInf[0].RtrdInstdAmt.Ccy":                                    "ReturnedInstructedAmount.Currency",
		"PmtRtr.TxInf[0].ChrgBr":                                              "ChargeBearer",
		"PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":         "InstructingAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":               "InstructingAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":         "InstructedAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":               "InstructedAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.Nm":                                "RtrChain.Debtor.Name",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.StrtNm":                    "RtrChain.Debtor.Address.StreetName",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.BldgNb":                    "RtrChain.Debtor.Address.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.PstCd":                     "RtrChain.Debtor.Address.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.TwnNm":                     "RtrChain.Debtor.Address.TownName",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.CtrySubDvsn":               "RtrChain.Debtor.Address.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Ctry":                      "RtrChain.Debtor.Address.Country",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.Nm":                                "RtrChain.Creditor.Name",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.StrtNm":                    "RtrChain.Creditor.Address.StreetName",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.BldgNb":                    "RtrChain.Creditor.Address.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.PstCd":                     "RtrChain.Creditor.Address.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.TwnNm":                     "RtrChain.Creditor.Address.TownName",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.CtrySubDvsn":               "RtrChain.Creditor.Address.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Ctry":                      "RtrChain.Creditor.Address.Country",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "RtrChain.DebtorAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":       "RtrChain.DebtorAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.Nm":                      "RtrChain.DebtorAgent.BankName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.StrtNm":          "RtrChain.DebtorAgent.PostalAddress.StreetName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.BldgNb":          "RtrChain.DebtorAgent.PostalAddress.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.PstCd":           "RtrChain.DebtorAgent.PostalAddress.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.TwnNm":           "RtrChain.DebtorAgent.PostalAddress.TownName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn":     "RtrChain.DebtorAgent.PostalAddress.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Ctry":            "RtrChain.DebtorAgent.PostalAddress.Country",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "RtrChain.CreditorAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":       "RtrChain.CreditorAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.Nm":                      "RtrChain.CreditorAgent.BankName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.StrtNm":          "RtrChain.CreditorAgent.PostalAddress.StreetName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.BldgNb":          "RtrChain.CreditorAgent.PostalAddress.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.PstCd":           "RtrChain.CreditorAgent.PostalAddress.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.TwnNm":           "RtrChain.CreditorAgent.PostalAddress.TownName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn":     "RtrChain.CreditorAgent.PostalAddress.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Ctry":            "RtrChain.CreditorAgent.PostalAddress.Country",
		"PmtRtr.TxInf[0].RtrRsnInf[0].Rsn.Cd":                                 "ReturnReasonInformation.Reason",
		"PmtRtr.TxInf[0].RtrRsnInf[0].AddtlInf[0]":                            "ReturnReasonInformation.AdditionalInfo",
		"PmtRtr.TxInf[0].OrgnlTxRef.PmtTpInf.LclInstrm.Prtry":                 "OriginalTransactionRef",
	}
}
func PathMapV9() map[string]string {
	return map[string]string{
		"PmtRtr.GrpHdr.MsgId":                                                 "MessageId",
		"PmtRtr.GrpHdr.CreDtTm":                                               "CreatedDateTime",
		"PmtRtr.GrpHdr.NbOfTxs":                                               "NumberOfTransactions",
		"PmtRtr.GrpHdr.SttlmInf.SttlmMtd":                                     "SettlementMethod",
		"PmtRtr.GrpHdr.SttlmInf.ClrSys.Cd":                                    "ClearingSystem",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgId":                              "OriginalMessageId",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgNmId":                            "OriginalMessageNameId",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlCreDtTm":                            "OriginalCreationDateTime",
		"PmtRtr.TxInf[0].OrgnlInstrId":                                        "OriginalInstructionId",
		"PmtRtr.TxInf[0].OrgnlEndToEndId":                                     "OriginalEndToEndId",
		"PmtRtr.TxInf[0].OrgnlUETR":                                           "OriginalUETR",
		"PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Value":                            "ReturnedInterbankSettlementAmount.Amount",
		"PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Ccy":                              "ReturnedInterbankSettlementAmount.Currency",
		"PmtRtr.TxInf[0].IntrBkSttlmDt":                                       "InterbankSettlementDate",
		"PmtRtr.TxInf[0].RtrdInstdAmt.Value":                                  "ReturnedInstructedAmount.Amount",
		"PmtRtr.TxInf[0].RtrdInstdAmt.Ccy":                                    "ReturnedInstructedAmount.Currency",
		"PmtRtr.TxInf[0].ChrgBr":                                              "ChargeBearer",
		"PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":         "InstructingAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":               "InstructingAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":         "InstructedAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":               "InstructedAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.Nm":                                "RtrChain.Debtor.Name",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.StrtNm":                    "RtrChain.Debtor.Address.StreetName",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.BldgNb":                    "RtrChain.Debtor.Address.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.BldgNm":                    "RtrChain.Debtor.Address.BuildingName",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Flr":                       "RtrChain.Debtor.Address.Floor",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Room":                      "RtrChain.Debtor.Address.RoomNumber",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.PstCd":                     "RtrChain.Debtor.Address.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.TwnNm":                     "RtrChain.Debtor.Address.TownName",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.CtrySubDvsn":               "RtrChain.Debtor.Address.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Ctry":                      "RtrChain.Debtor.Address.Country",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.Nm":                                "RtrChain.Creditor.Name",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.StrtNm":                    "RtrChain.Creditor.Address.StreetName",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.BldgNb":                    "RtrChain.Creditor.Address.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.BldgNm":                    "RtrChain.Creditor.Address.BuildingName",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Flr":                       "RtrChain.Creditor.Address.Floor",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Room":                      "RtrChain.Creditor.Address.RoomNumber",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.PstCd":                     "RtrChain.Creditor.Address.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.TwnNm":                     "RtrChain.Creditor.Address.TownName",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.CtrySubDvsn":               "RtrChain.Creditor.Address.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Ctry":                      "RtrChain.Creditor.Address.Country",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "RtrChain.DebtorAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":       "RtrChain.DebtorAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.Nm":                      "RtrChain.DebtorAgent.BankName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.StrtNm":          "RtrChain.DebtorAgent.PostalAddress.StreetName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.BldgNb":          "RtrChain.DebtorAgent.PostalAddress.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.BldgNm":          "RtrChain.DebtorAgent.PostalAddress.BuildingName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Flr":             "RtrChain.DebtorAgent.PostalAddress.Floor",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Room":            "RtrChain.DebtorAgent.PostalAddress.RoomNumber",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.PstCd":           "RtrChain.DebtorAgent.PostalAddress.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.TwnNm":           "RtrChain.DebtorAgent.PostalAddress.TownName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn":     "RtrChain.DebtorAgent.PostalAddress.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Ctry":            "RtrChain.DebtorAgent.PostalAddress.Country",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "RtrChain.CreditorAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":       "RtrChain.CreditorAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.Nm":                      "RtrChain.CreditorAgent.BankName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.StrtNm":          "RtrChain.CreditorAgent.PostalAddress.StreetName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.BldgNb":          "RtrChain.CreditorAgent.PostalAddress.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.BldgNm":          "RtrChain.CreditorAgent.PostalAddress.BuildingName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Flr":             "RtrChain.CreditorAgent.PostalAddress.Floor",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Room":            "RtrChain.CreditorAgent.PostalAddress.RoomNumber",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.PstCd":           "RtrChain.CreditorAgent.PostalAddress.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.TwnNm":           "RtrChain.CreditorAgent.PostalAddress.TownName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn":     "RtrChain.CreditorAgent.PostalAddress.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Ctry":            "RtrChain.CreditorAgent.PostalAddress.Country",
		"PmtRtr.TxInf[0].RtrRsnInf[0].Rsn.Cd":                                 "ReturnReasonInformation.Reason",
		"PmtRtr.TxInf[0].RtrRsnInf[0].AddtlInf[0]":                            "ReturnReasonInformation.AdditionalInfo",
		"PmtRtr.TxInf[0].OrgnlTxRef.PmtTpInf.LclInstrm.Prtry":                 "OriginalTransactionRef",
	}
}
func PathMapV10() map[string]string {
	return map[string]string{
		"PmtRtr.GrpHdr.MsgId":                                                 "MessageId",
		"PmtRtr.GrpHdr.CreDtTm":                                               "CreatedDateTime",
		"PmtRtr.GrpHdr.NbOfTxs":                                               "NumberOfTransactions",
		"PmtRtr.GrpHdr.SttlmInf.SttlmMtd":                                     "SettlementMethod",
		"PmtRtr.GrpHdr.SttlmInf.ClrSys.Cd":                                    "ClearingSystem",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgId":                              "OriginalMessageId",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlMsgNmId":                            "OriginalMessageNameId",
		"PmtRtr.TxInf[0].OrgnlGrpInf.OrgnlCreDtTm":                            "OriginalCreationDateTime",
		"PmtRtr.TxInf[0].OrgnlInstrId":                                        "OriginalInstructionId",
		"PmtRtr.TxInf[0].OrgnlEndToEndId":                                     "OriginalEndToEndId",
		"PmtRtr.TxInf[0].OrgnlUETR":                                           "OriginalUETR",
		"PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Value":                            "ReturnedInterbankSettlementAmount.Amount",
		"PmtRtr.TxInf[0].RtrdIntrBkSttlmAmt.Ccy":                              "ReturnedInterbankSettlementAmount.Currency",
		"PmtRtr.TxInf[0].IntrBkSttlmDt":                                       "InterbankSettlementDate",
		"PmtRtr.TxInf[0].RtrdInstdAmt.Value":                                  "ReturnedInstructedAmount.Amount",
		"PmtRtr.TxInf[0].RtrdInstdAmt.Ccy":                                    "ReturnedInstructedAmount.Currency",
		"PmtRtr.TxInf[0].ChrgBr":                                              "ChargeBearer",
		"PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":         "InstructingAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":               "InstructingAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":         "InstructedAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":               "InstructedAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.Nm":                                "RtrChain.Debtor.Name",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.StrtNm":                    "RtrChain.Debtor.Address.StreetName",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.BldgNb":                    "RtrChain.Debtor.Address.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.BldgNm":                    "RtrChain.Debtor.Address.BuildingName",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Flr":                       "RtrChain.Debtor.Address.Floor",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Room":                      "RtrChain.Debtor.Address.RoomNumber",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.PstCd":                     "RtrChain.Debtor.Address.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.TwnNm":                     "RtrChain.Debtor.Address.TownName",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.CtrySubDvsn":               "RtrChain.Debtor.Address.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.Dbtr.Pty.PstlAdr.Ctry":                      "RtrChain.Debtor.Address.Country",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.Nm":                                "RtrChain.Creditor.Name",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.StrtNm":                    "RtrChain.Creditor.Address.StreetName",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.BldgNb":                    "RtrChain.Creditor.Address.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.BldgNm":                    "RtrChain.Creditor.Address.BuildingName",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Flr":                       "RtrChain.Creditor.Address.Floor",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Room":                      "RtrChain.Creditor.Address.RoomNumber",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.PstCd":                     "RtrChain.Creditor.Address.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.TwnNm":                     "RtrChain.Creditor.Address.TownName",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.CtrySubDvsn":               "RtrChain.Creditor.Address.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.Cdtr.Pty.PstlAdr.Ctry":                      "RtrChain.Creditor.Address.Country",
		"PmtRtr.TxInf[0].RtrChain.DbtrAcct.Id.Othr.Id":                        "RtrChain.DebtorOtherTypeId",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "RtrChain.DebtorAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":       "RtrChain.DebtorAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.Nm":                      "RtrChain.DebtorAgent.BankName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.StrtNm":          "RtrChain.DebtorAgent.PostalAddress.StreetName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.BldgNb":          "RtrChain.DebtorAgent.PostalAddress.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.BldgNm":          "RtrChain.DebtorAgent.PostalAddress.BuildingName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Flr":             "RtrChain.DebtorAgent.PostalAddress.Floor",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Room":            "RtrChain.DebtorAgent.PostalAddress.RoomNumber",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.PstCd":           "RtrChain.DebtorAgent.PostalAddress.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.TwnNm":           "RtrChain.DebtorAgent.PostalAddress.TownName",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn":     "RtrChain.DebtorAgent.PostalAddress.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.DbtrAgt.FinInstnId.PstlAdr.Ctry":            "RtrChain.DebtorAgent.PostalAddress.Country",
		"PmtRtr.TxInf[0].RtrChain.CdtrAcct.Id.Othr.Id":                        "RtrChain.CreditorAccountOtherTypeId",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "RtrChain.CreditorAgent.PaymentSysCode",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":       "RtrChain.CreditorAgent.PaymentSysMemberId",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.Nm":                      "RtrChain.CreditorAgent.BankName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.StrtNm":          "RtrChain.CreditorAgent.PostalAddress.StreetName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.BldgNb":          "RtrChain.CreditorAgent.PostalAddress.BuildingNumber",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.BldgNm":          "RtrChain.CreditorAgent.PostalAddress.BuildingName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Flr":             "RtrChain.CreditorAgent.PostalAddress.Floor",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Room":            "RtrChain.CreditorAgent.PostalAddress.RoomNumber",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.PstCd":           "RtrChain.CreditorAgent.PostalAddress.PostalCode",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.TwnNm":           "RtrChain.CreditorAgent.PostalAddress.TownName",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn":     "RtrChain.CreditorAgent.PostalAddress.Subdivision",
		"PmtRtr.TxInf[0].RtrChain.CdtrAgt.FinInstnId.PstlAdr.Ctry":            "RtrChain.CreditorAgent.PostalAddress.Country",
		"PmtRtr.TxInf[0].RtrRsnInf[0].Rsn.Cd":                                 "ReturnReasonInformation.Reason",
		"PmtRtr.TxInf[0].RtrRsnInf[0].AddtlInf[0]":                            "ReturnReasonInformation.AdditionalInfo",
		"PmtRtr.TxInf[0].OrgnlTxRef.PmtTpInf.LclInstrm.Prtry":                 "OriginalTransactionRef",
	}
}
func PathMapV11() map[string]string {
	return PathMapV10()
}
func PathMapV12() map[string]string {
	return PathMapV10()
}
func PathMapV13() map[string]string {
	return PathMapV10()
}
