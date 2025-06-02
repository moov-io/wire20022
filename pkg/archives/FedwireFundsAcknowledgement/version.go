package FedwireFundsAcknowledgement

type ADMI_007_001_VESION string

const (
	ADMI_007_001_01 ADMI_007_001_VESION = "ADMI.007.001.01"
)

var VersionNameSpaceMap = map[ADMI_007_001_VESION]string{
	ADMI_007_001_01: "urn:iso:std:iso:20022:tech:xsd:admi.007.001.01",
}
var NameSpaceVersonMap = map[string]ADMI_007_001_VESION{
	"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01": ADMI_007_001_01,
}
var VersionPathMap = map[ADMI_007_001_VESION]map[string]any{
	ADMI_007_001_01: PathMapV1(),
}
