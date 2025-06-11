package DrawdownRequest

type PAIN_013_001_VESION string

const (
	PAIN_013_001_01 PAIN_013_001_VESION = "pain.013.001.01"
	PAIN_013_001_02 PAIN_013_001_VESION = "pain.013.001.02"
	PAIN_013_001_03 PAIN_013_001_VESION = "pain.013.001.03"
	PAIN_013_001_04 PAIN_013_001_VESION = "pain.013.001.04"
	PAIN_013_001_05 PAIN_013_001_VESION = "pain.013.001.05"
	PAIN_013_001_06 PAIN_013_001_VESION = "pain.013.001.06"
	PAIN_013_001_07 PAIN_013_001_VESION = "pain.013.001.07"
	PAIN_013_001_08 PAIN_013_001_VESION = "pain.013.001.08"
	PAIN_013_001_09 PAIN_013_001_VESION = "pain.013.001.09"
	PAIN_013_001_10 PAIN_013_001_VESION = "pain.013.001.10"
)

var VersionNameSpaceMap = map[PAIN_013_001_VESION]string{
	PAIN_013_001_01: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.01",
	PAIN_013_001_02: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.02",
	PAIN_013_001_03: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.03",
	PAIN_013_001_04: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.04",
	PAIN_013_001_05: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.05",
	PAIN_013_001_06: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.06",
	PAIN_013_001_07: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.07",
	PAIN_013_001_08: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.08",
	PAIN_013_001_09: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.09",
	PAIN_013_001_10: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.10",
}
var NameSpaceVersonMap = map[string]PAIN_013_001_VESION{
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.01": PAIN_013_001_01,
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.02": PAIN_013_001_02,
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.03": PAIN_013_001_03,
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.04": PAIN_013_001_04,
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.05": PAIN_013_001_05,
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.06": PAIN_013_001_06,
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.07": PAIN_013_001_07,
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08": PAIN_013_001_08,
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.09": PAIN_013_001_09,
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.10": PAIN_013_001_10,
}

var VersionPathMap = map[PAIN_013_001_VESION]map[string]string{
	PAIN_013_001_01: PathMapV1(),
	PAIN_013_001_02: PathMapV2(),
	PAIN_013_001_03: PathMapV3(),
	PAIN_013_001_04: PathMapV4(),
	PAIN_013_001_05: PathMapV5(),
	PAIN_013_001_06: PathMapV6(),
	PAIN_013_001_07: PathMapV7(),
	PAIN_013_001_08: PathMapV8(),
	PAIN_013_001_09: PathMapV9(),
	PAIN_013_001_10: PathMapV10(),
}
