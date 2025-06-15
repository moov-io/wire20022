package FedwireFundsSystemResponse

type ADMI_011_001_VERSION string

const (
	ADMI_011_001_01 ADMI_011_001_VERSION = "ADMI.011.001.01"
)

var VersionNameSpaceMap = map[ADMI_011_001_VERSION]string{
	ADMI_011_001_01: "urn:iso:std:iso:20022:tech:xsd:admi.011.001.01",
}
var NameSpaceVersionMap = map[string]ADMI_011_001_VERSION{
	"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01": ADMI_011_001_01,
}
var VersionPathMap = map[ADMI_011_001_VERSION]map[string]any{
	ADMI_011_001_01: PathMapV1(),
}
