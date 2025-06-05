package FedwireFundsSystemResponse

type ADMI_011_001_VESION string

const (
	ADMI_011_001_01 ADMI_011_001_VESION = "ADMI.011.001.01"
)

var VersionNameSpaceMap = map[ADMI_011_001_VESION]string{
	ADMI_011_001_01: "urn:iso:std:iso:20022:tech:xsd:admi.011.001.01",
}
var NameSpaceVersonMap = map[string]ADMI_011_001_VESION{
	"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01": ADMI_011_001_01,
}
var VersionPathMap = map[ADMI_011_001_VESION]map[string]any{
	ADMI_011_001_01: PathMapV1(),
}
