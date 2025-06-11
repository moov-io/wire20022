package FedwireFundsAcknowledgement

func PathMapV1() map[string]any {
	return map[string]any{
		"RctAck.MsgId.MsgId":          "MessageId",
		"RctAck.MsgId.CreDtTm":        "CreatedDateTime",
		"RctAck.Rpt[0].RltdRef.Ref":   "RelationReference",
		"RctAck.Rpt[0].RltdRef.MsgNm": "ReferenceName",
		"RctAck.Rpt[0].ReqHdlg.StsCd": "RequestHandling",
	}
}
