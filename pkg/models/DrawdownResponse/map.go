package DrawdownResponse

func PathMapV1() map[string]any {
	return PathMapV6()
}
func PathMapV2() map[string]any {
	return PathMapV6()
}
func PathMapV3() map[string]any {
	return PathMapV6()
}
func PathMapV4() map[string]any {
	return PathMapV6()
}
func PathMapV5() map[string]any {
	return PathMapV6()
}
func PathMapV6() map[string]any {
	return map[string]any{
		"CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId":                                            "MessageId",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.CreDtTm":                                          "CreatedDateTime",
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
func PathMapV7() map[string]any {
	return map[string]any{
		"CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId":                                            "MessageId",
		"CdtrPmtActvtnReqStsRpt.GrpHdr.CreDtTm":                                          "CreatedDateTime",
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
func PathMapV8() map[string]any {
	return PathMapV7()
}
func PathMapV9() map[string]any {
	return PathMapV7()
}
func PathMapV10() map[string]any {
	return PathMapV7()
}
