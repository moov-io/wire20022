package CustomerCreditTransfer

type PACS_008_001_VERSION string

const (
	PACS_008_001_02 PACS_008_001_VERSION = "pacs.008.001.02"
	PACS_008_001_03 PACS_008_001_VERSION = "pacs.008.001.03"
	PACS_008_001_04 PACS_008_001_VERSION = "pacs.008.001.04"
	PACS_008_001_05 PACS_008_001_VERSION = "pacs.008.001.05"
	PACS_008_001_06 PACS_008_001_VERSION = "pacs.008.001.06"
	PACS_008_001_07 PACS_008_001_VERSION = "pacs.008.001.07"
	PACS_008_001_08 PACS_008_001_VERSION = "pacs.008.001.08"
	PACS_008_001_09 PACS_008_001_VERSION = "pacs.008.001.09"
	PACS_008_001_10 PACS_008_001_VERSION = "pacs.008.001.10"
	PACS_008_001_11 PACS_008_001_VERSION = "pacs.008.001.11"
	PACS_008_001_12 PACS_008_001_VERSION = "pacs.008.001.12"
)

var VersionNameSpaceMap = map[PACS_008_001_VERSION]string{
	PACS_008_001_02: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.02",
	PACS_008_001_03: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.03",
	PACS_008_001_04: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.04",
	PACS_008_001_05: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.05",
	PACS_008_001_06: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.06",
	PACS_008_001_07: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.07",
	PACS_008_001_08: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08",
	PACS_008_001_09: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09",
	PACS_008_001_10: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.10",
	PACS_008_001_11: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.11",
	PACS_008_001_12: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.12",
}
var NameSpaceVersionMap = map[string]PACS_008_001_VERSION{
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.02": PACS_008_001_02,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.03": PACS_008_001_03,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.04": PACS_008_001_04,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.05": PACS_008_001_05,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.06": PACS_008_001_06,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.07": PACS_008_001_07,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08": PACS_008_001_08,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09": PACS_008_001_09,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.10": PACS_008_001_10,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.11": PACS_008_001_11,
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.12": PACS_008_001_12,
}

var VersionPathMap = map[PACS_008_001_VERSION]map[string]any{
	PACS_008_001_02: pathMapV2(),
	PACS_008_001_03: pathMapV3(),
	PACS_008_001_04: pathMapV4(),
	PACS_008_001_05: pathMapV5(),
	PACS_008_001_06: pathMapV6(),
	PACS_008_001_07: pathMapV7(),
	PACS_008_001_08: pathMapV8(),
	PACS_008_001_09: pathMapV9(),
	PACS_008_001_10: pathMapV10(),
	PACS_008_001_11: pathMapV11(),
	PACS_008_001_12: pathMapV12(),
}
