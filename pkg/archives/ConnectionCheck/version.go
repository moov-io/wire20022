package ConnectionCheck

type ADMI_004_001_VESION string

const (
	ADMI_004_001_01 ADMI_004_001_VESION = "admi.004.001.01"
	ADMI_004_001_02 ADMI_004_001_VESION = "admi.004.001.02"
)

var VersionNameSpaceMap = map[ADMI_004_001_VESION]string{
	ADMI_004_001_01: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.01",
	ADMI_004_001_02: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02",
}

var NameSpaceVersonMap = map[string]ADMI_004_001_VESION{
	"urn:iso:std:iso:20022:tech:xsd:admi.004.001.01": ADMI_004_001_01,
	"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02": ADMI_004_001_02,
}

var VersionPathMap = map[ADMI_004_001_VESION]map[string]string{
	ADMI_004_001_01: PathMapV1(),
	ADMI_004_001_02: PathMapV2(),
}
