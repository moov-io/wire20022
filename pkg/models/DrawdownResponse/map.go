package DrawdownResponse

func pathMapV1() map[string]any {
	return pathMapV6()
}
func pathMapV2() map[string]any {
	return pathMapV6()
}
func pathMapV3() map[string]any {
	return pathMapV6()
}
func pathMapV4() map[string]any {
	return pathMapV6()
}
func pathMapV5() map[string]any {
	return pathMapV6()
}
func pathMapV6() map[string]any {
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
func pathMapV7() map[string]any {
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
func pathMapV8() map[string]any {
	return pathMapV7()
}
func pathMapV9() map[string]any {
	return pathMapV7()
}
func pathMapV10() map[string]any {
	return pathMapV7()
}
