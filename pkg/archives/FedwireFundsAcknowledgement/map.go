package FedwireFundsAcknowledgement

import (
	"time"

	"github.com/moov-io/fedwire20022/gen/FedwireFundsAcknowledgement/admi_007_001_01"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type Message struct {
	Data MessageModel
	Doc  admi_007_001_01.Document
}

func Convert() {
	msg := Message{}
	msg.Data.MessageId = string(msg.Doc.RctAck.MsgId.MsgId)
	msg.Data.CreatedDateTime = time.Time(*msg.Doc.RctAck.MsgId.CreDtTm)
	msg.Data.RelationReference = string(msg.Doc.RctAck.Rpt[0].RltdRef.Ref)
	msg.Data.ReferenceName = string(*msg.Doc.RctAck.Rpt[0].RltdRef.MsgNm)
	msg.Data.RequestHandling = Archive.RelatedStatusCode(msg.Doc.RctAck.Rpt[0].ReqHdlg.StsCd)
}

func PathMapV1() map[string]any {
	return map[string]any{
		"RctAck.MsgId.MsgId":          "MessageId",
		"RctAck.MsgId.CreDtTm":        "CreatedDateTime",
		"RctAck.Rpt[0].RltdRef.Ref":   "RelationReference",
		"RctAck.Rpt[0].RltdRef.MsgNm": "ReferenceName",
		"RctAck.Rpt[0].ReqHdlg.StsCd": "RequestHandling",
	}
}
