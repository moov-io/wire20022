package FedwireFundsAcknowledgement

type ADMI_007_001_VERSION string

const (
	ADMI_007_001_01 ADMI_007_001_VERSION = "ADMI.007.001.01"
)

var VersionNameSpaceMap = map[ADMI_007_001_VERSION]string{
	ADMI_007_001_01: "urn:iso:std:iso:20022:tech:xsd:admi.007.001.01",
}
var NameSpaceVersionMap = map[string]ADMI_007_001_VERSION{
	"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01": ADMI_007_001_01,
}
var VersionPathMap = map[ADMI_007_001_VERSION]map[string]any{
	ADMI_007_001_01: pathMapV1(),
}
