package ReturnRequestResponse

func PathMapV3() map[string]any {
	return PathMapV8()
}
func PathMapV4() map[string]any {
	return PathMapV8()
}
func PathMapV5() map[string]any {
	return PathMapV8()
}
func PathMapV6() map[string]any {
	return PathMapV8()
}
func PathMapV7() map[string]any {
	return PathMapV8()
}
func PathMapV8() map[string]any {
	return map[string]any{
		"RsltnOfInvstgtn.Assgnmt.Id": "AssignmentId",
		"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Assigner.PaymentSysCode",
		"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId":        "Assigner.PaymentSysMemberId",
		"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Assignee.PaymentSysCode",
		"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId":        "Assignee.PaymentSysMemberId",
		"RsltnOfInvstgtn.Assgnmt.CreDtTm":                                        "AssignmentCreateTime",
		"RsltnOfInvstgtn.RslvdCase.Id":                                           "ResolvedCaseId",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Creator.PaymentSysCode",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Creator.PaymentSysMemberId",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.Nm":                      "Creator.BankName",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.StrtNm":          "Creator.PostalAddress.StreetName",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.BldgNb":          "Creator.PostalAddress.BuildingNumber",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.PstCd":           "Creator.PostalAddress.PostalCode",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.TwnNm":           "Creator.PostalAddress.TownName",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.CtrySubDvsn":     "Creator.PostalAddress.Subdivision",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.Ctry":            "Creator.PostalAddress.Country",
		"RsltnOfInvstgtn.Sts.Conf":                                               "Status",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgId":       "OriginalMessageId",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgNmId":     "OriginalMessageNameId",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlGrpInf.OrgnlCreDtTm":     "OriginalMessageCreateTime",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlInstrId":                 "OriginalInstructionId",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlEndToEndId":              "OriginalEndToEndId",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].CxlStsRsnInf[0].Orgtr.Nm":     "CancellationStatusReasonInfo.Originator",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].CxlStsRsnInf[0].Rsn.Cd":       "CancellationStatusReasonInfo.Reason",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].CxlStsRsnInf[0].AddtlInf[0]":  "CancellationStatusReasonInfo.AdditionalInfo",
	}
}
func PathMapV9() map[string]any {
	return map[string]any{
		"RsltnOfInvstgtn.Assgnmt.Id": "AssignmentId",
		"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Assigner.PaymentSysCode",
		"RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId":        "Assigner.PaymentSysMemberId",
		"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":  "Assignee.PaymentSysCode",
		"RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId":        "Assignee.PaymentSysMemberId",
		"RsltnOfInvstgtn.Assgnmt.CreDtTm":                                        "AssignmentCreateTime",
		"RsltnOfInvstgtn.RslvdCase.Id":                                           "ResolvedCaseId",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Creator.PaymentSysCode",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "Creator.PaymentSysMemberId",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.Nm":                      "Creator.BankName",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.StrtNm":          "Creator.PostalAddress.StreetName",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.BldgNb":          "Creator.PostalAddress.BuildingNumber",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.BldgNm":          "Creator.PostalAddress.BuildingName",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.Flr":             "Creator.PostalAddress.Floor",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.Room":            "Creator.PostalAddress.RoomNumber",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.PstCd":           "Creator.PostalAddress.PostalCode",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.TwnNm":           "Creator.PostalAddress.TownName",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.CtrySubDvsn":     "Creator.PostalAddress.Subdivision",
		"RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.PstlAdr.Ctry":            "Creator.PostalAddress.Country",
		"RsltnOfInvstgtn.Sts.Conf":                                               "Status",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgId":       "OriginalMessageId",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgNmId":     "OriginalMessageNameId",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlGrpInf.OrgnlCreDtTm":     "OriginalMessageCreateTime",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlInstrId":                 "OriginalInstructionId",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlEndToEndId":              "OriginalEndToEndId",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].OrgnlUETR":                    "EnhancedTransaction.OriginalUETR",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].CxlStsRsnInf[0].Orgtr.Nm":     "CancellationStatusReasonInfo.Originator",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].CxlStsRsnInf[0].Rsn.Cd":       "CancellationStatusReasonInfo.Reason",
		"RsltnOfInvstgtn.CxlDtls[0].TxInfAndSts[0].CxlStsRsnInf[0].AddtlInf[0]":  "CancellationStatusReasonInfo.AdditionalInfo",
	}
}
func PathMapV10() map[string]any {
	return PathMapV9()
}
func PathMapV11() map[string]any {
	return PathMapV9()
}
func PathMapV12() map[string]any {
	return PathMapV9()
}
