package PaymentReturn

type PACS_004_001_VERSION string

const (
	PACS_004_001_02 PACS_004_001_VERSION = "PACS.004.001.02"
	PACS_004_001_03 PACS_004_001_VERSION = "PACS.004.001.03"
	PACS_004_001_04 PACS_004_001_VERSION = "PACS.004.001.04"
	PACS_004_001_05 PACS_004_001_VERSION = "PACS.004.001.05"
	PACS_004_001_06 PACS_004_001_VERSION = "PACS.004.001.06"
	PACS_004_001_07 PACS_004_001_VERSION = "PACS.004.001.07"
	PACS_004_001_08 PACS_004_001_VERSION = "PACS.004.001.08"
	PACS_004_001_09 PACS_004_001_VERSION = "PACS.004.001.09"
	PACS_004_001_10 PACS_004_001_VERSION = "PACS.004.001.10"
	PACS_004_001_11 PACS_004_001_VERSION = "PACS.004.001.11"
	PACS_004_001_12 PACS_004_001_VERSION = "PACS.004.001.12"
	PACS_004_001_13 PACS_004_001_VERSION = "PACS.004.001.13"
)

var VersionNameSpaceMap = map[PACS_004_001_VERSION]string{
	PACS_004_001_02: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02",
	PACS_004_001_03: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.03",
	PACS_004_001_04: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.04",
	PACS_004_001_05: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.05",
	PACS_004_001_06: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.06",
	PACS_004_001_07: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.07",
	PACS_004_001_08: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.08",
	PACS_004_001_09: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.09",
	PACS_004_001_10: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10",
	PACS_004_001_11: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.11",
	PACS_004_001_12: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.12",
	PACS_004_001_13: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.13",
}
var NameSpaceVersionMap = map[string]PACS_004_001_VERSION{
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02": PACS_004_001_02,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.03": PACS_004_001_03,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.04": PACS_004_001_04,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.05": PACS_004_001_05,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.06": PACS_004_001_06,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.07": PACS_004_001_07,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.08": PACS_004_001_08,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.09": PACS_004_001_09,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10": PACS_004_001_10,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.11": PACS_004_001_11,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.12": PACS_004_001_12,
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.13": PACS_004_001_13,
}
var VersionPathMap = map[PACS_004_001_VERSION]map[string]any{
	PACS_004_001_02: PathMapV2(),
	PACS_004_001_03: PathMapV3(),
	PACS_004_001_04: PathMapV4(),
	PACS_004_001_05: PathMapV5(),
	PACS_004_001_06: PathMapV6(),
	PACS_004_001_07: PathMapV7(),
	PACS_004_001_08: PathMapV8(),
	PACS_004_001_09: PathMapV9(),
	PACS_004_001_10: PathMapV10(),
	PACS_004_001_11: PathMapV11(),
	PACS_004_001_12: PathMapV12(),
	PACS_004_001_13: PathMapV13(),
}
