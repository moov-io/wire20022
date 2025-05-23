package CustomerCreditTransfer

import (
	"time"

	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_02"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type Message struct {
	Data   MessageModel
	Doc    pacs_008_001_02.Document
}
func Convert() {
	msg := Message{}
	msg.Data.MessageId = string(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.MsgId)
	msg.Data.CreatedDateTime = time.Time(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.CreDtTm)
	msg.Data.NumberOfTransactions = string(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.NbOfTxs)
	msg.Data.SettlementMethod = Archive.SettlementMethodType(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.SttlmInf.SttlmMtd)
	msg.Data.CommonClearingSysCode = Archive.CommonClearingSysCodeType(*msg.Doc.FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd)
	msg.Data.InstructionId = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.InstrId)
	msg.Data.EndToEndId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.EndToEndId)
	msg.Data.TaxId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.TxId)
	msg.Data.SericeLevel = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.SvcLvl.Cd)
	msg.Data.InstrumentPropCode = Archive.InstrumentPropCodeType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.LclInstrm.Prtry)
	msg.Data.InterBankSettAmount.Amount = float64(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Value)
	msg.Data.InterBankSettAmount.Currency = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Ccy)
	msg.Data.InterBankSettDate = *msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmDt
	msg.Data.InstructedAmount.Amount = float64(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Value)
	msg.Data.InstructedAmount.Currency = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Ccy)
	msg.Data.ExchangeRate = float64(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].XchgRate)
	msg.Data.ChargeBearer = Archive.ChargeBearerType(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgBr)

	///msg.Data.ChargesInfo[] : msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.ChrgsInf[]
	msg.Data.ChargesInfo[0].Amount.Amount = float64(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgsInf[0].Amt.Value)
	msg.Data.ChargesInfo[0].Amount.Currency = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgsInf[0].Amt.Ccy)

	msg.Data.InstructingAgents.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.InstructingAgents.PaymentSysMemberId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.InstructedAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.InstructedAgent.PaymentSysMemberId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.IntermediaryAgent1Id = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.DebtorName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.Nm)
	msg.Data.DebtorAddress.StreetName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.StrtNm)
	msg.Data.DebtorAddress.BuildingNumber = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.BldgNb)
	msg.Data.DebtorAddress.PostalCode = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.PstCd)
	msg.Data.DebtorAddress.TownName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.TwnNm)
	msg.Data.DebtorAddress.Subdivision = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.CtrySubDvsn)
	msg.Data.DebtorAddress.Country = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Ctry)
	msg.Data.DebtorIBAN = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.IBAN)
	msg.Data.DebtorOtherTypeId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.Othr.Id)
	msg.Data.DebtorAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.DebtorAgent.PaymentSysMemberId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.DebtorAgent.BankName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.Nm)
	msg.Data.DebtorAgent.PostalAddress.StreetName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.StrtNm)
	msg.Data.DebtorAgent.PostalAddress.BuildingNumber = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.BldgNb)
	msg.Data.DebtorAgent.PostalAddress.PostalCode = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.PstCd)
	msg.Data.DebtorAgent.PostalAddress.TownName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.TwnNm)
	msg.Data.DebtorAgent.PostalAddress.Subdivision = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn)
	msg.Data.DebtorAgent.PostalAddress.Country = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Ctry)
	msg.Data.CreditorAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.CreditorAgent.PaymentSysMemberId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.CreditorAgent.BankName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.Nm)
	msg.Data.CreditorAgent.PostalAddress.StreetName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.StrtNm)
	msg.Data.CreditorAgent.PostalAddress.BuildingNumber = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.BldgNb)
	msg.Data.CreditorAgent.PostalAddress.PostalCode = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.PstCd)
	msg.Data.CreditorAgent.PostalAddress.TownName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.TwnNm)
	msg.Data.CreditorAgent.PostalAddress.Subdivision = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn)
	msg.Data.CreditorAgent.PostalAddress.Country = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Ctry)
	msg.Data.CreditorName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.Nm)
	msg.Data.CreditorPostalAddress.StreetName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.StrtNm)
	msg.Data.CreditorPostalAddress.BuildingNumber = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.BldgNb)
	msg.Data.CreditorPostalAddress.PostalCode = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.PstCd)
	msg.Data.CreditorPostalAddress.TownName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.TwnNm)
	msg.Data.CreditorPostalAddress.Subdivision = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.CtrySubDvsn)
	msg.Data.CreditorPostalAddress.Country = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Ctry)
	msg.Data.UltimateCreditorName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.Nm)
	msg.Data.UltimateCreditorAddress.StreetName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.StrtNm)
	msg.Data.UltimateCreditorAddress.BuildingNumber = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.BldgNb)
	msg.Data.UltimateCreditorAddress.PostalCode = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.PstCd)
	msg.Data.UltimateCreditorAddress.TownName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.TwnNm)
	msg.Data.UltimateCreditorAddress.Subdivision = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.CtrySubDvsn)
	msg.Data.UltimateCreditorAddress.Country = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Ctry)
	msg.Data.CreditorIBAN = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.IBAN)
	msg.Data.CreditorOtherTypeId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.Othr.Id)
	msg.Data.PurposeOfPayment = Archive.PurposeOfPaymentType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].Purp.Cd)
	msg.Data.RelatedRemittanceInfo.RemittanceId = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtId)
	msg.Data.RelatedRemittanceInfo.Method = Archive.RemittanceDeliveryMethod(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnMtd)
	msg.Data.RelatedRemittanceInfo.ElectronicAddress = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnElctrncAdr)
	msg.Data.RemittanceInfor.UnstructuredRemitInfo = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Ustrd[0])
	msg.Data.RemittanceInfor.CodeOrProprietary = Archive.CodeOrProprietaryType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd)
	msg.Data.RemittanceInfor.Number = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Nb)
	msg.Data.RemittanceInfor.RelatedDate = *msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt
}

