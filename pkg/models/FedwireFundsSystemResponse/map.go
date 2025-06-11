package FedwireFundsSystemResponse

func PathMapV1() map[string]any {
	return map[string]any{
		"SysEvtAck.MsgId":               "MessageId",
		"SysEvtAck.AckDtls.EvtCd":       "EventCode",
		"SysEvtAck.AckDtls.EvtParam[0]": "EventParam",
		"SysEvtAck.AckDtls.EvtTm":       "EventTime",
	}
}
