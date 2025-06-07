package FedwireFundsSystemResponse

import (
	"time"

	"github.com/moov-io/fedwire20022/gen/FedwireFundsSystemResponse/admi_011_001_01"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type Message struct {
	Data MessageModel
	Doc  admi_011_001_01.Document
}

func Convert() {
	msg := Message{}
	msg.Data.MessageId = string(msg.Doc.SysEvtAck.MsgId)
	msg.Data.EventCode = Archive.FundEventType(msg.Doc.SysEvtAck.AckDtls.EvtCd)
	msg.Data.EventParam = string(*msg.Doc.SysEvtAck.AckDtls.EvtParam[0])
	msg.Data.EventTime = time.Time(*msg.Doc.SysEvtAck.AckDtls.EvtTm)
}

func PathMapV1() map[string]any {
	return map[string]any{
		"SysEvtAck.MsgId":               "MessageId",
		"SysEvtAck.AckDtls.EvtCd":       "EventCode",
		"SysEvtAck.AckDtls.EvtParam[0]": "EventParam",
		"SysEvtAck.AckDtls.EvtTm":       "EventTime",
	}
}