func PathMapV2() map[string]any {
	return map[string]any{
		"FIToFICstmrCdtTrf.GrpHdr.MsgId" : "MessageId",
		"FIToFICstmrCdtTrf.GrpHdr.CreDtTm" : "CreatedDateTime",
		"FIToFICstmrCdtTrf.GrpHdr.NbOfTxs" : "NumberOfTransactions",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.SttlmMtd" : "SettlementMethod",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd" : "CommonClearingSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.InstrId" : "InstructionId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.EndToEndId" : "EndToEndId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.TxId" : "TaxId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.SvcLvl.Cd" : "SericeLevel",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.LclInstrm.Prtry" : "InstrumentPropCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Value" : "InterBankSettAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Ccy" : "InterBankSettAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmDt" : "InterBankSettDate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Value" : "InstructedAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Ccy" : "InstructedAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].XchgRate" : "ExchangeRate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgBr" : "ChargeBearer",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgsInf : ChargesInfo": map[string]string{
			"Amt.Value": "Amount.Amount",
			"Amt.Ccy":   "Amount.Currency",
		},
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructingAgents.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructingAgents.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructedAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructedAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId" : "IntermediaryAgent1Id",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.Nm" : "DebtorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.StrtNm" : "DebtorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.BldgNb" : "DebtorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.PstCd" : "DebtorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.TwnNm" : "DebtorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.CtrySubDvsn" : "DebtorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Ctry" : "DebtorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.IBAN" : "DebtorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.Othr.Id" : "DebtorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "DebtorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "DebtorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.Nm" : "DebtorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.StrtNm" : "DebtorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.BldgNb" : "DebtorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.PstCd" : "DebtorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.TwnNm" : "DebtorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "DebtorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Ctry" : "DebtorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "CreditorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "CreditorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.Nm" : "CreditorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.StrtNm" : "CreditorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.BldgNb" : "CreditorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.PstCd" : "CreditorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.TwnNm" : "CreditorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "CreditorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Ctry" : "CreditorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.Nm" : "CreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.StrtNm" : "CreditorPostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.BldgNb" : "CreditorPostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.PstCd" : "CreditorPostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.TwnNm" : "CreditorPostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.CtrySubDvsn" : "CreditorPostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Ctry" : "CreditorPostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.Nm" : "UltimateCreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.StrtNm" : "UltimateCreditorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.BldgNb" : "UltimateCreditorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.PstCd" : "UltimateCreditorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.TwnNm" : "UltimateCreditorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.CtrySubDvsn" : "UltimateCreditorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Ctry" : "UltimateCreditorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.IBAN" : "CreditorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.Othr.Id" : "CreditorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Purp.Cd" : "PurposeOfPayment",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtId" : "RelatedRemittanceInfo.RemittanceId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnMtd" : "RelatedRemittanceInfo.Method",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnElctrncAdr" : "RelatedRemittanceInfo.ElectronicAddress",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Ustrd[0]" : "RemittanceInfor.UnstructuredRemitInfo",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd" : "RemittanceInfor.CodeOrProprietary",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Nb" : "RemittanceInfor.Number",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt" : "RemittanceInfor.RelatedDate",
	}
}
func PathMapV3() map[string]any {
	return PathMapV4()
}
func PathMapV4() map[string]any {
	return map[string]any{
		"FIToFICstmrCdtTrf.GrpHdr.MsgId" : "MessageId",
		"FIToFICstmrCdtTrf.GrpHdr.CreDtTm" : "CreatedDateTime",
		"FIToFICstmrCdtTrf.GrpHdr.NbOfTxs" : "NumberOfTransactions",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.SttlmMtd" : "SettlementMethod",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd" : "CommonClearingSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.InstrId" : "InstructionId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.EndToEndId" : "EndToEndId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.TxId" : "TaxId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.SvcLvl.Cd" : "SericeLevel",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.LclInstrm.Prtry" : "InstrumentPropCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Value" : "InterBankSettAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Ccy" : "InterBankSettAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmDt" : "InterBankSettDate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Value" : "InstructedAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Ccy" : "InstructedAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].XchgRate" : "ExchangeRate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgBr" : "ChargeBearer",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgsInf : ChargesInfo": map[string]string{
			"Amt.Value": "Amount.Amount",
			"Amt.Ccy":   "Amount.Currency",
			"Agt.FinInstnId.BICFI": "BusinessIdCode",
		},
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructingAgents.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructingAgents.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructedAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructedAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId" : "IntermediaryAgent1Id",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.Nm" : "DebtorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.StrtNm" : "DebtorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.BldgNb" : "DebtorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.PstCd" : "DebtorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.TwnNm" : "DebtorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.CtrySubDvsn" : "DebtorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Ctry" : "DebtorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.IBAN" : "DebtorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.Othr.Id" : "DebtorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "DebtorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "DebtorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.Nm" : "DebtorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.StrtNm" : "DebtorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.BldgNb" : "DebtorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.PstCd" : "DebtorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.TwnNm" : "DebtorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "DebtorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Ctry" : "DebtorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "CreditorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "CreditorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.Nm" : "CreditorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.StrtNm" : "CreditorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.BldgNb" : "CreditorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.PstCd" : "CreditorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.TwnNm" : "CreditorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "CreditorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Ctry" : "CreditorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.Nm" : "CreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.StrtNm" : "CreditorPostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.BldgNb" : "CreditorPostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.PstCd" : "CreditorPostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.TwnNm" : "CreditorPostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.CtrySubDvsn" : "CreditorPostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Ctry" : "CreditorPostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.Nm" : "UltimateCreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.StrtNm" : "UltimateCreditorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.BldgNb" : "UltimateCreditorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.PstCd" : "UltimateCreditorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.TwnNm" : "UltimateCreditorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.CtrySubDvsn" : "UltimateCreditorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Ctry" : "UltimateCreditorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.IBAN" : "CreditorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.Othr.Id" : "CreditorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Purp.Cd" : "PurposeOfPayment",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtId" : "RelatedRemittanceInfo.RemittanceId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnMtd" : "RelatedRemittanceInfo.Method",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnElctrncAdr" : "RelatedRemittanceInfo.ElectronicAddress",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Ustrd[0]" : "RemittanceInfor.UnstructuredRemitInfo",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd" : "RemittanceInfor.CodeOrProprietary",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Nb" : "RemittanceInfor.Number",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt" : "RemittanceInfor.RelatedDate",
	}
}
func PathMapV5() map[string]any {
	return PathMapV7()
}
func PathMapV6() map[string]any {
	return PathMapV7()
}
func PathMapV7() map[string]any {
	return map[string]any{
		"FIToFICstmrCdtTrf.GrpHdr.MsgId" : "MessageId",
		"FIToFICstmrCdtTrf.GrpHdr.CreDtTm" : "CreatedDateTime",
		"FIToFICstmrCdtTrf.GrpHdr.NbOfTxs" : "NumberOfTransactions",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.SttlmMtd" : "SettlementMethod",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd" : "CommonClearingSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.InstrId" : "InstructionId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.EndToEndId" : "EndToEndId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.TxId" : "TaxId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.SvcLvl.Cd" : "SericeLevel",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.LclInstrm.Prtry" : "InstrumentPropCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Value" : "InterBankSettAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Ccy" : "InterBankSettAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmDt" : "InterBankSettDate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Value" : "InstructedAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Ccy" : "InstructedAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].XchgRate" : "ExchangeRate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgBr" : "ChargeBearer",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgsInf : ChargesInfo": map[string]string{
			"Amt.Value": "Amount.Amount",
			"Amt.Ccy":   "Amount.Currency",
			"Agt.FinInstnId.BICFI": "BusinessIdCode",
		},
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructingAgents.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructingAgents.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructedAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructedAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId" : "IntermediaryAgent1Id",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.Nm" : "DebtorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.StrtNm" : "DebtorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.BldgNb" : "DebtorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.PstCd" : "DebtorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.TwnNm" : "DebtorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.CtrySubDvsn" : "DebtorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Ctry" : "DebtorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.IBAN" : "DebtorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.Othr.Id" : "DebtorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "DebtorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "DebtorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.Nm" : "DebtorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.StrtNm" : "DebtorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.BldgNb" : "DebtorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.PstCd" : "DebtorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.TwnNm" : "DebtorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "DebtorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Ctry" : "DebtorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "CreditorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "CreditorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.Nm" : "CreditorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.StrtNm" : "CreditorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.BldgNb" : "CreditorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.PstCd" : "CreditorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.TwnNm" : "CreditorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "CreditorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Ctry" : "CreditorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.Nm" : "CreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.StrtNm" : "CreditorPostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.BldgNb" : "CreditorPostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.PstCd" : "CreditorPostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.TwnNm" : "CreditorPostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.CtrySubDvsn" : "CreditorPostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Ctry" : "CreditorPostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.Nm" : "UltimateCreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.StrtNm" : "UltimateCreditorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.BldgNb" : "UltimateCreditorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.PstCd" : "UltimateCreditorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.TwnNm" : "UltimateCreditorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.CtrySubDvsn" : "UltimateCreditorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Ctry" : "UltimateCreditorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.IBAN" : "CreditorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.Othr.Id" : "CreditorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Purp.Cd" : "PurposeOfPayment",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtId" : "RelatedRemittanceInfo.RemittanceId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnDtls[0].Mtd" : "RelatedRemittanceInfo.Method",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnDtls[0].ElctrncAdr" : "RelatedRemittanceInfo.ElectronicAddress",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Ustrd[0]" : "RemittanceInfor.UnstructuredRemitInfo",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd" : "RemittanceInfor.CodeOrProprietary",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Nb" : "RemittanceInfor.Number",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt" : "RemittanceInfor.RelatedDate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Cdtr.TaxId" : "RemittanceInfor.TaxDetail.TaxId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Tp" : "RemittanceInfor.TaxDetail.TaxTypeCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd.Yr" : "RemittanceInfor.TaxDetail.TaxPeriodYear",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd.Tp" : "RemittanceInfor.TaxDetail.TaxperiodTimeFrame",
	}
}
func PathMapV8() map[string]any {
	return map[string]any{
		"FIToFICstmrCdtTrf.GrpHdr.MsgId" : "MessageId",
		"FIToFICstmrCdtTrf.GrpHdr.CreDtTm" : "CreatedDateTime",
		"FIToFICstmrCdtTrf.GrpHdr.NbOfTxs" : "NumberOfTransactions",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.SttlmMtd" : "SettlementMethod",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd" : "CommonClearingSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.InstrId" : "InstructionId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.EndToEndId" : "EndToEndId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.UETR" : "UniqueEndToEndTransactionRef",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.SvcLvl[0].Cd" : "SericeLevel",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.LclInstrm.Prtry" : "InstrumentPropCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Value" : "InterBankSettAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Ccy" : "InterBankSettAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmDt" : "InterBankSettDate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Value" : "InstructedAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Ccy" : "InstructedAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].XchgRate" : "ExchangeRate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgBr" : "ChargeBearer",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgsInf : ChargesInfo": map[string]string{
			"Amt.Value": "Amount.Amount",
			"Amt.Ccy":   "Amount.Currency",
			"Agt.FinInstnId.BICFI": "BusinessIdCode",
		},
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructingAgents.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructingAgents.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructedAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructedAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId" : "IntermediaryAgent1Id",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.Nm" : "DebtorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.StrtNm" : "DebtorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.BldgNb" : "DebtorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.BldgNm" : "DebtorAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Flr" : "DebtorAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Room" : "DebtorAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.PstCd" : "DebtorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.TwnNm" : "DebtorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.CtrySubDvsn" : "DebtorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Ctry" : "DebtorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.IBAN" : "DebtorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.Othr.Id" : "DebtorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "DebtorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "DebtorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.Nm" : "DebtorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.StrtNm" : "DebtorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.BldgNb" : "DebtorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.BldgNm" : "DebtorAgent.PostalAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Flr" : "DebtorAgent.PostalAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Room" : "DebtorAgent.PostalAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.PstCd" : "DebtorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.TwnNm" : "DebtorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "DebtorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Ctry" : "DebtorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "CreditorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "CreditorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.Nm" : "CreditorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.StrtNm" : "CreditorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.BldgNb" : "CreditorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.BldgNm" : "CreditorAgent.PostalAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Flr" : "CreditorAgent.PostalAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Room" : "CreditorAgent.PostalAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.PstCd" : "CreditorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.TwnNm" : "CreditorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "CreditorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Ctry" : "CreditorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.Nm" : "CreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.StrtNm" : "CreditorPostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.BldgNb" : "CreditorPostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.BldgNm" : "CreditorPostalAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Flr" : "CreditorPostalAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Room" : "CreditorPostalAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.PstCd" : "CreditorPostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.TwnNm" : "CreditorPostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.CtrySubDvsn" : "CreditorPostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Ctry" : "CreditorPostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.Nm" : "UltimateCreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.StrtNm" : "UltimateCreditorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.BldgNb" : "UltimateCreditorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.BldgNm" : "UltimateCreditorAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Flr" : "UltimateCreditorAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Room" : "UltimateCreditorAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.PstCd" : "UltimateCreditorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.TwnNm" : "UltimateCreditorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.CtrySubDvsn" : "UltimateCreditorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Ctry" : "UltimateCreditorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.IBAN" : "CreditorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.Othr.Id" : "CreditorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Purp.Cd" : "PurposeOfPayment",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtId" : "RelatedRemittanceInfo.RemittanceId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnDtls[0].Mtd" : "RelatedRemittanceInfo.Method",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnDtls[0].ElctrncAdr" : "RelatedRemittanceInfo.ElectronicAddress",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Ustrd[0]" : "RemittanceInfor.UnstructuredRemitInfo",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd" : "RemittanceInfor.CodeOrProprietary",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Nb" : "RemittanceInfor.Number",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt" : "RemittanceInfor.RelatedDate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Cdtr.TaxId" : "RemittanceInfor.TaxDetail.TaxId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Tp" : "RemittanceInfor.TaxDetail.TaxTypeCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd.Yr" : "RemittanceInfor.TaxDetail.TaxPeriodYear",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd.Tp" : "RemittanceInfor.TaxDetail.TaxperiodTimeFrame",
	}
}
func PathMapV9() map[string]any {
	return PathMapV8()
}
func PathMapV10() map[string]any {
	return PathMapV8()
}
func PathMapV11() map[string]any {
	return PathMapV8()
}
func PathMapV12() map[string]any {
	return map[string]any{
		"FIToFICstmrCdtTrf.GrpHdr.MsgId" : "MessageId",
		"FIToFICstmrCdtTrf.GrpHdr.CreDtTm" : "CreatedDateTime",
		"FIToFICstmrCdtTrf.GrpHdr.NbOfTxs" : "NumberOfTransactions",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.SttlmMtd" : "SettlementMethod",
		"FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd" : "CommonClearingSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.InstrId" : "InstructionId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.EndToEndId" : "EndToEndId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtId.UETR" : "UniqueEndToEndTransactionRef",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.SvcLvl[0].Cd" : "SericeLevel",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].PmtTpInf.LclInstrm.Prtry" : "InstrumentPropCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Value" : "InterBankSettAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmAmt.Ccy" : "InterBankSettAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrBkSttlmDt" : "InterBankSettDate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Value" : "InstructedAmount.Amount",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAmt.Ccy" : "InstructedAmount.Currency",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].XchgRate" : "ExchangeRate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgBr" : "ChargeBearer",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].ChrgsInf : ChargesInfo": map[string]string{
			"Amt.Value": "Amount.Amount",
			"Amt.Ccy":   "Amount.Currency",
			"Agt.FinInstnId.BICFI": "BusinessIdCode",
		},
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructingAgents.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructingAgents.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "InstructedAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId" : "InstructedAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId" : "IntermediaryAgent1Id",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.Nm" : "DebtorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.StrtNm" : "DebtorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.BldgNb" : "DebtorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.BldgNm" : "DebtorAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Flr" : "DebtorAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Room" : "DebtorAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.PstCd" : "DebtorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.TwnNm" : "DebtorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.CtrySubDvsn" : "DebtorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Dbtr.PstlAdr.Ctry" : "DebtorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.IBAN" : "DebtorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAcct.Id.Othr.Id" : "DebtorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "DebtorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "DebtorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.Nm" : "DebtorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.StrtNm" : "DebtorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.BldgNb" : "DebtorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.BldgNm" : "DebtorAgent.PostalAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Flr" : "DebtorAgent.PostalAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Room" : "DebtorAgent.PostalAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.PstCd" : "DebtorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.TwnNm" : "DebtorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "DebtorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].DbtrAgt.FinInstnId.PstlAdr.Ctry" : "DebtorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd" : "CreditorAgent.PaymentSysCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId" : "CreditorAgent.PaymentSysMemberId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.Nm" : "CreditorAgent.BankName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.StrtNm" : "CreditorAgent.PostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.BldgNb" : "CreditorAgent.PostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.BldgNm" : "CreditorAgent.PostalAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Flr" : "CreditorAgent.PostalAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Room" : "CreditorAgent.PostalAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.PstCd" : "CreditorAgent.PostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.TwnNm" : "CreditorAgent.PostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.CtrySubDvsn" : "CreditorAgent.PostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAgt.FinInstnId.PstlAdr.Ctry" : "CreditorAgent.PostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.Nm" : "CreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.StrtNm" : "CreditorPostalAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.BldgNb" : "CreditorPostalAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.BldgNm" : "CreditorPostalAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Flr" : "CreditorPostalAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Room" : "CreditorPostalAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.PstCd" : "CreditorPostalAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.TwnNm" : "CreditorPostalAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.CtrySubDvsn" : "CreditorPostalAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Cdtr.PstlAdr.Ctry" : "CreditorPostalAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.Nm" : "UltimateCreditorName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.StrtNm" : "UltimateCreditorAddress.StreetName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.BldgNb" : "UltimateCreditorAddress.BuildingNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.BldgNm" : "UltimateCreditorAddress.BuildingName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Flr" : "UltimateCreditorAddress.Floor",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Room" : "UltimateCreditorAddress.RoomNumber",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.PstCd" : "UltimateCreditorAddress.PostalCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.TwnNm" : "UltimateCreditorAddress.TownName",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.CtrySubDvsn" : "UltimateCreditorAddress.Subdivision",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].UltmtCdtr.PstlAdr.Ctry" : "UltimateCreditorAddress.Country",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.IBAN" : "CreditorIBAN",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].CdtrAcct.Id.Othr.Id" : "CreditorOtherTypeId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].Purp.Cd" : "PurposeOfPayment",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtId" : "RelatedRemittanceInfo.RemittanceId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnDtls[0].Mtd" : "RelatedRemittanceInfo.Method",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RltdRmtInf[0].RmtLctnDtls[0].ElctrncAdr" : "RelatedRemittanceInfo.ElectronicAddress",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Ustrd[0]" : "RemittanceInfor.UnstructuredRemitInfo",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd" : "RemittanceInfor.CodeOrProprietary",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].Nb" : "RemittanceInfor.Number",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt.Dt" : "RemittanceInfor.RelatedDate",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Cdtr.TaxId" : "RemittanceInfor.TaxDetail.TaxId",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Tp" : "RemittanceInfor.TaxDetail.TaxTypeCode",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd.Yr" : "RemittanceInfor.TaxDetail.TaxPeriodYear",
		"FIToFICstmrCdtTrf.CdtTrfTxInf[0].RmtInf.Strd[0].TaxRmt.Rcrd[0].Prd.Tp" : "RemittanceInfor.TaxDetail.TaxperiodTimeFrame",
	}
}
