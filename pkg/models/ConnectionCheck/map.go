package ConnectionCheck

func PathMapV1() map[string]string {
	return map[string]string{
		"Admi00400101.EvtInf.EvtCd":       "EventType",
		"Admi00400101.EvtInf.EvtParam[0]": "EventParam",
		"Admi00400101.EvtInf.EvtTm":       "EventTime",
	}
}

func PathMapV2() map[string]string {
	return map[string]string{
		"SysEvtNtfctn.EvtInf.EvtCd":       "EventType",
		"SysEvtNtfctn.EvtInf.EvtParam[0]": "EventParam",
		"SysEvtNtfctn.EvtInf.EvtTm":       "EventTime",
	}
}
